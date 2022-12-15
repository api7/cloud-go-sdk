package cloud

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func configureTokenFromFile(tokenPath string) (*AccessToken, error) {
	var content struct {
		User struct {
			AccessToken string `yaml:"access_token"`
		} `yaml:"user"`
	}

	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &content); err != nil {
		return nil, errors.Wrap(err, "invalid token file")
	}

	if content.User.AccessToken == "" {
		return nil, ErrEmptyToken
	}

	return &AccessToken{
		Token: content.User.AccessToken,
	}, nil
}

func getLocalIPs() ([]net.IP, error) {
	var ips []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP)
		}
	}
	return ips, nil
}

func sumIPs(ips []net.IP) uint16 {
	total := 0
	for _, ip := range ips {
		for i := range ip {
			total += int(ip[i])
		}
	}
	return uint16(total)
}

func mergePagination(paging *Pagination) Pagination {
	if paging != nil {
		if paging.Page == 0 {
			paging.Page = DefaultPagination.Page
		}
		if paging.PageSize == 0 {
			paging.PageSize = DefaultPagination.PageSize
		}
		return *paging
	} else {
		return DefaultPagination
	}
}

func formatJSONData(raw []byte) (string, error) {
	js, err := simplejson.NewJson(raw)
	if err != nil {
		return "", errors.Wrap(err, "invalid json")
	}

	for _, resName := range []string{"routes", "services", "upstreams", "certificates", "consumers"} {
		res, ok := js.CheckGet(resName)
		if !ok {
			continue
		}

		for i := 0; i < len(res.MustArray()); i++ {
			var structualValue map[string]interface{}
			value := res.GetIndex(i).Get("value").MustString()

			// value is a JSON string, and we want to show it structurally, so
			// here we unmarshal and reset it.
			if err := json.Unmarshal([]byte(value), &structualValue); err != nil {
				return "", errors.Wrap(err, fmt.Sprintf("unmarshal %s", value))
			}

			res.GetIndex(i).Set("value", structualValue)
		}
	}
	newData, err := js.MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(newData), nil
}
