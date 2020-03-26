package ipgetter

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func init()  {
	rand.Seed(time.Now().Unix())
}

var serverList = []string{
	"http://ip.dnsexit.com",
	"http://ifconfig.me/ip",
	"http://ipecho.net/plain",
	"http://checkip.dyndns.org/plain",
	"http://ipogre.com/linux.php",
	"http://whatismyipaddress.com/",
	"http://ip.my-proxy.com/",
	"http://websiteipaddress.com/WhatIsMyIp",
	"http://getmyipaddress.org/",
	"http://showmyipaddress.com/",
	"http://www.my-ip-address.net/",
	"http://myexternalip.com/raw",
	"http://www.canyouseeme.org/",
	"http://www.trackip.net/",
	"http://myip.dnsdynamic.org/",
	"http://icanhazip.com/",
	"http://www.iplocation.net/",
	"http://www.howtofindmyipaddress.com/",
	"http://www.ipchicken.com/",
	"http://whatsmyip.net/",
	"http://www.ip-adress.com/",
	"http://checkmyip.com/",
	"http://www.tracemyip.org/",
	"http://checkmyip.net/",
	"http://www.lawrencegoetz.com/programs/ipinfo/",
	"http://www.findmyip.co/",
	"http://ip-lookup.net/",
	"http://www.dslreports.com/whois",
	"http://www.mon-ip.com/en/my-ip/",
	"http://www.myip.ru",
	"http://ipgoat.com/",
	"http://www.myipnumber.com/my-ip-address.asp",
	"http://www.whatsmyipaddress.net/",
	"http://formyip.com/",
	"https://check.torproject.org/",
	"http://www.displaymyip.com/",
	"http://www.bobborst.com/tools/whatsmyip/",
	"http://www.geoiptool.com/",
	"https://www.whatsmydns.net/whats-my-ip-address.html",
	"https://www.privateinternetaccess.com/pages/whats-my-ip/",
	"http://checkip.dyndns.com/",
	"http://myexternalip.com/",
	"http://www.ip-adress.eu/",
	"http://www.infosniper.net/",
	"http://wtfismyip.com/",
	"http://ipinfo.io/",
	"http://httpbin.org/ip",
}

func fetch(url string) []byte {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:70.0) Gecko/20100101 Firefox/70.0")
	response, err := client.Do(request)
	if err != nil {
		return nil
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body
}

func find(b []byte) (string, error) {
	re, _ := regexp.Compile("(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)")
	ip := string(re.Find(b))
	if ip == "" {
		return "", errors.New("not find ip")
	}
	return ip, nil
}

func random(s []string) []string {
	for i := len(s) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		s[i],s[num] = s[num],s[i]
	}
	return s
}

func Myip() string {
	s := random(serverList)
	myip := make(chan string)
	for _, url := range s {
		go func(url string) {
			body := fetch(url)
			ip, err := find(body)
			if err != nil {
				return
			}
			myip <- ip
		}(url)
	}
	return <-myip
}
