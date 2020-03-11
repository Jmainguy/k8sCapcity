# k8sCapcity
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/k8sCapcity)](https://goreportcard.com/report/github.com/Jmainguy/k8sCapcity)
[![Release](https://img.shields.io/github/release/Jmainguy/k8sCapcity.svg?style=flat-square)](https://github.com/Jmainguy/k8sCapcity/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/k8sCapcity/badge.svg?branch=master&service=github)](https://coveralls.io/github/Jmainguy/k8sCapcity?branch=master)

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

## Fields and their meaning
See [Fields](docs/fields.md)

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/k8sCapcity/releases)

## Build
```/bin/bash
export GO111MODULE=on
go get k8s.io/client-go@v12.0.0
go build
```
