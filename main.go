package main

import (
	//"k8s.io/api/core/v1"
	//"fmt"
	//"flag"
	"kube-operations/pkg/resource"
)


func main(){
	//todo: get pod log
	//options := &v1.PodLogOptions{}
	//podcontroller := resource.NewPodController("admin")
	//log := podcontroller.GetLog("res-imp-xusf-test-u6orzvrx-xrxcc", options)
	//fmt.Println(log)

	//todo: delete err pod
	podcontroller := resource.NewPodController("admin")
	podcontroller.DeleteErr()

	//todo: delete job
	//jobcontroller := resource.NewJobController("wy-project-test")
    //jobcontroller.Delete("res-imp-wyaaab-xunlian-vrpao7tb")

}