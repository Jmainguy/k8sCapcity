# Deploy to openshift
```/bin/bash
oc new-project k8scapcity
oc create -f clusterRole.yaml
oc create -f clusterRoleBinding.yaml
oc new-app https://github.com/Jmainguy/k8sCapcity
```

## Notes
This assumes Openshift, and that the namespace is k8scapcity

This should get adjusted to be more generic when possible
