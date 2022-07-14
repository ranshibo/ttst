/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-07-06 20:36:10
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-07-14 15:45:00
 * @FilePath: \beegoDemo\controllers\default.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controllers

import (
	_ "github.com/beego/beego/v2/client/cache/redis"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

type Result struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
type User struct {
	Id   int    `orm:"column(id);pk"`
	Name string `orm:"column(name)"`
}

func (c *MainController) GetTest() {
	c.Ctx.WriteString("hello")
}
