package myhttp

import (
	"math/rand"
	"time"
)

type UA struct {
	pc      []string
	android []string
	ios     []string
	wechat  []string
}

var userAgent *UA

func GetUA() *UA {
	if userAgent == nil {
		agent := getAgent()
		userAgent = &UA{
			pc:      agent["pc"],
			android: agent["android"],
			ios:     agent["ios"],
			wechat:  agent["wechat"],
		}
	}

	return userAgent
}

func (ua *UA) getPcAgent() string {
	return ua.randAgent(ua.pc)
}

func (ua *UA) getMobileAgent() string {
	agent := append(ua.android, ua.ios...)
	return ua.randAgent(agent)
}

func (ua *UA) getWechatAgent() string {
	return ua.randAgent(ua.wechat)
}

func (ua *UA) randAgent(agent []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	agentKey := len(agent)
	return agent[r.Intn(agentKey)]
}

func getAgent() map[string][]string {
	agent := make(map[string][]string)
	agent["pc"] = []string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	agent["android"] = []string{
		"Mozilla/5.0 (Linux; U; Android 8.1.0; zh-CN; 16th Plus Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.108 UCBrowser/12.3.0.1010 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 10; ELE-AL00 Build/HUAWEIELE-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.20 SP-engine/2.16.0 baiduboxapp/11.20.0.10 (Baidu; P1 10)",
		"Mozilla/5.0 (Linux; Android 9; V1813BA Build/PKQ1.181030.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.19 SP-engine/2.15.0 baiduboxapp/11.19.5.10 (Baidu; P1 9)",
		"Mozilla/5.0 (Linux; Android 9; MI 8 Build/PKQ1.180729.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/045120 Mobile Safari/537.36 V1_AND_SQ_8.2.6_1320_YYB_D QQ/8.2.6.4370 NetType/WIFI WebP/0.3.0 Pixel/1080 StatusBarHeight/89 SimpleUISwitch/1",
		"Mozilla/5.0 (Linux; Android 8.1.0; HLTE510T Build/OPM1.171019.026; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.19 SP-engine/2.15.0 baiduboxapp/11.19.5.10 (Baidu; P1 8.1.0)",
		"Mozilla/5.0 (Linux; Android 9; V1829A Build/PKQ1.181030.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.19 SP-engine/2.15.0 baiduboxapp/11.19.5.10 (Baidu; P1 9)",
		"Mozilla/5.0 (Linux; Android 8.0.0; STF-AL10 Build/HUAWEISTF-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.19 SP-engine/2.15.0 baiduboxapp/11.19.5.10 (Baidu; P1 8.0.0) NABar/1.0",
		"Mozilla/5.0 (Linux; Android 5.1.1; vmos Build/LMY48G; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/045120 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 9; PACM00 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/045120 Mobile Safari/537.36 V1_AND_SQ_8.2.7_1334_YYB_D QQ/8.2.7.4410 NetType/WIFI WebP/0.3.0 Pixel/1080 StatusBarHeight/84 SimpleUISwitch/0",
	}

	agent["ios"] = []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16A405 MicroMessenger/7.0.10(0x17000a21) NetType/4G Language/zh_CN",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_5 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) Mobile/15D60 MicroMessenger/7.0.10(0x17000a21) NetType/4G Language/zh_CN",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16A405 MicroMessenger/7.0.10(0x17000a21) NetType/WIFI Language/zh_CN",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.10(0x17000a21) NetType/4G Language/zh_HK",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.11(0x17000b21) NetType/WIFI Language/zh_CN",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E302 MicroMessenger/7.0.10(0x17000a21) NetType/4G Language/zh_CN",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.11(0x17000b21) NetType/WIFI Language/zh_CN",
	}

	agent["wechat"] = agent["ios"]
	return agent
}
