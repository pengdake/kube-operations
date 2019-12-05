package main

import (
	"k8s.io/api/core/v1"
	"fmt"

	"kube-operations/pkg/resource"
)

func main(){
	options := &v1.PodLogOptions{}
	podcontroller := resource.NewPodController("admin")
	log := podcontroller.GetLog("res-imp-xusf-test-u6orzvrx-xrxcc", options)
	fmt.Println(log)
}