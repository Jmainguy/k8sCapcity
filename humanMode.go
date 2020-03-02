package main

import (
	"fmt"

	resource "k8s.io/apimachinery/pkg/api/resource"
)

func toGib(rq resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Giga)) / 1.073741824)
	return result
}

func toMib(rq resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Mega)) / 1.048576)
	return result
}

func humanMode(clusterInfo ClusterInfo) {

	fmt.Printf("There are %d nodes in this cluster\n", len(clusterInfo.NodeInfo))

	for name, node := range clusterInfo.NodeInfo {
		if node.PrintOutput == true {
			fmt.Println("================")
			fmt.Printf("NodeName: %s\n", name)
			fmt.Printf("Allocatable CPU: %s\n", &node.AllocatableCPU)
			fmt.Printf("Allocatable Memory: %dGiB\n", toGib(node.AllocatableMemory))
			fmt.Printf("Allocatable Pods: %s\n", &node.AllocatablePods)
			fmt.Println("----------------")
			fmt.Printf("Used CPU: %s\n", &node.UsedCPU)
			fmt.Printf("Used Memory: %dGiB\n", toGib(node.UsedMemory))
			fmt.Printf("Used Pods: %d\n", node.UsedPods)
			fmt.Printf("Used CPU Requests: %s\n", &node.UsedCPURequests)
			fmt.Printf("Used Memory Requests: %dGiB\n", toGib(node.UsedMemoryRequests))
			fmt.Println("----------------")

			AvailbleCPURequests := resource.Quantity{}
			AvailableMemoryRequests := resource.Quantity{}

			AvailbleCPURequests = node.AllocatableCPU
			AvailbleCPURequests.Sub(node.UsedCPURequests)
			fmt.Printf("Available CPU Requests: %s\n", &AvailbleCPURequests)

			AvailableMemoryRequests = node.AllocatableMemory
			AvailableMemoryRequests.Sub(node.UsedMemoryRequests)
			fmt.Printf("Available Memory Requests: %dGiB\n", toGib(AvailableMemoryRequests))

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
	fmt.Printf("ClusterWide Allocatable Memory: %dGiB\n", toGib(clusterInfo.ClusterAllocatableMemory))
	fmt.Printf("ClusterWide Allocatable CPU: %s\n", &clusterInfo.ClusterAllocatableCPU)
	fmt.Printf("ClusterWide Allocatable Pods: %d\n", clusterInfo.ClusterAllocatablePods.Value())
	fmt.Println("================")
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.Memory: %dGiB\n", toGib(clusterInfo.RqclusterAllocatedLimitsMemory))
	fmt.Printf("ResourceQuota ClusterWide Allocated Limits.CPU: %d\n", clusterInfo.RqclusterAllocatedLimitsCPU.AsDec())
	fmt.Printf("ResourceQuota ClusterWide Allocated Pods: %d\n", clusterInfo.RqclusterAllocatedPods.Value())
	fmt.Println("================")
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.Memory: %dGiB\n", toGib(clusterInfo.RqclusterAllocatedRequestsMemory))
	fmt.Printf("ResourceQuota ClusterWide Allocated Requests.CPU: %d\n", clusterInfo.RqclusterAllocatedRequestsCPU.AsDec())
	fmt.Println("----------------")
	fmt.Printf("ClusterWide Used CPU: %d\n", clusterInfo.ClusterUsedCPU.Value())
	fmt.Printf("ClusterWide Used Memory: %dGiB\n", toGib(clusterInfo.ClusterUsedMemory))
	fmt.Printf("ClusterWide Used Pods: %d\n", clusterInfo.ClusterUsedPods)
	fmt.Printf("ClusterWide Used CPU Requests: %d\n", clusterInfo.ClusterUsedCPURequests.Value())
	fmt.Printf("ClusterWide Used Memory Requests: %dGiB\n", toGib(clusterInfo.ClusterUsedMemoryRequests))

}
