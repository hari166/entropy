package artifacts

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Cmap = &cobra.Command{
	Use:   "cmap",
	Short: "Delete a config map",
	Long:  `Remove a config map to observe effects`,
	Run: func(cmd *cobra.Command, args []string) {
		cmapName, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("ns")
		cMap(cmapName, namespace)
	},
}

func cMap(cmapName, namespace string) {
	client, err := Config()
	if err != nil {
		fmt.Println(err)
	}

	err = client.CoreV1().ConfigMaps(namespace).Delete(context.Background(), cmapName, v1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Config map %s deleted", cmapName)
	}
}
