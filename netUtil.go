package client

import (
	"io"
	"net/http"
	"net/url"
)

// Get 网络请求
func (c *Client) Get(interfaceId string, apiURL string, params url.Values) (rs []byte, err error) {
	// 如果启用了网关则使用网关地址
	if c.UseGateway {
		apiURL = c.GatewayHost
	}
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		LogErr(interfaceId, "解析url错误", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()

	// 创建一个具有自定义请求头的http.Client
	client := &http.Client{}

	// 创建一个HTTP请求对象并设置请求方法、URL和其他头部信息
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		LogErr(interfaceId, "创建GET类型请求对象错误", err)
		return nil, err
	}
	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", interfaceId)
	req.Header = headers

	res, err := client.Do(req)
	if err != nil {
		LogErr(interfaceId, "GET类型请求失败", err)
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Post 网络请求
func (c *Client) Post(interfaceId string, apiURL string, params url.Values) (rs []byte, err error) {
	// 如果启用了网关则使用网关地址
	if c.UseGateway {
		apiURL = c.GatewayHost
	}
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		LogErr(interfaceId, "解析url错误", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()

	// 创建一个具有自定义请求头的http.Client
	client := &http.Client{}

	// 创建一个HTTP请求对象并设置请求方法、URL和其他头部信息
	req, err := http.NewRequest("POST", Url.String(), nil)
	if err != nil {
		LogErr(interfaceId, "创建POST类型请求对象错误", err)
		return nil, err
	}
	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", interfaceId)
	req.Header = headers
	req.Header.Add("Content-Type", "application-json")

	res, err := client.Do(req)
	if err != nil {
		LogErr(interfaceId, "POST类型请求失败", err)
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
