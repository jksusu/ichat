package model

import "time"

type Friends struct {
	ID        int64     `json:"id" gorm:"comment:主键id"`
	UID       string    `json:"uid" gorm:"comment:我的uid"`
	To        string    `json:"to" gorm:"uniqueIndex;comment:联系人uid"`
	Remarks   string    `json:"remarks" gorm:"comment:备注"`
	Status    int       `json:"status" gorm:"comment:状态，1：正常，2：拉黑"`
	Topping   int       `json:"topping" gorm:"comment:置顶，0：否，1：是"`
	NoDisturb int       `json:"no_disturb" gorm:"comment:免打扰，0：否，1：是"`
	Extra     string    `json:"extra" gorm:"comment:附加字段"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
}

type RelationFriendsApply struct {
	MsgId       int64     `json:"msg_id" gorm:"comment:消息ID"`
	UID         string    `json:"uid" gorm:"comment:我的uid"`
	TO          string    `json:"to"`
	Remarks     string    `json:"remarks" `
	Status      int       `json:"status"`
	Extra       string    `json:"extra" gorm:"comment:附加字段"`
	ApplyTime   time.Time `json:"apply_time" gorm:"comment:发起时间"`
	ProcessTime time.Time `json:"process_time" gorm:"comment:处理时间"`
}

var FriendModel = new(Friends)
var RelationFriendsApplyModel = new(RelationFriendsApply)

// PageQuery 分页
func (f *Friends) PageQuery(userUid string, page, pageSize int) ([]Friends, int64, error) {
	var (
		friends []Friends
		total   int64
		err     error
	)
	offset := (page - 1) * pageSize
	//查询总数
	if err = DB.Model(&Friends{}).Where("user_uid = ?", userUid).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = DB.Model(&Friends{}).Where("user_uid= ?", userUid).Limit(pageSize).Offset(offset).Find(&friends).Error
	return friends, total, err
}

// AddFriend 添加一条数据
func (f *Friends) AddFriend(list *Friends) bool {
	if DB.Create(list).RowsAffected > 0 {
		return true
	}
	return false
}

// StatusUpdate 好友状态修改
func (f *Friends) StatusUpdate(userUid, friendUid string, status int) bool {
	result := DB.First(f, "user_uid = ? AND friend_uid = ?", userUid, friendUid)
	if result.RowsAffected > 0 {
		f.Status = status
		result.Save(f)
		return true
	}
	return false
}

func (r *RelationFriendsApply) Add(data *RelationFriendsApply) (ok bool, err error) {
	result := DB.Create(data)
	return result.RowsAffected > 0, result.Error
}

func (r *RelationFriendsApply) UpdateProcessStatus(msgId int64, status int) error {
	var relation = &RelationFriendsApply{
		Status:      status,
		ProcessTime: time.Now(),
	}
	res := DB.Where("msg_id = ?", msgId).Updates(&relation)
	if res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func (r *RelationFriendsApply) FirstByMsgId() (relation *RelationFriendsApply, err error) {
	err = DB.Where("msg_id = ?", r.MsgId).First(&relation).Error
	return
}
