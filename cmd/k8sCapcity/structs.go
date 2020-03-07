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
	Name                  string  `json:"name"`
	Pod                   string  `json:"pod"`
	CPURequestsMilliCores int64   `json:"cpu_requests.millicores"`
	CPULimitsMilliCores   int64   `json:"cpu_limits.millicores"`
	MemoryRequests        int64   `json:"memory_requests.bytes"`
	MemoryLimits          int64   `json:"memory_limits.bytes"`
	CPUUsedMilliCores     int64   `json:"cpu_used.millicores"`
	MemoryUsed            int64   `json:"memory_used.bytes"`
	MemoryRequestsMiB     int64   `json:"memory_requests.mebibytes"`
	MemoryLimitsMiB       int64   `json:"memory_limits.mebibytes"`
	MemoryUsedMiB         int64   `json:"memory_used.mebibytes"`
	CPURequestsCores      float64 `json:"cpu_requests.cores"`
	CPULimitsCores        float64 `json:"cpu_limits.cores"`
	CPUUsedCores          float64 `json:"cpu_used.cores"`
}

// DaemonLog : Json to print out about metrics we gathered
type DaemonLog struct {
	EventKind                                string             `json:"event.kind"`
	EventModule                              string             `json:"event.module"`
	EventProvider                            string             `json:"event.provider"`
	EventType                                string             `json:"event.type"`
	EventVersion                             string             `json:"event.version"`
	ResourceQuotaCPURequestCores             int64              `json:"k8s_quota.resource_quota.cpu_request.cores"`
	ResourceQuotaCPURequestMilliCores        int64              `json:"k8s_quota.resource_quota.cpu_request.millicores"`
	ResourceQuotaMemoryRequest               int64              `json:"k8s_quota.resource_quota.memory_request"`
	ResourceQuotaMemoryLimit                 int64              `json:"k8s_quota.resource_quota.memory_limit"`
	ResourceQuotaPods                        int64              `json:"k8s_quota.resource_quota.pods"`
	SubscriptionFactorMemoryRequestTotal     float64            `json:"k8s_quota.subscription_factor.memory.request.total"`
	SubscriptionFactorMemoryRequestNminusone float64            `json:"k8s_quota.subscription_factor.memory.request.nminusone"`
	SubscriptionFactorCPURequestTotal        float64            `json:"k8s_quota.subscription_factor.cpu.request.total"`
	SubscriptionFactorCPURequestNminusone    float64            `json:"k8s_quota.subscription_factor.cpu.request.nminusone"`
	SubscriptionFactorPodsTotal              float64            `json:"k8s_quota.subscription_factor.pods.total"`
	SubscriptionFactorPodsNminusone          float64            `json:"k8s_quota.subscription_factor.pods.nminusone"`
	AllocatableMemoryTotal                   int64              `json:"k8s_quota.alloctable.memory.total"`
	AllocatableMemoryNminusone               int64              `json:"k8s_quota.alloctable.memory.nminusone"`
	AllocatableCPUTotal                      int64              `json:"k8s_quota.alloctable.cpu.total"`
	AllocatableCPUNminusone                  int64              `json:"k8s_quota.alloctable.cpu.nminusone"`
	AllocatablePodsTotal                     int64              `json:"k8s_quota.alloctable.pods.total"`
	AllocatablePodsNminusone                 int64              `json:"k8s_quota.alloctable.pods.nminusone"`
	ContainerResourceCPURequestCores         int64              `json:"k8s_quota.container_resource.cpu_request.cores"`
	ContainerResourceCPURequestMilliCores    int64              `json:"k8s_quota.container_resource.cpu_request.millicores"`
	ContainerResourceMemoryRequest           int64              `json:"k8s_quota.container_resource.memory_request"`
	ContainerResourceMemoryLimit             int64              `json:"k8s_quota.container_resource.memory_limit"`
	ContainerResourcePods                    int64              `json:"k8s_quota.container_resource.pods"`
	NodeLabel                                string             `json:"k8s_quota.node_label"`
	UtilizationFactorPods                    map[string]float64 `json:"k8s_quota.utilization_factor.pods"`
	UtilizationFactorPodsTotal               float64            `json:"k8s_quota.utilization_factor.pods.total"`
	UtilizationFactorPodsNminusone           float64            `json:"k8s_quota.utilization_factor.pods.nminusone"`
	UtilizationFactorMemoryRequests          map[string]float64 `json:"k8s_quota.utilization_factor.memory_request"`
	UtilizationFactorMemoryRequestsTotal     float64            `json:"k8s_quota.utilization_factor.memory_request.total"`
	UtilizationFactorMemoryRequestsNminusone float64            `json:"k8s_quota.utilization_factor.memory_request.nminusone"`
	UtilizationFactorCPURequests             map[string]float64 `json:"k8s_quota.utilization_factor.cpu_request"`
	UtilizationFactorCPURequestsTotal        float64            `json:"k8s_quota.utilization_factor.cpu_request.total"`
	UtilizationFactorCPURequestsNminusone    float64            `json:"k8s_quota.utilization_factor.cpu_request.nminusone"`
	AvailableMemoryRequestTotal              int64              `json:"k8s_quota.available.memory_request.total"`
	AvailableMemoryRequestNminusone          int64              `json:"k8s_quota.available.memory_request.nminusone"`
	AvailableCPURequestTotal                 int64              `json:"k8s_quota.available.cpu_request.total"`
	AvailableCPURequestNminusone             int64              `json:"k8s_quota.available.cpu_request.nminusone"`
	AvailablePodsTotal                       int64              `json:"k8s_quota.available.pods.total"`
	AvailablePodsNminusone                   int64              `json:"k8s_quota.available.pods.nminusone"`
}

// NamespaceInfo : Information about the namespace
type NamespaceInfo struct {
	Name                           string          `json:"k8s_quota.namespace.name"`
	NamespacePods                  map[string]*Pod `json:"k8s_quota.namespace.pods"`
	NamespaceMemoryLimits          int64           `json:"k8s_quota.namespace.memory_limits.bytes"`
	NamespaceMemoryRequests        int64           `json:"k8s_quota.namespace.memory_requests.bytes"`
	NamespaceMemoryUsed            int64           `json:"k8s_quota.namespace.memory_used.bytes"`
	NamespaceCPULimitsMilliCores   int64           `json:"k8s_quota.namespace.cpu_limits.millicores"`
	NamespaceCPURequestsMilliCores int64           `json:"k8s_quota.namespace.cpu_requests.millicores"`
	NamespaceCPUUsedMilliCores     int64           `json:"k8s_quota.namespace.cpu_used.millicores"`
	NamespaceMemoryLimitsGiB       float64         `json:"k8s_quota.namespace.memory_limits.gibibytes"`
	NamespaceMemoryRequestsGiB     float64         `json:"k8s_quota.namespace.memory_requests.gibibytes"`
	NamespaceMemoryUsedGiB         float64         `json:"k8s_quota.namespace.memory_used.gibibytes"`
	NamespaceCPULimitsCores        float64         `json:"k8s_quota.namespace.cpu_limits.cores"`
	NamespaceCPURequestsCores      float64         `json:"k8s_quota.namespace.cpu_requests.cores"`
	NamespaceCPUUsedCores          float64         `json:"k8s_quota.namespace.cpu_used.cores"`
}

// Pod : A pod full of containers
type Pod struct {
	Containers map[string]ContainerInfo
}
