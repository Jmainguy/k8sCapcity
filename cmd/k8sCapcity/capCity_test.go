package main

import (
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func TestDaemonModeEmpty(t *testing.T) {
	clusterInfo := ClusterInfo{}
	getCapcity(clusterInfo)
}

func TestDaemonModeValidClusterInfo(t *testing.T) {
	clusterInfo := ClusterInfo{
		ClusterAllocatableMemory:         resource.MustParse("1Gi"),
		ClusterAllocatableCPU:            resource.MustParse("1"),
		ClusterAllocatablePods:           resource.MustParse("1000"),
		RqclusterAllocatedLimitsMemory:   resource.MustParse("1Gi"),
		RqclusterAllocatedLimitsCPU:      resource.MustParse("1"),
		RqclusterAllocatedPods:           resource.MustParse("100"),
		RqclusterAllocatedRequestsMemory: resource.MustParse("1Gi"),
		RqclusterAllocatedRequestsCPU:    resource.MustParse("1"),
		NminusCPU:                        resource.MustParse("1"),
		NminusMemory:                     resource.MustParse("1Gi"),
		NminusPods:                       resource.MustParse("10"),
		NodeLabel:                        "test=true",
		NodeInfo: map[string]NodeInfo{
			"test-node": {
				AllocatableCPU:     resource.MustParse("16"),
				AllocatableMemory:  resource.MustParse("256Gi"),
				AllocatablePods:    resource.MustParse("1000"),
				UsedPods:           10,
				UsedCPU:            resource.MustParse("1"),
				UsedMemory:         resource.MustParse("1Gi"),
				UsedMemoryRequests: resource.MustParse("10Gi"),
				UsedMemoryLimits:   resource.MustParse("10Gi"),
				UsedCPURequests:    resource.MustParse("1"),
				PrintOutput:        true,
			},
		},
	}
	getCapcity(clusterInfo)
}
