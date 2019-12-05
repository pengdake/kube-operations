package resource

import(
	"k8s.io/api/core/v1"
	client_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"kube-operations/pkg/common"
)

type PodController struct {
	Namespace  string
    Client client_v1.PodInterface
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