package main

import (
	"flag"
	"os"
	"path/filepath"

	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv1b1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"time"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getNodeMetrics(clientset *kubernetes.Clientset) (nodeMetricList *metricsv1b1.NodeMetricsList) {
	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &nodeMetricList); err != nil {
		panic(err)
	}
	return nodeMetricList
}

func getPodMetrics(clientset *kubernetes.Clientset) (podMetricList *metricsv1b1.PodMetricsList) {
	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &podMetricList); err != nil {
		panic(err)
	}
	return podMetricList
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
	daemonMode := flag.Bool("daemon", false, "Run in daemon mode")
	jsonMode := flag.Bool("json", false, "Output information in json format")
	flag.Parse()

	containerInfo := make(map[string]ContainerInfo)
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// no config, maybe we are inside a kubernetes cluster.
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// BreakOut to namespace if asked
	if *nameSpace != "" {
		namespaceMode(clientset, nameSpace, containerInfo)
	}

	// Gather info
	if *daemonMode {
		for {
			clusterInfo := gatherInfo(clientset, nodeLabel)
			runDaemonMode(clusterInfo)
			time.Sleep(300 * time.Second)
		}
	} else if *jsonMode {
		clusterInfo := gatherInfo(clientset, nodeLabel)
		runDaemonMode(clusterInfo)

	} else {
		clusterInfo := gatherInfo(clientset, nodeLabel)
		fmt.Println(len(clusterInfo.NodeInfo))
		humanMode(clusterInfo)
	}
}
