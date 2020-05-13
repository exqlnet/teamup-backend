package svc

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	"log"
	"sort"
	activity_pb "teamup/activity/proto"
	"teamup/activity/status"
	"teamup/db"
	"teamup/db/model"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type ActivityServiceImpl struct {}

func (as *ActivityServiceImpl) CreateActivity(ctx context.Context, req *activity_pb.CreateActivityReq, rsp *activity_pb.IntWrap) error {
	tx := db.Conn.Begin()
	// check request parameters validation
	// todo
	activity := &model.Activity{
		ActivityName:   req.ActivityName,
		Introduction:   req.Introduction,
		CreatorID:      int(req.CreatorId),
		Regulation:     "-",
		Roles:          req.Roles,
		AuthorityCode:  int(req.AuthorityCode),
		Illustration:   req.Illustration,
		StatusCode:     1,
		CurrentProcessID: sql.NullInt64{},
		StartTime:		time.Unix(req.StartTime, 0),
		EndTime:		time.Unix(req.EndTime, 0),
	}

	if len(req.Processes) == 0 {
		return status.NewError(status.EmptyProcess, "流程为空")
	}

	var processes []model.ActivityProcess
	for _, p := range req.Processes {
		var t time.Time
		t = time.Unix(p.StartTime, 0)
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
	*rsp = activity_pb.IntWrap{Val: int32(activity.ActivityID)}
	return nil
}

func (as *ActivityServiceImpl) DeleteActivity(ctx context.Context, req *activity_pb.IntWrap, _ empty.Empty) error {
	tx := db.Conn.Begin()
	act := &model.Activity{}
	if err := tx.Table("activity").Where("activity_id = ?", req.Val).First(act).Error; gorm.IsRecordNotFoundError(err) {
		log.Println(err)
		return status.NewError(status.ActivityNotFound,"activity not found")
	}
	act.CurrentProcessID = sql.NullInt64{}
	tx.Save(act)
	tx.Delete(act)
	tx.Commit()
	return nil
}

func (as *ActivityServiceImpl) UpdateActivity(ctx context.Context, req *activity_pb.UpdateActivityReq, _ *empty.Empty) error {
	tx := db.Conn.Begin()
	act := &model.Activity{}
	if err := tx.Table("activity").Where("activity_id = ?", req.ActivityId).First(act).Error; gorm.IsRecordNotFoundError(err) {
		return status.NewError(status.ActivityNotFound, "activity not found")
	}

	act.Introduction = req.Introduction
	act.Illustration = req.Illustration
	db.Conn.Save(act)

	return nil
}

func (as *ActivityServiceImpl) CreateTeam(ctx context.Context, req *activity_pb.CreateTeamReq, rsp *activity_pb.IntWrap) error {
	act := &model.Activity{}
	if err := db.Conn.Table("activity").Where("activity_id = ?", req.ActivityId).Preload("Teams").First(act).Error; gorm.IsRecordNotFoundError(err) {
		return status.NewError(status.ActivityNotFound, "activity not found")
	}

	team := &model.ActivityTeam{
		ActivityID: int(req.ActivityId),
		TeamName:   req.TeamName,
		Slogan:     req.Slogan,
	}

	act.Teams = append(act.Teams, team)
	err := db.Conn.Save(act).Error
	if err != nil {
		log.Println(err)
		return status.NewError(status.InternalServerError, "create failed")
	}

	rsp.Val = int32(team.TeamID)
	return nil
}

func (as *ActivityServiceImpl) GetTeamListByActivityID(ctx context.Context, req *activity_pb.IntWrap, rsp *activity_pb.TeamList) error {
	act := &model.Activity{}
	if err := db.Conn.Table("activity").Where("activity_id = ?", req.Val).Preload("Teams").First(act).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("activity not found")
	}

	teamList := activity_pb.TeamList{
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
			return teams
		}(),
	}
	*rsp = teamList
	return nil
}

func (as *ActivityServiceImpl) DeleteTeam(ctx context.Context, req *activity_pb.IntWrap, _ *empty.Empty) error {
	team := &model.ActivityTeam{}
	if err := db.Conn.Table("activity_team").Where("team_id = ?", req.Val).First(team).Error; gorm.IsRecordNotFoundError(err) {
		return status.NewError(status.TeamNotFound, "team not found")
	}

	db.Conn.Delete(team)

	return nil
}

func (as *ActivityServiceImpl) UpdateTeam(ctx context.Context, req *activity_pb.UpdateTeamReq, _ *empty.Empty) error {
	team := &model.ActivityTeam{}
	if err := db.Conn.Table("activity_team").Where("team_id = ?", req.TeamId).First(team).Error; gorm.IsRecordNotFoundError(err) {
		return status.NewError(status.TeamNotFound, "team not found")
	}
	if team.TeamName == req.TeamName {
		return nil
	}

	// 检查有无队伍重名
	var check int
	if err := db.Conn.Table("activity_team").Where("team_name = ?", req.TeamName).Count(&check).Error; err != nil {
		log.Printf("%v", err)
		return errors.New("error while checking repeated team name")
	}

	if check > 0 {
		return status.NewError(status.RepeatedTeamName, "repeated team name")
	}

	team.Slogan = req.Slogan
	team.TeamName = req.TeamName
	db.Conn.Save(team)

	return nil
}

func (as *ActivityServiceImpl) GetTeamByID(ctx context.Context, req *activity_pb.IntWrap, rsp *activity_pb.Team) error {
	team := &model.ActivityTeam{}
	if err := db.Conn.Table("activity_team").Where("team_id = ?", req.Val).First(team).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("team not found")
	}

	*rsp = activity_pb.Team{
		TeamId:               int32(team.TeamID),
		ActivityId:           int32(team.ActivityID),
		TeamName:             team.TeamName,
		Slogan:               team.Slogan,
	}

	return nil
}

func (as *ActivityServiceImpl) GetActivityJoinByUserID(ctx context.Context, req *activity_pb.IntWrap, rsp *activity_pb.ActivityJoinList) error {
	user := &model.User{}
	if err := db.Conn.Table("user").Where("user_id = ?", req.Val).Preload("JoinedActivities").First(user).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("user not found")
	}

	*rsp = activity_pb.ActivityJoinList{
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
	}
	return nil
}

func (as *ActivityServiceImpl) GetCreatedActivityByUserID(ctx context.Context, req *activity_pb.IntWrap, rsp *activity_pb.ActivityBriefList) error {
	user := &model.User{}
	if err := db.Conn.Table("user").Where("user_id = ?", req.Val).Preload("CreatedActivities").First(user).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("user not found")
	}
	*rsp = activity_pb.ActivityBriefList{
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
	}
	return nil
}

func (as *ActivityServiceImpl) GetHotActivities(ctx context.Context, req *empty.Empty, rsp *activity_pb.ActivityList) error {
	panic("implement me")
}

func (as *ActivityServiceImpl) GetActivityByID(ctx context.Context, req *activity_pb.IntWrap, rsp *activity_pb.Activity) error {

	activity := &model.Activity{}
	if err := db.Conn.Table("activity").Where("activity_id = ?", req.Val).First(activity).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("activity not found")
	}

	db.Conn.Table("activity_task").Where("process_id = ?")
	var processesPb []*activity_pb.Process
	for _, process := range activity.Processes {
		processesPb = append(processesPb, &activity_pb.Process{
			ProcessId:            int32(process.ProcessID),
			ProcessName:          process.ProcessName,
			StartTime:            process.StartTime.Unix(),
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

	*rsp = activity_pb.Activity{
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

	return nil
}
