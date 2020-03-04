#!/bin/sh
tar zxvf /usr/sbin/k8sCapcity_Linux_x86_64.tar.gz -C /usr/sbin/
if [ -z "$NODELABEL" ]; then
    /usr/sbin/k8sCapcity --daemon
else
    /usr/sbin/k8sCapcity --daemon --nodelabel $NODELABEL
fi
