package main

import (
	"flag"
	"os"
	"path/filepath"

	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
	// Support gcp and other authentication schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	nodeLabel := flag.String("nodelabel", "", "Label to match for nodes, if blank grab all nodes")
	nameSpace := flag.String("namespace", "", "Namespace to grab capacity usage from")
	namespaceList := flag.String("namespaceList", "", "Filepath containing a list of namespaces, one per line")
	daemonMode := flag.Bool("daemon", false, "Run in daemon mode")
	jsonMode := flag.Bool("json", false, "Output information in json format")
	checkMode := flag.Bool("check", false, "Check kubernetes connection")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// no config, maybe we are inside a kubernetes cluster.
		config, err = rest.InClusterConfig()
		check(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	check(err)

	if *checkMode {
		_, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		check(err)
		fmt.Println("ok")
		return
	}

	// BreakOut to namespace if asked
	if *nameSpace != "" {
		nsInfo := gatherNamespaceInfo(clientset, nameSpace)
		if *jsonMode {
			result, err := json.Marshal(nsInfo)
			check(err)
			fmt.Println(string(result))
			return
		}
		output := namespaceHumanMode(nsInfo)
		for _, line := range output {
			fmt.Println(line)
		}
		return
	}

	// Gather info
	if *daemonMode {
		for {
			clusterInfo := gatherInfo(clientset, nodeLabel)
			getCapcity(clusterInfo)
			time.Sleep(300 * time.Second)
		}
	} else if *jsonMode {
		clusterInfo := gatherInfo(clientset, nodeLabel)
		getCapcity(clusterInfo)

	} else if *namespaceList != "" {
		nsList := getNamespaceListFromFile(*namespaceList)
		var totalMemoryRequests int64
		var totalMemoryLimits int64
		var totalMemoryUsed int64
		var totalCPURequestsMilliCores int64
		var totalCPULimitsMilliCores int64
		var totalCPUUsedMilliCores int64
		for _, namespace := range nsList {
			fmt.Println("=========================")
			nsInfo := gatherNamespaceInfo(clientset, &namespace)
			fmt.Println(namespace)
			fmt.Printf("NamespaceMemoryRequestsGib %v\n", toGibFromByte(nsInfo.NamespaceMemoryRequests))
			totalMemoryRequests = totalMemoryRequests + nsInfo.NamespaceMemoryRequests
			fmt.Printf("NamespaceMemoryLimitsGib %v\n", toGibFromByte(nsInfo.NamespaceMemoryLimits))
			totalMemoryLimits = totalMemoryLimits + nsInfo.NamespaceMemoryLimits
			fmt.Printf("NamespaceMemoryUsedGib %v\n", toGibFromByte(nsInfo.NamespaceMemoryUsed))
			totalMemoryUsed = totalMemoryUsed + nsInfo.NamespaceMemoryUsed
			fmt.Printf("NamespaceCPURequestsMilliCores %v\n", nsInfo.NamespaceCPURequestsMilliCores)
			totalCPURequestsMilliCores = totalCPURequestsMilliCores + nsInfo.NamespaceCPURequestsMilliCores
			fmt.Printf("NamespaceCPULimitsMilliCores %v\n", nsInfo.NamespaceCPULimitsMilliCores)
			totalCPULimitsMilliCores = totalCPULimitsMilliCores + nsInfo.NamespaceCPULimitsMilliCores
			fmt.Printf("NamespaceCPUUsedMilliCores %v\n", nsInfo.NamespaceCPUUsedMilliCores)
			totalCPUUsedMilliCores = totalCPUUsedMilliCores + nsInfo.NamespaceCPUUsedMilliCores
		}
		fmt.Println("=========================")
		fmt.Println("Total for all namespaces combined")
		fmt.Printf("NamespaceMemoryRequestsGiB %v\n", toGibFromByte(totalMemoryRequests))
		fmt.Printf("NamespaceMemoryLimitsGiB %v\n", toGibFromByte(totalMemoryLimits))
		fmt.Printf("NamespaceMemoryUsedGiB %v\n", toGibFromByte(totalMemoryUsed))
		fmt.Printf("NamespaceCPURequestsCores %v\n", float64(totalCPURequestsMilliCores)/1000)
		fmt.Printf("NamespaceCPULimitsCores %v\n", float64(totalCPULimitsMilliCores)/1000)
		fmt.Printf("NamespaceCPUUsedCores %v\n", float64(totalCPUUsedMilliCores)/1000)
		fmt.Println("=========================")

	} else {
		clusterInfo := gatherInfo(clientset, nodeLabel)
		humanMode(clusterInfo)
	}
}
