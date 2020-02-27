package main

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func gatherInfo(clientset *kubernetes.Clientset, nodeLabel *string) (clusterInfo ClusterInfo) {
	nodeInfo := make(map[string]NodeInfo)
	labelSlice := strings.Split(*nodeLabel, "=")
	nodeLabelKey := labelSlice[0]
	nodeLabelValue := ""
	if nodeLabelKey != "" {
		nodeLabelValue = labelSlice[1]
	}

	// List all nodes
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
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
			if cpu.Value() > clusterInfo.NminusCPU.Value() {
				clusterInfo.NminusCPU = *cpu
				clusterInfo.NminusMemory = *mem
				clusterInfo.NminusPods = *pods
			}
			clusterInfo.ClusterAllocatableMemory.Add(*mem)
			clusterInfo.ClusterAllocatableCPU.Add(*cpu)
			clusterInfo.ClusterAllocatablePods.Add(*pods)
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
	// Add all the quotas up
	for _, v := range quotas.Items {
		limitmem := v.Spec.Hard[corev1.ResourceLimitsMemory]
		limitcpu := v.Spec.Hard[corev1.ResourceLimitsCPU]
		requestmem := v.Spec.Hard[corev1.ResourceRequestsMemory]
		requestcpu := v.Spec.Hard[corev1.ResourceRequestsCPU]
		pods := v.Spec.Hard[corev1.ResourcePods]
		clusterInfo.RqclusterAllocatedLimitsMemory.Add(limitmem)
		clusterInfo.RqclusterAllocatedLimitsCPU.Add(limitcpu)
		clusterInfo.RqclusterAllocatedPods.Add(pods)
		clusterInfo.RqclusterAllocatedRequestsMemory.Add(requestmem)
		clusterInfo.RqclusterAllocatedRequestsCPU.Add(requestcpu)
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
	clusterInfo.NodeInfo = nodeInfo
	return clusterInfo

}
