package main

import (
	_ "testing-golang/routers"
	"github.com/astaxie/beego"
	"fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:root@/perpusku_gc?charset=utf8")
}

func main() {
	o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
	beego.Run()
}

