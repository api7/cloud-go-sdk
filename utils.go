package cloud

import (
	"net"
	"os"

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
