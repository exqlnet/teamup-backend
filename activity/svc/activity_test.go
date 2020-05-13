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
	teamID := &activity_pb.IntWrap{}
	err := cli.CreateTeam(context.Background(), &activity_pb.CreateTeamReq{
		ActivityId:int32(1),
		TeamName:             "Best Team",
		Slogan:               "Very Good",
	}, teamID)

	if err != nil {
		t.Fatal(err)
	}

	if teamID.Val < 1 {
		t.Fatal("create failed")
	}

}

func TestGetTeamListByActivityID(t *testing.T)  {
	var err error
	actList := &activity_pb.TeamList{}
	err = cli.GetTeamListByActivityID(context.Background(), &activity_pb.IntWrap{Val: 1}, actList)
	if err != nil {
		t.Fatal(err)
	}
	if len(actList.TeamList) < 1 {
		t.Fatal("team size too small")
	}
	t.Logf("team size: %d", len(actList.TeamList))
}

func TestDeleteTeam(t *testing.T)  {
	teamID := &activity_pb.IntWrap{}
	var err error
	err = cli.CreateTeam(context.Background(), &activity_pb.CreateTeamReq{
		ActivityId:int32(1),
		TeamName:             "Best Team",
		Slogan:               "Very Good",
	}, teamID)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("team id: %d", teamID.Val)
	err = cli.DeleteTeam(context.Background(), &activity_pb.IntWrap{Val: teamID.Val}, &empty.Empty{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateTeam(t *testing.T)  {
	req := &activity_pb.UpdateTeamReq{
		TeamId:               1,
		TeamName:             "修改版老八冠军队",
		Slogan:               "奥力给！！！",
	}

	err := cli.UpdateTeam(context.Background(), req, &empty.Empty{})
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetTeamByID(t *testing.T)  {
	team := &activity_pb.Team{}
	err := cli.GetTeamByID(context.Background(), &activity_pb.IntWrap{Val: int32(1)}, team)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("team_id=%d, team_name=%s", team.TeamId, team.TeamName)
}

func TestGetActivityJoinByUserID(t *testing.T)  {
	joinList := &activity_pb.ActivityJoinList{}
	err := cli.GetActivityJoinByUserID(context.Background(), &activity_pb.IntWrap{Val: int32(1)}, joinList)
	if err != nil {
		t.Fatal(err)
	}

	if len(joinList.ActivityList) < 1 {
		t.Fatal("result size < 1")
	}
	t.Logf("joinList size: %d", len(joinList.ActivityList))
	t.Logf("first el is: %d %s", joinList.ActivityList[0].ActivityId, joinList.ActivityList[0].ActivityName)
}

func TestGetCreatedActivityByUserID(t *testing.T)  {
	list := &activity_pb.ActivityBriefList{}
	err := cli.GetCreatedActivityByUserID(context.Background(), &activity_pb.IntWrap{Val: int32(1)}, list)
	if err != nil {
		t.Fatal(err)
	}

	if len(list.ActivityBriefList) < 1 {
		t.Fatal("result size < 1")
	}

	t.Logf("list size: %d", len(list.ActivityBriefList))
	t.Logf("first el is: %d %s", list.ActivityBriefList[0].ActivityId, list.ActivityBriefList[0].ActivityName)
}