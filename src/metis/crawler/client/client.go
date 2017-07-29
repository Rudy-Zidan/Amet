package client

import (
	"log"
	"math/rand"

	resty "gopkg.in/resty.v0"
)

type Client struct{}

func (client *Client) Fetch(url string) (resp *resty.Response, errs error) {
	proxyIP := client.pickProxy()
	log.Println("Make a request using:", proxyIP)
	log.Println("Fetching :", url)
	resty.SetProxy("http://" + proxyIP + ":3128")
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	return resty.R().Get(url)
}

func (client *Client) pickProxy() string {
	proxyServers := []string{
		"149.56.89.166",
		"5.189.135.164",
		"92.222.217.116",
		"91.93.132.138",
		"46.101.99.191",
		"191.101.1.190",
		"163.172.86.64",
	}
	return proxyServers[rand.Intn(len(proxyServers))]
}

func (client *Client) pickUserAgent() string {
	agents := []string{
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3) Gecko/20090305 Firefox/3.1b3 GTB5",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; ko; rv:1.9.1b2) Gecko/20081201 Firefox/3.1b2",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2) Gecko/20100308 Ubuntu/10.04 (lucid) Firefox/3.6 GTB7.1",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/603.1.13 (KHTML, like Gecko) Version/10.1 Safari/603.1.13",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11) AppleWebKit/601.1.39 (KHTML, like Gecko) Version/9.0 Safari/601.1.39",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/600.3.10 (KHTML, like Gecko) Version/8.0.3 Safari/600.3.10",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/6.1.3 Safari/537.75.14",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/536.26.17 (KHTML, like Gecko) Version/6.0.2 Safari/536.26.17",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.82 Safari/537.36 Edge/14.14359",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.82 Safari/537.36 Edge/14.14359",
	}
	return agents[rand.Intn(len(agents))]
}
