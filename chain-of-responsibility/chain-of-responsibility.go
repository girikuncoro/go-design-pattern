package main

import "fmt"

type Logger interface {
	write(message string) string
	LogMessage(logger Logger, message string)
}

type AbstractLogger struct {
	Logger
	level      int
	nextLogger Logger
}

func (self *AbstractLogger) SetNextLogger(nextLogger Logger) {
	self.nextLogger = nextLogger
}

func (self *AbstractLogger) LogMessage(logger Logger, message string) {
	fmt.Println(logger.write(message))
	if self.nextLogger != nil {
		self.nextLogger.LogMessage(self.nextLogger, message)
	}
}

type ConsoleLogger struct {
	*AbstractLogger
}

func (self *ConsoleLogger) write(message string) string {
	return "Standard Console::Logger: " + message
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{&AbstractLogger{}}
}

type ErrorLogger struct {
	*AbstractLogger
}

func (self *ErrorLogger) write(message string) string {
	return "Error Console::Logger: " + message
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{&AbstractLogger{}}
}

type FileLogger struct {
	*AbstractLogger
}

func (self *FileLogger) write(message string) string {
	return "File::Logger: " + message
}

func NewFileLogger() *FileLogger {
	return &FileLogger{&AbstractLogger{}}
}

func main() {
	errorLogger := NewErrorLogger()
	fileLogger := NewFileLogger()
	consoleLogger := NewConsoleLogger()
	errorLogger.SetNextLogger(fileLogger)
	fileLogger.SetNextLogger(consoleLogger)

	consoleLogger.LogMessage(consoleLogger, "This is an information")
	fileLogger.LogMessage(fileLogger, "This is a debug level information")
	errorLogger.LogMessage(errorLogger, "This is an error information")
}
