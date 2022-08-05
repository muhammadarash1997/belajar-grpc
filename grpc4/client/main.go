package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/muhammadarash1997/grpc4/pb"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Fatalln("Error dial")
	}
	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "dimas@gmail.com")
	getDataStudentByEmail(client, "arash@gmail.com")
	getDataAllStudents(client)
}

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("Error find student by email")
	}

	fmt.Println(student)
}

func getDataAllStudents(client pb.DataStudentClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	students, err := client.GetAllStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalln("Error get all students")
	}

	fmt.Println(students.GetStudents())
}
