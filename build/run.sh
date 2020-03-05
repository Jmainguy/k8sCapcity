#!/bin/sh
tar zxvf /opt/k8sCapcity_Linux_x86_64.tar.gz -C /opt/
if [ -z "$NODELABEL" ]; then
    /opt/k8sCapcity --daemon
else
    /opt/k8sCapcity --daemon --nodelabel $NODELABEL
fi
