package interpreters

import (
	"DummyAlerts/messages"
	"encoding/json"
	"fmt"
)

type XOInterpreter struct {
	SupportedEvents []string
}

type XOEvent struct {
	CallID    string                 `json:"callId"`
	Method    string                 `json:"method"`
	Type      string                 `json:"type"`
	Error     XOError                `json:"error,omitempty"`
	Result    bool                   `json:"result,omitempty"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Timestamp int64                  `json:"timestamp"`
	UserID    string                 `json:"userId,omitempty"`
	UserName  string                 `json:"userName,omitempty"`
}

type XOError struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Stack   string `json:"stack"`
}

func (xe *XOError) Error() string {
	return xe.Message
}

func NewXOInterpreter() *XOInterpreter {
	return &XOInterpreter{
		SupportedEvents: []string{},
	}
}

func (xi *XOInterpreter) Interpret(input []byte) (*messages.Message, error) {
	var event XOEvent
	event.Result = true
	err := json.Unmarshal(input, &event)
	if err != nil {
		return &messages.Message{}, err
	}
	message := &messages.Message{
		Source:      "Xen Orchestra",
		Interpreter: "xo",
	}

	switch event.Method {
	case "backupNg.runJob":
		return handleBackupRunJob(event, message)
	default:
		return nil, fmt.Errorf("unhandled event method: %s", event.Method)
	}
}

func handleBackupRunJob(event XOEvent, message *messages.Message) (*messages.Message, error) {
	jsonStr, err := json.Marshal(event.Params)
	if err != nil {
		return nil, err
	}
	message.Content = string(jsonStr)
	message.Event = event.Method

	if event.Type == "pre" {
		message.Title = "**Backup job started**"
		message.Level = messages.INFO
	}

	result := "succeeded"
	if !event.Result {
		result = "failed"
		message.Level = messages.ERROR
		message.Content = fmt.Sprintf("%s ```%s```", event.Error.Message, message.Content)
	}
	if event.Type == "post" {
		message.Title = fmt.Sprintf("**Backup job %s**", result)
	}

	return message, nil
}
