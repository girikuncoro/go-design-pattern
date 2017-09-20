package main

type Logger string

const (
	Info  Logger = "info"
	Debug        = "debug"
	Error        = "error"
)

type Log interface {
	write(message string)
}

type AbstractLogger struct {
	level      Logger
	nextLogger *AbstractLogger
}

func (self *AbstractLogger) SetNextLogger(nextLogger *AbstractLogger) {
	self.nextLogger = nextLogger
}

func (self *AbstractLogger) LogMessage(level Logger, message string) {
	if self.level <= level {
		self.write(message)
	}
	if self.nextLogger != nil {
		self.nextLogger.LogMessage(level, message)
	}
}

func (self *AbstractLogger) write(message string) {}

// TODO: implement pattern usage
func main() {

}
