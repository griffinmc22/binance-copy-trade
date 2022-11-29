package general

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func HttpClient(selectedProxy string) (client *http.Client) {
	if selectedProxy == "" {
		client = &http.Client{}
		fmt.Println("No Proxies Found")
	} else {
		proxy, _ := url.Parse(selectedProxy)
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy), MaxIdleConnsPerHost: 1, DisableKeepAlives: true, TLSHandshakeTimeout: 700 * time.Millisecond}, Timeout: 5 * time.Second}
	}
	return client
}
func ImportProxy() (proxyList []string) {
	file, err := os.Open("./config/proxies.txt")
	if err != nil {
		log.Fatal("Error Importing Proxies, make sure proxies.txt exists")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		if len(s) == 2 {
			proxyList = append(proxyList, "http://"+s[0]+":"+s[1])
		} else if len(s) == 4 {
			proxyList = append(proxyList, "http://"+s[2]+":"+s[3]+"@"+s[0]+":"+s[1])
		}
	}
	return proxyList
}

func ImportUrl() (URLlist []string) {
	file, err := os.Open("./config/url.txt")
	if err != nil {
		log.Fatal("Error Importing URL, make sure url.txt exists")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		URLlist = append(URLlist, scanner.Text())
	}
	return URLlist
}

func RandomString(options []string) string {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int() % len(options)
	return options[randNum]
}
func timestamp() (timestamp string) {
	now := time.Now()
	nanos := now.UnixNano()
	milis := nanos / 1000000
	timestamp = strconv.Itoa(int(milis))
	return timestamp
}
