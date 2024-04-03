package node

import (
	"context"
	"fmt"

	"github.com/hari166/entropy/chaos/manifests/artifacts"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	drain "k8s.io/kubectl/pkg/drain"
)

var Cordon = &cobra.Command{
	Use:   "cord",
	Short: "Cordon a node",
	Long:  `Cordon a node to mark it unschedulable`,
	Run: func(cmd *cobra.Command, args []string) {
		nodeName, _ := cmd.Flags().GetString("name")
		desired, _ := cmd.Flags().GetBool("desired")
		CordonNode(nodeName, desired)
	},
}

func CordonNode(nodeName string, desired bool) {
	client, err := artifacts.Config()
	if err != nil {
		fmt.Println()
	}

	nodeInfo, err := client.CoreV1().Nodes().Get(context.Background(), nodeName, v1.GetOptions{})
	if err != nil {
		fmt.Println()
	}

	helper := &drain.Helper{
		Client:              client,
		Ctx:                 context.Background(),
		Force:               true,
		IgnoreAllDaemonSets: true,
	}
	if desired {
		err = drain.RunCordonOrUncordon(helper, nodeInfo, true)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Node %s cordoned", nodeInfo.Name)
		}
	} else {
		err = drain.RunCordonOrUncordon(helper, nodeInfo, false)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Node %s uncordoned", nodeInfo.Name)
		}
	}

}
