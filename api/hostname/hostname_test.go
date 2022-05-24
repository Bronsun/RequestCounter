package hostname

import (
	"fmt"
	"os"
	"testing"
)

func TestOverwriteHostnameBasic(t *testing.T) {
	host, err := os.Hostname()
	fmt.Printf(host)
	if err != nil {
		t.Errorf("Something wrong with getting hostname")
	}
	ans := OverwriteHostname(host)
	if ans != "host0" {
		t.Errorf("OverwriteHostname(host) = %s; want host0", ans)
	}
}
