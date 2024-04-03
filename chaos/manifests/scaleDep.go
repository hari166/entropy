package manifests

import (
	"context"
	"fmt"

	"github.com/hari166/entropy/chaos/manifests/artifacts"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Scale = &cobra.Command{
	Use:   "scale",
	Short: "Scale a deployment",
	Long:  `Add/Reduce number of replicas to observe changes/effects`,
	Run: func(cmd *cobra.Command, args []string) {
		nameDep, _ := cmd.Flags().GetString("name")
		replicas, _ := cmd.Flags().GetInt("replicas")
		nameSpace, _ := cmd.Flags().GetString("ns")
		ScaleDep(nameDep, nameSpace, replicas)
	},
}

func ScaleDep(nameDep, nameSpace string, replicas int) {
	clientset, err := artifacts.Config()
	if err != nil {
		fmt.Println(err)
	}

	scaleInfo, err := clientset.AppsV1().Deployments(nameSpace).GetScale(context.Background(), nameDep, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	scaleInfo.Spec.Replicas = int32(replicas)

	_, err = clientset.AppsV1().Deployments(nameSpace).UpdateScale(context.Background(), nameDep, scaleInfo, metav1.UpdateOptions{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Deployment scaled to %v replicas", replicas)
	}

}
