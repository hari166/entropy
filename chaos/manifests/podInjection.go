package manifests

import (
	"context"
	"fmt"
	"time"

	"github.com/hari166/entropy/chaos/manifests/artifacts"
	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Inject = &cobra.Command{
	Use:   "inject",
	Short: "Exit shell script with status code 1",
	Long:  `Terminate shell script with status code 1 to indicate failure/error`,
	Run: func(cmd *cobra.Command, args []string) {
		podName, _ := cmd.Flags().GetString("name")
		contName, _ := cmd.Flags().GetString("cont")
		nameSpace, _ := cmd.Flags().GetString("ns")
		PodInjection(podName, nameSpace, contName)
	},
}

func PodInjection(podName, nameSpace, contName string) {
	client, err := artifacts.Config()
	if err != nil {
		fmt.Println(err)
	}

	podInfo, err := client.CoreV1().Pods(nameSpace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	spec := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podInfo.Name,
			Namespace: podInfo.Namespace,
		},
		Spec: podInfo.Spec,
	}
	
	for i := range spec.Spec.Containers {
		if spec.Spec.Containers[i].Name == contName {
			spec.Spec.Containers[i].Command = []string{"sh", "-c", "echo hello"}
			break
		}
	}
	deleteOption := metav1.DeletePropagationForeground
	err = client.CoreV1().Pods(nameSpace).Delete(context.Background(), podName, metav1.DeleteOptions{
		PropagationPolicy: &deleteOption,
	})
	if err != nil {
		fmt.Println(err)
	}
	//spec := podInfo.DeepCopy()
	time.Sleep(5 * time.Second)
	_, err = client.CoreV1().Pods(nameSpace).Create(context.Background(), spec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("created %s", spec.Name)
	}

}
