package main

import (
	"fmt"

	resource "k8s.io/apimachinery/pkg/api/resource"
)

func toGib(rq resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Giga)) / 1.074)
	return result
}

func toMib(rq resource.Quantity) (result int64) {
	result = int64(float64(rq.ScaledValue(resource.Mega)) / 1.074)
	return result
}

func humanMode(clusterInfo ClusterInfo) {

	fmt.Printf("There are %d nodes in this cluster\n", len(clusterInfo.NodeInfo))

	clusterUsedCPURequests := resource.Quantity{}
	clusterUsedCPU := resource.Quantity{}
	clusterUsedMemory := resource.Quantity{}
	clusterUsedMemoryRequests := resource.Quantity{}
	clusterUsedMemoryLimits := resource.Quantity{}
	var clusterUsedPods int64
	for node, info := range clusterInfo.NodeInfo {
		if info.PrintOutput == true {
			fmt.Println("================")
			fmt.Printf("NodeName: %s\n", node)
			fmt.Printf("Allocatable CPU: %s\n", &info.AllocatableCPU)
			fmt.Printf("Allocatable Memory: %dGiB\n", toGib(info.AllocatableMemory))
			fmt.Printf("Allocatable Pods: %s\n", &info.AllocatablePods)
			fmt.Println("----------------")
			fmt.Printf("Used CPU: %s\n", &info.UsedCPU)
			fmt.Printf("Used Memory: %dGiB\n", toGib(info.UsedMemory))
			fmt.Printf("Used Pods: %d\n", info.UsedPods)
			fmt.Printf("Used CPU Requests: %s\n", &info.UsedCPURequests)
			fmt.Printf("Used Memory Requests: %dGiB\n", toGib(info.UsedMemoryRequests))
			fmt.Println("----------------")

			AvailbleCPURequests := resource.Quantity{}
			AvailableMemoryRequests := resource.Quantity{}

			AvailbleCPURequests = info.AllocatableCPU
			AvailbleCPURequests.Sub(info.UsedCPURequests)
			fmt.Printf("Available CPU Requests: %s\n", &AvailbleCPURequests)

			AvailableMemoryRequests = info.AllocatableMemory
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
	fmt.Printf("ClusterWide Used CPU: %d\n", clusterUsedCPU.Value())
	fmt.Printf("ClusterWide Used Memory: %dGiB\n", toGib(clusterUsedMemory))
	fmt.Printf("ClusterWide Used Pods: %d\n", clusterUsedPods)
	fmt.Printf("ClusterWide Used CPU Requests: %d\n", clusterUsedCPURequests.Value())
	fmt.Printf("ClusterWide Used Memory Requests: %dGiB\n", toGib(clusterUsedMemoryRequests))

}
