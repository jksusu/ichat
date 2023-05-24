package contacts

import (
	"ichat/pkg/model"
	"ichat/pkg/response"
)

type Friend struct {
	FriendsModel *model.Friends
}

// ContactsService 关系服务
func ContactsService() *Friend {
	return &Friend{}
}

type FriendLists struct {
	FriendUid    string `json:"friend_uid"`
	Remarks      string `json:"remarks"`
	Status       int    `json:"status"`
	Extra        string `json:"extra"`
	UID          string `json:"uid"`
	Nickname     string `json:"nickname"`
	HeadPortrait string `json:"head_portrait"`
}

// FriendList 分页查询好友列表
func (f *Friend) FriendList(uid string, page, pageSize int) *response.Response {

	friends, total, _ := f.FriendsModel.PageQuery(uid, page, pageSize)
	if total <= 0 {
		response.Ok.WithData("联系人为空")
	}

	//组装好友uid
	var ids []string
	for _, u := range friends {
		ids = append(ids, u.FriendUid)
	}

	//获取用户信息
	var user model.Users
	users, err := user.FindInIds(ids)
	if err != nil {
		response.Fail.WithMsg("获取联系人失败")
	}

	var list = make([]*FriendLists, len(friends))

	for i, f := range friends {
		for _, u := range users {
			if f.FriendUid == u.UID {
				list[i] = &FriendLists{
					Remarks:      f.Remarks,
					Status:       f.Status,
					Extra:        f.Extra,
					UID:          u.UID,
					Nickname:     u.Nickname,
					HeadPortrait: u.HeadPortrait,
				}
			}
		}
	}
	return response.Ok.WithPageData(&response.PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}).WithMsg("查询成功")
}
