package cmd

import (
	"github.com/pyneda/nuclei-api/pkg/server"
	pb "github.com/pyneda/nuclei-api/pkg/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"log"
	"net"
)

var (
	listenAddress string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the GRPC server for nuclei-api",

	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", listenAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Started nuclei-api server on ", listenAddress)

		s := grpc.NewServer()
		pb.RegisterNucleiApiServer(s, &server.Server{})
		reflection.Register(s)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&listenAddress, "address", "a", ":8555", "Address to listen on")
}
