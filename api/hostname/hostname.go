package hostname

import (
	"fmt"
	"os"
	"strings"
)

// OverwriteHostname overwrites default hostname name for docker container
func OverwriteHostname(hosts string) (string, error) {
	host, err := os.Hostname()
	if err != nil {
		return "", err
	}

	h := strings.Split(hosts, ",")
	for i, value := range h {
		if value == host {
			host = "host" + fmt.Sprint(i+1)
		}
	}

	return host, nil
}
