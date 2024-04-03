package artifacts

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Secret = &cobra.Command{
	Use:   "secret",
	Short: "Delete a secret",
	Long:  `Remove a secret to observe effects`,
	Run: func(cmd *cobra.Command, args []string) {
		secretName, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("ns")
		secretFunc(secretName, namespace)
	},
}

func secretFunc(secretName, namespace string) {
	client, err := Config()
	if err != nil {
		fmt.Println(err)
	}

	err = client.CoreV1().Secrets(namespace).Delete(context.Background(), secretName, v1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Secret %s deleted", secretName)
	}
}
