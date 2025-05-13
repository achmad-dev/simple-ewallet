package pkg

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func InitLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// format for the file path and line number
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logger.SetFormatter(formatter)
	return logger
}
