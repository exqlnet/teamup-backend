package main

import (
	"github.com/micro/go-micro/v2"
	"log"
	activity_pb "teamup/activity/proto"
	"teamup/activity/svc"
)

const SvcName  = "go.micro.teamup.svc.activity"

func main() {
	service := micro.NewService(
		micro.Name(SvcName),
	)

	activity_pb.RegisterActivityServiceHandler(service.Server(), &svc.ActivityServiceImpl{})
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}