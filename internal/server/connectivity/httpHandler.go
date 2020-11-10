package connectivity

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Ping web http handler for connectivity
func Ping(c *gin.Context) {
	svc := NewServer()
	res, err := svc.Ping()
	if err != nil {
		log.Errorf("[connectivity][httpHandler][Ping] svc.Ping error:%v", err)
		_ = c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}
