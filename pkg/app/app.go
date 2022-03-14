// -*- coding:utf-8 -*-
// @Time : 2022/3/15 1:33 上午
// @Author: MJ-CJM
// @File : go-template/app
package app

import (
	"github.com/spf13/cobra"
)

// App is the main structure of a cli application.
// It is recommended that an app be created with the app.NewApp() function.
type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	commands    []*Command
	args        cobra.PositionalArgs
	cmd         *cobra.Command
}
