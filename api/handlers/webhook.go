package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleWebhook(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Error(NewAPIError(http.StatusBadRequest, "failed to read request body", err.Error()))
		return
	}
	logrus.Infof("Body: %s", string(body))

	ctx.Status(http.StatusOK)
}
