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
	UsedMemoryLimits   resource.Quantity
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

type DaemonLog struct {
	EventKind                         string `json:"event.kind"`
	EventModule                       string `json:"event.module"`
	EventProvider                     string `json:"event.provider"`
	EventType                         string `json:"event.type"`
	EventVersion                      string `json:"event.version"`
	ResourceQuotaCPURequestCores      int64  `json:"k8s_quota.resource_quota.cpu_request.cores"`
	ResourceQuotaCPURequestMilliCores int64  `json:"k8s_quota.resource_quota.cpu_request.millicores"`
	ResourceQuotaMemoryRequest        int64  `json:"k8s_quota.resource_quota.memory_request"`
	ResourceQuotaMemoryLimit          int64  `json:"k8s_quota.resource_quota.memory_limit"`
	ResourceQuotaPods                 int64  `json:"k8s_quota.resource_quota.pods"`
	//OversubscriptionFactor                float64 `json:"k8s_quota.oversubscription_factor"`
	AllocatableMemory                     int64  `json:"k8s_quota.alloctable.memory"`
	AllocatableCPU                        int64  `json:"k8s_quota.alloctable.cpu"`
	AllocatablePods                       int64  `json:"k8s_quota.alloctable.pods"`
	ContainerResourceCPURequestCores      int64  `json:"k8s_quota.container_resource.cpu_request.cores"`
	ContainerResourceCPURequestMilliCores int64  `json:"k8s_quota.container_resource.cpu_request.millicores"`
	ContainerResourceMemoryRequest        int64  `json:"k8s_quota.container_resource.memory_request"`
	ContainerResourceMemoryLimit          int64  `json:"k8s_quota.container_resource.memory_limit"`
	ContainerResourcePods                 int64  `json:"k8s_quota.container_resource.pods"`
	NodeLabel                             string `json:"k8s_quota.node_label"`
}
