package qqbotapi

import (
	"strconv"
	"strings"
)

type RobotClient struct {
	client    *apiClient
	currentQQ string
}

func NewRobotClient(hostport string, currentQQ int64) (c *RobotClient) {
	c = &RobotClient{
		client:    NewApiclient(hostport),
		currentQQ: strconv.FormatInt(currentQQ, 10),
	}

	return
}

func (c *RobotClient) InstallService() (err error) {
	return
}

func (c *RobotClient) SendMsg(param *SendMsgParam) (err error) {
	resp := &RespHead{}
	action := "LuaApiCaller?qq=CurrentQQ&funcname=SendMsg&timeout=10"
	action = strings.Replace(action, "CurrentQQ", c.currentQQ, -1)
	err = c.client.Post("v1", "", action, param, resp)
	if err != nil {
		LogError("send req failed, err: %v", err)
		return
	}
	if resp.Ret != 0 {
		LogError("ret is not 0, msg:%v", resp.Msg)
		return
	}
	return
}

func (c *RobotClient) GetQQUserList(param *GetQQUserListParam) (err error, resp *GetQQUserListResp) {
	resp = &GetQQUserListResp{}
	action := "LuaApiCaller?qq=CurrentQQ&funcname=GetQQUserList&timeout=10"
	action = strings.Replace(action, "CurrentQQ", c.currentQQ, -1)
	err = c.client.Post("v1", "", action, param, resp)
	if err != nil {
		LogError("send req failed, err: %v", err)
		return
	}
	return
}

func (c *RobotClient) SearchGroup(param *SearchGroupParam) (err error, resp *SearchGroupResp) {
	resp = &SearchGroupResp{}
	action := "LuaApiCaller?qq=CurrentQQ&funcname=SearchGroup&timeout=10"
	action = strings.Replace(action, "CurrentQQ", c.currentQQ, -1)
	err = c.client.Post("v1", "", action, param, resp)
	if err != nil {
		LogError("send req failed, err: %v", err)
		return
	}
	return
}

func (c *RobotClient) GetGroupList() (err error) {
	return
}

func (c *RobotClient) AddQQUser() (err error) {
	return
}

func (c *RobotClient) RevokeMsg() (err error) {
	return
}

func (c *RobotClient) ShutUp() (err error) {
	return
}

func (c *RobotClient) LogOut() (err error) {
	return
}
