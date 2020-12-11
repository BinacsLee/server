package commands

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/binacsgo/log"

	"github.com/BinacsLee/server/config"
)

var (
	configFile string
	zaplogger  *zap.Logger
	logger     log.Logger
	cfg        *config.Config
)

func init() {
	rootCmdFlags(RootCmd)
}

func rootCmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&configFile, "configFile", "config.toml", "config file (default is ./config.toml)")
}

var (
	// RootCmd the root command
	RootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root Command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if cmd.Name() != StartCmd.Name() {
				return nil
			}

			fmt.Println("LoadFromFile: ", configFile)
			if cfg, err = config.LoadFromFile(configFile); err != nil {
				fmt.Println("LoadFromFile err: ", err)
			}

			zaplogger = initLogger(cfg.WorkSpace, cfg.LogConfig)
			logger = log.NewZapLoggerWrapper(zaplogger.Sugar())
			logger.Info("Init finished")

			return nil
		},
	}
)

func initLogger(rootPath string, cfg config.LogConfig) *zap.Logger {
	logpath := cfg.File
	if !path.IsAbs(logpath) {
		logpath = path.Join(rootPath, cfg.File)
	}
	fmt.Printf("Log path : %s\n", logpath)
	hook := lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    cfg.Maxsize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.Maxage,
		Compress:   true,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "_time",
		LevelKey:       "_level",
		NameKey:        "_logger",
		CallerKey:      "_caller",
		MessageKey:     "_message",
		StacktraceKey:  "_stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(stringToXZapLoggerLevel(cfg.Level))
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&hook), atomicLevel)
	logger := zap.New(core)
	return logger
}

func stringToXZapLoggerLevel(level string) zapcore.Level {
	lower := strings.ToLower(level)
	switch lower {
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}
