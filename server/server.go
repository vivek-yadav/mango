package server

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${green}${time}\t| ${pid}\t| ${locals:requestid}\t|${cyan} ${status}\t|${black} ${method}\t| ${bytesReceived}\t| ${reqHeaders}\t| ${path}\t| ${queryParams}\t| ${body}\t| ${latency}\t| RESPONSE:\t| ${bytesSent}\t| ${respHeader:Server}\t| ${resBody}\t| ${error} \n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Kolkata",
		Output:     os.Stdout,
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: version.AppName() + " Metrics Page"}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	utils.MangoLog.Info(app.Listen(fmt.Sprint("0.0.0.0:", conf.Serve.Port)))
}
