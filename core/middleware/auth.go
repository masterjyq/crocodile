package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/labulaka521/crocodile/common/jwt"
	"github.com/labulaka521/crocodile/common/log"
	"github.com/labulaka521/crocodile/common/rediscli"
	"github.com/labulaka521/crocodile/core/config"
	"github.com/labulaka521/crocodile/core/model"
	"github.com/labulaka521/crocodile/core/utils/resp"
)

const (
	tokenpre = "Bearer "
)

// CheckToken check token is valid
func CheckToken(token string) (string, string, bool) {
	claims, err := jwt.ParseToken(token)
	if err != nil || claims.UID == "" {
		log.Error("ParseToken failed", zap.Error(err))
		return "", "", false
	}
	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		log.Error("Token is Expire", zap.String("token", token))
		return "", "", false
	}

	return claims.UID, claims.UserName, true
}

// CheckToken check token is valid
func CheckCloudToken(token string) (string, bool) {
	// todo 校验用户
	client := rediscli.GetRedisClient()
	val, err := client.Get("LOGGED_TOKEN_" + strings.ToUpper(token)).Result()
	if err != nil {
		log.Error("Check Cloud User Err", zap.Error(err))
		return "", false
	}
	// 解析用户。获取用户ID
	if val == "" {
		log.Error("Check Cloud User Redis Not Fund")
		return "", false
	}
	cloudUser := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(val), &cloudUser)
	if err != nil {
		log.Error("Convert Cloud User Err", zap.Error(err))
		return "", false
	}
	account, ok := cloudUser["account"]
	if !ok {
		log.Error("Convert Cloud not fund account")
		return "", false
	}
	return account.(string), true
}

// 权限检查
func checkAuth(c *gin.Context) (pass bool, err error) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), tokenpre)

	if token == "" {
		err = errors.New("invalid token")
		return
	}

	uid, username, pass := CheckToken(token)
	if !pass {
		// 校验是不是cloud的用户
		username, pass = CheckCloudToken(token)
		if !pass {
			log.Error("CheckToken failed")
			return false, errors.New("CheckToken failed")
		}
		// uid 查找
		sql := fmt.Sprintf("crocodile_user where name = '%s'", username)
		userArray, err := model.GetNameID(c, sql)
		if err != nil || len(userArray) == 0 {
			log.Error("crocodile not has user", zap.Error(err))
			return false, errors.New("crocodile not has user:" + username)
		}
		// 得到ID
		uid = userArray[0].Value
	}

	c.Set("uid", uid)
	c.Set("username", username)
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	ok, err := model.Check(ctx, model.TBUser, model.UID, uid)
	if err != nil {
		log.Error("Check failed", zap.Error(err))
		return false, err
	}
	if !ok {
		log.Error("Check UID not exist", zap.String("uid", uid))
		return false, nil
	}

	role, err := model.QueryUserRule(ctx, uid)
	if err != nil {
		log.Error("QueryUserRule failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	c.Set("role", role)

	requrl := c.Request.URL.Path
	method := c.Request.Method
	enforcer := model.GetEnforcer()
	return enforcer.Enforce(uid, requrl, method)
}

var excludepath = []string{"login", "logout", "install", "websocket"}

// PermissionControl 权限控制middle
func PermissionControl() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			code = resp.Success
			err  error
		)
		if c.Request.URL.Path == "/" {
			c.Next()
			return
		}
		for _, url := range excludepath {
			if strings.Contains(c.Request.URL.Path, url) {
				c.Next()
				return
			}
		}
		defer func() {
			c.Set("statuscode", code)
		}()

		pass, err := checkAuth(c)
		if err != nil {
			log.Error("checkAuth failed", zap.Error(err))
			code = resp.ErrUnauthorized
			goto ERR
		}
		if !pass {
			log.Error("checkAuth not pass ")
			code = resp.ErrUnauthorized
			goto ERR
		}

		c.Next()
		return

	ERR:
		// 解析失败返回错误
		c.Writer.Header().Add("WWW-Authenticate", fmt.Sprintf("Bearer realm='%s'", resp.GetMsg(code)))
		resp.JSON(c, resp.ErrUnauthorized, nil)
		c.Abort()
	}
}
