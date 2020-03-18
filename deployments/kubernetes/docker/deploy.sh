#!/bin/bash
VERSION=v0.2.3
docker build -t=push.soh.re/k8scapcity:$VERSION .
docker tag push.soh.re/k8scapcity:$VERSION push.soh.re/k8scapcity:latest
docker push push.soh.re/k8scapcity:$VERSION
docker push push.soh.re/k8scapcity:latest
