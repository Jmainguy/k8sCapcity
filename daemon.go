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
	daemonLog := DaemonLog{}
	daemonLog.UtilizationFactorPods = make(map[string]float64)
	daemonLog.UtilizationFactorMemoryRequests = make(map[string]float64)
	daemonLog.UtilizationFactorCPURequests = make(map[string]float64)

	for name, node := range clusterInfo.NodeInfo {
		if node.PrintOutput == true {
			clusterUsedCPURequests.Add(node.UsedCPURequests)
			clusterUsedCPU.Add(node.UsedCPU)
			clusterUsedMemoryRequests.Add(node.UsedMemoryRequests)
			clusterUsedMemory.Add(node.UsedMemory)
			clusterUsedPods = clusterUsedPods + node.UsedPods
			clusterUsedMemoryLimits.Add(node.UsedMemoryLimits)
			daemonLog.UtilizationFactorPods[name] = float64(node.UsedPods) / float64(node.AllocatablePods.Value())
			daemonLog.UtilizationFactorMemoryRequests[name] = float64(node.UsedMemoryRequests.Value()) / float64(node.AllocatableMemory.Value())
			daemonLog.UtilizationFactorCPURequests[name] = float64(node.UsedCPURequests.Value()) / float64(node.AllocatableCPU.Value())

		}
	}

	daemonLog.EventKind = "metric"
	daemonLog.EventModule = "k8s_quota"
	daemonLog.EventProvider = "k8sCapcity"
	daemonLog.EventType = "info"
	daemonLog.EventVersion = "03/02/2020-01"
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
