package server

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/vivek-yadav/mango/config"
	"github.com/vivek-yadav/mango/utils"
	"github.com/vivek-yadav/mango/version"
)

func Start(conf config.Config) {
	fconf := fiber.Config{}
	fconf.AppName = version.AppName()
	fconf.ServerHeader = version.AppName()

	app := fiber.New(fconf)
	app.Use(pprof.New())
	app.Use(requestid.New())
	app.Use(logger.New(getLogConfig(config.CurrentConfig)))
	if conf.Log.LogLevel <= utils.LogLevelDebug {
		app.Use(recover.New(recover.Config{
			EnableStackTrace: true,
		}))
	} else {
		app.Use(recover.New(recover.Config{
			EnableStackTrace: false,
		}))
	}

	app.Get("/metrics-dashboard", monitor.New(monitor.Config{Title: version.AppName() + " Metrics Page"}))
	prometheus := fiberprometheus.New(version.AppName())
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Get("/", func(c *fiber.Ctx) error {
		panic("failed panic")
		return c.SendString("Hello, World!")
	})
	utils.MangoLog.Info(app.Listen(fmt.Sprint("0.0.0.0:", conf.Serve.Port)))
}

func getLogConfig(config config.Config) (lc logger.Config) {
	lc = logger.Config{
		// For more options, see the Config section
		//Format:     "${green}${time}\t| ${pid}\t| ${locals:requestid}\t|${cyan} ${status}\t|${black} ${method}\t| ${bytesReceived}\t| ${reqHeaders}\t| ${path}\t| ${queryParams}\t| ${body}\t| ${latency}\t| RESPONSE:\t| ${bytesSent}\t| ${respHeader:Server}\t| ${resBody}\t| ${error} \n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Kolkata",
		Output:     os.Stdout,
	}
	if config.Log.LogLevel <= utils.LogLevelDebug {
		lc.Format = generateLogFormat(
			"${magenta}"+strings.ToUpper(config.Log.LogLevel.String())+"${reset}",
			"time", "pid", "locals:requestid", "status", "method", "path", "latency", "error", "bytesReceived", "bytesSent", "queryParams", "body", "resBody",
		)
	} else {
		lc.Format = generateLogFormat(
			"${green}"+strings.ToUpper(config.Log.LogLevel.String())+"${reset}",
			"time", "pid", "locals:requestid", "status", "method", "path", "latency", "error", "bytesReceived", "bytesSent",
		)
	}
	return lc
}

func generateLogFormat(prefix string, fields ...string) string {
	color := []string{
		"red", "green", "yellow", "blue", "magenta", "cyan", "white",
	}
	logFormat := "| " + prefix + " | "
	for i, v := range fields {
		c := color[i%7]
		logFormat += "${" + c + "}${" + v + "}${reset}  |  "
	}
	logFormat += "\n"
	return logFormat
}
