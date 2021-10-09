package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type (
	//UserInfo 用户信息
	UserInfo struct {
		UUID           uint64     `gorm:"primary_key"`
		CreatedAt      time.Time  `gorm:"index"`
		UpdatedAt      time.Time  `gorm:"index"`
		DeletedAt      *time.Time `gorm:"index"`
		AccountID      string     `gorm:"index"`
		Nike           string     `gorm:"index"`
		Icon           string
		IconFrame      int32
		Course         int32
		AccountType    int // 0 是游客  1 怪力猫账号 2 bilibili账号
		ServerID       int
		ServerNum      int  `gorm:"default:'1'"`
		Banned         bool // 是否封号状态
		DefaultWarRole uint64
		DefaultWeapon  uint64
		DefaultPet1    uint64
		DefaultPet2    uint64
		DefaultPet3    uint64
		// DefaultShip    int32
		LastUpdateTime uint64
		Age            int32
		GameTime       int32
		MonthPay       int32
		Body           []byte // 对应OtherData结构体的二进制串
		ClickEffect    int32  //点击特效
	}
)

func main() {
	dbConfig := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=disable client_encoding=UTF8 timezone=Asia/Shanghai",
		"127.0.0.1",
		5432,
		"postgres",
		"ginisuger",
		"ginibong",
	)

	postgresDb, err := gorm.Open("postgres", dbConfig)
	if nil != err {
		panic(err)
	}
	postgresDb.AutoMigrate([]interface{}{
		&UserInfo{},
	}...)
	postgresDb.Model(&UserInfo{}).RemoveIndex("uix_user_infos_account_id")
	postgresDb.Model(&UserInfo{}).AddIndex("idx_account_id", "account_id")

	return
}
