# k8sCapcity Fields Documentation

<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [k8sCapcity Fields Documentation](#k8scapcity-fields-documentation)   
   - [Event and Node Label](#event-and-node-label)   
   - [Allocatable Resources and Allocatable N-1 Resources](#allocatable-resources-and-allocatable-n-1-resources)   
   - [ResourceQuota Resources](#resourcequota-resources)   
   - [Pod/Container Resources](#podcontainer-resources)   
   - [Utilization Factor](#utilization-factor)   
   - [Subscription Factor](#subscription-factor)   
   - [Available Resources](#available-resources)   
   - [Example Data](#example-data)   

<!-- /MDTOC -->

## Event and Node Label

| Metric Name          | Unit   | Description                                                                                   |
| -------------------- | ------ | --------------------------------------------------------------------------------------------- |
| event.kind           | string | Should always be "metric"                                                                     |
| event.module         | string | Should always be "k8s_quota"                                                                  |
| event.provider       | string | Should always be "k8sCapcity"                                                                 |
| event.type           | string | Should always be "info"                                                                       |
| event.version        | string |                                                                                               |
| k8s_quota.node_label | string | The value passed into k8sCap[acity for label, scopes examination to specific nodes in cluster |

## Allocatable Resources and Allocatable N-1 Resources

Allocatable resources are what kubernetes uses for scheduling pods into nodes. Allocatable N-1 resources are those resources while providing redundancy for the largest node. We use the term N-1 here since we are subtracting the largest node.  This kind of high availability is also referred to as N+1.

| Metric Name                           | Unit  | Formula / Description                                                 |
| ------------------------------------- | ----- | --------------------------------------------------------------------- |
| k8s_quota.alloctable.pods.total       | none  | AppNodes_count * max_pods                                             |
| k8s_quota.alloctable.cpu.total        | cores | AppNode1_allocatable_cpu_cores + ... + AppNodeN_allocatable_cpu_cores |
| k8s_quota.alloctable.memory.total     | bytes | AppNode1_allocatable_memory + ... + AppNodeN_allocatable_memory       |
| k8s_quota.alloctable.pods.nminusone   | none  | k8s_quota.alloctable.pods.total - Largest_node_max_pods               |
| k8s_quota.alloctable.cpu.nminusone    | cores | k8s_quota.alloctable.cpu.total - Largest_node_allocatable_cpu         |
| k8s_quota.alloctable.memory.nminusone | bytes | k8s_quota.alloctable.memory.total - Largest_node_allocatable_memory   |

## ResourceQuota Resources

ResourceQuota that has been handed out

| Metric Name                                     | Unit       | Formula / Description                                                       |
| ----------------------------------------------- | ---------- | --------------------------------------------------------------------------- |
| k8s_quota.resource_quota.pods                   | none       | ResourceQuota1_pods + ... + ResourceQuotaN_pods                             |
| k8s_quota.resource_quota.cpu_request.cores      | cores      | ResourceQuota1_requests_cpu_cores + ... + ResourceQuotaN_requests_cpu_cores |
| k8s_quota.resource_quota.cpu_request.millicores | millicores | ResourceQuota1_requests_cpu_cores + ... + ResourceQuotaN_requests_cpu_cores |
| k8s_quota.resource_quota.memory_request         | bytes      | ResourceQuota1_requests_memory + ... + ResourceQuotaN_requests_memory       |
| k8s_quota.resource_quota.memory_limit           | bytes      | ResourceQuota1_limits_memory + ... + ResourceQuotaN_limits_memory           |

## Pod/Container Resources

Count of non-terminated pods on nodes and the resources consuming according to their associated container specs.

| Metric Name                                         | Unit       | Formula / Description                                              |
| --------------------------------------------------- | ---------- | ------------------------------------------------------------------ |
| k8s_quota.container_resource.pods                   | none       | Count of non-terminated pods on App nodes                          |
| k8s_quota.container_resource.cpu_request.cores      | cores      | Sum of non-terminated pods on App Nodes requests.cpu               |
| k8s_quota.container_resource.cpu_request.millicores | millicores | Sum of non-terminated pods on App Nodes requests.cpu in millicores |
| k8s_quota.container_resource.memory_request         | bytes      | Sum of non-terminated pods on App Nodes requests.memory            |
| k8s_quota.container_resource.memory_limit           | bytes      | Sum of non-terminated pods on App Nodes limits.memory              |

## Utilization Factor

The utilization factor is the percentage (0-1) of allocatable resources in use from the various objects in kubernetes that consume resources.  Essentially it is the sum of all containers in a pod manifest/spec by resource component divided by the allocatable resource components. (This is not actual percent usage of say cpu)

| Metric Name                                           | Unit    | Formula / Description                                                               |
| ----------------------------------------------------- | ------- | ----------------------------------------------------------------------------------- |
| k8s_quota.utilization_factor.pods.total               | percent | k8s_quota.container_resource.pods / k8s_quota.alloctable.pods.total                 |
| k8s_quota.utilization_factor.cpu_request.total        | percent | k8s_quota.container_resource.cpu_request.cores / k8s_quota.alloctable.cpu.total     |
| k8s_quota.utilization_factor.memory_request.total     | percent | k8s_quota.container_resource.memory_request / k8s_quota.alloctable.memory.total     |
| k8s_quota.utilization_factor.pods.nminusone           | percent | k8s_quota.container_resource.pods / k8s_quota.alloctable.pods.nminusone             |
| k8s_quota.utilization_factor.cpu_request.nminusone    | percent | k8s_quota.container_resource.cpu_request.cores / k8s_quota.alloctable.cpu.nminusone |
| k8s_quota.utilization_factor.memory_request.nminusone | percent | k8s_quota.container_resource.memory_request / k8s_quota.alloctable.memory.nminusone |

There are also per node utilization factors to quickly see completely full nodes per resource component.

## Subscription Factor

The subscription factor is the "percentage" or ratio in which resourcequota has been distributed compared to the actual allocatable resources.  Essentially it is the sum of all resourcequotas divided by the allocatable resources.  In a "perfect" cluster with every deployment being exactly blue/green (A deployment requiring 2*N where N is the number of resources required) a "full" cluster would have a subscription factor of 2.

| Metric Name                                            | Unit    | Formula / Description                                                           |
| ------------------------------------------------------ | ------- | ------------------------------------------------------------------------------- |
| k8s_quota.subscription_factor.pods.total               | percent | k8s_quota.resource_quota.pods / k8s_quota.alloctable.pods.total                 |
| k8s_quota.subscription_factor.cpu_request.total        | percent | k8s_quota.resource_quota.cpu_request.cores / k8s_quota.alloctable.cpu.total     |
| k8s_quota.subscription_factor.memory_request.total     | percent | k8s_quota.resource_quota.memory_request / k8s_quota.alloctable.memory.total     |
| k8s_quota.subscription_factor.pods.total.nminusone     | percent | k8s_quota.resource_quota.pods / k8s_quota.alloctable.pods.nminusone             |
| k8s_quota.subscription_factor.cpu_request.nminusone    | percent | k8s_quota.resource_quota.cpu_request.cores / k8s_quota.alloctable.cpu.nminusone |
| k8s_quota.subscription_factor.memory_request.nminusone | percent | k8s_quota.resource_quota.memory_request / k8s_quota.alloctable.memory.nminusone |

## Available Resources

The remaining amount of a resource in aggregate across a cluster irregardless of the actual usage and irregardless of distributed resourcequota.

| Metric Name                                  | Unit  | Formula / Description                                                               |
| -------------------------------------------- | ----- | ----------------------------------------------------------------------------------- |
| k8s_quota.available.pods.total               | none  | k8s_quota.alloctable.pods.total - k8s_quota.container_resource.pods                 |
| k8s_quota.available.cpu_request.total        | cores | k8s_quota.alloctable.cpu.total - k8s_quota.container_resource.cpu_request.cores     |
| k8s_quota.available.memory_request.total     | bytes | k8s_quota.alloctable.memory.total - k8s_quota.container_resource.memory_request     |
| k8s_quota.available.pods.nminusone           | none  | k8s_quota.alloctable.pods.nminusone - k8s_quota.container_resource.pods             |
| k8s_quota.available.cpu_request.nminusone    | cores | k8s_quota.alloctable.cpu.nminusone - k8s_quota.container_resource.cpu_request.cores |
| k8s_quota.available.memory_request.nminusone | bytes | k8s_quota.alloctable.memory.nminusone - k8s_quota.container_resource.memory_request |


## Example Data

```
$ k8sCapcity -json -nodelabel 'node-role.kubernetes.io/compute=true' | jq '.'
{
  "event.kind": "metric",
  "event.module": "k8s_quota",
  "event.provider": "k8sCapcity",
  "event.type": "info",
  "event.version": "03/04/2020-01",
  "k8s_quota.resource_quota.cpu_request.cores": 3073,
  "k8s_quota.resource_quota.cpu_request.millicores": 3073000,
  "k8s_quota.resource_quota.memory_request": 6515465388032,
  "k8s_quota.resource_quota.memory_limit": 11065983238144,
  "k8s_quota.resource_quota.pods": 13476,
  "k8s_quota.subscription_factor.memory.request.total": 0.8875476401450748,
  "k8s_quota.subscription_factor.memory.request.nminusone": 0.9213110353719374,
  "k8s_quota.subscription_factor.cpu.request.total": 3.2144351464435146,
  "k8s_quota.subscription_factor.cpu.request.nminusone": 3.4683972911963883,
  "k8s_quota.subscription_factor.pods.total": 6.097737556561086,
  "k8s_quota.subscription_factor.pods.nminusone": 6.4788461538461535,
  "k8s_quota.alloctable.memory.total": 7340975394816,
  "k8s_quota.alloctable.memory.nminusone": 7071949795328,
  "k8s_quota.alloctable.cpu.total": 956,
  "k8s_quota.alloctable.cpu.nminusone": 886,
  "k8s_quota.alloctable.pods.total": 2210,
  "k8s_quota.alloctable.pods.nminusone": 2080,
  "k8s_quota.container_resource.cpu_request.cores": 730,
  "k8s_quota.container_resource.cpu_request.millicores": 729644,
  "k8s_quota.container_resource.memory_request": 2430308188237,
  "k8s_quota.container_resource.memory_limit": 3468113481293,
  "k8s_quota.container_resource.pods": 1503,
  "k8s_quota.node_label": "node-role.kubernetes.io/compute",
  "k8s_quota.utilization_factor.pods": {
    "ocp-app-01a.lab1.bandwidthclec.local": 0.4846153846153846,
    "ocp-app-01b.lab1.bandwidthclec.local": 0.47692307692307695,
    "ocp-app-01c.lab1.bandwidthclec.local": 0.5,
    "ocp-app-01d.lab1.bandwidthclec.local": 1,
    "ocp-app-01e.lab1.bandwidthclec.local": 0.6846153846153846,
    "ocp-app-01f.lab1.bandwidthclec.local": 0.5615384615384615,
    "ocp-app-01g.lab1.bandwidthclec.local": 0.6153846153846154,
    "ocp-app-01h.lab1.bandwidthclec.local": 0.6923076923076923,
    "ocp-app-01i.lab1.bandwidthclec.local": 0.6692307692307692,
    "ocp-app-01j.lab1.bandwidthclec.local": 0.8307692307692308,
    "ocp-app-01k.lab1.bandwidthclec.local": 0.8615384615384616,
    "ocp-app-01l.lab1.bandwidthclec.local": 0.5692307692307692,
    "ocp-app-01m.lab1.bandwidthclec.local": 0.7461538461538462,
    "ocp-app-01n.lab1.bwnet.us": 0.6846153846153846,
    "ocp-app-01o.lab1.bwnet.us": 0.8692307692307693,
    "ocp-app-01p.lab1.bwnet.us": 0.6076923076923076,
    "ocp-app-01q.lab1.bwnet.us": 0.7076923076923077
  },
  "k8s_quota.utilization_factor.pods.total": 0.6800904977375566,
  "k8s_quota.utilization_factor.pods.nminusone": 0.7225961538461538,
  "k8s_quota.utilization_factor.memory_request": {
    "ocp-app-01a.lab1.bandwidthclec.local": 0.1604980278677086,
    "ocp-app-01b.lab1.bandwidthclec.local": 0.20450092541493772,
    "ocp-app-01c.lab1.bandwidthclec.local": 0.39484849406629124,
    "ocp-app-01d.lab1.bandwidthclec.local": 0.6911502559231126,
    "ocp-app-01e.lab1.bandwidthclec.local": 0.39758870522938355,
    "ocp-app-01f.lab1.bandwidthclec.local": 0.4972231489253001,
    "ocp-app-01g.lab1.bandwidthclec.local": 0.2905845208429357,
    "ocp-app-01h.lab1.bandwidthclec.local": 0.2522023408040453,
    "ocp-app-01i.lab1.bandwidthclec.local": 0.32816339928020316,
    "ocp-app-01j.lab1.bandwidthclec.local": 0.33478116556956766,
    "ocp-app-01k.lab1.bandwidthclec.local": 0.42469723420388217,
    "ocp-app-01l.lab1.bandwidthclec.local": 0.23481917228652785,
    "ocp-app-01m.lab1.bandwidthclec.local": 0.2662394860777249,
    "ocp-app-01n.lab1.bwnet.us": 0.4060353058207202,
    "ocp-app-01o.lab1.bwnet.us": 0.39657278765181314,
    "ocp-app-01p.lab1.bwnet.us": 0.3563309947144189,
    "ocp-app-01q.lab1.bwnet.us": 0.1722227230809698
  },
  "k8s_quota.utilization_factor.memory_request.total": 0.3310606639484473,
  "k8s_quota.utilization_factor.memory_request.nminusone": 0.3436546155690407,
  "k8s_quota.utilization_factor.cpu_request": {
    "ocp-app-01a.lab1.bandwidthclec.local": 0.75,
    "ocp-app-01b.lab1.bandwidthclec.local": 0.6875,
    "ocp-app-01c.lab1.bandwidthclec.local": 0.8125,
    "ocp-app-01d.lab1.bandwidthclec.local": 0.6,
    "ocp-app-01e.lab1.bandwidthclec.local": 0.7727272727272727,
    "ocp-app-01f.lab1.bandwidthclec.local": 0.8913043478260869,
    "ocp-app-01g.lab1.bandwidthclec.local": 0.7857142857142857,
    "ocp-app-01h.lab1.bandwidthclec.local": 0.7428571428571429,
    "ocp-app-01i.lab1.bandwidthclec.local": 0.7857142857142857,
    "ocp-app-01j.lab1.bandwidthclec.local": 0.8,
    "ocp-app-01k.lab1.bandwidthclec.local": 0.8428571428571429,
    "ocp-app-01l.lab1.bandwidthclec.local": 0.7285714285714285,
    "ocp-app-01m.lab1.bandwidthclec.local": 0.8428571428571429,
    "ocp-app-01n.lab1.bwnet.us": 0.7571428571428571,
    "ocp-app-01o.lab1.bwnet.us": 0.8,
    "ocp-app-01p.lab1.bwnet.us": 0.8,
    "ocp-app-01q.lab1.bwnet.us": 0.7142857142857143
  },
  "k8s_quota.utilization_factor.cpu_request.total": 0.7635983263598326,
  "k8s_quota.utilization_factor.cpu_request.nminusone": 0.8239277652370203,
  "k8s_quota.available.memory_request.total": 4910667206579,
  "k8s_quota.available.memory_request.nminusone": 4641641607091,
  "k8s_quota.available.cpu_request.total": 226,
  "k8s_quota.available.cpu_request.nminusone": 156,
  "k8s_quota.available.pods.total": 707,
  "k8s_quota.available.pods.nminusone": 577
}
```
