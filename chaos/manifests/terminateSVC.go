package manifests

import (
	"context"
	"fmt"

	"github.com/hari166/entropy/chaos/manifests/artifacts"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Service = &cobra.Command{
	Use:   "service",
	Short: "Terminate a service",
	Long:  `Kill a service to observe induced effects`,
	Run: func(cmd *cobra.Command, args []string) {
		nameSVC, _ := cmd.Flags().GetString("name")
		nameSpace, _ := cmd.Flags().GetString("ns")
		TerminateSVC(nameSVC, nameSpace)
	},
}

func TerminateSVC(nameSVC, nameSpace string) {
	client, err := artifacts.Config()
	if err != nil {
		fmt.Println(err)
	}

	svcInfo, err := client.CoreV1().Services(nameSpace).Get(context.Background(), nameSVC, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	err = client.CoreV1().Services(nameSpace).Delete(context.Background(), nameSVC, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		svc := svcInfo.Name
		fmt.Printf("Deleted service %s", svc)
	}

}
