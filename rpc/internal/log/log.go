package log

import (
	"strings"

	mll "github.com/go-micro/plugins/v4/logger/logrus"
	mlog "github.com/jinmukeji/go-pkg/v2/log"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	ml "go-micro.dev/v4/logger"
)

const (
	defaultLogLevel = "INFO"
)

// MicroCliFlags 返回 micro cli 的 flags
func MicroCliFlags() []cli.Flag {
	return []cli.Flag{
		// 日志相关
		&cli.StringFlag{
			Name:    "log_format",
			Usage:   "Log format. Empty string or LOGSTASH.",
			EnvVars: []string{"LOG_FORMAT"},
			Value:   "",
		},

		&cli.StringFlag{
			Name:    "log_level",
			Usage:   "Log level. [TRACE, DEBUG, INFO, WARN, ERROR, PANIC, FATAL]",
			EnvVars: []string{"LOG_LEVEL", "MICRO_LOG_LEVEL"},
			Value:   defaultLogLevel,
		},
	}
}

// SetupLogger 设置 Logger
func SetupLogger(c *cli.Context, svc string) {
	std := mlog.StandardLogger()

	// Setup Log level
	lv := mlog.GetLevel()
	if logLevel := c.String("log_level"); logLevel != "" {
		if level, err := logrus.ParseLevel(logLevel); err != nil {
			std.Fatal(err.Error())
		} else {
			lv = level
			std.SetLevel(lv)
		}
	}
	std.Infof("Log Level: %s", lv)

	if logFormat := c.String("log_format"); strings.ToLower(logFormat) == "logstash" {
		// logstash 日志形式下注入 svc 字段，用来输出当前 service 的名称
		f := mlog.NewLogstashFormatter(logrus.Fields{
			"svc": svc,
		})

		std.SetFormatter(f)
	}

	ml.DefaultLogger = mll.NewLogger(
		mll.WithLogger(std),
	)
}
