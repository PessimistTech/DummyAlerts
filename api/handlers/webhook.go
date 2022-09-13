package handlers

import (
	"DummyAlerts/config"
	"DummyAlerts/interpreters"
	"DummyAlerts/notifiers"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleWebhook(ctx *gin.Context) {
	interpreterStr := ctx.Param("interpreter")
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Error(NewAPIError(http.StatusBadRequest, "failed to read request body", err.Error()))
		return
	}
	logrus.Infof("Body: %s", string(body))

	go interpretBody(body, interpreterStr)

	ctx.Status(http.StatusOK)
}

func interpretBody(body []byte, interpreterStr string) {
	interpreter, err := interpreters.GetInterpreter(interpreterStr)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get interpreter")
		return
	}

	message, err := interpreter.Interpret(body)
	if err != nil {
		logrus.WithError(err).Errorf("failed to interpret message")
		return
	}

	cfg := config.GetConfig()

	logrus.Infof("Message: %+v", message)
	for name := range cfg.Notifiers {
		notifier, err := notifiers.GetNotifier(name)
		if err != nil {
			logrus.WithError(err).Errorf("failed to get notifier")
		}
		err = notifier.Notify(message)
		if err != nil {
			logrus.WithError(err).Errorf("failed to send notification")
		}
	}
}
