package crawler

import (
	"fmt"
	"holder_contract/pkg/log"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	crawlClient http.Client
	tryLimit    int
)

func init() {
	crawlClient = http.Client{}
	tryLimit = 5
}

/*
Split dom html scan website
Return holders, logo, website, socials
*/
func CrawlHoldersETH(address string) (int64, error) {
	holders := int64(0)
	var req *http.Request
	req, err := http.NewRequest("GET", "https://etherscan.io/token/"+address, nil)
	if err != nil {
		log.Println(log.LogLevelError, "CrawlHoldersETH : http.NewRequestl ethereum"+address, err)
		return 0, err
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	req.Header.Add("Content-Type", "application/json")

	for i := 0; i < tryLimit; i++ {
		res, err := crawlClient.Do(req)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersETH : client.Do(req) ethereum "+address, err)
			return 0, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersETH : io.ReadAll(res.Body) ethereum "+address, err)
			return 0, err
		}

		domString := string(body)
		indexHolders := strings.Index(domString, `number of holders`)
		if indexHolders == -1 {
			fmt.Println("not index", address)
			return 0, nil
		}
		domString = domString[indexHolders+len("number of holders "):]
		indexFirstSpace := strings.Index(domString, " ")
		if indexHolders == -1 {
			time.Sleep(10 * time.Second)
			continue
		}
		holderHTML := domString[:indexFirstSpace]
		holderHTML = strings.ReplaceAll(holderHTML, ",", "")
		intVar, err := strconv.ParseInt(holderHTML, 0, 64)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersETH: strconv.ParseInt(holderHTML, 0, 64)", err)
			holders = 0
		}
		holders = intVar
		if holders != 0 {
			break
		}
	}
	return holders, nil
}

func CrawlHoldersBSC(address string) (int64, error) {
	holders := int64(0)
	client := &http.Client{}
	failCount := 10
	for {
		req, err := http.NewRequest("GET", "https://bscscan.com/token/"+address, nil)
		req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
		req.Header.Add("Content-Type", "application/json")

		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersBSC : http.NewRequestl "+address, err)
			failCount--
			time.Sleep(3 * time.Second)
			if failCount == 0 {
				return 0, err
			}
			continue
		}

		res, err := client.Do(req)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersBSC : client.Do(req) "+address, err)
			failCount--
			time.Sleep(3 * time.Second)
			if failCount == 0 {
				return 0, err
			}
			continue
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersBSC : io.ReadAll(res.Body) "+address, err)
			failCount--
			time.Sleep(3 * time.Second)
			if failCount == 0 {
				return 0, err
			}
			continue
		}

		domString := string(body)
		indexHolders := strings.Index(domString, `number of holders`)
		if indexHolders == -1 {
			failCount--
			time.Sleep(3 * time.Second)
			if failCount == 0 {
				return 0, err
			}
			continue
		}
		domString = domString[indexHolders+len("number of holders "):]
		indexFirstSpace := strings.Index(domString, " ")
		holderHTML := domString[:indexFirstSpace]
		holderHTML = strings.ReplaceAll(holderHTML, ",", "")
		intVar, err := strconv.ParseInt(holderHTML, 0, 64)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlHoldersBSC: strconv.ParseInt(holderHTML, 0, 64) "+holderHTML, err)
			return 0, err
		}
		holders = intVar
		return holders, nil
	}
}
