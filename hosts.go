// hosts

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var platform string
var checkBegin, _ = regexp.Compile("<pre>")
var checkEnd, _ = regexp.Compile("</pre>")

func detectPlatform() string {
	//not implemented, fake windows for now
	return "windows"
}
func main() {
	//starting, prompt something
	fmt.Println("Doing 事儿")
	resp, err := http.Get("http://goo.gl/hbhlU")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	test, _ := os.Create("resp.Body")
	defer test.Close()
	test.Write(body)
	//this section checks hosts info embedded within
	//a <pre> block, and grab the starting- and ending
	//index of the response body.
	begin := checkBegin.FindIndex(body)
	end := checkEnd.FindIndex(body)
	iBegin := begin[1]
	iEnd := end[0]

	hosts := body[iBegin:iEnd]
	//test hosts
	file, err := os.Create("parsed.Content")
	defer file.Close()
	file.Write(hosts)

	if platform != "nix" {
		syshosts, err := os.Open("%systemroot%/drivers/etc/hosts")
		if err != nil {
			log.Fatal(err)
		}
		syshosts.Seek(1, 2)
		syshosts.Write(hosts)
	} else {
		syshosts, err := os.Open("/etc/hosts")
		if err != nil {
			log.Fatal(err)
		}
		syshosts.Seek(1, 2)
		syshosts.Write(hosts)
	}
}
