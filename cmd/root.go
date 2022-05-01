package cmd

import (
	"fmt"
	"github.com/avinashb98/ask/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var input = Input{}

var rootCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask is cli-based http client built with go",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
			log.Fatalln(err)
		}

		request, err := ParseInputToRequest(input)
		if err != nil {
			log.Fatalln(err)
		}

		request.Print()

		response, err := lib.ExecuteRequest(request)
		if err != nil {
			log.Fatalln(err)
		}
		response.Print()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input.Headers, "headers", "H", "", "headers for the request. format: 'key:value key:value'")
	rootCmd.PersistentFlags().StringVarP(&input.Body, "body", "d", "", "body for the request. format: json")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
