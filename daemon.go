package main

import (
	"encoding/json"
	"fmt"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

func runDaemonMode(clusterInfo ClusterInfo) {
	clusterUsedCPURequests := &resource.Quantity{}
	clusterUsedCPU := &resource.Quantity{}
	clusterUsedMemory := &resource.Quantity{}
	clusterUsedMemoryRequests := &resource.Quantity{}
	clusterUsedMemoryLimits := &resource.Quantity{}
	var clusterUsedPods int64

	for _, info := range clusterInfo.NodeInfo {
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
	daemonLog.NodeLabel = clusterInfo.NodeLabel
	daemonLog.ResourceQuotaCPURequestCores = clusterInfo.RqclusterAllocatedRequestsCPU.Value()
	daemonLog.ResourceQuotaCPURequestMilliCores = clusterInfo.RqclusterAllocatedRequestsCPU.ScaledValue(resource.Milli)
	daemonLog.ResourceQuotaMemoryRequest = clusterInfo.RqclusterAllocatedRequestsMemory.Value()
	daemonLog.ResourceQuotaMemoryLimit = clusterInfo.RqclusterAllocatedLimitsMemory.Value()
	daemonLog.ResourceQuotaPods = clusterInfo.RqclusterAllocatedPods.Value()
	daemonLog.ContainerResourceCPURequestCores = clusterUsedCPURequests.Value()
	daemonLog.ContainerResourceCPURequestMilliCores = clusterUsedCPURequests.ScaledValue(resource.Milli)
	daemonLog.ContainerResourceMemoryRequest = clusterUsedMemoryRequests.Value()
	daemonLog.ContainerResourceMemoryLimit = clusterUsedMemoryLimits.Value()
	daemonLog.ContainerResourcePods = clusterUsedPods
	daemonLog.AllocatableMemory = clusterInfo.ClusterAllocatableMemory.Value()
	daemonLog.AllocatableMemoryNminusone = clusterInfo.ClusterAllocatableMemory.Value() - clusterInfo.NminusMemory.Value()
	daemonLog.AllocatableCPU = clusterInfo.ClusterAllocatableCPU.Value()
	daemonLog.AllocatableCPUNminusone = clusterInfo.ClusterAllocatableCPU.Value() - clusterInfo.NminusCPU.Value()
	daemonLog.AllocatablePods = clusterInfo.ClusterAllocatablePods.Value()
	daemonLog.AllocatablePodsNminusone = clusterInfo.ClusterAllocatablePods.Value() - clusterInfo.NminusPods.Value()
	daemonLog.OversubscriptionFactorMemoryRequest = float64(daemonLog.ResourceQuotaMemoryRequest) / float64(daemonLog.AllocatableMemory)
	daemonLog.OversubscriptionFactorMemoryRequestNminusone = float64(daemonLog.ResourceQuotaMemoryRequest) / float64(daemonLog.AllocatableMemoryNminusone)
	daemonLog.OversubscriptionFactorCPURequest = float64(daemonLog.ResourceQuotaCPURequestMilliCores) / float64(clusterInfo.ClusterAllocatableCPU.ScaledValue(resource.Milli))
	daemonLog.OversubscriptionFactorCPURequestNminusone = float64(daemonLog.ResourceQuotaCPURequestMilliCores) / float64(clusterInfo.ClusterAllocatableCPU.ScaledValue(resource.Milli)-clusterInfo.NminusCPU.ScaledValue(resource.Milli))
	daemonLog.OversubscriptionFactorPods = float64(daemonLog.ResourceQuotaPods) / float64(daemonLog.AllocatablePods)
	daemonLog.OversubscriptionFactorPodsNminusone = float64(daemonLog.ResourceQuotaPods) / float64(daemonLog.AllocatablePodsNminusone)
	result, _ := json.Marshal(daemonLog)
	fmt.Println(string(result))
}
