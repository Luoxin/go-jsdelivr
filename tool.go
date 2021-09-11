package go_jsdelivr

import (
	"errors"
	"fmt"
	"github.com/elliotchance/pie/pie"
	"net/url"
	"strings"
)

var NotSupportedErr = errors.New("not support")

func Cover(urlRaw string) (string, error) {
	u, err := url.Parse(urlRaw)
	if err != nil {
		return "", err
	}

	switch u.Hostname() {
	case "raw.githubusercontent.com":
		nu := url.URL{
			Scheme: "https",
			Host:   "cdn.jsdelivr.net",
		}

		pi := pie.Strings(strings.Split(u.Path, "/"))

		nu.Path = fmt.Sprintf("/gh%s@%s", pi.Top(3).Join("/"), pi.DropTop(3).Join("/"))

		return nu.String(), nil
	default:
		return "", NotSupportedErr
	}
}
