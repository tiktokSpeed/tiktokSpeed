package initialize

import (
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	GlobalLogger *log.Logger
)

// InitLogger to init logrus
func InitLogger() {
	// Customizable output directory.
	logFilePath := consts.HlogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}
	//全局日志
	file, _ := os.OpenFile("globalogfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	multiWriter := io.MultiWriter(os.Stdout, file)
	GlobalLogger = log.New(multiWriter, "MYAPP: ", log.Ldate|log.Ltime|log.Lshortfile)
	GlobalLogger.Println("This is a log message")

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := hertzlogrus.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	if runtime.GOOS == "linux" {
		logger.SetOutput(lumberjackLogger)
	}
	logger.SetLevel(hlog.LevelDebug)

	hlog.SetLogger(logger)
}
