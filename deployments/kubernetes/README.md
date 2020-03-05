# Deploy to openshift
```/bin/bash
oc new-project k8scapcity
oc create -f clusterRole.yaml
oc create -f clusterRoleBinding.yaml
oc create -f deployment.yaml
```

## Notes
This assumes Openshift, and that the namespace is k8scapcity

This should get adjusted to be more generic when possible

The following in deployment.yaml is optional, only if you want to limit metric gathering to compute nodes
```/bin/bash
          env:
          - name: NODELABEL
            value: node-role.kubernetes.io/compute=true
```
