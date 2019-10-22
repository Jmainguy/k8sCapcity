package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"reflect"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	nodeLabel := flag.String("nodelabel", "", "Label to match for nodes, if blank grab all nodes")
	flag.Parse()

	labelSlice := strings.Split(*nodeLabel, "=")
	nodeLabelKey := labelSlice[0]
	nodeLabelValue := ""
	if nodeLabelKey != "" {
		nodeLabelValue = labelSlice[1]
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// List all nodes
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	clusterAllocatableMemory := &resource.Quantity{}
	clusterAllocatableCPU := &resource.Quantity{}
	clusterAllocatablePods := &resource.Quantity{}
	var nodesWeCareAbout []string
	if nodeLabelKey != "" {
		for _, v := range nodes.Items {
			for label, value := range v.ObjectMeta.Labels {
				if label == nodeLabelKey {
					if value == nodeLabelValue {
						nodesWeCareAbout = append(nodesWeCareAbout, v.Name)
					}
				}
			}
		}
		fmt.Printf("There are %d nodes in this cluster\n", len(nodesWeCareAbout))
	} else {
		fmt.Printf("There are %d nodes in this cluster\n", len(nodes.Items))
	}
	for _, v := range nodes.Items {
		if nodeLabelKey != "" {
			for label, value := range v.ObjectMeta.Labels {
				if label == nodeLabelKey {
					if value == nodeLabelValue {
						fmt.Println("================")
						fmt.Printf("Node: %s\n", v.Name)
						cpu := v.Status.Allocatable.Cpu()
						mem := v.Status.Allocatable.Memory()
						pods := v.Status.Allocatable.Pods()
						am := mem.ScaledValue(resource.Giga)
						fmt.Printf("Allocatable CPU: %s\n", cpu)
						fmt.Printf("Allocatable Memory: %s (%dGB)\n", mem, am)
						fmt.Printf("Allocatable Pods: %s\n", pods)
						clusterAllocatableMemory.Add(*mem)
						clusterAllocatableCPU.Add(*cpu)
						clusterAllocatablePods.Add(*pods)
					}
				}
			}
		} else {
			fmt.Println("================")
			fmt.Printf("Node: %s\n", v.Name)
			cpu := v.Status.Allocatable.Cpu()
			mem := v.Status.Allocatable.Memory()
			pods := v.Status.Allocatable.Pods()
			am := mem.ScaledValue(resource.Giga)
			fmt.Printf("Allocatable CPU: %s\n", cpu)
			fmt.Printf("Allocatable Memory: %s (%dGB)\n", mem, am)
			fmt.Printf("Allocatable Pods: %s\n", pods)
			clusterAllocatableMemory.Add(*mem)
			clusterAllocatableCPU.Add(*cpu)
			clusterAllocatablePods.Add(*pods)
		}

	}

	// List quotas
	quotas, err := clientset.CoreV1().ResourceQuotas("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	clusterAllocatedLimitsMemory := &resource.Quantity{}
	clusterAllocatedLimitsCPU := &resource.Quantity{}
	clusterAllocatedPods := &resource.Quantity{}
	clusterAllocatedRequestsMemory := &resource.Quantity{}
	clusterAllocatedRequestsCPU := &resource.Quantity{}
	// Add all the quotas up
	for _, v := range quotas.Items {
		limitmem := v.Spec.Hard[corev1.ResourceLimitsMemory]
		limitcpu := v.Spec.Hard[corev1.ResourceLimitsCPU]
		requestmem := v.Spec.Hard[corev1.ResourceRequestsMemory]
		requestcpu := v.Spec.Hard[corev1.ResourceRequestsCPU]
		pods := v.Spec.Hard[corev1.ResourcePods]
		clusterAllocatedLimitsMemory.Add(limitmem)
		clusterAllocatedLimitsCPU.Add(limitcpu)
		clusterAllocatedPods.Add(pods)
		clusterAllocatedRequestsMemory.Add(requestmem)
		clusterAllocatedRequestsCPU.Add(requestcpu)
	}

	fmt.Println("================")
	cwam := clusterAllocatableMemory.ScaledValue(resource.Giga)
	fmt.Printf("ClusterWide Allocatable Memory: %s (%dGB)\n", clusterAllocatableMemory, cwam)
	fmt.Printf("ClusterWide Allocatable CPU: %s\n", clusterAllocatableCPU)
	fmt.Printf("ClusterWide Allocatable Pods: %s\n", clusterAllocatablePods)
	fmt.Println("================")
	cwalm := clusterAllocatedLimitsMemory.ScaledValue(resource.Giga)
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.Memory: %s (%dGB)\n", clusterAllocatedLimitsMemory, cwalm)
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.CPU: %d\n", clusterAllocatedLimitsCPU.AsDec())
	fmt.Printf("ResourceQuota ClusterWide Allocated Pods: %d\n", clusterAllocatedPods.AsDec())
	fmt.Println("================")
	cwarm := clusterAllocatedRequestsMemory.ScaledValue(resource.Giga)
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.Memory: %s (%dGB)\n", clusterAllocatedRequestsMemory, cwarm)
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.CPU: %d\n", clusterAllocatedRequestsCPU.AsDec())

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
