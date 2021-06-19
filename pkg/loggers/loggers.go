package loggers

// to import use
// logs "loggers"
// sample usage
// logs.Warning("This is a Warning")
// logs.Info("This is an Info")
// logs.Error("This is an Error")

import (
	"log"
	"os"
)

var (
	warningLogger = log.New(os.Stderr, "WARNING: ", log.Lmsgprefix|log.LstdFlags|log.Llongfile)
	infoLogger = log.New(os.Stderr, "INFO: ", log.Lmsgprefix|log.LstdFlags|log.Llongfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Lmsgprefix|log.LstdFlags|log.Llongfile)
	Warning = warningLogger.Print
	Info = infoLogger.Print
	Error = errorLogger.Print
)
