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
$ ./k8sCapcity --json | jq .

{
  "event.kind": "metric",
  "event.module": "k8s_quota",
  "event.provider": "k8sCapcity",
  "event.type": "info",
  "event.version": "03/02/2020-01",
  "k8s_quota.resource_quota.cpu_request.cores": 3087,
  "k8s_quota.resource_quota.cpu_request.millicores": 3087000,
  "k8s_quota.resource_quota.memory_request": 6546603900928,
  "k8s_quota.resource_quota.memory_limit": 11164767485952,
  "k8s_quota.resource_quota.pods": 13601,
  "k8s_quota.oversubscription_factor.memory.request": 0.8776558012147581,
  "k8s_quota.oversubscription_factor.memory.request.nminusone": 0.9104939632423018,
  "k8s_quota.oversubscription_factor.cpu.request": 3.130831643002028,
  "k8s_quota.oversubscription_factor.cpu.request.nminusone": 3.3700873362445414,
  "k8s_quota.oversubscription_factor.pods": 4.5488294314381275,
  "k8s_quota.oversubscription_factor.pods.nminusone": 4.7555944055944055,
  "k8s_quota.alloctable.memory": 7459192877056,
  "k8s_quota.alloctable.memory.nminusone": 7190167277568,
  "k8s_quota.alloctable.cpu": 986,
  "k8s_quota.alloctable.cpu.nminusone": 916,
  "k8s_quota.alloctable.pods": 2990,
  "k8s_quota.alloctable.pods.nminusone": 2860,
  "k8s_quota.container_resource.cpu_request.cores": 746,
  "k8s_quota.container_resource.cpu_request.millicores": 745509,
  "k8s_quota.container_resource.memory_request": 2472648286432,
  "k8s_quota.container_resource.memory_limit": 3499247592448,
  "k8s_quota.container_resource.pods": 1633,
  "k8s_quota.node_label": "",
  "k8s_quota.utilization_factor.pods": {
    "ocp-app-01a.lab1.soh.re": 0.47692307692307695,
    "ocp-app-01b.lab1.soh.re": 0.5076923076923077,
    "ocp-app-01c.lab1.soh.re": 0.38461538461538464,
    "ocp-app-01d.lab1.soh.re": 1,
    "ocp-app-01e.lab1.soh.re": 0.7153846153846154,
    "ocp-app-01f.lab1.soh.re": 0.5153846153846153,
    "ocp-app-01g.lab1.soh.re": 0.5769230769230769,
    "ocp-app-01h.lab1.soh.re": 0.7461538461538462,
    "ocp-app-01i.lab1.soh.re": 0.676923076923077,
    "ocp-app-01j.lab1.soh.re": 0.8846153846153846,
    "ocp-app-01k.lab1.soh.re": 0.9384615384615385,
    "ocp-app-01l.lab1.soh.re": 0.6384615384615384,
    "ocp-app-01m.lab1.soh.re": 0.7615384615384615,
    "ocp-app-01n.lab1.soh.re": 0.6461538461538462,
    "ocp-app-01o.lab1.soh.re": 0.9846153846153847,
    "ocp-app-01p.lab1.soh.re": 0.5769230769230769,
    "ocp-app-01q.lab1.soh.re": 0.7,
    "ocp-infra-01a.lab1.soh.re": 0.13846153846153847,
    "ocp-infra-01b.lab1.soh.re": 0.16153846153846155,
    "ocp-infra-01c.lab1.soh.re": 0.09230769230769231,
    "ocp-master-01a.lab1.soh.re": 0.15384615384615385,
    "ocp-master-01b.lab1.soh.re": 0.15384615384615385,
    "ocp-master-01c.lab1.soh.re": 0.13076923076923078
  },
  "k8s_quota.utilization_factor.memory_request": {
    "ocp-app-01a.lab1.soh.re": 0.15041221378138223,
    "ocp-app-01b.lab1.soh.re": 0.23655541483079373,
    "ocp-app-01c.lab1.soh.re": 0.3092939861405514,
    "ocp-app-01d.lab1.soh.re": 0.6607483431848239,
    "ocp-app-01e.lab1.soh.re": 0.38664043822989524,
    "ocp-app-01f.lab1.soh.re": 0.4581393928217036,
    "ocp-app-01g.lab1.soh.re": 0.2859302354145597,
    "ocp-app-01h.lab1.soh.re": 0.27971907274773783,
    "ocp-app-01i.lab1.soh.re": 0.34289753956464597,
    "ocp-app-01j.lab1.soh.re": 0.3504009985519957,
    "ocp-app-01k.lab1.soh.re": 0.4368416942789993,
    "ocp-app-01l.lab1.soh.re": 0.2669152453421946,
    "ocp-app-01m.lab1.soh.re": 0.25797085703838823,
    "ocp-app-01n.lab1.soh.re": 0.4150510790544038,
    "ocp-app-01o.lab1.soh.re": 0.3987434036983667,
    "ocp-app-01p.lab1.soh.re": 0.33615749145738183,
    "ocp-app-01q.lab1.soh.re": 0.18366319224389813,
    "ocp-infra-01a.lab1.soh.re": 0.18804349735493622,
    "ocp-infra-01b.lab1.soh.re": 0.35910406500940384,
    "ocp-infra-01c.lab1.soh.re": 0.14915958721780517,
    "ocp-master-01a.lab1.soh.re": 0.14642597022918205,
    "ocp-master-01b.lab1.soh.re": 0.15150136239442144,
    "ocp-master-01c.lab1.soh.re": 0.1355355591492472
  },
  "k8s_quota.utilization_factor.cpu_request": {
    "ocp-app-01a.lab1.soh.re": 0.75,
    "ocp-app-01b.lab1.soh.re": 0.8125,
    "ocp-app-01c.lab1.soh.re": 0.8125,
    "ocp-app-01d.lab1.soh.re": 0.5714285714285714,
    "ocp-app-01e.lab1.soh.re": 0.8181818181818182,
    "ocp-app-01f.lab1.soh.re": 0.8695652173913043,
    "ocp-app-01g.lab1.soh.re": 0.7714285714285715,
    "ocp-app-01h.lab1.soh.re": 0.7714285714285715,
    "ocp-app-01i.lab1.soh.re": 0.7857142857142857,
    "ocp-app-01j.lab1.soh.re": 0.7857142857142857,
    "ocp-app-01k.lab1.soh.re": 0.8571428571428571,
    "ocp-app-01l.lab1.soh.re": 0.7714285714285715,
    "ocp-app-01m.lab1.soh.re": 0.7571428571428571,
    "ocp-app-01n.lab1.soh.re": 0.8142857142857143,
    "ocp-app-01o.lab1.soh.re": 0.8,
    "ocp-app-01p.lab1.soh.re": 0.7428571428571429,
    "ocp-app-01q.lab1.soh.re": 0.7714285714285715,
    "ocp-infra-01a.lab1.soh.re": 0.5,
    "ocp-infra-01b.lab1.soh.re": 0.5,
    "ocp-infra-01c.lab1.soh.re": 0.5,
    "ocp-master-01a.lab1.soh.re": 0.75,
    "ocp-master-01b.lab1.soh.re": 0.75,
    "ocp-master-01c.lab1.soh.re": 0.75
  }
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
