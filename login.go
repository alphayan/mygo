package main

import "github.com/labstack/echo"

// LoginData 登录数据结构体
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	l := LoginData{}
	r := Result{}
	c.Bind(&l)
	if l.Username == "alpha" && l.Password == "123456" {
		r.Code = "success"
		r.Message = "恭喜，登录成功"
		return c.JSON(200, r)
	}
	r.Code = "fail"
	r.Message = "用户名密码错误"
	return c.JSON(400, r)
}
