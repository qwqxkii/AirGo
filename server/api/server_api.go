package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 主题配置
func GetThemeConfig(ctx *gin.Context) {
	response.OK("GetThemeConfig success", global.Theme, ctx)
}

// 更新主题
func UpdateThemeConfig(ctx *gin.Context) {
	var theme model.Theme
	err := ctx.ShouldBind(&theme)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("UpdateThemeConfig error:"+err.Error(), nil, ctx)
		return
	}
	err = service.CommonSqlSave[model.Theme](theme)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("UpdateThemeConfig error:"+err.Error(), nil, ctx)
		return
	}
	//重新加载主题设置
	global.Theme = theme
	response.OK("UpdateThemeConfig success", nil, ctx)
}

// 获取系统设置
func GetSetting(ctx *gin.Context) {
	res, _, err := service.CommonSqlFind[model.Server, string, model.Server]("id = 1")
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetSetting error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetSetting success", res, ctx)
}

// 获取公共系统设置
func GetPublicSetting(ctx *gin.Context) {
	var ps = model.PublicSystem{
		EnableRegister:          global.Server.Subscribe.EnableRegister,
		EnableEmailCode:         global.Server.Subscribe.EnableEmailCode,
		EnableLoginEmailCode:    global.Server.Subscribe.EnableLoginEmailCode,
		RebateRate:              global.Server.Subscribe.RebateRate,
		BackendUrl:              global.Server.Subscribe.BackendUrl,
		EnabledClockIn:          global.Server.Subscribe.EnabledClockIn,
		AcceptableEmailSuffixes: global.Server.Subscribe.AcceptableEmailSuffixes,
	}
	response.OK("GetPublicSetting success", ps, ctx)
}

// 更新系统设置
func UpdateSetting(ctx *gin.Context) {
	var setting model.Server
	err := ctx.ShouldBind(&setting)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("UpdateSetting error:"+err.Error(), nil, ctx)
		return
	}
	service.Show(setting)
	err = service.UpdateSetting(&setting)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("UpdateSetting error:"+err.Error(), nil, ctx)
		return
	}

	response.OK("UpdateSetting success", nil, ctx)
}
