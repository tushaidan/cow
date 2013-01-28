package main

import (
	"net"
	"testing"
)

func TestCalcDigest(t *testing.T) {
	a1 := md5sum("cyf" + ":" + authRealm + ":" + "wlx")
	auth := map[string]string{
		"nonce":  "50ed159c3b707061418bbb14",
		"nc":     "00000001",
		"cnonce": "6c46874228c087eb",
		"uri":    "/",
	}
	const targetDigest = "bad1cb3526e4b257a62cda10f7c25aad"

	digest := calcRequestDigest(auth, a1, "GET")
	if digest != targetDigest {
		t.Errorf("authentication digest calculation wrong, got: %x, should be: %s\n", digest, targetDigest)
	}
}

func TestParseAllowedClient(t *testing.T) {
	parseAllowedClient("") // this should not cause error

	parseAllowedClient("192.168.1.1/16, 192.169.1.2")

	na := &auth.allowedClient[0]
	if !na.ip.Equal(net.ParseIP("192.168.0.0")) {
		t.Error("ParseAllowedClient 192.168.1.1/16 ip error, got ip:", na.ip)
	}
	mask := []byte(na.mask)
	if mask[0] != 0xff || mask[1] != 0xff || mask[2] != 0 || mask[3] != 0 {
		t.Error("ParseAllowedClient 192.168.1.1/16 mask error")
	}

	na = &auth.allowedClient[1]
	if !na.ip.Equal(net.ParseIP("192.169.1.2")) {
		t.Error("ParseAllowedClient 192.169.1.2 ip error")
	}
	mask = []byte(na.mask)
	if mask[0] != 0xff || mask[1] != 0xff || mask[2] != 0xff || mask[3] != 0xff {
		t.Error("ParseAllowedClient 192.169.1.2 mask error")
	}
}
