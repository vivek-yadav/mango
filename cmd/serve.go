/*
Copyright © 2022 Vivek Yadav <vivekyadav.jit@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vivek-yadav/mango/config"
	"github.com/vivek-yadav/mango/server"
	"github.com/vivek-yadav/mango/utils"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		fmt.Println(config.CurrentConfig.Show())
		utils.MangoLog.Debug("Test Debug")
		utils.MangoLog.Info("Test Info")
		utils.MangoLog.Warn("Test Warn")
		utils.MangoLog.Error("Test Error")
		// defer func() {
		// 	if err := recover(); err != nil {
		// 		utils.MangoLog.Error("panic occurred:", err)
		// 	}
		// }()
		// utils.CheckFatal(errors.New("Custom Error Check"))
		// utils.MangoLog.Fatal("Test Fatal-Error")
		server.Start(config.CurrentConfig)

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
