package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewHelloCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints the hello message",
		Long:  `This is a longer description of the command, which can span multiple lines and contain examples.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}

	return cmd
}
