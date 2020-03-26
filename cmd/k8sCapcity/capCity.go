package main

import (
	"encoding/json"
	"fmt"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

func getCapcity(clusterInfo ClusterInfo) {
	capCity := Capcity{}
	capCity.UtilizationFactorPods = make(map[string]float64)
	capCity.UtilizationFactorMemoryRequests = make(map[string]float64)
	capCity.UtilizationFactorCPURequests = make(map[string]float64)

	for name, node := range clusterInfo.NodeInfo {
		if node.PrintOutput {
			clusterInfo.ClusterUsedCPURequests.Add(node.UsedCPURequests)
			clusterInfo.ClusterUsedCPU.Add(node.UsedCPU)
			clusterInfo.ClusterUsedMemoryRequests.Add(node.UsedMemoryRequests)
			clusterInfo.ClusterUsedMemory.Add(node.UsedMemory)
			clusterInfo.ClusterUsedPods = clusterInfo.ClusterUsedPods + node.UsedPods
			clusterInfo.ClusterUsedMemoryLimits.Add(node.UsedMemoryLimits)
			capCity.UtilizationFactorPods[name] = float64(node.UsedPods) / float64(node.AllocatablePods.Value())
			capCity.UtilizationFactorMemoryRequests[name] = float64(node.UsedMemoryRequests.Value()) / float64(node.AllocatableMemory.Value())
			capCity.UtilizationFactorCPURequests[name] = float64(node.UsedCPURequests.Value()) / float64(node.AllocatableCPU.Value())

		}
	}

	capCity.EventKind = "metric"
	capCity.EventModule = "k8s_quota"
	capCity.EventProvider = "k8sCapcity"
	capCity.EventType = "info"
	capCity.EventVersion = "03/06/2020-01"
	capCity.NodeLabel = clusterInfo.NodeLabel
	capCity.ResourceQuotaCPURequestCores = clusterInfo.RqclusterAllocatedRequestsCPU.Value()
	capCity.ResourceQuotaCPURequestMilliCores = clusterInfo.RqclusterAllocatedRequestsCPU.ScaledValue(resource.Milli)
	capCity.ResourceQuotaMemoryRequest = clusterInfo.RqclusterAllocatedRequestsMemory.Value()
	capCity.ResourceQuotaMemoryLimit = clusterInfo.RqclusterAllocatedLimitsMemory.Value()
	capCity.ResourceQuotaPods = clusterInfo.RqclusterAllocatedPods.Value()
	capCity.ContainerResourceCPURequestCores = clusterInfo.ClusterUsedCPURequests.Value()
	capCity.ContainerResourceCPURequestMilliCores = clusterInfo.ClusterUsedCPURequests.ScaledValue(resource.Milli)
	capCity.ContainerResourceMemoryRequest = clusterInfo.ClusterUsedMemoryRequests.Value()
	capCity.ContainerResourceMemoryLimit = clusterInfo.ClusterUsedMemoryLimits.Value()
	capCity.ContainerResourcePods = clusterInfo.ClusterUsedPods
	capCity.AllocatableMemoryTotal = clusterInfo.ClusterAllocatableMemory.Value()
	capCity.AllocatableMemoryNminusone = clusterInfo.ClusterAllocatableMemory.Value() - clusterInfo.NminusMemory.Value()
	capCity.AllocatableCPUTotal = clusterInfo.ClusterAllocatableCPU.Value()
	capCity.AllocatableCPUNminusone = clusterInfo.ClusterAllocatableCPU.Value() - clusterInfo.NminusCPU.Value()
	capCity.AllocatablePodsTotal = clusterInfo.ClusterAllocatablePods.Value()
	capCity.AllocatablePodsNminusone = clusterInfo.ClusterAllocatablePods.Value() - clusterInfo.NminusPods.Value()
	if float64(capCity.AllocatableMemoryTotal) == 0 {
		capCity.SubscriptionFactorMemoryRequestTotal = 0
	} else {
		capCity.SubscriptionFactorMemoryRequestTotal = float64(capCity.ResourceQuotaMemoryRequest) / float64(capCity.AllocatableMemoryTotal)
	}
	if capCity.AllocatableMemoryNminusone == 0 {
		capCity.SubscriptionFactorMemoryRequestNminusone = 0
	} else {
		capCity.SubscriptionFactorMemoryRequestNminusone = float64(capCity.ResourceQuotaMemoryRequest) / float64(capCity.AllocatableMemoryNminusone)
	}
	if float64(clusterInfo.ClusterAllocatableCPU.ScaledValue(resource.Milli)) == 0 {
		capCity.SubscriptionFactorCPURequestTotal = 0
	} else {
		capCity.SubscriptionFactorCPURequestTotal = float64(capCity.ResourceQuotaCPURequestMilliCores) / float64(clusterInfo.ClusterAllocatableCPU.ScaledValue(resource.Milli))
	}
	if float64(clusterInfo.ClusterAllocatableCPU.ScaledValue(resource.Milli)-clusterInfo.NminusCPU.ScaledValue(resource.Milli)) == 0 {
		capCity.SubscriptionFactorCPURequestNminusone = 0
	} else {
		capCity.SubscriptionFactorCPURequestNminusone = float64(capCity.ResourceQuotaCPURequestMilliCores) / float64(clusterInfo.ClusterAllocatableCPU.ScaledValue(resource.Milli)-clusterInfo.NminusCPU.ScaledValue(resource.Milli))
	}
	if float64(capCity.AllocatablePodsTotal) == 0 {
		capCity.SubscriptionFactorPodsTotal = 0
	} else {
		capCity.SubscriptionFactorPodsTotal = float64(capCity.ResourceQuotaPods) / float64(capCity.AllocatablePodsTotal)
	}
	if float64(capCity.AllocatablePodsNminusone) == 0 {
		capCity.SubscriptionFactorPodsNminusone = 0
	} else {
		capCity.SubscriptionFactorPodsNminusone = float64(capCity.ResourceQuotaPods) / float64(capCity.AllocatablePodsNminusone)
	}
	if float64(capCity.AllocatablePodsTotal) == 0 {
		capCity.UtilizationFactorPodsTotal = 0
	} else {
		capCity.UtilizationFactorPodsTotal = float64(clusterInfo.ClusterUsedPods) / float64(capCity.AllocatablePodsTotal)
	}
	if float64(capCity.AllocatablePodsNminusone) == 0 {
		capCity.UtilizationFactorPodsNminusone = 0
	} else {
		capCity.UtilizationFactorPodsNminusone = float64(clusterInfo.ClusterUsedPods) / float64(capCity.AllocatablePodsNminusone)
	}
	if float64(capCity.AllocatableMemoryTotal) == 0 {
		capCity.UtilizationFactorMemoryRequestsTotal = 0
	} else {
		capCity.UtilizationFactorMemoryRequestsTotal = float64(capCity.ContainerResourceMemoryRequest) / float64(capCity.AllocatableMemoryTotal)
	}
	if float64(capCity.AllocatableMemoryNminusone) == 0 {
		capCity.UtilizationFactorMemoryRequestsNminusone = 0
	} else {
		capCity.UtilizationFactorMemoryRequestsNminusone = float64(capCity.ContainerResourceMemoryRequest) / float64(capCity.AllocatableMemoryNminusone)
	}
	if float64(capCity.AllocatableCPUTotal) == 0 {
		capCity.UtilizationFactorCPURequestsTotal = 0
	} else {
		capCity.UtilizationFactorCPURequestsTotal = float64(clusterInfo.ClusterUsedCPURequests.Value()) / float64(capCity.AllocatableCPUTotal)
	}
	if float64(capCity.AllocatableCPUNminusone) == 0 {
		capCity.UtilizationFactorCPURequestsNminusone = 0
	} else {
		capCity.UtilizationFactorCPURequestsNminusone = float64(clusterInfo.ClusterUsedCPURequests.Value()) / float64(capCity.AllocatableCPUNminusone)
	}
	capCity.AvailableMemoryRequestTotal = capCity.AllocatableMemoryTotal - capCity.ContainerResourceMemoryRequest
	capCity.AvailableMemoryRequestNminusone = capCity.AllocatableMemoryNminusone - capCity.ContainerResourceMemoryRequest
	capCity.AvailableCPURequestTotal = capCity.AllocatableCPUTotal - capCity.ContainerResourceCPURequestCores
	capCity.AvailableCPURequestNminusone = capCity.AllocatableCPUNminusone - capCity.ContainerResourceCPURequestCores
	capCity.AvailablePodsTotal = capCity.AllocatablePodsTotal - capCity.ContainerResourcePods
	capCity.AvailablePodsNminusone = capCity.AllocatablePodsNminusone - capCity.ContainerResourcePods
	result, err := json.Marshal(capCity)
	if err != nil {
		fmt.Printf("There was an error during json.Marshal, Error: %s\n", err)
		panic(err)
	}
	fmt.Println(string(result))
}
