package manifests

import (
	"fmt"
	"os"

	"github.com/hari166/entropy/chaos/manifests/artifacts"
	"github.com/spf13/cobra"
)

var Artifact = &cobra.Command{
	Use:   "artifact",
	Short: "Experiment with config maps and secrets",
	Long:  `artifact deals with config maps and secrets`,
}

func init() {
	Artifact.AddCommand(artifacts.Cmap)
	artifacts.Cmap.Flags().String("name", "", "Enter ConfigMap name")
	artifacts.Cmap.Flags().String("ns", "default", "Enter namespace")

	Artifact.AddCommand(artifacts.Secret)
	artifacts.Secret.Flags().String("name", "", "Enter Secret name")
	artifacts.Secret.Flags().String("ns", "default", "Enter namespace")

}

func Execute() {
	if err := Artifact.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
