package main

import (
	"context"
	"log"
	"net"

	"github.com/muhammadarash1997/grpcmoviescrudapi/pb"
	"google.golang.org/grpc"
)

var movies []*pb.MovieInfo

type Server struct {
	pb.UnimplementedMovieServer
}

func (this *Server) GetMovies(in *pb.Empty, stream pb.Movie_GetMoviesClient) error {
	log.Printf("Received: %v", in)
	for _, movie := range movies {
		err := stream.Send(movie)
		if err != nil {
			return err
		}
	}
}

func (this *Server) GetMovie(ctx context.Context, )

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	// Create gRPC Server
	gRPCServer := grpc.NewServer()

	// Register Server API (Service) into gRPC Server
	pb.RegisterMovieServer(gRPCServer, &Server{})

	log.Printf("Server listening at %v", lis.Addr())

}

func initMovies() {
	movie1 := &pb.MovieInfo{
		Id:    "1",
		Isbn:  "010101",
		Title: "The Batman",
		Director: &pb.Director{
			Firstname: "Matt",
			Lastname:  "Reeves",
		},
	}

	movie2 := &pb.MovieInfo{
		Id:    "2",
		Isbn:  "020202",
		Title: "Doctor Strange",
		Director: &pb.Director{
			Firstname: "Sam",
			Lastname:  "Raimi",
		},
	}

	movies = append(movies, movie1)
	movies = append(movies, movie2)
}
