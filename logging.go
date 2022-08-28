package main

import (
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
)

var url = "http://hoge.com"

func main() {
	// log standard package
	/*
		SetupLogging("sample.log")
		log.Println("log.Println")
		log.Fatalln("log.Fatalln")
		log.Println("log.Println")
	*/

	// zap pattern 1(Logger)
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Error("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// zap pattern 2(SugeredLogger)
	/*
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		sugar := logger.Sugar()
		sugar.Infow("failed to fetch URL",
			// Structured context as loosely typed key-value pairs.
			"url", url,
			"attempt", 3,
			"backoff", time.Second,
		)
		sugar.Infof("Failed to fetch URL: %s", url)
	*/
}

func SetupLogging(logFilePath string) {
	logFile, _ := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
