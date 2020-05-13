package svc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	activity_pb "teamup/activity/proto"
	"testing"
)

var cli *ActivityServiceImpl

func init() {
	cli = &ActivityServiceImpl{}
}


func TestGetActivityByID(t *testing.T) {
	act := &activity_pb.Activity{}
	err := cli.GetActivityByID(context.Background(), &activity_pb.IntWrap{Val: int32(1)}, act)
	if err != nil {
		t.Fatal(err)
	}
	println(act.ActivityId, act.ActivityName)
}

func TestCreateActivity(t *testing.T)  {
	var processes []*activity_pb.Process
	processes = append(processes, &activity_pb.Process{
		ProcessName:          "产品设计",
		StartTime:            int64(1589378760),
		Tasks:                []*activity_pb.Task{},
	})
	req := &activity_pb.CreateActivityReq{
		ActivityName:         "test-act",
		Introduction:         "活动介绍",
		Roles:                "产品经理,研发",
		CreatorId:            1,
		Illustration:         "封面",
		AuthorityCode:        1,
		Processes:            processes,
		StartTime:            int64(1589378760),
		EndTime:int64(1589999999),
	}
	rsp := &activity_pb.IntWrap{}
	err := cli.CreateActivity(context.Background(), req, rsp)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.Val)
}

func TestDeleteActivity(t *testing.T)  {
	var processes []*activity_pb.Process
	var err error
	processes = append(processes, &activity_pb.Process{
		ProcessName:          "产品设计",
		StartTime:            int64(1589378760),
		Tasks:                []*activity_pb.Task{},
	})
	req := &activity_pb.CreateActivityReq{
		ActivityName:         "test-act",
		Introduction:         "活动介绍",
		Roles:                "产品经理,研发",
		CreatorId:            1,
		Illustration:         "封面",
		AuthorityCode:        1,
		Processes:            processes,
		StartTime:            int64(1589378760),
		EndTime:int64(1589999999),
	}
	rsp := &activity_pb.IntWrap{}
	err = cli.CreateActivity(context.Background(), req, rsp)

	if err != nil {
		t.Fatal(err)
	}

	// 拿到act id，删除act
	activityId := rsp.Val
	err = cli.DeleteActivity(context.Background(), &activity_pb.IntWrap{Val: activityId}, empty.Empty{})
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateActivity(t *testing.T)  {
	err := cli.UpdateActivity(context.Background(), &activity_pb.UpdateActivityReq{
		ActivityId:           1,
		Introduction:         "已经修改的introduction",
		Illustration:         "已经修改的illustration",
	}, &empty.Empty{})

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateTeam(t *testing.T)  {
	err := cli.CreateTeam(context.Background(), &activity_pb.CreateTeamReq{
		ActivityId:           1,
		TeamName:             "Best Team",
		Slogan:               "Very Good",
	}, &empty.Empty{})

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTeamListByActivityID(t *testing.T)  {
	actList := &activity_pb.TeamList{}
	err := cli.GetTeamListByActivityID(context.Background(), &activity_pb.IntWrap{Val: int32(1)}, actList)
	if err != nil {
		t.Fatal(err)
	}
	if len(actList.TeamList) < 1 {
		t.Fatal("team size too small")
	}
	t.Logf("team size: %d", len(actList.TeamList))
}

