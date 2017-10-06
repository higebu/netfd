package netfd

import (
	"net"
	"testing"
)

func TestGetFdFromConn(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	fd := GetFdFromConn(conn)
	if fd <= 0 {
		t.Fatal("failed to get fd")
	}
	t.Logf("fd: %d", fd)
}

func TestGetFdFromListener(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	fd := GetFdFromListener(l)
	if fd <= 0 {
		t.Fatal("failed to get fd")
	}
	t.Logf("fd: %d", fd)
}
