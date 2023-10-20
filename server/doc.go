// Package server 包含HTTP,WebSocket,反向WebSocket请求处理的相关函数与结构体
package server

import "github.com/rhwong/go-cqhttp-dev/modules/servers"

// 注册
func init() {
	servers.Register("http", runHTTP)
	servers.Register("ws", runWSServer)
	servers.Register("ws-reverse", runWSClient)
	servers.Register("lambda", runLambda)
}
