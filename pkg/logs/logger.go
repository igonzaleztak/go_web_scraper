package logs

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"intelygenz/pkg/config"
	"intelygenz/pkg/enums"
	"os"
)

func InitLogger() error {
	pe := zap.NewProductionEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level := zap.InfoLevel
	if config.CmdFlags.Verbose == enums.VerboseModeLog {
		level = zap.DebugLevel
	}

	fileName := "record.log"
	pathToLog := fmt.Sprintf("%s/%s", config.AppConfig.LogsPath, fileName)

	// create the log folder if it does not exist
	if _, err := os.Stat(pathToLog); os.IsNotExist(err) {
		os.Mkdir(config.AppConfig.LogsPath, os.ModePerm)
	}

	f, err := os.OpenFile(pathToLog, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	Logger = zap.New(core).Sugar()
	return nil
}

var Logger *zap.SugaredLogger
