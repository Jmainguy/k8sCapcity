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
cd cmd/k8sCapcity
export GO111MODULE=on
go build
```

## Example
```/bin/bash
$ cd cmd/k8sCapcity
$ export GO111MODULE=on
$ go build
go: downloading k8s.io/api v0.0.0-20191010143144-fbf594f18f80
go: downloading k8s.io/apimachinery v0.0.0-20191014065749-fb3eea214746
go: downloading k8s.io/client-go v0.0.0-20191014070654-bd505ee787b2
go: downloading github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
go: extracting github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
go: extracting k8s.io/client-go v0.0.0-20191014070654-bd505ee787b2
go: extracting k8s.io/apimachinery v0.0.0-20191014065749-fb3eea214746
go: downloading golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f
go: downloading gopkg.in/inf.v0 v0.9.0
go: downloading github.com/google/gofuzz v1.0.0
go: downloading golang.org/x/net v0.0.0-20190812203447-cdfb69ac37fc
go: extracting gopkg.in/inf.v0 v0.9.0
go: extracting github.com/google/gofuzz v1.0.0
go: extracting k8s.io/api v0.0.0-20191010143144-fbf594f18f80
go: extracting golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f
go: extracting golang.org/x/net v0.0.0-20190812203447-cdfb69ac37fc
go: downloading github.com/json-iterator/go v1.1.7
go: downloading golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
go: downloading golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
go: downloading github.com/golang/protobuf v1.3.1
go: downloading github.com/spf13/pflag v1.0.3
go: downloading github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d
go: downloading k8s.io/utils v0.0.0-20191010214722-8d271d903fe4
go: extracting golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
go: downloading sigs.k8s.io/yaml v1.1.0
go: extracting github.com/json-iterator/go v1.1.7
go: extracting sigs.k8s.io/yaml v1.1.0
go: extracting github.com/spf13/pflag v1.0.3
go: extracting k8s.io/utils v0.0.0-20191010214722-8d271d903fe4
go: downloading gopkg.in/yaml.v2 v2.2.4
go: extracting golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
go: extracting github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d
go: extracting github.com/golang/protobuf v1.3.1
go: extracting gopkg.in/yaml.v2 v2.2.4
go: finding k8s.io/client-go v0.0.0-20191014070654-bd505ee787b2
go: finding k8s.io/apimachinery v0.0.0-20191014065749-fb3eea214746
go: finding k8s.io/api v0.0.0-20191010143144-fbf594f18f80
go: finding github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
go: finding golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f
go: finding github.com/spf13/pflag v1.0.3
go: finding golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
go: finding github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d
go: finding github.com/golang/protobuf v1.3.1
go: finding golang.org/x/net v0.0.0-20190812203447-cdfb69ac37fc
go: finding gopkg.in/inf.v0 v0.9.0
go: finding github.com/google/gofuzz v1.0.0
go: finding sigs.k8s.io/yaml v1.1.0
go: finding github.com/json-iterator/go v1.1.7
go: finding golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
go: finding k8s.io/utils v0.0.0-20191010214722-8d271d903fe4
go: finding gopkg.in/yaml.v2 v2.2.4
```
