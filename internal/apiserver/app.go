// -*- coding:utf-8 -*-
// @Time : 2022/3/15 1:16 上午
// @Author: MJ-CJM
// @File : go-template/app
package apiserver

import "go-template/pkg/app"

const commandDesc = `The API Server services REST operations to do the api objects management`

// NewApp creates a App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}