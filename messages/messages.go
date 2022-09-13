package messages

const (
	INFO  = "INFO"
	ERROR = "ERROR"
)

type Message struct {
	Source          string
	Interpreter     string
	Content         string
	Title           string
	Level           string
	Event           string
	CustomStructure interface{}
}
