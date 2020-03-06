#!/bin/bash
VERSION=0.2.1
docker build -t=push.soh.re/k8scapcity:0.2.1
docker tag push.soh.re/k8scapcity:0.2.1 push.soh.re/k8scapcity:latest
docker push push.soh.re/k8scapcity:0.2.1
docker push push.soh.re/k8scapcity:0.2.1
