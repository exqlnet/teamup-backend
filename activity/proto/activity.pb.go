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
	StartTime            int64      `protobuf:"varint,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64      `protobuf:"varint,9,opt,name=endTime,proto3" json:"endTime,omitempty"`
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

func (m *CreateActivityReq) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *CreateActivityReq) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type Process struct {
	ProcessId            int32    `protobuf:"varint,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProcessName          string   `protobuf:"bytes,2,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
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

func (m *Process) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
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
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x56, 0xe2, 0xfc, 0x9e, 0xfc, 0x6c, 0x33, 0x2c, 0xad, 0x9b, 0xb2, 0xcb, 0x62, 0x84, 0xba,
	0x37, 0xa4, 0xd2, 0x22, 0x71, 0x51, 0x89, 0x15, 0x6c, 0x7f, 0xb2, 0xa9, 0xa0, 0x54, 0xee, 0x56,
	0x5c, 0x55, 0xab, 0xd9, 0xf8, 0x34, 0x58, 0xb5, 0x3d, 0xc1, 0x9e, 0x54, 0x8a, 0xc4, 0x05, 0x8f,
	0x80, 0xc4, 0x43, 0xf0, 0x2c, 0x48, 0x3c, 0x14, 0x9a, 0x63, 0x4f, 0xfc, 0x13, 0x3b, 0x0d, 0x08,
	0xf5, 0x2e, 0x73, 0xe6, 0xcc, 0x77, 0xbe, 0xf9, 0xbe, 0x33, 0x27, 0x86, 0x21, 0x9f, 0x4b, 0xf7,
	0x9d, 0x2b, 0xd7, 0x93, 0x65, 0x28, 0xa4, 0x60, 0x3d, 0xbd, 0xbe, 0x5e, 0xde, 0x8c, 0xef, 0x2d,
	0x84, 0x58, 0x78, 0xf8, 0x80, 0xb6, 0x6e, 0x56, 0x6f, 0x1e, 0xa0, 0xbf, 0xd4, 0x99, 0xe3, 0xfe,
	0x5c, 0xf8, 0xbe, 0x08, 0xe2, 0x95, 0xf5, 0x9b, 0x01, 0x9d, 0xef, 0x92, 0xa3, 0xec, 0x53, 0x48,
	0x61, 0x5c, 0xc7, 0xac, 0x9d, 0xd4, 0x4e, 0x9b, 0x36, 0xe8, 0xd0, 0xcc, 0x61, 0x9f, 0xc3, 0x60,
	0x93, 0x10, 0x70, 0x1f, 0xcd, 0xfa, 0x49, 0xed, 0xb4, 0x6b, 0xf7, 0x75, 0xf0, 0x39, 0xf7, 0x91,
	0x59, 0xd0, 0x77, 0x03, 0x19, 0x0a, 0x67, 0x35, 0x97, 0xae, 0x08, 0x4c, 0x23, 0xce, 0xc9, 0xc6,
	0xd8, 0x11, 0xc0, 0x3c, 0x44, 0x2e, 0x45, 0xa8, 0x0a, 0x35, 0xa8, 0x50, 0x37, 0x89, 0xcc, 0x1c,
	0x76, 0x0c, 0x10, 0xe2, 0x62, 0xe5, 0x71, 0x02, 0x68, 0x12, 0x40, 0x26, 0xc2, 0x0e, 0xa1, 0x19,
	0x0a, 0x0f, 0x23, 0xb3, 0x45, 0x5b, 0xf1, 0x82, 0x7d, 0x02, 0x5d, 0xbe, 0x92, 0x3f, 0x8b, 0xd0,
	0x95, 0x6b, 0xb3, 0x4d, 0x3b, 0x69, 0x80, 0x68, 0x79, 0xde, 0x2a, 0x92, 0x61, 0x8c, 0xda, 0x49,
	0x68, 0x65, 0x62, 0xec, 0x3e, 0x1c, 0xcc, 0x57, 0x61, 0x88, 0x81, 0xbc, 0x5e, 0x86, 0x62, 0x8e,
	0x51, 0x64, 0x76, 0x29, 0x6d, 0x98, 0x84, 0x5f, 0xc4, 0x51, 0x76, 0x1b, 0x5a, 0x91, 0xe4, 0x72,
	0x15, 0x99, 0x40, 0xfb, 0xc9, 0x8a, 0x9d, 0x41, 0x37, 0x39, 0x88, 0x91, 0xd9, 0x3b, 0x31, 0x4e,
	0x7b, 0x67, 0x87, 0x93, 0x8c, 0x35, 0x93, 0x04, 0xc0, 0x4e, 0xd3, 0xac, 0x67, 0xd0, 0xd7, 0x0e,
	0x7c, 0xef, 0x46, 0x92, 0x3d, 0xcc, 0x88, 0xec, 0xb9, 0x91, 0x34, 0x6b, 0x84, 0xf3, 0x71, 0x0e,
	0x47, 0x9f, 0x48, 0xb5, 0x57, 0x67, 0xad, 0xbf, 0xea, 0x30, 0x7a, 0xa4, 0x64, 0xc4, 0x4d, 0x02,
	0xfe, 0xb2, 0x6d, 0x5b, 0x6d, 0x0f, 0xdb, 0xea, 0x25, 0xb6, 0x6d, 0x74, 0x37, 0xb2, 0xba, 0xbf,
	0xc7, 0xcc, 0xa2, 0xf0, 0xcd, 0x12, 0xe1, 0xbf, 0x80, 0xe1, 0xc6, 0xa9, 0xeb, 0xb9, 0x70, 0x90,
	0x9c, 0x6d, 0xda, 0x83, 0x4d, 0xf4, 0x91, 0x70, 0x30, 0x2f, 0x6f, 0x7b, 0x2f, 0x79, 0x15, 0xbb,
	0x48, 0xf2, 0x50, 0x5e, 0x4b, 0xd7, 0x47, 0x72, 0xdd, 0xb0, 0xbb, 0x14, 0xb9, 0x72, 0x7d, 0x64,
	0x26, 0xb4, 0x31, 0x70, 0xd4, 0x4f, 0xb2, 0xda, 0xb0, 0xf5, 0xd2, 0xfa, 0xa3, 0x06, 0x6d, 0xed,
	0xf7, 0x11, 0x40, 0x82, 0x98, 0x3e, 0x0c, 0x5d, 0x63, 0xe6, 0xb0, 0xcf, 0xa0, 0xaf, 0xb7, 0x33,
	0xcf, 0xa2, 0x97, 0xc4, 0x48, 0xde, 0x3c, 0x0d, 0xa3, 0x48, 0xe3, 0x3e, 0x34, 0x25, 0x8f, 0xde,
	0x46, 0x66, 0x83, 0x6e, 0x35, 0xca, 0xdd, 0xea, 0x8a, 0x47, 0x6f, 0xed, 0x78, 0xdf, 0x7a, 0x01,
	0x0d, 0xb5, 0x64, 0x77, 0xa0, 0xad, 0x02, 0x29, 0x9d, 0x96, 0x5a, 0xce, 0x1c, 0x76, 0x0f, 0xba,
	0xb4, 0x91, 0x21, 0xd2, 0x51, 0x01, 0x62, 0xc1, 0xa0, 0xa1, 0x3c, 0x4b, 0xfc, 0xa3, 0xdf, 0xd6,
	0x0f, 0xf0, 0xd1, 0x14, 0x75, 0x67, 0xbf, 0xa4, 0x3e, 0x56, 0x4d, 0xa3, 0x0a, 0x20, 0xf7, 0xb3,
	0x05, 0x90, 0xfb, 0x33, 0xa7, 0xa0, 0x45, 0xbd, 0xa0, 0x85, 0xf5, 0x7b, 0x2d, 0xed, 0xe7, 0x67,
	0xc2, 0x0d, 0x3e, 0xe0, 0x54, 0xa9, 0x78, 0x95, 0xd6, 0x9f, 0x35, 0x18, 0x68, 0x4a, 0x17, 0xa1,
	0x8b, 0x6f, 0xfe, 0x27, 0x4e, 0xdb, 0x9e, 0x76, 0xb3, 0x9e, 0xde, 0x85, 0x0e, 0x06, 0x4e, 0xbc,
	0xd9, 0xa0, 0x4d, 0xdd, 0x5b, 0x19, 0xa6, 0xcd, 0x1c, 0xd3, 0xd7, 0x30, 0xca, 0x11, 0xa5, 0x81,
	0x70, 0x09, 0x23, 0x5e, 0x0c, 0x26, 0x43, 0x61, 0x5c, 0x3a, 0x14, 0x28, 0xcb, 0xde, 0x3e, 0x64,
	0x2d, 0xe0, 0x56, 0xd6, 0x1a, 0x42, 0x3f, 0x84, 0xa6, 0x14, 0x92, 0x7b, 0x89, 0x08, 0xf1, 0x82,
	0x7d, 0x03, 0xb9, 0xc1, 0x62, 0xd6, 0xa9, 0xdc, 0xdd, 0xd2, 0x72, 0x0a, 0xaa, 0x30, 0x87, 0x7e,
	0x85, 0xd1, 0xab, 0xa5, 0x53, 0x18, 0x43, 0xef, 0x15, 0x7d, 0x9f, 0x11, 0x54, 0x9c, 0x26, 0xc6,
	0xf6, 0x34, 0xb1, 0x10, 0x06, 0xf1, 0x10, 0xbc, 0x42, 0xee, 0xef, 0x55, 0x59, 0x3d, 0x1a, 0xd5,
	0xec, 0xb9, 0x47, 0x83, 0xdc, 0x27, 0x9b, 0x95, 0x59, 0x9e, 0x58, 0x70, 0x5d, 0x2c, 0x59, 0x59,
	0xaf, 0x61, 0x10, 0x5f, 0x52, 0x97, 0xa9, 0x7c, 0x32, 0xff, 0x09, 0x7e, 0x05, 0x0d, 0x05, 0x5c,
	0x8d, 0x5a, 0xb8, 0x55, 0x7d, 0xf7, 0xad, 0x8c, 0xca, 0xb2, 0x8d, 0x5c, 0xd9, 0x1f, 0xa1, 0xa3,
	0xca, 0xee, 0xe8, 0x8d, 0x2f, 0x81, 0x50, 0x32, 0x7d, 0x51, 0x18, 0x57, 0x4a, 0x8e, 0x4d, 0xca,
	0xd9, 0xdf, 0x2d, 0x38, 0xd0, 0x6d, 0xf0, 0x12, 0xc3, 0x77, 0xee, 0x1c, 0xd9, 0x53, 0x18, 0xe6,
	0xff, 0xa6, 0xd8, 0x71, 0x0e, 0x62, 0xeb, 0x3f, 0x6c, 0x9c, 0x9f, 0xf3, 0xb3, 0x40, 0xfe, 0x14,
	0xf2, 0x25, 0x3b, 0x87, 0xe1, 0x63, 0xf4, 0x30, 0x83, 0x53, 0x9a, 0x37, 0xbe, 0x3d, 0x89, 0x3f,
	0x89, 0x26, 0xfa, 0x93, 0x68, 0xf2, 0x44, 0x7d, 0x12, 0xb1, 0x4b, 0x18, 0xe6, 0xfb, 0xb4, 0xc0,
	0x63, 0xab, 0x89, 0x2b, 0x91, 0xce, 0xe1, 0x60, 0x8a, 0x72, 0xf3, 0x02, 0xd7, 0xb3, 0xc7, 0x15,
	0x54, 0xca, 0xff, 0xc7, 0xd9, 0x39, 0x40, 0xda, 0xb3, 0x6c, 0x5c, 0xa2, 0x46, 0xd2, 0x65, 0x15,
	0x4a, 0x3c, 0x85, 0x3b, 0x53, 0x94, 0xda, 0xb9, 0x8b, 0xb5, 0x06, 0xde, 0x93, 0xc7, 0xc6, 0xf2,
	0x87, 0x00, 0xb1, 0xa2, 0xc4, 0xe3, 0xdf, 0xa9, 0xf9, 0x2d, 0x40, 0xfa, 0x20, 0x0a, 0x77, 0xc8,
	0xbd, 0x94, 0x4a, 0x84, 0xaf, 0xa1, 0x97, 0xdc, 0x62, 0x87, 0x82, 0xdb, 0xdd, 0xc6, 0x9e, 0xd3,
	0xed, 0xb3, 0x03, 0xe9, 0x62, 0xfd, 0x2a, 0xc2, 0xb0, 0x12, 0xe3, 0xa8, 0x72, 0x92, 0x91, 0x0a,
	0x36, 0x8c, 0xa7, 0x28, 0x63, 0xdd, 0x9d, 0xd4, 0xd4, 0x9d, 0x90, 0xc7, 0xd5, 0xb3, 0x98, 0x30,
	0x9f, 0xc0, 0xad, 0x29, 0xca, 0x4b, 0xa1, 0x69, 0xba, 0x18, 0xb1, 0x0a, 0x1d, 0xc6, 0xe5, 0x83,
	0x56, 0xc1, 0xdc, 0xb4, 0x28, 0xf5, 0xab, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xce, 0x5a,
	0xa9, 0x02, 0x0c, 0x00, 0x00,
}
