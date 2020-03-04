package main

import (
	"fmt"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func namespaceMode(clientset *kubernetes.Clientset, nameSpace *string, containerInfo map[string]ContainerInfo) {
	namespaceMemoryLimits := resource.Quantity{}
	namespaceMemoryRequests := resource.Quantity{}
	namespaceMemoryUsed := resource.Quantity{}
	namespaceCPULimits := resource.Quantity{}
	namespaceCPURequests := resource.Quantity{}
	namespaceCPUUsed := resource.Quantity{}

	podMetricList := getPodMetrics(clientset)
	for _, metricPod := range podMetricList.Items {
		if *nameSpace == metricPod.Namespace {
			pods, err := clientset.CoreV1().Pods(*nameSpace).List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			for _, pod := range pods.Items {
				if pod.Name == metricPod.Name {
					if pod.Status.Phase != "Failed" {
						if pod.Status.Phase != "Succeeded" {
							for _, container := range pod.Spec.Containers {
								uniqueContainerName := fmt.Sprintf("%s-%s", pod.Name, container.Name)
								containerStats := containerInfo[uniqueContainerName]
								crrm := container.Resources.Requests.Memory()
								crrc := container.Resources.Requests.Cpu()
								crlm := container.Resources.Limits.Memory()
								crlc := container.Resources.Limits.Cpu()
								containerStats.MemoryRequests = *crrm
								containerStats.MemoryLimits = *crlm
								containerStats.CPURequests = *crrc
								containerStats.CPULimits = *crlc
								containerStats.Name = container.Name
								containerStats.Pod = pod.Name
								containerInfo[uniqueContainerName] = containerStats
								// Add up for the namespace
								namespaceMemoryLimits.Add(*crlm)
								namespaceMemoryRequests.Add(*crrm)
								namespaceCPULimits.Add(*crlc)
								namespaceCPURequests.Add(*crrc)
							}
						}
					}
				}
			}

			fmt.Println("")
			fmt.Println("================")
			fmt.Printf("****Pod Name: %s****\n", metricPod.Name)
			for _, container := range metricPod.Containers {
				uniqueContainerName := fmt.Sprintf("%s-%s", metricPod.Name, container.Name)
				containerStats := containerInfo[uniqueContainerName]
				containerStats.UsedMemory = *container.Usage.Memory()
				containerStats.UsedCPU = *container.Usage.Cpu()
				containerInfo[uniqueContainerName] = containerStats
				// Add up for the namespace
				namespaceMemoryUsed.Add(*container.Usage.Memory())
				namespaceCPUUsed.Add(*container.Usage.Cpu())
			}
			for _, container := range containerInfo {
				if metricPod.Name == container.Pod {
					fmt.Println("================")
					fmt.Printf("Container Name: %s\n", container.Name)
					fmt.Println("----------------")
					fmt.Printf("CPURequests: %s\n", &container.CPURequests)
					fmt.Printf("MemoryRequests: %dMiB\n", toMib(container.MemoryRequests))
					fmt.Printf("CPULimits: %s\n", &container.CPULimits)
					fmt.Printf("MemoryLimits: %dMiB\n", toMib(container.MemoryLimits))
					fmt.Println("----------------")
					fmt.Printf("Used CPU: %s\n", &container.UsedCPU)
					fmt.Printf("Used Memory: %dMiB\n", toMib(container.UsedMemory))
				}
			}
		}
	}
	fmt.Printf("<><><><><>Sum Total for Namespace: %s<><><><><>\n", *nameSpace)
	fmt.Println("----------------")
	fmt.Printf("Namespace Total CPURequests: %s\n", &namespaceCPURequests)
	fmt.Printf("Namespace Total MemoryRequests: %dMiB\n", toMib(namespaceMemoryRequests))
	fmt.Printf("Namespace Total CPULimits: %s\n", &namespaceCPULimits)
	fmt.Printf("Namespace Total MemoryLimits: %dMiB\n", toMib(namespaceMemoryLimits))
	fmt.Println("----------------")
	fmt.Printf("Namespace Total Used CPU: %s\n", &namespaceCPUUsed)
	fmt.Printf("Namespace Total Used Memory: %dMiB\n", toMib(namespaceMemoryUsed))

	os.Exit(0)
}
