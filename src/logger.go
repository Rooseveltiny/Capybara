package capybara

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var CapyLogger *logrus.Logger

func InitLogger() {

	// init logger instance to use it across an application
	f, err := os.OpenFile("system_logging.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "capybara.log")
		panic(err)
	}
	CapyLogger = &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\r\n",
		},
	}
}

func init() {
	InitLogger()
}
