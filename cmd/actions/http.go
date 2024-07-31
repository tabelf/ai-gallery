package actions

import (
	"log"

	"ai-gallery/service"

	"github.com/spf13/cobra"
)

func init() {
	command := &cobra.Command{
		Use: "start",
		Run: Run,
	}
	rootCmd.AddCommand(command)
}

func Run(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		log.Fatalf("get env failed: %v\n", err)
	}
	service.StartHTTP(env)
}
