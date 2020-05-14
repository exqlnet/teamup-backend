package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"strconv"
	"strings"
	activity_pb "teamup/activity/proto"
	"teamup/api/svc"
	"teamup/api/util"
)


func ActivityCreateHandler(c *gin.Context)  {
	params := &struct {
		ActivityName string `json:"activity_name";binding:"required"`
		Introduction string `binding:"required"`
		Roles		[]string `binding:"required,dive,required"`
		Illustration string `binding:"required"`
		AuthorityCode int32 `json:"authority_code";binding:"required"`
		StartTime int64	`json:"start_time";binding:"required"`
		EndTime int64 `json:"end_time";binding:"required"`
		Processes []struct{
			ProcessName string `json:"process_name";binding:"required"`
			StartTime int64	`json:"start_time";binding:"required"`
			Tasks	[]struct{
				TaskName string	`json:"task_name";binding:"required"`
				Role string	`binding:"required";json:"role"`
			}
		}	`binding:"required"`
	}{}

	err := c.ShouldBind(params); if err != nil {
		log.Printf("%v", err)
		util.BadRequest(c)
		return
	}

	val, e := c.Get("userID"); if !e {
		c.AbortWithStatus(500)
		return
	}

	userID := val.(int32)
	activityID, err := svc.ActivitySvc.CreateActivity(context.Background(), &activity_pb.CreateActivityReq{
		ActivityName:         params.ActivityName,
		Introduction:         params.Introduction,
		Roles: strings.Join(params.Roles, ","),
		CreatorId:            userID,
		Illustration:         params.Illustration,
		AuthorityCode:        params.AuthorityCode,
		Processes: func()[]*activity_pb.Process {
			var processes []*activity_pb.Process
			for _, pro := range params.Processes {
				processes = append(processes, &activity_pb.Process{
					ProcessName:          pro.ProcessName,
					StartTime:            pro.StartTime,
					Tasks: func() []*activity_pb.Task {
						var tasks []*activity_pb.Task
						for _, task := range pro.Tasks {
							tasks = append(tasks, &activity_pb.Task{
								TaskName:             task.TaskName,
								Role:                 task.Role,
							})
						}
						return tasks
					}(),
				})
			}
			return processes
		}(),
		StartTime:            params.StartTime,
		EndTime:              params.EndTime,
	})

	if err != nil {
		log.Printf("%v", err)
		util.InternalError(c)
		return
	}

	act, err := svc.ActivitySvc.GetActivityByID(context.Background(), activityID); if err != nil {
		log.Printf("%v", err)
		util.InternalError(c)
		return
	}

	util.Success(c, act)
}

func GetJoinedActivityHandler(c *gin.Context)  {
	var userID int32
	val, e := c.Get("userID")
	if !e {
		util.Unauthorize(c)
		return
	}
	userID = val.(int32)
	actList := &activity_pb.ActivityJoinList{}
	log.Println(userID)
	actList, err := svc.ActivitySvc.GetActivityJoinByUserID(context.Background(), &activity_pb.IntWrap{Val: userID})
	if err != nil {
		log.Printf("%v", err)
		util.InternalError(c)
		return
	}

	if actList.Total == 0 {
		actList.ActivityList = []*activity_pb.ActivityJoin{}
	}
	util.Success(c, actList.ActivityList)
}

func GetCreatedActivityHandler(c *gin.Context)  {
	val, e := c.Get("userID")
	log.Println(val)
	if !e {
		util.Unauthorize(c)
		return
	}
	userID := val.(int32)

	actList, err := svc.ActivitySvc.GetCreatedActivityByUserID(context.Background(), &activity_pb.IntWrap{Val: userID})
	if err != nil {
		util.Failed(c, 1, "获取失败")
		return
	}

	if actList.ActivityBriefList == nil {
		actList.ActivityBriefList = []*activity_pb.ActivityBrief{}
	}
	util.SuccessWithTotal(c, actList.ActivityBriefList, len(actList.ActivityBriefList))
}

func DeleteActivityHandler(c *gin.Context) {
	params := &struct {
		ActivityID int32 `json:"activity_id";binding:"required"`
	}{}

	var err error
	err = c.ShouldBind(params); if err != nil {
		log.Printf("%v", err)
		util.BadRequest(c)
		return
	}

	act, err := svc.ActivitySvc.GetActivityByID(context.Background(), &activity_pb.IntWrap{Val: params.ActivityID})
	if err != nil {
		log.Printf("%v", err)
		util.Failed(c, 404, "任务不存在")
		return
	}

	val, e := c.Get("userID")
	if !e {
		util.Unauthorize(c)
		return
	}
	userID := val.(int32)
	if act.CreatorId != userID {
		util.Failed(c, 403, "任务不属于你")
		return
	}

	_, err = svc.ActivitySvc.DeleteActivity(c, &activity_pb.IntWrap{Val: int32(params.ActivityID)})
	if err != nil {
		util.Failed(c, 500, err.Error())
		return
	}

	util.Success(c, nil)
}

func GetActivityDetailByIDHandler(c *gin.Context)  {
	actID, err := strconv.Atoi(c.Request.URL.Query().Get("activity_id"))
	if err != nil || actID < 1 {
		util.BadRequest(c)
		return
	}

	act, err := svc.ActivitySvc.GetActivityByID(context.Background(), &activity_pb.IntWrap{Val: int32(actID)})
	if err != nil {
		log.Printf("%v", act)
		util.Failed(c, 404, "activity not found")
		return
	}

	util.Success(c, act)
}

func UpdateActivityHandler(c *gin.Context)  {
	params := &struct {
		ActivityID int `json:"activity_id";binding:"required"`
		Introduction string `json:"introduction";binding:"required"`
	}{}
	err := c.ShouldBind(params)
	if err != nil {
		util.BadRequest(c)
		return
	}

	val, e := c.Get("userID")
	if !e {
		util.Unauthorize(c)
		return
	}

	userID := val.(int32)

	act, err := svc.ActivitySvc.GetActivityByID(context.Background(), &activity_pb.IntWrap{Val: int32(params.ActivityID)})
	if err != nil {
		util.Failed(c, 404, "activity not found")
		return
	}

	// 是否是该用户创建的活动
	if act.CreatorId != userID {
		util.Failed(c, 403, "activity not belongs to you")
		return
	}

	_, err = svc.ActivitySvc.UpdateActivity(context.Background(), &activity_pb.UpdateActivityReq{
		ActivityId:           int32(params.ActivityID),
		Introduction:         params.Introduction,
		Illustration:         "",
	})

	if err != nil {
		util.Failed(c, 1, "更新失败")
		return
	}

	util.Success(c, nil)
}

func CreateTeamHandler(c *gin.Context)  {
	params := &struct {
		ActivityID int32 `json:"activity_id";binding:"required"`
		TeamName string `json:"team_name";binding:"required"`
		Slogan string `json:"slogan"`
	}{}
	err := c.ShouldBind(params)
	if err != nil {
		util.BadRequest(c)
		return
	}

	_, err = svc.ActivitySvc.CreateTeam(context.Background(), &activity_pb.CreateTeamReq{
		ActivityId:           params.ActivityID,
		TeamName:             params.TeamName,
		Slogan:               params.Slogan,
	})

	if err != nil {
		util.Failed(c, 1, "创建失败")
		return
	}

	util.Success(c, nil)
}

func GetTeamListByActivityIDHandler(c *gin.Context)  {

	actID, err := strconv.Atoi(c.Query("activity_id"))
	if err != nil {
		util.BadRequest(c)
		return
	}

	teamList, err := svc.ActivitySvc.GetTeamListByActivityID(context.Background(), &activity_pb.IntWrap{Val: int32(actID)})

	if err != nil {
		log.Printf("%v", err)
		util.Failed(c, 1, "获取队伍列表失败")
		return
	}

	if teamList.Total == 0 {
		teamList.TeamList = []*activity_pb.Team{}
	}
	util.SuccessWithTotal(c, teamList.TeamList, int(teamList.Total))
}

func GetHotActivityListHandler(c *gin.Context) {
	activityList, err := svc.ActivitySvc.GetHotActivities(context.Background(), &empty.Empty{})
	if err != nil {
		util.Failed(c, 1, "获取活动列表失败")
		return
	}

	if activityList.ActivityList == nil {
		activityList.ActivityList = []*activity_pb.Activity{}
	}
	util.SuccessWithTotal(c, activityList, len(activityList.ActivityList))
}

func GetTeamByIDHandler(c *gin.Context) {

	teamID, err := strconv.Atoi(c.Query("team_id"))
	if err != nil {
		util.BadRequest(c)
		return
	}

	team, err := svc.ActivitySvc.GetTeamByID(context.Background(), &activity_pb.IntWrap{Val: int32(teamID)})
	if err != nil {
		util.Failed(c, 404, "未找到队伍")
		return
	}

	util.Success(c, team)
}
