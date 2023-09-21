package get_port_token

//package get_port_token

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"czw_lol_query_tools/lcu"
)

var (
	cli = http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
)

const (
	apiUrlFmt = "https://riot:%s@127.0.0.1:%d"
)

var (
	portFlag = flag.Int("port", 8098, "lcu 代理端口")
)

// 根据启动客户端的命令参数确定客户端获取的令牌和port，得到真正的能够获取数据的网址。
func Return_port_token() string {

	flag.Parse()
	//port := *portFlag
	lcuPort, lcuToken, err := lcu.GetLolClientApiInfo()

	if err != nil {

		log.Fatal(err)
	}
	proxyURL := fmt.Sprintf(apiUrlFmt, lcuToken, lcuPort)
	go func() { //不明白为什么要一直进行循环呢？是为了预防服务器临时卡顿吗
		ticker := time.NewTicker(time.Second * 3)
		for { //这是个会一直执行的循环
			<-ticker.C
			var lcuPort, lcuToken, err = lcu.GetLolClientApiInfo()

			if err != nil {
				continue
			}
			updateProxyURL := fmt.Sprintf(apiUrlFmt, lcuToken, lcuPort)

			if updateProxyURL == proxyURL {
				continue
			}
			proxyURL = updateProxyURL
			//log.Println("update lcu:", proxyURL)
		}
	}()
	//log.Printf("listen on :%d, lcu api:%s\n", port, proxyURL)

	return proxyURL

}
