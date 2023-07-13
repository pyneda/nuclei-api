package main

import (
	"context"
	pb "github.com/pyneda/nuclei-api/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"time"
)

var (
	scanRequest = &pb.ScanRequest{}
	apiAddress  string
)

// ClientCmd represents the client command
var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "An example client for nuclei-api",
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
	ClientCmd.Flags().StringVar(&apiAddress, "address", "localhost:8555", "Nuclei API address")
	ClientCmd.Flags().StringSliceVarP(&scanRequest.Targets, "target", "t", []string{}, "Target to scan")
	ClientCmd.MarkFlagRequired("target")
	ClientCmd.Flags().BoolVarP(&scanRequest.AutomaticScan, "automatic-scan", "a", false, "Automatic scan")
	ClientCmd.Flags().StringSliceVar(&scanRequest.IncludeIds, "ids", []string{}, "Include IDS")
	ClientCmd.Flags().StringSliceVar(&scanRequest.ExcludeIds, "exclude-ids", []string{}, "Exclude IDS")
	ClientCmd.Flags().StringSliceVar(&scanRequest.Tags, "tags", []string{}, "Tags to include")
	ClientCmd.Flags().StringSliceVar(&scanRequest.ExcludeTags, "exclude-tags", []string{}, "Tags to exclude")
	ClientCmd.Flags().StringSliceVar(&scanRequest.Workflows, "workflows", []string{}, "Workflows to include")
	ClientCmd.Flags().StringSliceVar(&scanRequest.WorkflowUrls, "workflow-urls", []string{}, "Workflows to include")
	ClientCmd.Flags().StringSliceVar(&scanRequest.Templates, "templates", []string{}, "Templates to include")
	ClientCmd.Flags().StringSliceVar(&scanRequest.TemplateUrls, "template-urls", []string{}, "Templates to include")
	ClientCmd.Flags().StringSliceVar(&scanRequest.ExcludedTemplates, "excluded-templates", []string{}, "Templates to exclude")
	ClientCmd.Flags().StringSliceVar(&scanRequest.Authors, "authors", []string{}, "Authors to include")
	ClientCmd.Flags().StringSliceVar(&scanRequest.ExcludeMatchers, "exclude-matchers", []string{}, "Matchers to exclude")
}

func main() {
	err := ClientCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
