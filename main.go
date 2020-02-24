package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsv1b1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"strings"
)

func toGib(rq *resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Giga)) / 1.074)
	return result
}

func toMib(rq *resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Mega)) / 1.074)
	return result
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getNodeMetrics(clientset *kubernetes.Clientset) (nodeMetricList *metricsv1b1.NodeMetricsList) {
	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &nodeMetricList); err != nil {
		panic(err)
	}
	return nodeMetricList
}

func getPodMetrics(clientset *kubernetes.Clientset) (podMetricList *metricsv1b1.PodMetricsList) {
	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &podMetricList); err != nil {
		panic(err)
	}
	return podMetricList
}

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
	nameSpace := flag.String("namespace", "", "Namespace to grab capacity usage from")
	daemonMode := flag.Bool("daemon", false, "Run in daemon mode")
	flag.Parse()

	nodeInfo := make(map[string]NodeInfo)
	containerInfo := make(map[string]ContainerInfo)
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

	// BreakOut to namespace if asked
	if *nameSpace != "" {
		namespaceMode(clientset, nameSpace, containerInfo)
	}

	// List all nodes
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	clusterAllocatableMemory := &resource.Quantity{}
	clusterAllocatableCPU := &resource.Quantity{}
	clusterAllocatablePods := &resource.Quantity{}
	if nodeLabelKey != "" {
		for _, v := range nodes.Items {
			for label, value := range v.ObjectMeta.Labels {
				if label == nodeLabelKey {
					if value == nodeLabelValue {
						node := nodeInfo[v.Name]
						node.PrintOutput = true
						nodeInfo[v.Name] = node
					}
				}
			}
		}
	} else {
		for _, v := range nodes.Items {
			node := nodeInfo[v.Name]
			node.PrintOutput = true
			nodeInfo[v.Name] = node
		}
	}

	for _, v := range nodes.Items {
		if nodeInfo[v.Name].PrintOutput == true {
			cpu := v.Status.Allocatable.Cpu()
			mem := v.Status.Allocatable.Memory()
			pods := v.Status.Allocatable.Pods()
			clusterAllocatableMemory.Add(*mem)
			clusterAllocatableCPU.Add(*cpu)
			clusterAllocatablePods.Add(*pods)
			node := nodeInfo[v.Name]
			node.AllocatableCPU = *v.Status.Allocatable.Cpu()
			node.AllocatableMemory = *v.Status.Allocatable.Memory()
			node.AllocatablePods = *v.Status.Allocatable.Pods()
			nodeInfo[v.Name] = node
		}

	}

	// List quotas
	quotas, err := clientset.CoreV1().ResourceQuotas("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	rqclusterAllocatedLimitsMemory := &resource.Quantity{}
	rqclusterAllocatedLimitsCPU := &resource.Quantity{}
	rqclusterAllocatedPods := &resource.Quantity{}
	rqclusterAllocatedRequestsMemory := &resource.Quantity{}
	rqclusterAllocatedRequestsCPU := &resource.Quantity{}
	// Add all the quotas up
	for _, v := range quotas.Items {
		limitmem := v.Spec.Hard[corev1.ResourceLimitsMemory]
		limitcpu := v.Spec.Hard[corev1.ResourceLimitsCPU]
		requestmem := v.Spec.Hard[corev1.ResourceRequestsMemory]
		requestcpu := v.Spec.Hard[corev1.ResourceRequestsCPU]
		pods := v.Spec.Hard[corev1.ResourcePods]
		rqclusterAllocatedLimitsMemory.Add(limitmem)
		rqclusterAllocatedLimitsCPU.Add(limitcpu)
		rqclusterAllocatedPods.Add(pods)
		rqclusterAllocatedRequestsMemory.Add(requestmem)
		rqclusterAllocatedRequestsCPU.Add(requestcpu)
	}

	nodeMetricList := getNodeMetrics(clientset)
	for _, metricNode := range nodeMetricList.Items {
		cpuUsed := metricNode.Usage.Cpu()
		memUsed := metricNode.Usage.Memory()
		node := nodeInfo[metricNode.Name]
		node.UsedCPU = *cpuUsed
		node.UsedMemory = *memUsed
		nodeInfo[metricNode.Name] = node
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	for _, pod := range pods.Items {
		node := nodeInfo[pod.Spec.NodeName]
		if pod.Status.Phase != "Failed" {
			if pod.Status.Phase != "Succeeded" {
				for _, container := range pod.Spec.Containers {
					crrm := container.Resources.Requests.Memory()
					crlm := container.Resources.Limits.Memory()
					crrc := container.Resources.Requests.Cpu()
					UsedMemRequests := &resource.Quantity{}
					UsedMemLimits := &resource.Quantity{}
					UsedCPURequests := &resource.Quantity{}
					UsedMemRequests.Add(node.UsedMemoryRequests)
					UsedMemRequests.Add(*crrm)
					UsedMemLimits.Add(node.UsedMemoryLimits)
					UsedMemLimits.Add(*crlm)
					UsedCPURequests.Add(node.UsedCPURequests)
					UsedCPURequests.Add(*crrc)
					node.UsedMemoryRequests = *UsedMemRequests
					node.UsedMemoryLimits = *UsedMemLimits
					node.UsedCPURequests = *UsedCPURequests
				}
				node.UsedPods += 1
			}
		}
		nodeInfo[pod.Spec.NodeName] = node
	}

	if *daemonMode {
		runDaemonMode(nodeInfo, clusterAllocatableMemory, clusterAllocatableCPU, clusterAllocatablePods, rqclusterAllocatedLimitsMemory, rqclusterAllocatedLimitsCPU, rqclusterAllocatedPods, rqclusterAllocatedRequestsMemory, rqclusterAllocatedRequestsCPU, *nodeLabel)
		os.Exit(0)
	}
	fmt.Printf("There are %d nodes in this cluster\n", len(nodeInfo))

	clusterUsedCPURequests := &resource.Quantity{}
	clusterUsedCPU := &resource.Quantity{}
	clusterUsedMemory := &resource.Quantity{}
	clusterUsedMemoryRequests := &resource.Quantity{}
	clusterUsedMemoryLimits := &resource.Quantity{}
	var clusterUsedPods int64
	for node, info := range nodeInfo {
		if info.PrintOutput == true {
			fmt.Println("================")
			fmt.Printf("NodeName: %s\n", node)
			fmt.Printf("Allocatable CPU: %s\n", &info.AllocatableCPU)
			fmt.Printf("Allocatable Memory: %dGiB\n", toGib(&info.AllocatableMemory))
			fmt.Printf("Allocatable Pods: %s\n", &info.AllocatablePods)
			fmt.Println("----------------")
			fmt.Printf("Used CPU: %s\n", &info.UsedCPU)
			fmt.Printf("Used Memory: %dGiB\n", toGib(&info.UsedMemory))
			fmt.Printf("Used Pods: %d\n", info.UsedPods)
			fmt.Printf("Used CPU Requests: %s\n", &info.UsedCPURequests)
			fmt.Printf("Used Memory Requests: %dGiB\n", toGib(&info.UsedMemoryRequests))
			fmt.Println("----------------")

			AvailbleCPURequests := &resource.Quantity{}
			AvailableMemoryRequests := &resource.Quantity{}

			AvailbleCPURequests = &info.AllocatableCPU
			AvailbleCPURequests.Sub(info.UsedCPURequests)
			fmt.Printf("Available CPU Requests: %s\n", AvailbleCPURequests)

			AvailableMemoryRequests = &info.AllocatableMemory
			AvailableMemoryRequests.Sub(info.UsedMemoryRequests)
			fmt.Printf("Available Memory Requests: %dGiB\n", toGib(AvailableMemoryRequests))

			AvailablePods, _ := info.AllocatablePods.AsInt64()
			AvailablePods = AvailablePods - info.UsedPods
			fmt.Printf("Available Pods: %d\n", AvailablePods)
			// Add to cluster total
			clusterUsedCPURequests.Add(info.UsedCPURequests)
			clusterUsedCPU.Add(info.UsedCPU)
			clusterUsedMemoryRequests.Add(info.UsedMemoryRequests)
			clusterUsedMemoryLimits.Add(info.UsedMemoryLimits)
			clusterUsedMemory.Add(info.UsedMemory)
			clusterUsedPods = clusterUsedPods + info.UsedPods
		}
	}
	fmt.Println("================")
	fmt.Printf("ClusterWide Allocatable Memory: %dGiB\n", toGib(clusterAllocatableMemory))
	fmt.Printf("ClusterWide Allocatable CPU: %s\n", clusterAllocatableCPU)
	fmt.Printf("ClusterWide Allocatable Pods: %s\n", clusterAllocatablePods)
	fmt.Println("================")
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.Memory: %dGiB\n", toGib(rqclusterAllocatedLimitsMemory))
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.CPU: %d\n", rqclusterAllocatedLimitsCPU.AsDec())
	fmt.Printf("ResourceQuota ClusterWide Allocated Pods: %d\n", rqclusterAllocatedPods.AsDec())
	fmt.Println("================")
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.Memory: %dGiB\n", toGib(rqclusterAllocatedRequestsMemory))
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.CPU: %d\n", rqclusterAllocatedRequestsCPU.AsDec())
	fmt.Println("----------------")
	fmt.Printf("ClusterWide Used CPU: %s\n", clusterUsedCPU)
	fmt.Printf("ClusterWide Used Memory: %dGiB\n", toGib(clusterUsedMemory))
	fmt.Printf("ClusterWide Used Pods: %d\n", clusterUsedPods)
	fmt.Printf("ClusterWide Used CPU Requests: %s\n", clusterUsedCPURequests)
	fmt.Printf("ClusterWide Used Memory Requests: %dGiB\n", toGib(clusterUsedMemoryRequests))

}
