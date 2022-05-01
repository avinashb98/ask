package cmd

import (
	"github.com/avinashb98/ask/lib"
	"github.com/spf13/cobra"
	"os"
)

var input = Input{}

var rootCmd = &cobra.Command{
	Use:   "ask (GET)[POST] [URL] [-H] [-d]",
	Short: "Ask is cli-based http client built with go",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 2 {
			panic("more than 2 arguments passed")
		}
		if len(args) == 1 {
			input.URL = args[0]
		} else {
			input.Method = args[0]
			input.URL = args[1]
		}

		err := input.Validate()
		if err != nil {
			return err
		}

		request, err := ParseInputToRequest(input)
		if err != nil {
			return err
		}

		//request.Print()

		response, err := lib.ExecuteRequest(request)
		if err != nil {
			return err
		}
		response.Print()
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input.Headers, "headers", "H", "", "headers for the request. format: 'key:value key:value'")
	rootCmd.PersistentFlags().StringVarP(&input.Body, "body", "d", "", "body for the request. format: json")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
