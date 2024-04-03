package chaos

import (
	"fmt"
	"os"

	"github.com/hari166/entropy/chaos/manifests"
	"github.com/hari166/entropy/chaos/node"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "entropy",
	Short: "A chaos tool for a minikube cluster",
	Long:  `Entropy is a chaos simulator for Kubernetes cluster`,
}

func init() {
	manifests.Kill_Random.Flags().String("ns", "", "Enter namespace")
	RootCmd.AddCommand(manifests.Kill_Random)

	manifests.Inject.Flags().String("name", "", "Enter pod name")
	manifests.Inject.Flags().String("ns", "default", "Enter namespace")
	manifests.Inject.Flags().String("cont", "", "Enter container name")
	RootCmd.AddCommand(manifests.Inject)

	manifests.Scale.Flags().String("name", "", "Enter name of deployment")
	manifests.Scale.Flags().Int("replicas", 7, "Enter number of replicas")
	manifests.Scale.Flags().String("ns", "default", "Enter namespace")
	RootCmd.AddCommand(manifests.Scale)

	manifests.Service.Flags().String("name", "", "Enter name of service")
	manifests.Service.Flags().String("ns", "default", "Enter namespace")
	RootCmd.AddCommand(manifests.Service)

	manifests.Artifact.Flags().String("artifact", "", "Enter artifact type")
	RootCmd.AddCommand(manifests.Artifact)

	node.Cordon.Flags().String("name", "", "Enter node name")
	node.Cordon.Flags().Bool("desired", false, "Enter desired state (true for cordon and vice-versa)\nDefault false")
	RootCmd.AddCommand(node.Cordon)
}
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
