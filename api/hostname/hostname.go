package hostname

import (
	"fmt"
	"os"
	"strings"
)

// OverwriteHostname overwrites default hostname name for docker container
func OverwriteHostname(hosts string) string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	h := strings.Split(hosts, ",")
	for i, value := range h {
		if value == host {
			host = "host" + fmt.Sprint(i)
		}
	}

	return host
}
