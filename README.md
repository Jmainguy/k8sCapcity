# k8sCapcity
This command is designed to show assist in capacity planning by showing capacity

## Usage
Login to your kubernetes cluster, then run
```/bin/bash
./k8sCapcity
```
-nodelabel flag allows you to select on only the nodes you care about
```/bin/bash
./k8sCapcity -nodelabel node-role.kubernetes.io/compute=true
```
-namespace flag allows you to focus on a single namespaces usage
```/bin/bash
./k8sCapcity -namespace "aebot"
```

## Sample output
```/bin/bash
There are 19 nodes in this cluster
================
ClusterWide Allocatable Memory: 5178518744Ki (5303GB)
ClusterWide Allocatable CPU: 708
ClusterWide Allocatable Pods: 2470
================
ResourceQuota ClusterWide Allocated Limits.Memory: 7720Gi (8290GB)
ResourceQuota ClusterWide Allocated Limits.CPU: 7092
ResourceQuota ClusterWide Allocated Pods: 11306
================
ResourceQuota ClusterWide Allocated Requests.Memory: 4124Gi (4429GB)
ResourceQuota ClusterWide Allocated Requests.CPU: 2238
================
NodeName: ocp-app-01h.lab1.example.com
Allocatable CPU: 70
Allocatable Memory: 526487648Ki (540GB)
Allocatable Pods: 130
----------------
Used CPU: 5677m
Used Memory: 80510312Ki (83GB)
Used Pods: 105
Used CPU Requests: 61510m
Used Memory Requests: 116567542016 (117GB)
----------------
Available CPU Requests: 8490m
Available Memory Requests: 422555809536 (423GB)
Available Pods: 25
```

## Sample output of just a namespace
```/bin/bash
./k8sCapcity -namespace "aebot"

================
****Pod Name: aebot-65-c4zbf****
================
Container Name: aebot
----------------
CPURequests: 100m
MemoryRequests: 50Mi
CPULimits: 100m
MemoryLimits: 50Mi
----------------
Used CPU: 0
Used Memory: 15452Ki (16MB)
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/k8sCapcity/releases)

## Build
```/bin/bash
export GO111MODULE=on
go mod init
go get k8s.io/client-go@v12.0.0
go build
```

