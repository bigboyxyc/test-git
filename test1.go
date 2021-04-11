package main

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func init(){
	db,err := gorm.Open("mysql","bigboom:XYC123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=local")
	if err != nil{
		panic(err)
	}
	db.SingularTable(true)
}

type Like struct {
	ID			int  	`gorm:"primary_key"`
	IP    		string  `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua    		string  `gorm:"type:varchar(256);not null;"`
	Title       string  `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash  		uint64	`gorm:"type:unique_index:hash_idx;"`
	CreatedAt	time.Time
}