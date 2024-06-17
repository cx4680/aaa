package web

import (
	"fmt"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/opentrx/seata-golang/v2/pkg/util/log"
	"guangfa-fund/internal/web/result"
	"net/http"
	"os"
)

type GinEngineRouterFunc func(engine *gin.Engine)

func InitGinEngine(routerFunc GinEngineRouterFunc) {
	gin.SetMode(stringsx.DefaultIfEmpty(os.Getenv(gin.EnvGinMode), "debug"))

	engine := gin.Default()

	// 如果中间件出现问题，返回500
	engine.Use(gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		log.Error(fmt.Errorf("%v", recovered), "")
		result.Failure(context, 500, "系统异常")
	}))
	// 主入口文件 注册 session 中间件
	store := cookie.NewStore([]byte("secret")) // 设置 Session 密钥
	engine.Use(sessions.Sessions("mysession", store))

	//探活
	engine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	if routerFunc != nil {
		routerFunc(engine)
	}

	port := stringsx.DefaultIfEmpty(os.Getenv("PORT"), "8080")
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Error(err, "start guangfa-fund server error")
	}
}
