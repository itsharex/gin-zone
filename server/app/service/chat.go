package service

import (
	"fmt"

	db "gitee.com/jiang-xia/gin-zone/server/app/database"
	"gitee.com/jiang-xia/gin-zone/server/app/model"
)

type chat struct {
}

// Chat 结构体实例化
var Chat *chat

type ChatFriends struct {
	model.ChatFriends `gorm:"embedded"`
	model.User        `gorm:"embedded"` // 合并成一个结构体
	model.ChatGroup   `gorm:"embedded"`
}

// ChatFriends 获取好友列表
func (u *chat) ChatFriends(userId string) []ChatFriends {
	var friends []ChatFriends
	db.Mysql.Where("user_id = ?", userId).Or("group_id = ?", userId).Find(&friends)
	db.Mysql.Model(&friends)
	//去根据好友id查询用户信息
	for i, friend := range friends {
		var user model.User
		db.Mysql.Where("user_id = ?", friend.FriendId).Find(&user)
		//需要取切片里的元素，不能使用循环中的friend
		friends[i].User = user
		// 是群组查询群信息
		if friend.GroupId != 0 {
			var chatGroup model.ChatGroup
			db.Mysql.Where("id = ?", friend.GroupId).Find(&chatGroup)
			friends[i].ChatGroup = chatGroup
		}
		// fmt.Printf("user===============================:%+v", friend.User)
	}
	// fmt.Printf("user===============================:%+v", friends)
	return friends
}

// CreateChatFriends 新增
func (u *chat) CreateChatFriends(model *model.ChatFriends) (err error) {
	res := db.Mysql.Create(model)
	if res.Error != nil { //判断是否插入数据出错
		fmt.Println(res.Error)
	}
	return
}

// DeleteChatFriends 删除
func (u *chat) DeleteChatFriends(id int) bool {
	db.Mysql.Where("id = ?", id).Delete(&model.ChatFriends{})
	return true
}

// ChatLogList 获取聊天记录
// https://gorm.io/zh_CN/docs/query.html#%E6%9D%A1%E4%BB%B6
func (u *chat) ChatLogList(Page int, PageSize int, maps interface{}) ([]model.ChatLog, int64) {
	// Where可以使用Struct或者Map作为条件
	var list []model.ChatLog
	var total int64
	db.Mysql.Where(maps).Offset((Page - 1) * PageSize).Limit(PageSize).Find(&list)
	// 条件统计
	db.Mysql.Model(&list).Where(maps).Count(&total)
	// fmt.Printf("查询参数: %+v", maps)
	return list, total
}

// CreateChatLog 新增
func (u *chat) CreateChatLog(model *model.ChatLog) (err error) {
	res := db.Mysql.Create(model)
	if res.Error != nil { //判断是否插入数据出错
		fmt.Println(res.Error)
	}
	return
}

// DeleteChatLog 删除
func (u *chat) DeleteChatLog(id int) bool {
	db.Mysql.Where("id = ?", id).Delete(&model.ChatLog{})
	return true
}

// 群组
func (u *chat) ChatGroup(userId string) []model.ChatGroup {
	var list []model.ChatGroup
	db.Mysql.Where("user_id = ?", userId).Find(&list)
	return list
}

// CreateGroup 新增
func (u *chat) CreateGroup(model *model.ChatGroup) (err error) {
	res := db.Mysql.Create(model)
	if res.Error != nil { //判断是否插入数据出错
		fmt.Println(res.Error)
	}
	return
}

// DeleteGroup 删除
func (u *chat) DeleteGroup(id int) bool {
	db.Mysql.Where("id = ?", id).Delete(&model.ChatGroup{})
	return true
}

// ChatGroupMember 群成员
func (u *chat) ChatGroupMember(groupId int) []model.ChatGroupMember {
	var list []model.ChatGroupMember
	db.Mysql.Where("group_id = ?", groupId).Find(&list)
	return list
}

// CreateChatGroupMember 新增
func (u *chat) CreateChatGroupMember(model *model.ChatGroupMember) (err error) {
	res := db.Mysql.Create(model)
	if res.Error != nil { //判断是否插入数据出错
		fmt.Println(res.Error)
	}
	return
}

// DeleteChatGroupMember 删除
func (u *chat) DeleteChatGroupMember(id int) bool {
	db.Mysql.Where("id = ?", id).Delete(&model.ChatGroupMember{})
	return true
}