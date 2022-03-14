// -*- coding:utf-8 -*-
// @Time : 2022/3/15 1:14 上午
// @Author: MJ-CJM
// @File : go-template/apiserver
package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"go-template/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("iam-apiserver").Run()
}

