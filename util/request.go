package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)


func HandRequest(req *http.Request) *gjson.Result {
	jsonString := RequestJSON(req)
	//jsonResult := gjson.Result
	if gjson.Valid(jsonString) {
		jsonResult := gjson.Parse(jsonString)
		return &jsonResult
	}
	return nil
}

// RequestJSON 直接获取 请求参数是 JSON 的字符串
func RequestJSON(req *http.Request) string {
	if req != nil {
		vars := mux.Vars(req)
		if vars != nil && len(vars) != 0  {
			return MapToJsonString(vars)
		} else {
			result, err := ioutil.ReadAll(req.Body)	
			if err == nil {
				return string(result)
			}
		}
		
	}
	return ""
}

func MapToJsonString(m map[string]string) string{
	buf := bytes.Buffer{}
	for k, v := range m {
		buf.WriteString(fmt.Sprintf("%q:%q,",k,v))
	}	
	return "{"+strings.TrimRight(buf.String(), ",") + "}"
}

// GetRequestJSON 注释说明
func GetRequestJSON(req *http.Request) string {
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") || strings.Contains(contentType, "text/plain") {
		return RequestJSON(req)
	} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return RequestFormToJSON(req)
	} else if strings.Contains(contentType, "multipart/form-data") {
		return RequestMultipartFormToJSON(req)
	} else {
		return RequestJSON(req)
	}
}



// RequestFormToJSON 把 x-www-form-urlencoded 键值对 格式的请求转换为 JSON 的字符串
func RequestFormToJSON(req *http.Request) string {
	if req == nil {
		return ""
	}

	buf := bytes.Buffer{}
	req.ParseForm()
	for k, v := range req.PostForm {
		if nil != v && len(v) > 0 {
			buf.WriteString(fmt.Sprintf(`"%s":%v,`, k, v[0]))
		}
	}
	return `{` + strings.TrimRight(buf.String(), ",") + `}`
}

// RequestMultipartFormToJSON 把 multipart/form-data 格式的请求转换为 JSON 的字符串
func RequestMultipartFormToJSON(req *http.Request) string {
	if req == nil {
		return ""
	}

	req.ParseMultipartForm(32 << 20)
	values := req.MultipartForm
	if values != nil {
		return ""
	}

	buf := bytes.Buffer{}
	for k, v := range values.Value {
		if len(v) > 0 && v[0] != "" {
			buf.WriteString(fmt.Sprintf(`"%s":%v,`, k, v[0]))
		}
	}

	for k, v := range values.File {
		if len(v) > 0 && v[0] != nil {
			mutipartFile, err := v[0].Open()
			defer mutipartFile.Close()
			if err == nil {
				bytArray, err := ioutil.ReadAll(mutipartFile)
				if err == nil {
					buf.WriteString(fmt.Sprintf(`"%s":%v,`, k, bytArray))
				}
			}
		}
	}

	return `{` + strings.TrimRight(buf.String(), ",") + `}`
}

//GenURLQueryString 把字典参数 参数化，如：url?a=1&b=2&c=3
func GenURLQueryString(url string, m map[string]string) string {
	return url + "?" + MapToQueryString(m)
}

// MapToQueryString 把 map 格式参数转换为 url query 格式字符串，如 a=1&b=2&c=3
func MapToQueryString(m map[string]string) string {
	buf := bytes.Buffer{}
	for k, v := range m {
		if v != "" {
			buf.WriteString(k + "=" + v + "&")
		}
	}
	return strings.TrimRight(buf.String(), "&")
}

//GetRemoteAddr 获取客户端的ip地址，兼容 nginx 代理的
func GetRemoteAddr(req *http.Request) (remoteAddr string) {
	remoteAddr = req.Header.Get("Remote_addr")
	if len(remoteAddr) == 0 {
		remoteAddr = strings.Split(strings.Replace(req.RemoteAddr, "[::1]", "127.0.0.1", 1), ":")[0]
	}

	return
}

//GetDomain 获取自身的域名
func GetDomain(req *http.Request) string {
	scheme := req.Header.Get("X-Scheme")
	if scheme == "" {
		scheme = "http"
	}
	host := req.Header.Get("Host")
	if host == "" {
		host = req.Host
	}
	return scheme + "://" + host
}

// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "127.0.0.1"
}
