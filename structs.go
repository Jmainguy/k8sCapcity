package main

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

type NodeInfo struct {
	UsedPods           int64
	AllocatableCPU     resource.Quantity
	AllocatableMemory  resource.Quantity
	AllocatablePods    resource.Quantity
	UsedCPU            resource.Quantity
	UsedMemory         resource.Quantity
	UsedMemoryRequests resource.Quantity
	UsedCPURequests    resource.Quantity
	PrintOutput        bool
}

type ContainerInfo struct {
	Name           string
	Pod            string
	CPURequests    resource.Quantity
	CPULimits      resource.Quantity
	MemoryRequests resource.Quantity
	MemoryLimits   resource.Quantity
	UsedCPU        resource.Quantity
	UsedMemory     resource.Quantity
}
