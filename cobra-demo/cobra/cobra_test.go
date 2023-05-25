package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"testing"
)

func TestCobra(t *testing.T) {
	var cmdHello = &cobra.Command{
		Use:   "hello [name]",
		Short: "Say hello to someone",
		Long:  "This command will greet the person you provide by name.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, " + args[0] + "!")
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdHello)
	rootCmd.Execute()
}
