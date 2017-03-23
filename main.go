package main

import "fmt"
import "reflect"

// Version is the current version of Goventus
var (
	version string
)

func main() {
	fmt.Println("Goventus", version)

	Subscribe(LogNothing, Log)
	Notify(Log, LogEvent{Message: "Message 1"})
	Notify(Log, LogEvent{Message: "Message 2"})
}

// LogEvent is an event type for logging
type LogEvent struct {
	Message string
}

// LogNothing literally logs nothing
func LogNothing(e Event) {
	fmt.Println("ValueOf", reflect.ValueOf(e))

	if e, ok := e.(LogEvent); ok {
		fmt.Println(e.Message)
	}
}
