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
	ans, _ := OverwriteHostname(host)
	if ans != "host1" {
		t.Errorf("OverwriteHostname(host) = %s; want host1", ans)
	}
}
