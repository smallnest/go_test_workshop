package s2

import (
	"net"
	"testing"
)

func testConn(t *testing.T) net.Listener {
	t.Helper()
	// ln, err := net.Listen("tcp", "1.2.3.4:0")
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	return ln
}

func TestDial(t *testing.T) {
	ln := testConn(t)
	defer ln.Close()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("want dial success but get %v", err)
	}

	err = conn.Close()
	if err != nil {
		t.Fatalf("want close success but get %v", err)
	}
}
