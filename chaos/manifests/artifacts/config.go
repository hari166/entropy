package artifacts

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Config() (k *kubernetes.Clientset, err error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})

	config, err := kConfig.ClientConfig()
	if err != nil {
		fmt.Println(err)
	}
	clients, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	return clients, nil
}
