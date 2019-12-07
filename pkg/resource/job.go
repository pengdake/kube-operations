package resource

import(
	batch_v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	"kube-operations/pkg/common"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobController struct {
	Namespace string
	Client batch_v1.JobInterface
}

func NewJobController(namespace string)*JobController{
	kubeclient, err := common.KubeClient()
	if err != nil{
		panic(err.Error())
	}
	jobclient := kubeclient.BatchV1().Jobs(namespace)
	return &JobController{
		Namespace: namespace,
		Client:    jobclient,
	}
}

func(jc *JobController)Delete(job string){
	deletePolicy := meta_v1.DeletePropagationBackground
	operation := &meta_v1.DeleteOptions{PropagationPolicy: &deletePolicy}
	err := jc.Client.Delete(job, operation)
	if err != nil{
		panic(err.Error())
	}
}