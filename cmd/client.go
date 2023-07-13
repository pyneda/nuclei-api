/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	pb "github.com/pyneda/nuclei-api/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

var (
	scanRequest = &pb.ScanRequest{}
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(scanRequest.Targets) == 0 {
			log.Fatal("Target is required")
			return
		}

		log.Println("Requesting scan for targets: ", scanRequest.Targets)

		conn, err := grpc.Dial("localhost:8555", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewNucleiApiClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
		defer cancel()
		stream, err := c.Scan(ctx, scanRequest)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("client.Scan failed: %v", err)
			}
			log.Printf("Result %v", result)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringSliceVarP(&scanRequest.Targets, "target", "t", []string{}, "Target to scan")
	clientCmd.MarkFlagRequired("target")
	clientCmd.Flags().BoolVarP(&scanRequest.AutomaticScan, "automatic-scan", "a", false, "Automatic scan")
	clientCmd.Flags().StringSliceVar(&scanRequest.IncludeIds, "ids", []string{}, "Include IDS")
	clientCmd.Flags().StringSliceVar(&scanRequest.ExcludeIds, "exclude-ids", []string{}, "Exclude IDS")
	clientCmd.Flags().StringSliceVar(&scanRequest.Tags, "tags", []string{}, "Tags to include")
	clientCmd.Flags().StringSliceVar(&scanRequest.ExcludeTags, "exclude-tags", []string{}, "Tags to exclude")
	clientCmd.Flags().StringSliceVar(&scanRequest.Workflows, "workflows", []string{}, "Workflows to include")
	clientCmd.Flags().StringSliceVar(&scanRequest.WorkflowUrls, "workflow-urls", []string{}, "Workflows to include")
	clientCmd.Flags().StringSliceVar(&scanRequest.Templates, "templates", []string{}, "Templates to include")
	clientCmd.Flags().StringSliceVar(&scanRequest.TemplateUrls, "template-urls", []string{}, "Templates to include")
	clientCmd.Flags().StringSliceVar(&scanRequest.ExcludedTemplates, "excluded-templates", []string{}, "Templates to exclude")
	clientCmd.Flags().StringSliceVar(&scanRequest.Authors, "authors", []string{}, "Authors to include")
	clientCmd.Flags().StringSliceVar(&scanRequest.ExcludeMatchers, "exclude-matchers", []string{}, "Matchers to exclude")

}
