package main

import (
	"fmt"
	//"log"
	"net/http"
	"strings"
	//"mytest"
	"time"
	"strconv"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Println("UnixTime:" + strconv.Itoa(int(time.Now().Unix())))
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Web!~")
}

func main() {
	//http.HandleFunc("/", sayHelloName)
	//http.HandleFunc("/a", mytest.TestA)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}

	t := time.Now().Unix()
	fmt.Println(t)
	fmt.Println(time.Now().Format("2006-01-02_15:04"))
	fmt.Println(time.Unix(t, 0).String())
	fmt.Println(time.Now().String())
	channel := map[string]bool{"i":true, "n":true}
	var channels []string
	for k := range channel {
		channels = append(channels, k)
	}
	fmt.Println(GetMinLastRangeIdx(channels))
}

func GetMinLastRangeIdx(channels []string) uint32 {
	// 获取每个渠道配置的最末档位
	minLastRangeIdx := 10000
	if channels == nil {
		minLastRangeIdx = 0
	} else {
		for _, v := range channels {
			channelPosRangesLen := len(GetChannelPosRanges(v))
			if channelPosRangesLen > 0 && minLastRangeIdx > channelPosRangesLen {
				minLastRangeIdx = channelPosRangesLen - 1
			}
		}
		if minLastRangeIdx == 10000 {
			minLastRangeIdx = 0
		}
	}

	return uint32(minLastRangeIdx)
}

// 获取指定渠道的位置配置信息
func GetChannelPosRanges(channel string) []string {
	PosRangesConf := map[string][]string{
		"i": {"0_9", "10_19", "20_29", "30_39", "40_49", "50_59"},
		"n": {"0_19", "20_39", "40_59", "60_79"},
	}
	fmt.Println(PosRangesConf[channel])
	return PosRangesConf[channel]
	//return settings.NetSessionConf().ChannelDspPosRanges[channel]
}
