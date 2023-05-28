package cobra

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestCobra(t *testing.T) {
	cmdHello := &cobra.Command{
		Use:   "hello [name]",
		Short: "Say hello to someone",
		Long:  "This command will greet the person you provide by name.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, " + args[0] + "!")
		},
	}

	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdHello)
	rootCmd.Execute()
}
