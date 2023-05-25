package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

//go run main.go hello Bob

func main() {
	var cmdVersion = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("1.0")
		},
	}

	var cmdHello = &cobra.Command{
		Use: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello")
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "app [name]",
		Short: "Say hello to someone",
		Long:  "This command will greet the person you provide by name.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	rootCmd.AddCommand(cmdVersion)
	rootCmd.AddCommand(cmdHello)
	rootCmd.Execute()
}
