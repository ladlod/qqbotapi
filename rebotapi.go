package rebotapi

import (
	"strconv"
	"strings"
)

type rebotClient struct {
	client    *apiClient
	currentQQ string
}

func NewRebotClient(hostport string, currentQQ int64) (c *rebotClient) {
	c = &rebotClient{
		client:    NewApiclient(hostport),
		currentQQ: strconv.FormatInt(currentQQ, 10),
	}

	return
}

func (c *rebotClient) InstallService() (err error) {
	return
}

func (c *rebotClient) SendMsg(param *SendMsgParam) (err error) {
	resp := &RespHead{}
	action := "LuaApiCaller?qq=CurrentQQ&funcname=SendMsg&timeout=10"
	strings.Replace(action, "CurrentQQ", c.currentQQ, -1)
	err = c.client.Post("v1", "", action, param, resp)
	if err != nil {
		logError("send msg failed, err: %v", err)
		return
	}
	if resp.Ret != 0 {
		logError("ret is not 0, msg:%v", resp.Msg)
		return
	}
	return
}

func (c *rebotClient) GetQQUserList() (err error) {
	return
}

func (c *rebotClient) SearchGroup(param *SearchGroupParam) (err error, resp *SearchGroupResp) {
	resp = &SearchGroupResp{}
	action := "LuaApiCaller?qq=CurrentQQ&funcname=SearchGroup&timeout=10"
	strings.Replace(action, "CurrentQQ", c.currentQQ, -1)
	err = c.client.Post("v1", "", action, param, resp)
	if err != nil {
		logError("send msg failed, err: %v", err)
		return
	}
	return
}

func (c *rebotClient) GetGroupList() (err error) {
	return
}

func (c *rebotClient) AddQQUser() (err error) {
	return
}

func (c *rebotClient) RemoveMsg() (err error) {
	return
}

func (c *rebotClient) ShutUp() (err error) {
	return
}

func (c *rebotClient) LogOut() (err error) {
	return
}
