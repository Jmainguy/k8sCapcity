package main

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ClusterInfo : Information about the cluster
type ClusterInfo struct {
	NodeInfo                         map[string]NodeInfo
	ClusterAllocatableMemory         resource.Quantity
	ClusterAllocatableCPU            resource.Quantity
	ClusterAllocatablePods           resource.Quantity
	ClusterUsedCPURequests           resource.Quantity
	ClusterUsedCPU                   resource.Quantity
	ClusterUsedMemory                resource.Quantity
	ClusterUsedMemoryRequests        resource.Quantity
	ClusterUsedMemoryLimits          resource.Quantity
	ClusterUsedPods                  int64
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

// NodeInfo : Information about the node
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

// ContainerInfo : Information about the container
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

// DaemonLog : Json to print out about metrics we gathered
type DaemonLog struct {
	EventKind                                     string             `json:"event.kind"`
	EventModule                                   string             `json:"event.module"`
	EventProvider                                 string             `json:"event.provider"`
	EventType                                     string             `json:"event.type"`
	EventVersion                                  string             `json:"event.version"`
	ResourceQuotaCPURequestCores                  int64              `json:"k8s_quota.resource_quota.cpu_request.cores"`
	ResourceQuotaCPURequestMilliCores             int64              `json:"k8s_quota.resource_quota.cpu_request.millicores"`
	ResourceQuotaMemoryRequest                    int64              `json:"k8s_quota.resource_quota.memory_request"`
	ResourceQuotaMemoryLimit                      int64              `json:"k8s_quota.resource_quota.memory_limit"`
	ResourceQuotaPods                             int64              `json:"k8s_quota.resource_quota.pods"`
	SubscriptionFactorMemoryRequest               float64            `json:"k8s_quota.subscription_factor.memory.request"`
	SubscriptionFactorMemoryRequestNminusone      float64            `json:"k8s_quota.subscription_factor.memory.request.nminusone"`
	SubscriptionFactorCPURequest                  float64            `json:"k8s_quota.subscription_factor.cpu.request"`
	SubscriptionFactorCPURequestNminusone         float64            `json:"k8s_quota.subscription_factor.cpu.request.nminusone"`
	SubscriptionFactorPods                        float64            `json:"k8s_quota.subscription_factor.pods"`
	SubscriptionFactorPodsNminusone               float64            `json:"k8s_quota.subscription_factor.pods.nminusone"`
	AllocatableMemory                             int64              `json:"k8s_quota.alloctable.memory"`
	AllocatableMemoryNminusone                    int64              `json:"k8s_quota.alloctable.memory.nminusone"`
	AllocatableCPU                                int64              `json:"k8s_quota.alloctable.cpu"`
	AllocatableCPUNminusone                       int64              `json:"k8s_quota.alloctable.cpu.nminusone"`
	AllocatablePods                               int64              `json:"k8s_quota.alloctable.pods"`
	AllocatablePodsNminusone                      int64              `json:"k8s_quota.alloctable.pods.nminusone"`
	ContainerResourceCPURequestCores              int64              `json:"k8s_quota.container_resource.cpu_request.cores"`
	ContainerResourceCPURequestMilliCores         int64              `json:"k8s_quota.container_resource.cpu_request.millicores"`
	ContainerResourceMemoryRequest                int64              `json:"k8s_quota.container_resource.memory_request"`
	ContainerResourceMemoryLimit                  int64              `json:"k8s_quota.container_resource.memory_limit"`
	ContainerResourcePods                         int64              `json:"k8s_quota.container_resource.pods"`
	NodeLabel                                     string             `json:"k8s_quota.node_label"`
	UtilizationFactorPods                         map[string]float64 `json:"k8s_quota.utilization_factor.pods"`
	UtilizationFactorPodsTotal                    float64            `json:"k8s_quota.utilization_factor.pods.total"`
	UtilizationFactorPodsTotalNminusone           float64            `json:"k8s_quota.utilization_factor.pods.total.nminusone"`
	UtilizationFactorMemoryRequests               map[string]float64 `json:"k8s_quota.utilization_factor.memory_request"`
	UtilizationFactorMemoryRequestsTotal          float64            `json:"k8s_quota.utilization_factor.memory_request.total"`
	UtilizationFactorMemoryRequestsTotalNminusone float64            `json:"k8s_quota.utilization_factor.memory_request.total.nminusone"`
	UtilizationFactorCPURequests                  map[string]float64 `json:"k8s_quota.utilization_factor.cpu_request"`
	UtilizationFactorCPURequestsTotal             float64            `json:"k8s_quota.utilization_factor.cpu_request.total"`
	UtilizationFactorCPURequestsTotalNminusone    float64            `json:"k8s_quota.utilization_factor.cpu_request.total.nminusone"`
	AvailableMemoryRequest                        int64              `json:"k8s_quota.available.memory_request"`
	AvailableMemoryRequestNminusone               int64              `json:"k8s_quota.available.memory_request.nminusone"`
	AvailableCPURequest                           int64              `json:"k8s_quota.available.cpu_request"`
	AvailableCPURequestNminusone                  int64              `json:"k8s_quota.available.cpu_request.nminusone"`
	AvailablePods                                 int64              `json:"k8s_quota.available.pods"`
	AvailablePodsNminusone                        int64              `json:"k8s_quota.available.pods.nminusone"`
}
