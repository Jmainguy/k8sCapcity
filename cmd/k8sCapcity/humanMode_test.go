package main

import (
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func TestHumanModeEmpty(t *testing.T) {
	clusterInfo := ClusterInfo{}
	humanMode(clusterInfo)
}

func TestHumanModeValidClusterInfo(t *testing.T) {
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
	}
	humanMode(clusterInfo)
}

func TestToGib(t *testing.T) {
	memory := resource.MustParse("1024Mi")
	memoryGib := toGib(memory)
	if memoryGib != 1 {
		t.Errorf("Expected 1, got %d", memoryGib)
	}
}

func TestToMib(t *testing.T) {
	memory := resource.MustParse("1Gi")
	memoryMib := toMib(memory)
	if memoryMib != 1024 {
		t.Errorf("Expected 1024, got %d", memoryMib)
	}
}
