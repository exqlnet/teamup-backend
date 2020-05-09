package test

import (
	"database/sql"
	"log"
	"sort"
	activity_pb "teamup/activity/proto"
	"teamup/db"
	"teamup/db/model"
	"testing"
	"time"
)

func TestCreateActivity(tst *testing.T)  {
	in := &activity_pb.CreateActivityReq{
		ActivityName:         "test ac",
		Introduction:         "ac ac ac",
		Roles:                "dev",
		CreatorId:            1,
		Illustration:         "-",
		AuthorityCode:			1,
		Processes:            []*activity_pb.Process{{
			ProcessName: "test pro",
			Tasks: []*activity_pb.Task{{
				TaskName: "tsk 1",
				Role:     "dev",
			}},
			StartTime: time.Now().Format("2006-01-02 15:04:05"),
		}, {
			ProcessName: "test pro",
			Tasks: []*activity_pb.Task{{
				TaskName: "tsk 2",
				Role:     "dev",
			}},
			StartTime: time.Now().Format("2006-01-02 15:03:05"),
		}},
	}
	tx := db.Conn.Begin()
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


	var processes []model.ActivityProcess
	for _, p := range in.Processes {
		var t time.Time
		var err error
		if t, err = time.Parse("2006-01-02 15:04:05", p.StartTime); err != nil {
			log.Fatal("time parse err", p.StartTime)
			return
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
	tx.Debug().Create(activity)
	//activity.CurrentProcess = processes[0]
	tst.Log("pro id: ", processes[0].ProcessID)
	//tx.Save(activity)
	tx.Commit()
}

func TestTimeNow(t *testing.T)  {
	t.Log(time.Now().Format("2006-01-02 15:04:05"))
}