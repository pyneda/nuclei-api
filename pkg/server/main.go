package server

import (
	"github.com/pyneda/nuclei-api/pkg/nuclei"
	pb "github.com/pyneda/nuclei-api/pkg/service"
	"log"
)

type Server struct {
	pb.UnimplementedNucleiApiServer
}

func (s *Server) Scan(in *pb.ScanRequest, stream pb.NucleiApi_ScanServer) error {
	log.Println("Got a request to scan: ", in.Targets)
	nuclei.Scan(in, stream)
	return nil
}
