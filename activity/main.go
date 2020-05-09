package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"log"
	"sort"
	activity_pb "teamup/activity/proto"
	"teamup/activity/status"
	"teamup/db"
	"teamup/db/model"
	"time"
)

const SvcName  = "go.micro.teamup.svc.activity"
const TimeFormat = "2006-01-02 15:04:05"

type ActivityService struct {}

func (as *ActivityService) DeleteActivity(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.CommonResp, error) {
	tx := db.Conn.Begin()
	act := &model.Activity{}
	if err := tx.Table("activity").Where("activity_id = ?", in.Val).First(act).Error; gorm.IsRecordNotFoundError(err) {
		log.Println(err)
		return &activity_pb.CommonResp{}, errors.New("activity not found")
	}
	act.CurrentProcessID = sql.NullInt64{}
	tx.Save(act)
	tx.Delete(act)
	return &activity_pb.CommonResp{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (as *ActivityService) UpdateActivity(ctx context.Context, in *activity_pb.UpdateActivityReq, opts ...client.CallOption) (*activity_pb.CommonResp, error) {
	tx := db.Conn.Begin()
	act := &model.Activity{}
	if err := tx.Table("activity").Where("activity_id = ?", in.ActivityId).First(act).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.CommonResp{}, errors.New("activity not found")
	}

	act.Introduction = in.Introduction
	act.Illustration = in.Illustration

	return &activity_pb.CommonResp{
		Code:                 0,
		Msg:                  "success",
	}, nil
}

func (as *ActivityService) CreateTeam(ctx context.Context, in *activity_pb.CreateTeamReq, opts ...client.CallOption) (*activity_pb.CommonResp, error) {
	act := &model.Activity{}
	if err := db.Conn.Table("activity").Where("activity_id = ?", in.ActivityId).First(act).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.CommonResp{}, errors.New("activity not found")
	}

	team := model.Team{
		ActivityID: int(in.ActivityId),
		TeamName:   in.TeamName,
		Slogan:     in.Slogan,
	}

	act.Teams = append(act.Teams, team)
	err := db.Conn.Save(act).Error
	if err != nil {
		log.Println(err)
		return &activity_pb.CommonResp{}, errors.New("create failed")
	}

	return &activity_pb.CommonResp{
		Code:                 0,
		Msg:                  "success",
	}, nil



}

func (as *ActivityService) GetTeamListByActivityID(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.TeamList, error) {
	act := &model.Activity{}
	if err := db.Conn.Table("activity").Where("activity_id = ?", in.Val).First(act).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.TeamList{}, errors.New("activity not found")
	}
	teamList := &activity_pb.TeamList{
		Total: int32(len(act.Teams)),
		TeamList: func()[]*activity_pb.Team {
			var teams []*activity_pb.Team
			for _, team := range act.Teams {
				teams = append(teams, &activity_pb.Team{
					TeamId:               int32(team.TeamID),
					ActivityId:           int32(team.ActivityID),
					TeamName:             team.TeamName,
					Slogan:               team.Slogan,
				})
			}
			return nil
		}(),
	}
	return teamList, nil
}

func (as *ActivityService) DeleteTeam(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.CommonResp, error) {
	team := &model.Team{}
	if err := db.Conn.Table("activity_team").Where("team_id = ?", in.Val).First(team).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.CommonResp{}, errors.New("team not found")
	}

	db.Conn.Delete(team)
	return &activity_pb.CommonResp{
		Code:                 0,
		Msg:                  "success",
	}, nil
}

func (as *ActivityService) UpdateTeam(ctx context.Context, in *activity_pb.UpdateTeamReq, opts ...client.CallOption) (*activity_pb.CommonResp, error) {
	team := &model.Team{}
	if err := db.Conn.Table("activity_team").Where("team_id = ?", in.TeamId).First(team).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.CommonResp{}, errors.New("team not found")
	}

	// 检查有无队伍重名
	var check int
	if err := db.Conn.Table("activity_team").Where("team_id").Count(&check); err != nil {
		log.Println(err)
		return &activity_pb.CommonResp{}, errors.New("error while checking repeated team name")
	}
	if check > 0 {
		return &activity_pb.CommonResp{}, errors.New("repeated team name")
	}

	team.Slogan = in.Slogan
	team.TeamName = in.TeamName

	return &activity_pb.CommonResp{
		Code:                 0,
		Msg:                  "success",
	}, nil
}

func (as *ActivityService) GetTeamByID(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.Team, error) {
	team := &model.Team{}
	if err := db.Conn.Table("activity_team").Where("team_id = ?", in.TeamId).First(team).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.Team{}, errors.New("team not found")
	}

	return &activity_pb.Team{
		TeamId:               int32(team.TeamID),
		ActivityId:           int32(team.ActivityID),
		TeamName:             team.TeamName,
		Slogan:               team.Slogan,
	}, nil
}

func (as *ActivityService) GetActivityJoinByUserID(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.ActivityJoinList, error) {
	user := &model.User{}
	if err := db.Conn.Table("user").Where("user_id = ?", in.Val).First(user).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.ActivityJoinList{}, errors.New("user not found")
	}

	return &activity_pb.ActivityJoinList{
		Total:                0,
		ActivityList: func() []*activity_pb.ActivityJoin{
			var list []*activity_pb.ActivityJoin
			for _, act := range user.JoinedActivities {
				list = append(list, &activity_pb.ActivityJoin{
					ActivityId:           int32(act.ActivityID),
					ActivityName:         act.ActivityName,
					Introduction:         act.Introduction,
					Status:               act.Status.Name,
				})
			}
			return list
		}(),
	}, nil
}

func (as *ActivityService) GetCreatedActivityByUserID(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.ActivityBriefList, error) {
	user := &model.User{}
	if err := db.Conn.Table("user").Where("user_id = ?", in.Val).First(user).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.ActivityBriefList{}, errors.New("user not found")
	}
	return &activity_pb.ActivityBriefList{
		ActivityBriefList: func() []*activity_pb.ActivityBrief{
			var list []*activity_pb.ActivityBrief
			for _, act := range user.CreatedActivities {
				list = append(list, &activity_pb.ActivityBrief{
					ActivityId:           int32(act.ActivityID),
					ActivityName:         act.ActivityName,
					StartTime:            act.StartTime.Format(TimeFormat),
					EndTime:              act.EndTime.Format(TimeFormat),
					Status:               act.Status.Name,
				})
			}
			return list
		}(),
	}, nil
}

func (as *ActivityService) GetHotActivities(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*activity_pb.ActivityList, error) {
	panic("implement me")
}

func (as *ActivityService) CreateActivity(ctx context.Context, in *activity_pb.CreateActivityReq, opts ...client.CallOption) (*activity_pb.CommonResp, error) {
	tx := db.Conn.Begin()
	// check request parameters validation
	// todo
	activity := &model.Activity{
		ActivityName:   in.ActivityName,
		Introduction:   in.Introduction,
		CreatorID:      int(in.CreatorId),
		Regulation:     "-",
		Roles:          in.Roles,
		AuthorityCode:  int(in.AuthorityCode),
		Illustration:   in.Illustration,
		StatusCode:     1,
		CurrentProcessID: sql.NullInt64{},
	}

	if len(in.Processes) == 0 {
		return &activity_pb.CommonResp{
			Code:                 status.EmptyProcess,
			Msg:                  "process is empty",
		}, nil
	}

	var processes []model.ActivityProcess
	for _, p := range in.Processes {
		var t time.Time
		var err error
		if t, err = time.Parse("", p.StartTime); err != nil {
			return &activity_pb.CommonResp{
				Code:                 status.TimeFormatError,
				Msg:                  "time format error",
			}, nil
		}
		process := model.ActivityProcess{
			ProcessName: p.ProcessName,
			StartTime: t,
		}

		var tasks []model.ActivityTask
		for _, t := range p.Tasks {
			task := model.ActivityTask{
				TaskName:  t.TaskName,
				ProcessID: process.ProcessID,
				Role:      t.Role,
			}
			tasks = append(tasks, task)
		}
		process.Tasks = tasks
		processes = append(processes, process)
	}

	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].StartTime.After(processes[j].StartTime)
	})
	activity.Processes = processes
	tx.Create(activity)
	activity.CurrentProcess = processes[0]
	tx.Save(activity)
	tx.Commit()
	return &activity_pb.CommonResp{
		Code:                 0,
		Msg:                  "success",
	}, nil
}

func (as *ActivityService) GetActivityByID(ctx context.Context, in *activity_pb.IntWrap, opts ...client.CallOption) (*activity_pb.Activity, error) {

	activity := &model.Activity{}
	if err := db.Conn.Table("activity").Where("activity_id = ?", in.Val).First(activity).Error; gorm.IsRecordNotFoundError(err) {
		return &activity_pb.Activity{}, errors.New("activity not found")
	}

	db.Conn.Table("activity_task").Where("process_id = ?")
	var processesPb []*activity_pb.Process
	for _, process := range activity.Processes {
		processesPb = append(processesPb, &activity_pb.Process{
			ProcessId:            int32(process.ProcessID),
			ProcessName:          process.ProcessName,
			StartTime:            process.StartTime.Format(TimeFormat),
			Tasks: func() []*activity_pb.Task {
				var tasksPb []*activity_pb.Task
				for _, task := range process.Tasks {
					tasksPb = append(tasksPb, &activity_pb.Task{
						TaskId:               int32(task.TaskID),
						TaskName:             task.TaskName,
						Role:                 task.Role,
					})
				}
				return tasksPb
			}(),
		})
	}

	result := &activity_pb.Activity{
		ActivityId:           int32(activity.ActivityID),
		ActivityName:         activity.ActivityName,
		Introduction:         activity.Introduction,
		CreatorId:            int32(activity.CreatorID),
		Regulation:           activity.Regulation,
		Roles:                activity.Roles,
		Authority:            activity.Authority.Name,
		Illustration:         activity.Illustration,
		CurrentProcess:       activity.CurrentProcess.ProcessName,
		Status:               activity.Status.Name,
		Processes:            processesPb,
	}

	return result, nil
}



func main() {
	service := micro.NewService(
		micro.Name(SvcName),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}