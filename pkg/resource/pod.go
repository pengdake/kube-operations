package resource

import(
	"k8s.io/api/core/v1"
	core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-operations/pkg/common"
	"fmt"
	"sync"
)

type PodController struct {
	Namespace  string
    Client core_v1.PodInterface
}

type PodLog string

func NewPodController(namespace string)*PodController{
	kubeclient, err := common.KubeClient()
	if err != nil{
		panic(err.Error())
	}
	podclient := kubeclient.CoreV1().Pods(namespace)
	return &PodController{
		Namespace: namespace,
		Client: podclient,
	}
}

func(pc *PodController)GetLog(pod string, options *v1.PodLogOptions)PodLog{
	req := pc.Client.GetLogs(pod, options)

	res, err := req.DoRaw()
	if err != nil{
		panic(err.Error())
	}
	return PodLog(res)
}

func(pc *PodController)Delete(pod string){
    options := &meta_v1.DeleteOptions{}
    err := pc.Client.Delete(pod,options)
    if err != nil{
    	panic(err.Error())
	}
	fmt.Printf("success to remove pod:%s\n", pod)
}

func(pc *PodController)List()*v1.PodList{
	options := meta_v1.ListOptions{}
	podList, err := pc.Client.List(options)
	if err != nil{
		panic(err.Error())
	}
	return podList
}

func(pc *PodController)DeleteErr(){
    podList := pc.List()
    wg := sync.WaitGroup{}
    for _, pod := range podList.Items{
        if pod.Status.Phase == v1.PodFailed{
        	fmt.Printf("find err pod:%s\n", pod.Name)
        	wg.Add(1)
			go func(podname string){
				defer wg.Done()
				pc.Delete(podname)
			}(pod.Name)
		}
	}
	wg.Wait()
}