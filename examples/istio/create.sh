#!/bin/sh

kubectl create -f quick_start.yaml
kubectl create -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml
