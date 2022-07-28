package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"github.com/muhammadarash1997/grpc4/pb"
	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu      sync.Mutex
	student []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.student {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, errors.New("Student not found")
}

func (d *dataStudentServer) loadData() {
	dataByte, err := ioutil.ReadFile("./data/datas.json")
	if err != nil {
		log.Fatalln("Error reading data", err.Error())
	}

	err = json.Unmarshal(dataByte, &d.student)
	if err != nil {
		log.Fatalln("Error unmashalling data json", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error listening", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln("Error serve grpc", err.Error())
	}
}
