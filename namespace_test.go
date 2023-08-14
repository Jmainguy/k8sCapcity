package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func compareString(actual, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestNamespaceHumanModeEmpty(t *testing.T) {
	nsInfo := NamespaceInfo{}
	output := namespaceHumanMode(nsInfo)
	compareString(output[0], "", t)
	compareString(output[1], "================", t)
	compareString(output[2], "<><><><><>Sum Total for Namespace: <><><><><>", t)
	compareString(output[3], "----------------", t)
	compareString(output[4], "Namespace Total CPURequests: 0", t)
	compareString(output[5], "Namespace Total MemoryRequests: 0MiB (0.0GiB)", t)
	compareString(output[6], "Namespace Total CPULimits: 0", t)
	compareString(output[7], "Namespace Total MemoryLimits: 0MiB (0.0GiB)", t)
	compareString(output[8], "----------------", t)
	compareString(output[9], "Namespace Total CPU Used: 0", t)
	compareString(output[10], "Namespace Total Memory Used: 0MiB (0.0GiB)", t)
}

func TestGatherPodSpecInfo(t *testing.T) {
	nsInfo := NamespaceInfo{}
	nsInfo.NamespacePods = make(map[string]*Pod)
	containerArray := make(map[string]ContainerInfo)
	c := &ContainerInfo{
		Name:              "testContainer1",
		MemoryUsed:        1024 * 1024 * 1024,
		MemoryUsedMiB:     toMibFromByte(1024 * 1024 * 1024),
		CPUUsedMilliCores: 1000,
		CPUUsedCores:      1,
	}
	containerArray["testContainer1"] = *c

	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "testPod",
		},
		Spec: corev1.PodSpec{},
	}
	nsInfo.NamespacePods["testPod"] = &Pod{Containers: containerArray}
	nsInfo = gatherPodSpecInfo(pod, nsInfo)
	output := namespaceHumanMode(nsInfo)
	compareString(output[0], "", t)
	compareString(output[11], "CPU Used: 1000m", t)
	compareString(output[12], "Memory Used: 1024MiB", t)
	for _, line := range output {
		fmt.Println(line)
	}

}
