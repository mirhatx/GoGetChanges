package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func check(url string) bool {
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		return true
	}
	return false
}

func has_changed(html *[]string, target *string, interval *int) bool {

	time.Sleep(time.Duration(*interval) * time.Second)
	resp, err := http.Get(*target)
	if err != nil {
		fmt.Println("an error occured.")
	}
	response, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println("error")
	}
	*html = append(*html, string(response))

	for j := 1; j < len(*html); j++ {
		if (*html)[0] != (*html)[j] {
			return true
		}
	}
	return false
}

func main() {
	html := []string{}
	flag.Parse()
	target := flag.Arg(0)
	interval := flag.Arg(1)
	if target == "" || interval == "" {
		fmt.Println("[+] Usage ./check hackerone.com 60, this will check every 60 to detect a change")
	}
	exists := check(target)
	fmt.Println("[+] Program will print out, if it detects any change and save it to a file")

	if exists == false {
		target = "https://" + target
	}
	req, err := http.Get(target)
	if err != nil {
		fmt.Println("there was an error. Please try again")
	}

	resp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("there was an error.")
	}
	finalResponse := string(resp)
	html = append(html, finalResponse)

	inter, errinter := strconv.Atoi(interval)

	if errinter != nil {
		fmt.Println("Error occured")
	}
	for {
		whater := has_changed(&html, &target, &inter)
		if whater == true {
			fmt.Println("Change is found!! Go to", html)
			ct := time.Now()
			new := ct.Format("01_02_2006_15:04_05")
			//	file, err := os.Create(new)
			//	defer file.Close()

			if err != nil {
				fmt.Println("error")
			}
			errr := ioutil.WriteFile(new, []byte(target), 0777)
			if errr != nil {
				fmt.Println("err")
			}
			// _, erri := file.Write([]byte(target))
			// if erri != nil {
			// 	fmt.Println("err")
			// }
			anErr := ioutil.WriteFile(new, []byte(html[1]), 0777)
			if anErr != nil {
				fmt.Println("error occured")
			}
		} else {
			continue
		}
	}
}
