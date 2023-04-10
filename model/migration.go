package model

import (
	"fmt"
)

//执行数据迁移

func migration() {
	// 自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
	if err != nil {
		fmt.Println("register table fail", err)
	}
	fmt.Println("register table success")
}
