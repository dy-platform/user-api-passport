package router

import (
	"github.com/dy-platform/user-api-passport/idl"
	info "github.com/dy-platform/user-api-passport/idl/platform/user/srv-info"
	passport "github.com/dy-platform/user-api-passport/idl/platform/user/srv-passport"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/sirupsen/logrus"
)

type WeChatSignInForm struct {
	DeviceID int64  ` json:"deviceID"  `
	Code     string ` json:"code"   binding:"required"`
	AppID    string `json:"appID"   binding:"required"`
}

func WeChatSignIn(c *gin.Context) {
	var req WeChatSignInForm

	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, err)
		logrus.Errorf("gin.Context.Bind error:%v", err)
		return
	}

	cl := passport.NewPassportService("platform.user.srv.passport", client.DefaultClient)
	req1 := &passport.WeChatSignInReq{
		DeviceID: req.DeviceID,
		Code:     req.Code,
		AppID:    req.AppID,
	}

	rsp1, err := cl.WeChatSignIn(c.Request.Context(), req1)
	if err != nil {
		logrus.Warnf("platform.user.srv.passport WeChatSignIn error:%v", err)
		c.JSON(500, err)
		return
	}

	if rsp1.BaseResp.Code != int32(base.CODE_OK) {
		c.JSON(500, rsp1.BaseResp.Msg)
		return
	}

	//第一次登陆, 添加用户信息
	if rsp1.IsFirstSignIn {
		cl2 := info.NewUserInfoService("platform.user.srv.info", client.DefaultClient)
		req2 := &info.CreateUserReq{
			UserId:    rsp1.UserID,
			NickName:  "",
			Gender:    base.Gender_Unkonw,
			AvatarUrl: "",
		}
		_, err = cl2.CreateUser(c.Request.Context(), req2)
		if err != nil {
			logrus.Warnf("platform.user.srv.info CreateUser error:%v", err)
			c.JSON(500, err)
			return
		}
	}

	c.JSON(200, gin.H{
		"status": 200,
		"userID": rsp1.UserID,
	})
}
