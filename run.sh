#!/bin/bash
if [ -z "$NODELABEL" ]; then
    /go/src/app/k8sCapcity --daemon
else
    /go/src/app/k8sCapcity --daemon --nodelabel $NODELABEL
fi
