# k8sCapcity
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/k8sCapcity)](https://goreportcard.com/report/github.com/Jmainguy/k8sCapcity)
[![Release](https://img.shields.io/github/release/Jmainguy/k8sCapcity.svg?style=flat-square)](https://github.com/Jmainguy/k8sCapcity/releases/latest)

This command is designed to assist in capacity planning, by showing capacity

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
-json flag allows you to output in json format
```/bin/bash
./k8sCapcity -json
```
-daemon flag allows you run in daemon mode, where it outputs info every 5 minutes, in json format
```/bin/bash
./k8sCapcity -daemon
```

## Sample output
```/bin/bash
$ ./k8sCapcity --json --nodelabel node-role.kubernetes.io/compute=true | jq .
{
  "event.kind": "metric",
  "event.module": "k8s_quota",
  "event.provider": "k8sCapcity",
  "event.type": "info",
  "event.version": "03/04/2020-01",
  "k8s_quota.resource_quota.cpu_request.cores": 3087,
  "k8s_quota.resource_quota.cpu_request.millicores": 3087000,
  "k8s_quota.resource_quota.memory_request": 6546603900928,
  "k8s_quota.resource_quota.memory_limit": 11164767485952,
  "k8s_quota.resource_quota.pods": 13616,
  "k8s_quota.subscription_factor.memory.request": 0.8917893806797168,
  "k8s_quota.subscription_factor.memory.request.nminusone": 0.9257141368923372,
  "k8s_quota.subscription_factor.cpu.request": 3.2290794979079496,
  "k8s_quota.subscription_factor.cpu.request.nminusone": 3.484198645598194,
  "k8s_quota.subscription_factor.pods": 6.161085972850679,
  "k8s_quota.subscription_factor.pods.nminusone": 6.546153846153846,
  "k8s_quota.alloctable.memory": 7340975394816,
  "k8s_quota.alloctable.memory.nminusone": 7071949795328,
  "k8s_quota.alloctable.cpu": 956,
  "k8s_quota.alloctable.cpu.nminusone": 886,
  "k8s_quota.alloctable.pods": 2210,
  "k8s_quota.alloctable.pods.nminusone": 2080,
  "k8s_quota.container_resource.cpu_request.cores": 728,
  "k8s_quota.container_resource.cpu_request.millicores": 727044,
  "k8s_quota.container_resource.memory_request": 2430993722701,
  "k8s_quota.container_resource.memory_limit": 3458915138381,
  "k8s_quota.container_resource.pods": 1533,
  "k8s_quota.node_label": "node-role.kubernetes.io/compute",
  "k8s_quota.utilization_factor.pods": {
    "ocp-app-01a.lab1.soh.re": 0.5230769230769231,
    "ocp-app-01b.lab1.soh.re": 0.5,
    "ocp-app-01c.lab1.soh.re": 0.4230769230769231,
    "ocp-app-01d.lab1.soh.re": 1,
    "ocp-app-01e.lab1.soh.re": 0.7153846153846154,
    "ocp-app-01f.lab1.soh.re": 0.6153846153846154,
    "ocp-app-01g.lab1.soh.re": 0.5538461538461539,
    "ocp-app-01h.lab1.soh.re": 0.7384615384615385,
    "ocp-app-01i.lab1.soh.re": 0.676923076923077,
    "ocp-app-01j.lab1.soh.re": 0.8615384615384616,
    "ocp-app-01k.lab1.soh.re": 0.9076923076923077,
    "ocp-app-01l.lab1.soh.re": 0.6,
    "ocp-app-01m.lab1.soh.re": 0.7153846153846154,
    "ocp-app-01n.lab1.soh.re": 0.8076923076923077,
    "ocp-app-01o.lab1.soh.re": 0.7384615384615385,
    "ocp-app-01p.lab1.soh.re": 0.6230769230769231,
    "ocp-app-01q.lab1.soh.re": 0.7923076923076923
  },
  "k8s_quota.utilization_factor.pods.total": 0.6936651583710407,
  "k8s_quota.utilization_factor.pods.total.nminusone": 0.7370192307692308,
  "k8s_quota.utilization_factor.memory_request": {
    "ocp-app-01a.lab1.soh.re": 0.15538375703223023,
    "ocp-app-01b.lab1.soh.re": 0.2092471424042161,
    "ocp-app-01c.lab1.soh.re": 0.3659245978281739,
    "ocp-app-01d.lab1.soh.re": 0.6507702795168727,
    "ocp-app-01e.lab1.soh.re": 0.41168864147353995,
    "ocp-app-01f.lab1.soh.re": 0.47184134886681184,
    "ocp-app-01g.lab1.soh.re": 0.2760629700597209,
    "ocp-app-01h.lab1.soh.re": 0.2805204157266052,
    "ocp-app-01i.lab1.soh.re": 0.3307502581692465,
    "ocp-app-01j.lab1.soh.re": 0.33521561671698297,
    "ocp-app-01k.lab1.soh.re": 0.43173560401002464,
    "ocp-app-01l.lab1.soh.re": 0.23961303425687833,
    "ocp-app-01m.lab1.soh.re": 0.25716686532315547,
    "ocp-app-01n.lab1.soh.re": 0.4287220701551472,
    "ocp-app-01o.lab1.soh.re": 0.3695995767794075,
    "ocp-app-01p.lab1.soh.re": 0.34826314940877007,
    "ocp-app-01q.lab1.soh.re": 0.20000259179325308
  },
  "k8s_quota.utilization_factor.memory_request.total": 0.33115404860472664,
  "k8s_quota.utilization_factor.memory_request.total.nminusone": 0.34375155269159396,
  "k8s_quota.utilization_factor.cpu_request": {
    "ocp-app-01a.lab1.soh.re": 0.75,
    "ocp-app-01b.lab1.soh.re": 0.6875,
    "ocp-app-01c.lab1.soh.re": 0.8125,
    "ocp-app-01d.lab1.soh.re": 0.6285714285714286,
    "ocp-app-01e.lab1.soh.re": 0.8181818181818182,
    "ocp-app-01f.lab1.soh.re": 0.8478260869565217,
    "ocp-app-01g.lab1.soh.re": 0.7857142857142857,
    "ocp-app-01h.lab1.soh.re": 0.7714285714285715,
    "ocp-app-01i.lab1.soh.re": 0.8,
    "ocp-app-01j.lab1.soh.re": 0.8,
    "ocp-app-01k.lab1.soh.re": 0.8571428571428571,
    "ocp-app-01l.lab1.soh.re": 0.7428571428571429,
    "ocp-app-01m.lab1.soh.re": 0.7428571428571429,
    "ocp-app-01n.lab1.soh.re": 0.7714285714285715,
    "ocp-app-01o.lab1.soh.re": 0.8,
    "ocp-app-01p.lab1.soh.re": 0.7857142857142857,
    "ocp-app-01q.lab1.soh.re": 0.7
  },
  "k8s_quota.utilization_factor.cpu_request.total": 0.7615062761506276,
  "k8s_quota.utilization_factor.cpu_request.total.nminusone": 0.8216704288939052,
  "k8s_quota.available.memory_request": 4909981672115,
  "k8s_quota.available.memory_request.nminusone": 4640956072627,
  "k8s_quota.available.cpu_request": 228,
  "k8s_quota.available.cpu_request.nminusone": 158,
  "k8s_quota.available.pods": 677,
  "k8s_quota.available.pods.nminusone": 547
}

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
go get k8s.io/client-go@v12.0.0
go build
```
