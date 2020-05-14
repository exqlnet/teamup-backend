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
		Activity_name string `binding:"required";json:"activity_name"`
		Introduction string `binding:"required"`
		Roles		[]string `binding:"required,dive,required"`
		Illustration string `binding:"required"`
		Authority_code int32 `binding:"required";json:"authority_code"`
		Start_time int64	`binding:"required";json:"start_time"`
		End_time int64 `binding:"required";json:"end_time"`
		Processes []struct{
			Process_name string `binding:"required";json:"process_name"`
			Start_time int64	`binding:"required";json:"start_time"`
			Tasks	[]struct{
				Task_name string	`binding:"required";json:"task_name"`
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
		ActivityName:         params.Activity_name,
		Introduction:         params.Introduction,
		Roles: strings.Join(params.Roles, ","),
		CreatorId:            userID,
		Illustration:         params.Illustration,
		AuthorityCode:        params.Authority_code,
		Processes: func()[]*activity_pb.Process {
			var processes []*activity_pb.Process
			for _, pro := range params.Processes {
				processes = append(processes, &activity_pb.Process{
					ProcessName:          pro.Process_name,
					StartTime:            pro.Start_time,
					Tasks: func() []*activity_pb.Task {
						var tasks []*activity_pb.Task
						for _, task := range pro.Tasks {
							tasks = append(tasks, &activity_pb.Task{
								TaskName:             task.Task_name,
								Role:                 task.Role,
							})
						}
						return tasks
					}(),
				})
			}
			return processes
		}(),
		StartTime:            params.Start_time,
		EndTime:              params.End_time,
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