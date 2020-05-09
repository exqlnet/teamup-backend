// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: activity.proto

package activity_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ActivityService service

func NewActivityServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ActivityService service

type ActivityService interface {
	// 发起活动
	CreateActivity(ctx context.Context, in *CreateActivityReq, opts ...client.CallOption) (*CommonResp, error)
	// 删除活动
	DeleteActivity(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*CommonResp, error)
	// 更新活动
	UpdateActivity(ctx context.Context, in *UpdateActivityReq, opts ...client.CallOption) (*CommonResp, error)
	// 获取活动详情
	GetActivityByID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*Activity, error)
	// 创建队伍
	CreateTeam(ctx context.Context, in *CreateTeamReq, opts ...client.CallOption) (*CommonResp, error)
	// activity的所有队伍
	GetTeamListByActivityID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*TeamList, error)
	// 删除队伍
	DeleteTeam(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*CommonResp, error)
	// 更新队伍
	UpdateTeam(ctx context.Context, in *UpdateTeamReq, opts ...client.CallOption) (*CommonResp, error)
	// 获取队伍
	GetTeamByID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*Team, error)
	// 获取已参加的活动（简要）
	GetActivityJoinByUserID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*ActivityJoinList, error)
	// 获取发起的活动（简要）
	GetCreatedActivityByUserID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*ActivityBriefList, error)
	// 获取热门活动
	GetHotActivities(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ActivityList, error)
}

type activityService struct {
	c    client.Client
	name string
}

func NewActivityService(name string, c client.Client) ActivityService {
	return &activityService{
		c:    c,
		name: name,
	}
}

func (c *activityService) CreateActivity(ctx context.Context, in *CreateActivityReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ActivityService.CreateActivity", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) DeleteActivity(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ActivityService.DeleteActivity", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) UpdateActivity(ctx context.Context, in *UpdateActivityReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ActivityService.UpdateActivity", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) GetActivityByID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*Activity, error) {
	req := c.c.NewRequest(c.name, "ActivityService.GetActivityByID", in)
	out := new(Activity)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) CreateTeam(ctx context.Context, in *CreateTeamReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ActivityService.CreateTeam", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) GetTeamListByActivityID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*TeamList, error) {
	req := c.c.NewRequest(c.name, "ActivityService.GetTeamListByActivityID", in)
	out := new(TeamList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) DeleteTeam(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ActivityService.DeleteTeam", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) UpdateTeam(ctx context.Context, in *UpdateTeamReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ActivityService.UpdateTeam", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) GetTeamByID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*Team, error) {
	req := c.c.NewRequest(c.name, "ActivityService.GetTeamByID", in)
	out := new(Team)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) GetActivityJoinByUserID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*ActivityJoinList, error) {
	req := c.c.NewRequest(c.name, "ActivityService.GetActivityJoinByUserID", in)
	out := new(ActivityJoinList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) GetCreatedActivityByUserID(ctx context.Context, in *IntWrap, opts ...client.CallOption) (*ActivityBriefList, error) {
	req := c.c.NewRequest(c.name, "ActivityService.GetCreatedActivityByUserID", in)
	out := new(ActivityBriefList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityService) GetHotActivities(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ActivityList, error) {
	req := c.c.NewRequest(c.name, "ActivityService.GetHotActivities", in)
	out := new(ActivityList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ActivityService service

type ActivityServiceHandler interface {
	// 发起活动
	CreateActivity(context.Context, *CreateActivityReq, *CommonResp) error
	// 删除活动
	DeleteActivity(context.Context, *IntWrap, *CommonResp) error
	// 更新活动
	UpdateActivity(context.Context, *UpdateActivityReq, *CommonResp) error
	// 获取活动详情
	GetActivityByID(context.Context, *IntWrap, *Activity) error
	// 创建队伍
	CreateTeam(context.Context, *CreateTeamReq, *CommonResp) error
	// activity的所有队伍
	GetTeamListByActivityID(context.Context, *IntWrap, *TeamList) error
	// 删除队伍
	DeleteTeam(context.Context, *IntWrap, *CommonResp) error
	// 更新队伍
	UpdateTeam(context.Context, *UpdateTeamReq, *CommonResp) error
	// 获取队伍
	GetTeamByID(context.Context, *IntWrap, *Team) error
	// 获取已参加的活动（简要）
	GetActivityJoinByUserID(context.Context, *IntWrap, *ActivityJoinList) error
	// 获取发起的活动（简要）
	GetCreatedActivityByUserID(context.Context, *IntWrap, *ActivityBriefList) error
	// 获取热门活动
	GetHotActivities(context.Context, *empty.Empty, *ActivityList) error
}

func RegisterActivityServiceHandler(s server.Server, hdlr ActivityServiceHandler, opts ...server.HandlerOption) error {
	type activityService interface {
		CreateActivity(ctx context.Context, in *CreateActivityReq, out *CommonResp) error
		DeleteActivity(ctx context.Context, in *IntWrap, out *CommonResp) error
		UpdateActivity(ctx context.Context, in *UpdateActivityReq, out *CommonResp) error
		GetActivityByID(ctx context.Context, in *IntWrap, out *Activity) error
		CreateTeam(ctx context.Context, in *CreateTeamReq, out *CommonResp) error
		GetTeamListByActivityID(ctx context.Context, in *IntWrap, out *TeamList) error
		DeleteTeam(ctx context.Context, in *IntWrap, out *CommonResp) error
		UpdateTeam(ctx context.Context, in *UpdateTeamReq, out *CommonResp) error
		GetTeamByID(ctx context.Context, in *IntWrap, out *Team) error
		GetActivityJoinByUserID(ctx context.Context, in *IntWrap, out *ActivityJoinList) error
		GetCreatedActivityByUserID(ctx context.Context, in *IntWrap, out *ActivityBriefList) error
		GetHotActivities(ctx context.Context, in *empty.Empty, out *ActivityList) error
	}
	type ActivityService struct {
		activityService
	}
	h := &activityServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ActivityService{h}, opts...))
}

type activityServiceHandler struct {
	ActivityServiceHandler
}

func (h *activityServiceHandler) CreateActivity(ctx context.Context, in *CreateActivityReq, out *CommonResp) error {
	return h.ActivityServiceHandler.CreateActivity(ctx, in, out)
}

func (h *activityServiceHandler) DeleteActivity(ctx context.Context, in *IntWrap, out *CommonResp) error {
	return h.ActivityServiceHandler.DeleteActivity(ctx, in, out)
}

func (h *activityServiceHandler) UpdateActivity(ctx context.Context, in *UpdateActivityReq, out *CommonResp) error {
	return h.ActivityServiceHandler.UpdateActivity(ctx, in, out)
}

func (h *activityServiceHandler) GetActivityByID(ctx context.Context, in *IntWrap, out *Activity) error {
	return h.ActivityServiceHandler.GetActivityByID(ctx, in, out)
}

func (h *activityServiceHandler) CreateTeam(ctx context.Context, in *CreateTeamReq, out *CommonResp) error {
	return h.ActivityServiceHandler.CreateTeam(ctx, in, out)
}

func (h *activityServiceHandler) GetTeamListByActivityID(ctx context.Context, in *IntWrap, out *TeamList) error {
	return h.ActivityServiceHandler.GetTeamListByActivityID(ctx, in, out)
}

func (h *activityServiceHandler) DeleteTeam(ctx context.Context, in *IntWrap, out *CommonResp) error {
	return h.ActivityServiceHandler.DeleteTeam(ctx, in, out)
}

func (h *activityServiceHandler) UpdateTeam(ctx context.Context, in *UpdateTeamReq, out *CommonResp) error {
	return h.ActivityServiceHandler.UpdateTeam(ctx, in, out)
}

func (h *activityServiceHandler) GetTeamByID(ctx context.Context, in *IntWrap, out *Team) error {
	return h.ActivityServiceHandler.GetTeamByID(ctx, in, out)
}

func (h *activityServiceHandler) GetActivityJoinByUserID(ctx context.Context, in *IntWrap, out *ActivityJoinList) error {
	return h.ActivityServiceHandler.GetActivityJoinByUserID(ctx, in, out)
}

func (h *activityServiceHandler) GetCreatedActivityByUserID(ctx context.Context, in *IntWrap, out *ActivityBriefList) error {
	return h.ActivityServiceHandler.GetCreatedActivityByUserID(ctx, in, out)
}

func (h *activityServiceHandler) GetHotActivities(ctx context.Context, in *empty.Empty, out *ActivityList) error {
	return h.ActivityServiceHandler.GetHotActivities(ctx, in, out)
}
