package version

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	VERSION  = "1.0.1"
	BUILD    = "2022-06-25T14:57:06+0530"
	APP_NAME = "Mango"
)

// Version returns the version of the app
func Version() string {
	return VERSION
}

func Build() (string, time.Time) {
	buildTime, err := time.Parse("2006-01-02T15:04:05-0700", BUILD)
	cobra.CheckErr(err)
	return BUILD, buildTime
}

func AppName() string {
	return APP_NAME
}
