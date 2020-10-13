package qqbotapi

type RespHead struct {
	Ret int64  `json:"Ret"`
	Msg string `json:"Msg"`
}

// 发送消息
type SendMsgParam struct {
	ToUser      int64  `json:"toUser"`
	SendToType  int64  `json:"sendToType"`  // 1.好友 2.群 3.私聊
	SendMsgType string `json:"sendMsgType"` // TextMsg, JsonMsg, XmlMsg, ReplayMsg, TeXiaoTextMsg, PicMsg, VoiceMsg, PhpneMsg
	GroupId     int64  `json:"groupid"`     // 私聊时传群id, 其它情况传0
	Content     string `json:"content"`
	AtUser      int64  `json:"atUser"` // default:0

	VoiceUrl       string      `json:"voiceUrl"`
	VoiceBase64Buf string      `json:"voiceBase64Buf"`
	PicUrl         string      `json:"picUrl"`
	PicBase64Buf   string      `json:"picBase64Buf"`
	ForwordBuf     string      `json:"forwordBuf"`
	ForwordField   string      `json:"forwordFiled"`
	FileMd5        string      `json:"fileMd5"`
	ReplyInfo      interface{} `json:"replyInfo"`
}

// 搜索群组
type SearchGroupParam struct {
	Content string `json:"Content"`
	Page    int64  `json:"Page"`
}

type SearchGroupResp struct {
	GroupData         string `json:"GroupData"`         // 群名
	GroupDes          string `json:"GroupDes"`          // 群介绍
	GroupID           int64  `json:"GroupID"`           // 群号
	GroupMaxMembers   int64  `json:"GroupMaxMembers"`   // 群容量
	GroupName         string `json:"GroupName"`         // 群名
	GroupNotice       string `json:"GroupNotice"`       // 群公告
	GroupOwner        int64  `json:"GroupOwner"`        // 群主qq
	GroupQuestion     string `json:"GroupQuestion"`     // 进群问题
	GroupTotalMembers int64  `json:"GroupTotalMembers"` // 当前群人数
}

// 获取好友列表
type GetQQUserListParam struct {
	StartIndex int64 `json:"StartIndex"`
}

type GetQQUserListResp struct {
	FriendCount int64       `json:"Friend_Count"`
	FriendList  []*FriendSt `json:"FriendList"`
}

type FriendSt struct {
	FriendUin int64  `json:"FriendUin"` // 好友账号
	IsRemark  bool   `json:"IsRemark"`  // 是否有备注
	NickName  string `json:"NickName"`  // 昵称
	OnlineStr string `json:"OnlineStr"` // 在线状态
	Remark    string `json:"Remark"`    // 备注
	Status    int64  `json:"Status"`    // 状态 10.在线 20.隐身
}
