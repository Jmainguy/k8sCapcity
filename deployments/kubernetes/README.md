# Deploy to openshift
```/bin/bash
oc new-project k8scapcity
oc create -f serviceAccount.yaml
oc create -f clusterRole.yaml
oc create -f clusterRoleBinding.yaml
oc create -f deployment.yaml
```

## Notes
This assumes that the namespace is k8scapcity

The following in deployment.yaml is optional, only if you want to limit metric gathering to compute nodes
```/bin/bash
          env:
          - name: NODELABEL
            value: node-role.kubernetes.io/compute=true
```
