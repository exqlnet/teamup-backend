package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	activity_pb "teamup/activity/proto"
	"teamup/api/svc"
	"teamup/api/util"
)

// swagger:operation POST /user/login user login
// ---
// summary: 创建活动（需要认证）
// parameters:
// - name:
//   in: path
//   description: 微信授权码
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/userLoginResp"
//   "400":
//     "$ref": "#/responses/badReq"
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

	val, e := c.Get("userId"); if !e {
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