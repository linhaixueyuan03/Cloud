package model

import (
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/Cloud?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panicf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}
