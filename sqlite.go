package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Xiaoshuo struct {
	HeadUrl   string `gorm:"primaryKey;column:head_url",json:"headurl"`
	Content   string `gorm:"content",json:"content"`
	Title     string `gorm:"title",json:"title"`
	NextPage  string `gorm:"nextpage",json:"nextpage"`
	BreakFlag string `gorm:"breakflag",json:"breakflag"`
}

var (
	db  *gorm.DB
	err error
)

func init() {

	db, err = gorm.Open(sqlite.Open("./xs.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Xiaoshuo{}) //AutoMigrate 方法会根据 User 结构体自动创建 users 表，如果表已经存在，则会更新表的结构以匹配结构体。

}

func getDb(headUrl string) (xiaoshuo Xiaoshuo) {
	db.Where("head_url = ?", headUrl).First(&xiaoshuo)
	return xiaoshuo

}
func insertDb(xiaoshuo Xiaoshuo) {
	db.Save(xiaoshuo)
}

func getCount(headUrl string) (bool, error) {
	var count int64

	err := db.Table("xiaoshuos").Where("head_url=?", headUrl).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return count > 0, err
}
