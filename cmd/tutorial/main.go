package tutorial

import (
	"fmt"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"os"
	"sigs.k8s.io/scheduler-plugins/pkg/trimaran/targetloadpacking"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(targetloadpacking.Name, targetloadpacking.New))
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}