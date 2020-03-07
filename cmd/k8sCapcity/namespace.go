package main

import (
	"encoding/json"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"os"
)

func gatherPodSpecInfo(pod corev1.Pod, nsInfo NamespaceInfo) NamespaceInfo {
	for _, container := range pod.Spec.Containers {
		uniqueContainerName := fmt.Sprintf("%s-%s", pod.Name, container.Name)
		containerStats := nsInfo.NamespacePods[pod.Name].Containers[uniqueContainerName]
		containerStats.MemoryRequests = *container.Resources.Requests.Memory()
		containerStats.MemoryLimits = *container.Resources.Limits.Memory()
		containerStats.CPURequestsMilliCores = container.Resources.Requests.Cpu().ScaledValue(resource.Milli)
		containerStats.CPURequestsCores = float64(container.Resources.Requests.Cpu().ScaledValue(resource.Milli)) / 1000
		containerStats.CPULimitsMilliCores = container.Resources.Limits.Cpu().ScaledValue(resource.Milli)
		containerStats.CPULimitsCores = float64(container.Resources.Limits.Cpu().ScaledValue(resource.Milli)) / 1000
		containerStats.Name = container.Name
		containerStats.Pod = pod.Name
		nsInfo.NamespacePods[pod.Name].Containers[uniqueContainerName] = containerStats
		// Add up for the namespace
		nsInfo.NamespaceMemoryLimits.Add(containerStats.MemoryLimits)
		nsInfo.NamespaceMemoryRequests.Add(containerStats.MemoryRequests)
		nsInfo.NamespaceCPULimitsMilliCores = nsInfo.NamespaceCPULimitsMilliCores + containerStats.CPULimitsMilliCores
		nsInfo.NamespaceCPURequestsMilliCores = nsInfo.NamespaceCPURequestsMilliCores + containerStats.CPURequestsMilliCores
	}
	return nsInfo
}

func gatherNamespaceInfo(clientset *kubernetes.Clientset, nameSpace *string) NamespaceInfo {

	nsInfo := NamespaceInfo{}
	podMetricList := getPodMetrics(clientset)
	podList := getPodList(clientset, nameSpace)
	nsInfo.NamespacePods = make(map[string]*Pod)
	for _, metricPod := range podMetricList.Items {
		if *nameSpace == metricPod.Namespace {
			containerArray := make(map[string]ContainerInfo)
			for _, container := range metricPod.Containers {
				uniqueContainerName := fmt.Sprintf("%s-%s", metricPod.Name, container.Name)
				c := &ContainerInfo{
					MemoryUsed:        *container.Usage.Memory(),
					CPUUsedMilliCores: container.Usage.Cpu().MilliValue(),
					CPUUsedCores:      float64(container.Usage.Cpu().MilliValue()) / 1000,
				}
				containerArray[uniqueContainerName] = *c
				// Add up for the namespace
				nsInfo.NamespaceMemoryUsed.Add(*container.Usage.Memory())
				nsInfo.NamespaceCPUUsedMilliCores = nsInfo.NamespaceCPUUsedMilliCores + container.Usage.Cpu().MilliValue()
			}
			nsInfo.NamespacePods[metricPod.Name] = &Pod{Containers: containerArray}
			for _, pod := range podList.Items {
				if pod.Name == metricPod.Name {
					if pod.Status.Phase != "Failed" {
						if pod.Status.Phase != "Succeeded" {
							nsInfo = gatherPodSpecInfo(pod, nsInfo)
						}
					}
				}
			}
		}
	}
	nsInfo.NamespaceCPUUsedCores = float64(nsInfo.NamespaceCPUUsedMilliCores) / 1000
	nsInfo.NamespaceCPULimitsCores = float64(nsInfo.NamespaceCPULimitsMilliCores) / 1000
	nsInfo.NamespaceCPURequestsCores = float64(nsInfo.NamespaceCPURequestsMilliCores) / 1000
	nsInfo.Name = *nameSpace
	return nsInfo
}

func namespaceJSONMode(nsInfo NamespaceInfo) {
	result, err := json.Marshal(nsInfo)
	check(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func namespaceHumanMode(nsInfo NamespaceInfo) {
	fmt.Println("")
	fmt.Println("================")
	for podName, pods := range nsInfo.NamespacePods {
		fmt.Printf("****Pod Name: %s****\n", podName)
		for _, container := range pods.Containers {
			fmt.Println("================")
			fmt.Printf("Container Name: %s\n", container.Name)
			fmt.Println("----------------")
			fmt.Printf("CPURequests: %v\n", container.CPURequestsCores)
			fmt.Printf("MemoryRequests: %vMiB\n", toMib(container.MemoryRequests))
			fmt.Printf("CPULimits: %v\n", container.CPULimitsCores)
			fmt.Printf("MemoryLimits: %vMiB\n", toMib(container.MemoryLimits))
			fmt.Println("----------------")
			fmt.Printf("CPU Used: %dm\n", container.CPUUsedMilliCores)
			fmt.Printf("Memory Used: %vMiB\n", toMib(container.MemoryUsed))
		}
	}
	fmt.Printf("<><><><><>Sum Total for Namespace: %s<><><><><>\n", nsInfo.Name)
	fmt.Println("----------------")
	fmt.Printf("Namespace Total CPURequests: %v\n", nsInfo.NamespaceCPURequestsCores)
	fmt.Printf("Namespace Total MemoryRequests: %vMiB (%.1fGiB)\n", toMib(nsInfo.NamespaceMemoryRequests), toGibFromMib(toMib(nsInfo.NamespaceMemoryRequests)))
	fmt.Printf("Namespace Total CPULimits: %v\n", nsInfo.NamespaceCPULimitsCores)
	fmt.Printf("Namespace Total MemoryLimits: %vMiB (%.1fGiB)\n", toMib(nsInfo.NamespaceMemoryLimits), toGibFromMib(toMib(nsInfo.NamespaceMemoryLimits)))
	fmt.Println("----------------")
	fmt.Printf("Namespace Total CPU Used: %v\n", nsInfo.NamespaceCPUUsedCores)
	fmt.Printf("Namespace Total Memory Used: %vMiB (%.1fGiB)\n", toMib(nsInfo.NamespaceMemoryUsed), toGibFromMib(toMib(nsInfo.NamespaceMemoryUsed)))

	os.Exit(0)
}
