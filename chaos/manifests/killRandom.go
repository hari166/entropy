package manifests

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hari166/entropy/chaos/manifests/artifacts"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Kill_Random = &cobra.Command{
	Use:   "killRandom",
	Short: "Kill a random pod",
	Long:  `Terminate a random pod from a namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		nameSpace, _ := cmd.Flags().GetString("ns")
		KillRandom(nameSpace)
	},
}

func KillRandom(nameSpace string) {
	cmd1 := exec.Command("kubectl", "get", "pod", "-o", "custom-columns=Name:.metadata.name", "--no-headers")
	cmd2 := exec.Command("shuf", "-n", "1")

	var output bytes.Buffer

	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd2.Stdout = &output

	cmd2.Start()
	cmd1.Run()
	cmd2.Wait()
	podName := strings.TrimSpace(output.String())

	client, err := artifacts.Config()
	if err != nil {
		fmt.Println(err)
	}
	if nameSpace == "" {
		fmt.Println("WARNING: Namespace required, even if default")
		os.Exit(0)
	}

	deleteOption := metav1.DeletePropagationForeground
	err = client.CoreV1().Pods(nameSpace).Delete(context.Background(), podName, metav1.DeleteOptions{
		PropagationPolicy: &deleteOption,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Pod terminated: %v\n", podName)

	}

}
