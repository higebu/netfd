// +build !go1.9

package netfd

import (
	"net"
	"reflect"
)

// GetFdFromConn get net.Conn's file descriptor.
func GetFdFromConn(l net.Conn) int {
	v := reflect.ValueOf(l)
	netFD := reflect.Indirect(reflect.Indirect(v).FieldByName("fd"))
	fd := int(netFD.FieldByName("sysfd").Int())
	return fd
}

// GetFdFromListener get net.Listener's file descriptor.
func GetFdFromListener(l net.Listener) int {
	v := reflect.ValueOf(l)
	netFD := reflect.Indirect(reflect.Indirect(v).FieldByName("fd"))
	fd := int(netFD.FieldByName("sysfd").Int())
	return fd
}
