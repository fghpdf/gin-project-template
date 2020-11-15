package connectivity

import (
	"context"
	"fghpdf.me/gin-project-template/internal/pkg/ctxType"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type header struct {
	RequestId string `header:"X-Request-Id"`
}

// Ping web http handler for connectivity
func Ping(c *gin.Context) {
	h := &header{}

	if err := c.BindHeader(h); err != nil {
		log.Errorf("[connectivity][Ping] bind header error:%v", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("http.default_timeout"))*time.Millisecond)
	defer cancel()
	ctx = ctxType.WithRequestId(ctx, h.RequestId)

	svc := NewServer()
	res, err := svc.Ping(ctx)
	if err != nil {
		log.Errorf("[connectivity][Ping] svc.Ping error:%v", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, res)
}
