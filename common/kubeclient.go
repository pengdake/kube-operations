package common

import(
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"os"
	"path/filepath"
)

func KubeClient()(*kubernetes.Clientset, error){
    home := os.Getenv("HOME")
    kubeconfigPath := filepath.Join(home,".kube","config")
    kubeconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
    if err != nil{
    	panic(err.Error())
	}
	return kubernetes.NewForConfig(kubeconfig)
}
