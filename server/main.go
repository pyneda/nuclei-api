package server

import (
	"github.com/pyneda/nuclei-api/nuclei"
	pb "github.com/pyneda/nuclei-api/service"
	"log"
)

type Server struct {
	pb.UnimplementedNucleiApiServer
}

func (s *Server) Scan(in *pb.ScanRequest, stream pb.NucleiApi_ScanServer) error {
	log.Println("scanning target: ", in.GetUrl())
	nuclei.Scan(in, stream)
	return nil
}
