package main

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

type ClusterInfo struct {
	NodeInfo                         map[string]NodeInfo
	ClusterAllocatableMemory         resource.Quantity
	ClusterAllocatableCPU            resource.Quantity
	ClusterAllocatablePods           resource.Quantity
	RqclusterAllocatedLimitsMemory   resource.Quantity
	RqclusterAllocatedLimitsCPU      resource.Quantity
	RqclusterAllocatedPods           resource.Quantity
	RqclusterAllocatedRequestsMemory resource.Quantity
	RqclusterAllocatedRequestsCPU    resource.Quantity
	NminusCPU                        resource.Quantity
	NminusMemory                     resource.Quantity
	NminusPods                       resource.Quantity
	NodeLabel                        string
}

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
	EventKind                                    string             `json:"event.kind"`
	EventModule                                  string             `json:"event.module"`
	EventProvider                                string             `json:"event.provider"`
	EventType                                    string             `json:"event.type"`
	EventVersion                                 string             `json:"event.version"`
	ResourceQuotaCPURequestCores                 int64              `json:"k8s_quota.resource_quota.cpu_request.cores"`
	ResourceQuotaCPURequestMilliCores            int64              `json:"k8s_quota.resource_quota.cpu_request.millicores"`
	ResourceQuotaMemoryRequest                   int64              `json:"k8s_quota.resource_quota.memory_request"`
	ResourceQuotaMemoryLimit                     int64              `json:"k8s_quota.resource_quota.memory_limit"`
	ResourceQuotaPods                            int64              `json:"k8s_quota.resource_quota.pods"`
	OversubscriptionFactorMemoryRequest          float64            `json:"k8s_quota.oversubscription_factor.memory.request"`
	OversubscriptionFactorMemoryRequestNminusone float64            `json:"k8s_quota.oversubscription_factor.memory.request.nminusone"`
	OversubscriptionFactorCPURequest             float64            `json:"k8s_quota.oversubscription_factor.cpu.request"`
	OversubscriptionFactorCPURequestNminusone    float64            `json:"k8s_quota.oversubscription_factor.cpu.request.nminusone"`
	OversubscriptionFactorPods                   float64            `json:"k8s_quota.oversubscription_factor.pods"`
	OversubscriptionFactorPodsNminusone          float64            `json:"k8s_quota.oversubscription_factor.pods.nminusone"`
	AllocatableMemory                            int64              `json:"k8s_quota.alloctable.memory"`
	AllocatableMemoryNminusone                   int64              `json:"k8s_quota.alloctable.memory.nminusone"`
	AllocatableCPU                               int64              `json:"k8s_quota.alloctable.cpu"`
	AllocatableCPUNminusone                      int64              `json:"k8s_quota.alloctable.cpu.nminusone"`
	AllocatablePods                              int64              `json:"k8s_quota.alloctable.pods"`
	AllocatablePodsNminusone                     int64              `json:"k8s_quota.alloctable.pods.nminusone"`
	ContainerResourceCPURequestCores             int64              `json:"k8s_quota.container_resource.cpu_request.cores"`
	ContainerResourceCPURequestMilliCores        int64              `json:"k8s_quota.container_resource.cpu_request.millicores"`
	ContainerResourceMemoryRequest               int64              `json:"k8s_quota.container_resource.memory_request"`
	ContainerResourceMemoryLimit                 int64              `json:"k8s_quota.container_resource.memory_limit"`
	ContainerResourcePods                        int64              `json:"k8s_quota.container_resource.pods"`
	NodeLabel                                    string             `json:"k8s_quota.node_label"`
	UtilizationFactorPods                        map[string]float64 `json:"k8s_quota.utilization_factor.pods"`
	UtilizationFactorMemoryRequests              map[string]float64 `json:"k8s_quota.utilization_factor.memory_request"`
	UtilizationFactorCPURequests                 map[string]float64 `json:"k8s_quota.utilization_factor.cpu_request"`
}
