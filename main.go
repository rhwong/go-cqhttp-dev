// Package main
package main

import (
	"github.com/rhwong/go-cqhttp-dev/cmd/gocq"
	"github.com/rhwong/go-cqhttp-dev/global/terminal"

	_ "github.com/rhwong/go-cqhttp-dev/db/leveldb"   // leveldb 数据库支持
	_ "github.com/rhwong/go-cqhttp-dev/modules/silk" // silk编码模块
	// 其他模块
	// _ "github.com/rhwong/go-cqhttp-dev/db/sqlite3"   // sqlite3 数据库支持
	// _ "github.com/rhwong/go-cqhttp-dev/db/mongodb"    // mongodb 数据库支持
	// _ "github.com/rhwong/go-cqhttp-dev/modules/pprof" // pprof 性能分析
)

func main() {
	terminal.SetTitle()
	gocq.InitBase()
	gocq.PrepareData()
	gocq.LoginInteract()
	_ = terminal.DisableQuickEdit()
	_ = terminal.EnableVT100()
	gocq.WaitSignal()
	_ = terminal.RestoreInputMode()
}
