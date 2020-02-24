package main

import (
	"encoding/json"
	"fmt"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

func runDaemonMode(nodeInfo map[string]NodeInfo, clusterAllocatableMemory, clusterAllocatableCPU, clusterAllocatablePods, rqclusterAllocatedLimitsMemory, rqclusterAllocatedLimitsCPU, rqclusterAllocatedPods, rqclusterAllocatedRequestsMemory, rqclusterAllocatedRequestsCPU *resource.Quantity, nodeLabel string) {
	clusterUsedCPURequests := &resource.Quantity{}
	clusterUsedCPU := &resource.Quantity{}
	clusterUsedMemory := &resource.Quantity{}
	clusterUsedMemoryRequests := &resource.Quantity{}
	clusterUsedMemoryLimits := &resource.Quantity{}
	var clusterUsedPods int64

	for _, info := range nodeInfo {
		if info.PrintOutput == true {
			clusterUsedCPURequests.Add(info.UsedCPURequests)
			clusterUsedCPU.Add(info.UsedCPU)
			clusterUsedMemoryRequests.Add(info.UsedMemoryRequests)
			clusterUsedMemory.Add(info.UsedMemory)
			clusterUsedPods = clusterUsedPods + info.UsedPods
			clusterUsedMemoryLimits.Add(info.UsedMemoryLimits)
		}
	}
	daemonLog := DaemonLog{}
	daemonLog.EventKind = "metric"
	daemonLog.EventModule = "k8s_quota"
	daemonLog.EventProvider = "k8sCapcity"
	daemonLog.EventType = "info"
	daemonLog.EventVersion = "02/24/2020-01"
	daemonLog.NodeLabel = nodeLabel
	daemonLog.ResourceQuotaCPURequestCores = rqclusterAllocatedRequestsCPU.Value()
	daemonLog.ResourceQuotaCPURequestMilliCores = rqclusterAllocatedRequestsCPU.ScaledValue(resource.Milli)
	daemonLog.ResourceQuotaMemoryRequest = rqclusterAllocatedRequestsMemory.Value()
	daemonLog.ResourceQuotaMemoryLimit = rqclusterAllocatedLimitsMemory.Value()
	daemonLog.ContainerResourceCPURequestCores = clusterUsedCPURequests.Value()
	daemonLog.ContainerResourceCPURequestMilliCores = clusterUsedCPURequests.ScaledValue(resource.Milli)
	daemonLog.ContainerResourceMemoryRequest = clusterUsedMemoryRequests.Value()
	daemonLog.ContainerResourceMemoryLimit = clusterUsedMemoryLimits.Value()
	daemonLog.AllocatableMemory = clusterAllocatableMemory.Value()
	daemonLog.AllocatableCPU = clusterAllocatableCPU.Value()
	daemonLog.AllocatablePods = clusterAllocatablePods.Value()
	daemonLog.ContainerResourcePods = clusterUsedPods
	result, _ := json.Marshal(daemonLog)
	fmt.Println(string(result))
}
