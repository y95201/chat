/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-07-04 12:56:06
 * @LastEditors: Y95201
 * @LastEditTime: 2022-07-05 12:49:53
 */
package service

import (
	"errors"
	"fmt"
	"log"

	"weChat/model"

	"github.com/go-xorm/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "chat:123456@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动User
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}
