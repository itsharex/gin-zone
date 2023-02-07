package service

import (
	"fmt"
	db "gitee.com/jiang-xia/gin-zone/server/app/database"
	"gitee.com/jiang-xia/gin-zone/server/app/model"
)

type moment struct {
}

// Moment 结构体实例化
var Moment *moment

// List 获取 moment 列表
func (m *moment) List(Page int, PageSize int, maps interface{}) ([]model.Moment, int64) {
	var moments []model.Moment
	var total int64
	//如果设置了关联会自动查询，这个设置了关联用户表，所以自动查询user信息
	db.Mysql.Where(maps).Offset((Page - 1) * PageSize).Limit(PageSize).Find(&moments)
	db.Mysql.Model(&moments).Count(&total)
	return moments, total
}

// CreateMoment 新增
func (m *moment) CreateMoment(moment *model.Moment) (err error) {
	// Omit跳过关联创建
	res := db.Mysql.Create(moment)
	if res.Error != nil { //判断是否插入数据出错
		fmt.Println(res.Error)
	}
	return
}

// Update 修改
func (m *moment) Update(id int, model *model.Moment) (err error) {
	err = db.Mysql.Model(&model).Where("id = ? ", id).Updates(model).Error
	return err
}
func (m *moment) Find(id int) (model *model.Moment, err error) {
	err = db.Mysql.Where("id = ? ", id).Find(&model).Error
	return model, err
}
