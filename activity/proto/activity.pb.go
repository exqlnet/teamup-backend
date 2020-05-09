// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activity.proto

package activity_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Activity struct {
	ActivityId           int32      `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityName         string     `protobuf:"bytes,2,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	Introduction         string     `protobuf:"bytes,3,opt,name=introduction,proto3" json:"introduction,omitempty"`
	CreatorId            int32      `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Regulation           string     `protobuf:"bytes,5,opt,name=regulation,proto3" json:"regulation,omitempty"`
	Roles                string     `protobuf:"bytes,6,opt,name=roles,proto3" json:"roles,omitempty"`
	Authority            string     `protobuf:"bytes,7,opt,name=authority,proto3" json:"authority,omitempty"`
	Illustration         string     `protobuf:"bytes,8,opt,name=illustration,proto3" json:"illustration,omitempty"`
	CurrentProcess       string     `protobuf:"bytes,9,opt,name=current_process,json=currentProcess,proto3" json:"current_process,omitempty"`
	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Processes            []*Process `protobuf:"bytes,11,rep,name=processes,proto3" json:"processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{0}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetActivityId() int32 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *Activity) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *Activity) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *Activity) GetCreatorId() int32 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

func (m *Activity) GetRegulation() string {
	if m != nil {
		return m.Regulation
	}
	return ""
}

func (m *Activity) GetRoles() string {
	if m != nil {
		return m.Roles
	}
	return ""
}

func (m *Activity) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *Activity) GetIllustration() string {
	if m != nil {
		return m.Illustration
	}
	return ""
}

func (m *Activity) GetCurrentProcess() string {
	if m != nil {
		return m.CurrentProcess
	}
	return ""
}

func (m *Activity) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Activity) GetProcesses() []*Process {
	if m != nil {
		return m.Processes
	}
	return nil
}

type ActivityList struct {
	ActivityList         []*Activity `protobuf:"bytes,1,rep,name=activity_list,json=activityList,proto3" json:"activity_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ActivityList) Reset()         { *m = ActivityList{} }
func (m *ActivityList) String() string { return proto.CompactTextString(m) }
func (*ActivityList) ProtoMessage()    {}
func (*ActivityList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{1}
}

func (m *ActivityList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityList.Unmarshal(m, b)
}
func (m *ActivityList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityList.Marshal(b, m, deterministic)
}
func (m *ActivityList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityList.Merge(m, src)
}
func (m *ActivityList) XXX_Size() int {
	return xxx_messageInfo_ActivityList.Size(m)
}
func (m *ActivityList) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityList.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityList proto.InternalMessageInfo

func (m *ActivityList) GetActivityList() []*Activity {
	if m != nil {
		return m.ActivityList
	}
	return nil
}

type CreateActivityReq struct {
	ActivityName         string     `protobuf:"bytes,1,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	Introduction         string     `protobuf:"bytes,2,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Roles                string     `protobuf:"bytes,3,opt,name=roles,proto3" json:"roles,omitempty"`
	CreatorId            int32      `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Illustration         string     `protobuf:"bytes,5,opt,name=illustration,proto3" json:"illustration,omitempty"`
	AuthorityCode        int32      `protobuf:"varint,6,opt,name=authority_code,json=authorityCode,proto3" json:"authority_code,omitempty"`
	Processes            []*Process `protobuf:"bytes,7,rep,name=processes,proto3" json:"processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateActivityReq) Reset()         { *m = CreateActivityReq{} }
func (m *CreateActivityReq) String() string { return proto.CompactTextString(m) }
func (*CreateActivityReq) ProtoMessage()    {}
func (*CreateActivityReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{2}
}

func (m *CreateActivityReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActivityReq.Unmarshal(m, b)
}
func (m *CreateActivityReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActivityReq.Marshal(b, m, deterministic)
}
func (m *CreateActivityReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActivityReq.Merge(m, src)
}
func (m *CreateActivityReq) XXX_Size() int {
	return xxx_messageInfo_CreateActivityReq.Size(m)
}
func (m *CreateActivityReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActivityReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActivityReq proto.InternalMessageInfo

func (m *CreateActivityReq) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *CreateActivityReq) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *CreateActivityReq) GetRoles() string {
	if m != nil {
		return m.Roles
	}
	return ""
}

func (m *CreateActivityReq) GetCreatorId() int32 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

func (m *CreateActivityReq) GetIllustration() string {
	if m != nil {
		return m.Illustration
	}
	return ""
}

func (m *CreateActivityReq) GetAuthorityCode() int32 {
	if m != nil {
		return m.AuthorityCode
	}
	return 0
}

func (m *CreateActivityReq) GetProcesses() []*Process {
	if m != nil {
		return m.Processes
	}
	return nil
}

type Process struct {
	ProcessId            int32    `protobuf:"varint,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProcessName          string   `protobuf:"bytes,2,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	StartTime            string   `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Tasks                []*Task  `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{3}
}

func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

func (m *Process) GetProcessId() int32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *Process) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Process) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Process) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Task struct {
	TaskId               int32    `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskName             string   `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{4}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *Task) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *Task) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type GetProcessStatusReq struct {
	TeamId               int32    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	ProcessId            int32    `protobuf:"varint,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProcessStatusReq) Reset()         { *m = GetProcessStatusReq{} }
func (m *GetProcessStatusReq) String() string { return proto.CompactTextString(m) }
func (*GetProcessStatusReq) ProtoMessage()    {}
func (*GetProcessStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{5}
}

func (m *GetProcessStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProcessStatusReq.Unmarshal(m, b)
}
func (m *GetProcessStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProcessStatusReq.Marshal(b, m, deterministic)
}
func (m *GetProcessStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProcessStatusReq.Merge(m, src)
}
func (m *GetProcessStatusReq) XXX_Size() int {
	return xxx_messageInfo_GetProcessStatusReq.Size(m)
}
func (m *GetProcessStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProcessStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProcessStatusReq proto.InternalMessageInfo

func (m *GetProcessStatusReq) GetTeamId() int32 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *GetProcessStatusReq) GetProcessId() int32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

type ActivityJoin struct {
	ActivityId           int32    `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityName         string   `protobuf:"bytes,2,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	Introduction         string   `protobuf:"bytes,3,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Status               string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityJoin) Reset()         { *m = ActivityJoin{} }
func (m *ActivityJoin) String() string { return proto.CompactTextString(m) }
func (*ActivityJoin) ProtoMessage()    {}
func (*ActivityJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{6}
}

func (m *ActivityJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityJoin.Unmarshal(m, b)
}
func (m *ActivityJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityJoin.Marshal(b, m, deterministic)
}
func (m *ActivityJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityJoin.Merge(m, src)
}
func (m *ActivityJoin) XXX_Size() int {
	return xxx_messageInfo_ActivityJoin.Size(m)
}
func (m *ActivityJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityJoin.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityJoin proto.InternalMessageInfo

func (m *ActivityJoin) GetActivityId() int32 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *ActivityJoin) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *ActivityJoin) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *ActivityJoin) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ActivityBrief struct {
	ActivityId           int32    `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityName         string   `protobuf:"bytes,2,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	StartTime            string   `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityBrief) Reset()         { *m = ActivityBrief{} }
func (m *ActivityBrief) String() string { return proto.CompactTextString(m) }
func (*ActivityBrief) ProtoMessage()    {}
func (*ActivityBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{7}
}

func (m *ActivityBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityBrief.Unmarshal(m, b)
}
func (m *ActivityBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityBrief.Marshal(b, m, deterministic)
}
func (m *ActivityBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityBrief.Merge(m, src)
}
func (m *ActivityBrief) XXX_Size() int {
	return xxx_messageInfo_ActivityBrief.Size(m)
}
func (m *ActivityBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityBrief.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityBrief proto.InternalMessageInfo

func (m *ActivityBrief) GetActivityId() int32 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *ActivityBrief) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *ActivityBrief) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *ActivityBrief) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *ActivityBrief) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ActivityBriefList struct {
	ActivityBriefList    []*ActivityBrief `protobuf:"bytes,1,rep,name=activityBriefList,proto3" json:"activityBriefList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ActivityBriefList) Reset()         { *m = ActivityBriefList{} }
func (m *ActivityBriefList) String() string { return proto.CompactTextString(m) }
func (*ActivityBriefList) ProtoMessage()    {}
func (*ActivityBriefList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{8}
}

func (m *ActivityBriefList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityBriefList.Unmarshal(m, b)
}
func (m *ActivityBriefList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityBriefList.Marshal(b, m, deterministic)
}
func (m *ActivityBriefList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityBriefList.Merge(m, src)
}
func (m *ActivityBriefList) XXX_Size() int {
	return xxx_messageInfo_ActivityBriefList.Size(m)
}
func (m *ActivityBriefList) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityBriefList.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityBriefList proto.InternalMessageInfo

func (m *ActivityBriefList) GetActivityBriefList() []*ActivityBrief {
	if m != nil {
		return m.ActivityBriefList
	}
	return nil
}

type ActivityJoinList struct {
	Total                int32           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	ActivityList         []*ActivityJoin `protobuf:"bytes,2,rep,name=activityList,proto3" json:"activityList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ActivityJoinList) Reset()         { *m = ActivityJoinList{} }
func (m *ActivityJoinList) String() string { return proto.CompactTextString(m) }
func (*ActivityJoinList) ProtoMessage()    {}
func (*ActivityJoinList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{9}
}

func (m *ActivityJoinList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityJoinList.Unmarshal(m, b)
}
func (m *ActivityJoinList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityJoinList.Marshal(b, m, deterministic)
}
func (m *ActivityJoinList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityJoinList.Merge(m, src)
}
func (m *ActivityJoinList) XXX_Size() int {
	return xxx_messageInfo_ActivityJoinList.Size(m)
}
func (m *ActivityJoinList) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityJoinList.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityJoinList proto.InternalMessageInfo

func (m *ActivityJoinList) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ActivityJoinList) GetActivityList() []*ActivityJoin {
	if m != nil {
		return m.ActivityList
	}
	return nil
}

type UpdateActivityReq struct {
	ActivityId           int32    `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	Introduction         string   `protobuf:"bytes,2,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Illustration         string   `protobuf:"bytes,3,opt,name=illustration,proto3" json:"illustration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateActivityReq) Reset()         { *m = UpdateActivityReq{} }
func (m *UpdateActivityReq) String() string { return proto.CompactTextString(m) }
func (*UpdateActivityReq) ProtoMessage()    {}
func (*UpdateActivityReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{10}
}

func (m *UpdateActivityReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActivityReq.Unmarshal(m, b)
}
func (m *UpdateActivityReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActivityReq.Marshal(b, m, deterministic)
}
func (m *UpdateActivityReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActivityReq.Merge(m, src)
}
func (m *UpdateActivityReq) XXX_Size() int {
	return xxx_messageInfo_UpdateActivityReq.Size(m)
}
func (m *UpdateActivityReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActivityReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActivityReq proto.InternalMessageInfo

func (m *UpdateActivityReq) GetActivityId() int32 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *UpdateActivityReq) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *UpdateActivityReq) GetIllustration() string {
	if m != nil {
		return m.Illustration
	}
	return ""
}

type CreateTeamReq struct {
	ActivityId           int32    `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	TeamName             string   `protobuf:"bytes,2,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Slogan               string   `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamReq) Reset()         { *m = CreateTeamReq{} }
func (m *CreateTeamReq) String() string { return proto.CompactTextString(m) }
func (*CreateTeamReq) ProtoMessage()    {}
func (*CreateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{11}
}

func (m *CreateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamReq.Unmarshal(m, b)
}
func (m *CreateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamReq.Marshal(b, m, deterministic)
}
func (m *CreateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamReq.Merge(m, src)
}
func (m *CreateTeamReq) XXX_Size() int {
	return xxx_messageInfo_CreateTeamReq.Size(m)
}
func (m *CreateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamReq proto.InternalMessageInfo

func (m *CreateTeamReq) GetActivityId() int32 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *CreateTeamReq) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *CreateTeamReq) GetSlogan() string {
	if m != nil {
		return m.Slogan
	}
	return ""
}

type UpdateTeamReq struct {
	TeamId               int32    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	TeamName             string   `protobuf:"bytes,2,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Slogan               string   `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamReq) Reset()         { *m = UpdateTeamReq{} }
func (m *UpdateTeamReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamReq) ProtoMessage()    {}
func (*UpdateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{12}
}

func (m *UpdateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamReq.Unmarshal(m, b)
}
func (m *UpdateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamReq.Marshal(b, m, deterministic)
}
func (m *UpdateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamReq.Merge(m, src)
}
func (m *UpdateTeamReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamReq.Size(m)
}
func (m *UpdateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamReq proto.InternalMessageInfo

func (m *UpdateTeamReq) GetTeamId() int32 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *UpdateTeamReq) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *UpdateTeamReq) GetSlogan() string {
	if m != nil {
		return m.Slogan
	}
	return ""
}

type Team struct {
	TeamId               int32    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	ActivityId           int32    `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	TeamName             string   `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Slogan               string   `protobuf:"bytes,4,opt,name=slogan,proto3" json:"slogan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{13}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetTeamId() int32 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *Team) GetActivityId() int32 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *Team) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *Team) GetSlogan() string {
	if m != nil {
		return m.Slogan
	}
	return ""
}

type TeamList struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	TeamList             []*Team  `protobuf:"bytes,2,rep,name=teamList,proto3" json:"teamList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamList) Reset()         { *m = TeamList{} }
func (m *TeamList) String() string { return proto.CompactTextString(m) }
func (*TeamList) ProtoMessage()    {}
func (*TeamList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{14}
}

func (m *TeamList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamList.Unmarshal(m, b)
}
func (m *TeamList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamList.Marshal(b, m, deterministic)
}
func (m *TeamList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamList.Merge(m, src)
}
func (m *TeamList) XXX_Size() int {
	return xxx_messageInfo_TeamList.Size(m)
}
func (m *TeamList) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamList.DiscardUnknown(m)
}

var xxx_messageInfo_TeamList proto.InternalMessageInfo

func (m *TeamList) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *TeamList) GetTeamList() []*Team {
	if m != nil {
		return m.TeamList
	}
	return nil
}

func init() {
	proto.RegisterType((*Activity)(nil), "activity_pb.Activity")
	proto.RegisterType((*ActivityList)(nil), "activity_pb.ActivityList")
	proto.RegisterType((*CreateActivityReq)(nil), "activity_pb.CreateActivityReq")
	proto.RegisterType((*Process)(nil), "activity_pb.Process")
	proto.RegisterType((*Task)(nil), "activity_pb.Task")
	proto.RegisterType((*GetProcessStatusReq)(nil), "activity_pb.GetProcessStatusReq")
	proto.RegisterType((*ActivityJoin)(nil), "activity_pb.ActivityJoin")
	proto.RegisterType((*ActivityBrief)(nil), "activity_pb.ActivityBrief")
	proto.RegisterType((*ActivityBriefList)(nil), "activity_pb.ActivityBriefList")
	proto.RegisterType((*ActivityJoinList)(nil), "activity_pb.ActivityJoinList")
	proto.RegisterType((*UpdateActivityReq)(nil), "activity_pb.UpdateActivityReq")
	proto.RegisterType((*CreateTeamReq)(nil), "activity_pb.CreateTeamReq")
	proto.RegisterType((*UpdateTeamReq)(nil), "activity_pb.UpdateTeamReq")
	proto.RegisterType((*Team)(nil), "activity_pb.Team")
	proto.RegisterType((*TeamList)(nil), "activity_pb.TeamList")
}

func init() { proto.RegisterFile("activity.proto", fileDescriptor_a684c9a0549e7832) }

var fileDescriptor_a684c9a0549e7832 = []byte{
	// 931 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdb, 0x8e, 0x1b, 0x45,
	0x10, 0x95, 0xef, 0x76, 0xf9, 0xb2, 0x71, 0xb3, 0x64, 0x27, 0x0e, 0xbb, 0x2c, 0x83, 0x50, 0xf6,
	0x05, 0x47, 0x5a, 0x24, 0x1e, 0x40, 0x80, 0x76, 0x93, 0xe0, 0x4c, 0x04, 0x21, 0x72, 0x36, 0xe2,
	0x29, 0x5a, 0xf5, 0x7a, 0x2a, 0x66, 0x94, 0x99, 0x69, 0x33, 0xd3, 0x8e, 0x64, 0x89, 0x07, 0x3e,
	0x21, 0x12, 0x1f, 0xc1, 0x1f, 0xf1, 0x3d, 0xa8, 0x6b, 0xba, 0x3d, 0x17, 0x7b, 0xbc, 0x06, 0x21,
	0xde, 0xdc, 0x55, 0xd5, 0xa7, 0x4e, 0x9f, 0xba, 0x78, 0x60, 0xc0, 0x67, 0xd2, 0x7b, 0xe7, 0xc9,
	0xd5, 0x78, 0x11, 0x09, 0x29, 0x58, 0xd7, 0x9c, 0xaf, 0x17, 0x37, 0xa3, 0xfb, 0x73, 0x21, 0xe6,
	0x3e, 0x3e, 0x24, 0xd7, 0xcd, 0xf2, 0xcd, 0x43, 0x0c, 0x16, 0x26, 0x72, 0xd4, 0x9b, 0x89, 0x20,
	0x10, 0x61, 0x72, 0xb2, 0x7f, 0xaf, 0x41, 0xfb, 0x42, 0x5f, 0x65, 0x1f, 0x43, 0x0a, 0xe3, 0xb9,
	0x56, 0xe5, 0xb4, 0x72, 0xd6, 0x98, 0x82, 0x31, 0x39, 0x2e, 0xfb, 0x14, 0xfa, 0xeb, 0x80, 0x90,
	0x07, 0x68, 0x55, 0x4f, 0x2b, 0x67, 0x9d, 0x69, 0xcf, 0x18, 0x9f, 0xf3, 0x00, 0x99, 0x0d, 0x3d,
	0x2f, 0x94, 0x91, 0x70, 0x97, 0x33, 0xe9, 0x89, 0xd0, 0xaa, 0x25, 0x31, 0x59, 0x1b, 0x3b, 0x06,
	0x98, 0x45, 0xc8, 0xa5, 0x88, 0x54, 0xa2, 0x3a, 0x25, 0xea, 0x68, 0x8b, 0xe3, 0xb2, 0x13, 0x80,
	0x08, 0xe7, 0x4b, 0x9f, 0x13, 0x40, 0x83, 0x00, 0x32, 0x16, 0x76, 0x08, 0x8d, 0x48, 0xf8, 0x18,
	0x5b, 0x4d, 0x72, 0x25, 0x07, 0xf6, 0x11, 0x74, 0xf8, 0x52, 0xfe, 0x22, 0x22, 0x4f, 0xae, 0xac,
	0x16, 0x79, 0x52, 0x03, 0xd1, 0xf2, 0xfd, 0x65, 0x2c, 0xa3, 0x04, 0xb5, 0xad, 0x69, 0x65, 0x6c,
	0xec, 0x01, 0x1c, 0xcc, 0x96, 0x51, 0x84, 0xa1, 0xbc, 0x5e, 0x44, 0x62, 0x86, 0x71, 0x6c, 0x75,
	0x28, 0x6c, 0xa0, 0xcd, 0x2f, 0x12, 0x2b, 0xbb, 0x0b, 0xcd, 0x58, 0x72, 0xb9, 0x8c, 0x2d, 0x20,
	0xbf, 0x3e, 0xb1, 0x73, 0xe8, 0xe8, 0x8b, 0x18, 0x5b, 0xdd, 0xd3, 0xda, 0x59, 0xf7, 0xfc, 0x70,
	0x9c, 0x29, 0xcd, 0x58, 0x03, 0x4c, 0xd3, 0x30, 0xfb, 0x19, 0xf4, 0x4c, 0x05, 0x7e, 0xf0, 0x62,
	0xc9, 0xbe, 0xca, 0x88, 0xec, 0x7b, 0xb1, 0xb4, 0x2a, 0x84, 0xf3, 0x61, 0x0e, 0xc7, 0xdc, 0x48,
	0xb5, 0x57, 0x77, 0xed, 0xf7, 0x55, 0x18, 0x3e, 0x52, 0x32, 0xe2, 0x3a, 0x00, 0x7f, 0xdd, 0x2c,
	0x5b, 0x65, 0x8f, 0xb2, 0x55, 0xb7, 0x94, 0x6d, 0xad, 0x7b, 0x2d, 0xab, 0xfb, 0x2d, 0xc5, 0x2c,
	0x0a, 0xdf, 0xd8, 0x22, 0xfc, 0x67, 0x30, 0x58, 0x57, 0xea, 0x7a, 0x26, 0x5c, 0xa4, 0xca, 0x36,
	0xa6, 0xfd, 0xb5, 0xf5, 0x91, 0x70, 0x31, 0x2f, 0x6f, 0x6b, 0x3f, 0x79, 0xff, 0xa8, 0x40, 0xcb,
	0x94, 0xed, 0x18, 0x40, 0x3b, 0xd2, 0xfe, 0x36, 0xa1, 0x8e, 0xcb, 0x3e, 0x81, 0x9e, 0x71, 0x67,
	0xba, 0xbb, 0xab, 0x6d, 0xa4, 0xd2, 0x31, 0x40, 0x2c, 0x79, 0x24, 0xaf, 0xa5, 0x17, 0xa0, 0x96,
	0xa1, 0x43, 0x96, 0x2b, 0x2f, 0x40, 0xf6, 0x00, 0x1a, 0x92, 0xc7, 0x6f, 0x63, 0xab, 0x4e, 0xe4,
	0x86, 0x39, 0x72, 0x57, 0x3c, 0x7e, 0x3b, 0x4d, 0xfc, 0xf6, 0x0b, 0xa8, 0xab, 0x23, 0x3b, 0x82,
	0x96, 0x32, 0xa4, 0x74, 0x9a, 0xea, 0xe8, 0xb8, 0xec, 0x3e, 0x74, 0xc8, 0x91, 0x21, 0xd2, 0x56,
	0x06, 0x62, 0xc1, 0xa0, 0xae, 0xa4, 0xd7, 0xf9, 0xe9, 0xb7, 0xfd, 0x23, 0x7c, 0x30, 0x41, 0xd3,
	0xa0, 0x2f, 0xa9, 0x1d, 0x55, 0xed, 0x55, 0x02, 0xe4, 0x41, 0x36, 0x01, 0xf2, 0xc0, 0x71, 0x0b,
	0x5a, 0x54, 0x0b, 0x5a, 0xd8, 0xef, 0x2b, 0x69, 0x5b, 0x3e, 0x13, 0x5e, 0xf8, 0x3f, 0x2e, 0x87,
	0x92, 0xe1, 0xb2, 0xff, 0xac, 0x40, 0xdf, 0x50, 0xba, 0x8c, 0x3c, 0x7c, 0xf3, 0x1f, 0x71, 0xba,
	0xa5, 0xa6, 0xf7, 0xa0, 0x8d, 0xa1, 0x9b, 0x38, 0xeb, 0xe4, 0x6c, 0x61, 0xe8, 0x92, 0x2b, 0x65,
	0xda, 0xc8, 0x31, 0x7d, 0x0d, 0xc3, 0x1c, 0x51, 0x9a, 0xeb, 0xa7, 0x30, 0xe4, 0x45, 0xa3, 0x9e,
	0xed, 0xd1, 0xd6, 0xd9, 0xa6, 0xa8, 0xe9, 0xe6, 0x25, 0x7b, 0x0e, 0x77, 0xb2, 0xa5, 0x21, 0xf4,
	0x43, 0x68, 0x48, 0x21, 0xb9, 0xaf, 0x45, 0x48, 0x0e, 0xec, 0x1b, 0xc8, 0xed, 0x07, 0xab, 0x4a,
	0xe9, 0xee, 0x6d, 0x4d, 0xa7, 0xa0, 0x0a, 0xeb, 0xe4, 0x37, 0x18, 0xbe, 0x5a, 0xb8, 0x85, 0x6d,
	0x72, 0xab, 0xe8, 0xfb, 0x6c, 0x92, 0xe2, 0x52, 0xa8, 0x6d, 0x2e, 0x05, 0x1b, 0xa1, 0x9f, 0xec,
	0xb2, 0x2b, 0xe4, 0xc1, 0x5e, 0x99, 0xd5, 0xd0, 0xa8, 0x66, 0xcf, 0x0d, 0x0d, 0xf2, 0x80, 0xca,
	0xac, 0x8a, 0xe5, 0x8b, 0x39, 0x37, 0xc9, 0xf4, 0xc9, 0x7e, 0x0d, 0xfd, 0xe4, 0x91, 0x26, 0x4d,
	0xe9, 0xc8, 0xfc, 0x2b, 0xf8, 0x25, 0xd4, 0x15, 0x70, 0x39, 0x6a, 0xe1, 0x55, 0xd5, 0xdd, 0xaf,
	0xaa, 0x95, 0xa6, 0xad, 0xe7, 0xd2, 0xfe, 0x04, 0x6d, 0x95, 0x76, 0x47, 0x6f, 0x7c, 0x0e, 0x84,
	0x92, 0xe9, 0x8b, 0xc2, 0xba, 0x52, 0x72, 0xac, 0x43, 0xce, 0xff, 0x6a, 0xc2, 0x81, 0x69, 0x83,
	0x97, 0x18, 0xbd, 0xf3, 0x66, 0xc8, 0x1c, 0x18, 0xe4, 0xff, 0x6d, 0xd8, 0x49, 0x0e, 0x62, 0xe3,
	0xaf, 0x68, 0x74, 0x94, 0xf7, 0xd3, 0xa7, 0xc8, 0x14, 0xe3, 0x05, 0xfb, 0x0e, 0x06, 0x8f, 0xd1,
	0xc7, 0x0c, 0x54, 0x7e, 0xb3, 0x3b, 0xa1, 0xfc, 0x39, 0xe2, 0x8b, 0x72, 0x00, 0x07, 0x06, 0xf9,
	0x5e, 0x2d, 0x70, 0xd9, 0x68, 0xe4, 0x72, 0xa8, 0x6f, 0xe1, 0x60, 0x82, 0x72, 0x3d, 0x86, 0x2b,
	0xe7, 0x71, 0x09, 0x99, 0xed, 0xff, 0xc9, 0xec, 0x02, 0x20, 0x6d, 0x5c, 0x36, 0xda, 0x22, 0x89,
	0x6e, 0xb5, 0x72, 0x0a, 0xdf, 0xc3, 0xd1, 0x04, 0xa5, 0xa9, 0xe0, 0xe5, 0xca, 0x60, 0xef, 0x49,
	0x65, 0x5d, 0xfa, 0xaf, 0x01, 0x12, 0x59, 0x89, 0xca, 0x3f, 0x94, 0xf4, 0x02, 0x20, 0x9d, 0x8c,
	0xc2, 0x3b, 0x72, 0x23, 0x53, 0x0e, 0xf1, 0x25, 0x74, 0xf5, 0x3b, 0x76, 0xc8, 0xb8, 0xd9, 0x77,
	0xec, 0x39, 0xbd, 0x3f, 0xbb, 0x9a, 0x2e, 0x57, 0xaf, 0x62, 0x8c, 0x4a, 0x31, 0x8e, 0x4b, 0x77,
	0x1a, 0xe9, 0x30, 0x85, 0xd1, 0x04, 0x65, 0x22, 0xbe, 0x9b, 0x56, 0x76, 0x27, 0xe4, 0x49, 0xf9,
	0x56, 0x26, 0xcc, 0x27, 0x70, 0x67, 0x82, 0xf2, 0xa9, 0x30, 0x34, 0x3d, 0x8c, 0xd9, 0xdd, 0x71,
	0xf2, 0xed, 0x3d, 0x36, 0xdf, 0xde, 0xe3, 0x27, 0xea, 0xdb, 0x7b, 0xb4, 0x7d, 0xe5, 0x2a, 0x98,
	0x9b, 0x26, 0x85, 0x7e, 0xf1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x20, 0x7f, 0xad, 0xd3,
	0x0b, 0x00, 0x00,
}
