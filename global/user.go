package global

import (
	"net"
)

// OnlineUsers 在线用户
var OnlineUsers = make(map[string]net.Conn)
