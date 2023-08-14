package main

import (
	"fmt"

	resource "k8s.io/apimachinery/pkg/api/resource"
)

func toGib(rq resource.Quantity) (result float64) {
	mib := toMib(rq)
	result = float64(mib) / 1024
	return result
}

func toMib(rq resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Mega)) / 1.048576)
	return result
}

func toMibFromByte(bytes int64) (mib int64) {
	kib := int64(float64(bytes) / 1024)
	mib = int64(float64(kib) / 1024)
	return mib
}

func toGibFromByte(bytes int64) (gib float64) {
	kib := int64(float64(bytes) / 1024)
	mib := int64(float64(kib) / 1024)
	gib = float64(mib) / 1024
	return gib
}

func humanMode(clusterInfo ClusterInfo) {

	fmt.Printf("There are %d nodes in this cluster\n", len(clusterInfo.NodeInfo))

	for name, node := range clusterInfo.NodeInfo {
		if node.PrintOutput {
			fmt.Println("================")
			fmt.Printf("NodeName: %s\n", name)
			fmt.Printf("Allocatable CPU: %s\n", &node.AllocatableCPU)
			fmt.Printf("Allocatable Memory: %.1fGiB\n", toGib(node.AllocatableMemory))
			fmt.Printf("Allocatable Pods: %s\n", &node.AllocatablePods)
			fmt.Println("----------------")
			fmt.Printf("Used CPU: %s\n", &node.UsedCPU)
			fmt.Printf("Used Memory: %.1fGiB\n", toGib(node.UsedMemory))
			fmt.Printf("Used Pods: %d\n", node.UsedPods)
			fmt.Printf("Used CPU Requests: %s\n", &node.UsedCPURequests)
			fmt.Printf("Used Memory Requests: %.1fGiB\n", toGib(node.UsedMemoryRequests))
			fmt.Println("----------------")

			AvailbleCPURequests := resource.Quantity{}
			AvailableMemoryRequests := resource.Quantity{}

			AvailbleCPURequests = node.AllocatableCPU
			AvailbleCPURequests.Sub(node.UsedCPURequests)
			fmt.Printf("Available CPU Requests: %s\n", &AvailbleCPURequests)

			AvailableMemoryRequests = node.AllocatableMemory
			AvailableMemoryRequests.Sub(node.UsedMemoryRequests)
			fmt.Printf("Available Memory Requests: %.1fGiB\n", toGib(AvailableMemoryRequests))

			AvailablePods, _ := node.AllocatablePods.AsInt64()
			AvailablePods = AvailablePods - node.UsedPods
			fmt.Printf("Available Pods: %d\n", AvailablePods)
			// Add to cluster total
			clusterInfo.ClusterUsedCPURequests.Add(node.UsedCPURequests)
			clusterInfo.ClusterUsedCPU.Add(node.UsedCPU)
			clusterInfo.ClusterUsedMemoryRequests.Add(node.UsedMemoryRequests)
			clusterInfo.ClusterUsedMemoryLimits.Add(node.UsedMemoryLimits)
			clusterInfo.ClusterUsedMemory.Add(node.UsedMemory)
			clusterInfo.ClusterUsedPods = clusterInfo.ClusterUsedPods + node.UsedPods
		}
	}
	fmt.Println("================")
	fmt.Printf("ClusterWide Allocatable Memory: %.1fGiB\n", toGib(clusterInfo.ClusterAllocatableMemory))
	fmt.Printf("ClusterWide Allocatable CPU: %s\n", &clusterInfo.ClusterAllocatableCPU)
	fmt.Printf("ClusterWide Allocatable Pods: %d\n", clusterInfo.ClusterAllocatablePods.Value())
	fmt.Println("================")
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.Memory: %.1fGiB\n", toGib(clusterInfo.RqclusterAllocatedLimitsMemory))
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.CPU: %d\n", clusterInfo.RqclusterAllocatedLimitsCPU.AsDec())
	fmt.Printf("ResourceQuota ClusterWide Allocated Pods: %d\n", clusterInfo.RqclusterAllocatedPods.Value())
	fmt.Println("================")
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.Memory: %.1fGiB\n", toGib(clusterInfo.RqclusterAllocatedRequestsMemory))
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.CPU: %d\n", clusterInfo.RqclusterAllocatedRequestsCPU.AsDec())
	fmt.Println("----------------")
	fmt.Printf("ClusterWide Used CPU: %d\n", clusterInfo.ClusterUsedCPU.Value())
	fmt.Printf("ClusterWide Used Memory: %.1fGiB\n", toGib(clusterInfo.ClusterUsedMemory))
	fmt.Printf("ClusterWide Used Pods: %d\n", clusterInfo.ClusterUsedPods)
	fmt.Printf("ClusterWide Used CPU Requests: %d\n", clusterInfo.ClusterUsedCPURequests.Value())
	fmt.Printf("ClusterWide Used Memory Requests: %.1fGiB\n", toGib(clusterInfo.ClusterUsedMemoryRequests))

}
