package cloud

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"strings"
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

func FormatTraceSeries(data *TraceSeries) string {
	req := data.Request
	output := strings.Builder{}
	output.WriteString(fmt.Sprintf("[%v] Send a %v request to %v \n", data.ID, req.Method, req.URL.String()))
	output.WriteString(fmt.Sprintf("[%v] With request header:  %v \n", data.ID, req.Header))

	if len(data.RequestBody) != 0 {
		output.WriteString(fmt.Sprintf("[%v] Send a request body: %s\n", data.ID, string(data.RequestBody)))
	}

	output.WriteString(fmt.Sprintf("[%v] Receive a response with status: %v\n", data.ID, data.Response.StatusCode))
	if len(data.ResponseBody) != 0 {
		output.WriteString(fmt.Sprintf("[%v] Receive a response body: %s\n", data.ID, string(data.ResponseBody)))
	}

	evts := data.Events
	output.WriteString(fmt.Sprintf("[%v] Dump %d events:\n", data.ID, len(evts)))
	for i, evt := range evts {
		output.WriteString(fmt.Sprintf("[%v] Event#%d %s : %s\n", data.ID, i, evt.HappenedAt.Format("2006-01-02 15:04:05"), evt.Message))
	}

	return output.String()
}

func ensureClusterID(s StoreInterface, opts ResourceCommonOpts) bool {
	if s != nil && s.GetGlobalClusterID() > 0 {
		return true
	}

	if opts == nil || opts.GetCluster() == nil || opts.GetCluster().ID <= 0 {
		return false
	}

	return true
}
