package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metricsv1b1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func getPodMetrics(clientset *kubernetes.Clientset) (podMetricList *metricsv1b1.PodMetricsList) {
	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.Background())
	check(err)
	err = json.Unmarshal(data, &podMetricList)
	check(err)
	return podMetricList
}

func getPodList(clientset *kubernetes.Clientset, nameSpace *string) (pods *corev1.PodList) {
	pods, err := clientset.CoreV1().Pods(*nameSpace).List(context.Background(), metav1.ListOptions{})
	check(err)
	return pods
}

func gatherPodSpecInfo(pod corev1.Pod, nsInfo NamespaceInfo) NamespaceInfo {
	for _, container := range pod.Spec.Containers {
		uniqueContainerName := fmt.Sprintf("%s-%s", pod.Name, container.Name)
		containerStats := nsInfo.NamespacePods[pod.Name].Containers[uniqueContainerName]
		containerStats.MemoryRequests = container.Resources.Requests.Memory().Value()
		containerStats.MemoryRequestsMiB = toMib(*container.Resources.Requests.Memory())
		containerStats.MemoryLimits = container.Resources.Limits.Memory().Value()
		containerStats.MemoryLimitsMiB = toMib(*container.Resources.Limits.Memory())
		containerStats.CPURequestsMilliCores = container.Resources.Requests.Cpu().ScaledValue(resource.Milli)
		containerStats.CPURequestsCores = float64(container.Resources.Requests.Cpu().ScaledValue(resource.Milli)) / 1000
		containerStats.CPULimitsMilliCores = container.Resources.Limits.Cpu().ScaledValue(resource.Milli)
		containerStats.CPULimitsCores = float64(container.Resources.Limits.Cpu().ScaledValue(resource.Milli)) / 1000
		containerStats.Name = container.Name
		containerStats.Pod = pod.Name
		nsInfo.NamespacePods[pod.Name].Containers[uniqueContainerName] = containerStats
		// Add up for the namespace
		nsInfo.NamespaceMemoryLimits = nsInfo.NamespaceMemoryLimits + containerStats.MemoryLimits
		nsInfo.NamespaceMemoryRequests = nsInfo.NamespaceMemoryRequests + containerStats.MemoryRequests
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
	namespacePods := make(map[string]bool)
	for _, metricPod := range podMetricList.Items {
		if *nameSpace == metricPod.Namespace {
			containerArray := make(map[string]ContainerInfo)
			for _, container := range metricPod.Containers {
				uniqueContainerName := fmt.Sprintf("%s-%s", metricPod.Name, container.Name)
				c := &ContainerInfo{
					MemoryUsed:        container.Usage.Memory().Value(),
					MemoryUsedMiB:     toMib(*container.Usage.Memory()),
					CPUUsedMilliCores: container.Usage.Cpu().MilliValue(),
					CPUUsedCores:      float64(container.Usage.Cpu().MilliValue()) / 1000,
				}
				containerArray[uniqueContainerName] = *c
				// Add up for the namespace
				nsInfo.NamespaceMemoryUsed = nsInfo.NamespaceMemoryUsed + container.Usage.Memory().Value()
				nsInfo.NamespaceCPUUsedMilliCores = nsInfo.NamespaceCPUUsedMilliCores + container.Usage.Cpu().MilliValue()
			}
			nsInfo.NamespacePods[metricPod.Name] = &Pod{Containers: containerArray}
			namespacePods[metricPod.Name] = true
		}
	}
	for _, pod := range podList.Items {
		if namespacePods[pod.Name] {
			if pod.Status.Phase != "Failed" {
				if pod.Status.Phase != "Succeeded" {
					nsInfo = gatherPodSpecInfo(pod, nsInfo)
				}
			}
		}
	}
	nsInfo.NamespaceCPUUsedCores = float64(nsInfo.NamespaceCPUUsedMilliCores) / 1000
	nsInfo.NamespaceCPULimitsCores = float64(nsInfo.NamespaceCPULimitsMilliCores) / 1000
	nsInfo.NamespaceCPURequestsCores = float64(nsInfo.NamespaceCPURequestsMilliCores) / 1000
	nsInfo.NamespaceMemoryLimitsGiB = toGibFromByte(nsInfo.NamespaceMemoryLimits)
	nsInfo.NamespaceMemoryRequestsGiB = toGibFromByte(nsInfo.NamespaceMemoryRequests)
	nsInfo.NamespaceMemoryUsedGiB = toGibFromByte(nsInfo.NamespaceMemoryUsed)
	nsInfo.Name = *nameSpace
	return nsInfo
}

func namespaceHumanMode(nsInfo NamespaceInfo) (output []string) {
	output = append(output, fmt.Sprintf(""))
	output = append(output, fmt.Sprintf("================"))
	for podName, pods := range nsInfo.NamespacePods {
		output = append(output, fmt.Sprintf("****Pod Name: %s****", podName))
		for _, container := range pods.Containers {
			output = append(output, fmt.Sprintf("================"))
			output = append(output, fmt.Sprintf("Container Name: %s", container.Name))
			output = append(output, fmt.Sprintf("----------------"))
			output = append(output, fmt.Sprintf("CPURequests: %v", container.CPURequestsCores))
			output = append(output, fmt.Sprintf("MemoryRequests: %dMiB", toMibFromByte(container.MemoryRequests)))
			output = append(output, fmt.Sprintf("CPULimits: %v", container.CPULimitsCores))
			output = append(output, fmt.Sprintf("MemoryLimits: %dMiB", toMibFromByte(container.MemoryLimits)))
			output = append(output, fmt.Sprintf("----------------"))
			output = append(output, fmt.Sprintf("CPU Used: %dm", container.CPUUsedMilliCores))
			output = append(output, fmt.Sprintf("Memory Used: %dMiB", toMibFromByte(container.MemoryUsed)))
			output = append(output, fmt.Sprintf("================"))
		}
	}
	output = append(output, fmt.Sprintf("<><><><><>Sum Total for Namespace: %s<><><><><>", nsInfo.Name))
	output = append(output, fmt.Sprintf("----------------"))
	output = append(output, fmt.Sprintf("Namespace Total CPURequests: %v", nsInfo.NamespaceCPURequestsCores))
	output = append(output, fmt.Sprintf("Namespace Total MemoryRequests: %vMiB (%.1fGiB)", toMibFromByte(nsInfo.NamespaceMemoryRequests), nsInfo.NamespaceMemoryRequestsGiB))
	output = append(output, fmt.Sprintf("Namespace Total CPULimits: %v", nsInfo.NamespaceCPULimitsCores))
	output = append(output, fmt.Sprintf("Namespace Total MemoryLimits: %vMiB (%.1fGiB)", toMibFromByte(nsInfo.NamespaceMemoryLimits), nsInfo.NamespaceMemoryLimitsGiB))
	output = append(output, fmt.Sprintf("----------------"))
	output = append(output, fmt.Sprintf("Namespace Total CPU Used: %v", nsInfo.NamespaceCPUUsedCores))
	output = append(output, fmt.Sprintf("Namespace Total Memory Used: %dMiB (%.1fGiB)", toMibFromByte(nsInfo.NamespaceMemoryUsed), nsInfo.NamespaceMemoryUsedGiB))

	return output
}

func getNamespaceListFromFile(namespaceList string) (namespaces []string) {
	file, err := os.Open(namespaceList)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		namespaces = append(namespaces, scanner.Text())
	}
	file.Close()
	return namespaces

}
