package main

import (
	"github.com/pengdake/kube-operations/common"
	"k8s.io/api/core/v1"
	"fmt"
)

func main(){
	kubeClient, err := common.KubeClient(restClientConf)
	if err != nil{
		panic(err)
	}
	req := kubeClient.CoreV1().Pods("admin").GetLogs("res-imp-xusf-test-nznhc9zr-snbw5", &v1.PodLogOptions{})
	res, err := req.DoRaw()
	if err != nil{
		panic(err)
	}
	fmt.Println(string(res))
}