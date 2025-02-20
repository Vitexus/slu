package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func KubernetesClient() (*kubernetes.Clientset, string, error) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	namespace, _, _ := kubeConfig.Namespace()
	restconfig, _ := kubeConfig.ClientConfig()
	clientset, _ := kubernetes.NewForConfig(restconfig)
	return clientset, namespace, nil
}
