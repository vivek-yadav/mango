package version

import (
	"time"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.1"
const BUILD_TIMESTAMP = "2022-06-25T14:57:06+0530"
const APP_NAME = "Mango"

// Version returns the version of the app
func Version() string {
	return VERSION
}

func BuildTimestamp() (string, time.Time) {
	buildTime, err := time.Parse("2006-01-02T15:04:05-0700", BUILD_TIMESTAMP)
	cobra.CheckErr(err)
	return BUILD_TIMESTAMP, buildTime
}

func AppName() string {
	return APP_NAME
}
