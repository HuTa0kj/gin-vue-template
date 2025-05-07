package global

import (
	"time"

	"gintemplate/app/config"
)

var JwtKey []byte
var EffectiveHours time.Duration
var RoutePrefix string

func InitGlobal() {
	JwtKey = []byte(config.ConfigInfo.Auth.JwtKey)
	EffectiveHours = config.ConfigInfo.Auth.EffectiveHours
	RoutePrefix = config.ConfigInfo.Server.Route.Prefix
}
