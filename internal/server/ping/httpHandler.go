package ping

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Ping web http handler for ping
func Ping(c *gin.Context) {
	svc := NewServer()
	res, err := svc.Ping()
	if err != nil {
		log.Errorf("[ping][httpHandler][Ping] svc.Ping error:%v", err)
	}

	c.JSON(http.StatusOK, res)
}
