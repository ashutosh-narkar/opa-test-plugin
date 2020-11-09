#!/bin/sh

kubectl delete -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml
kubectl delete -f quick_start.yaml
