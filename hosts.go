// hosts
package main

import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var checkBegin, _ = regexp.Compile("<pre>")
var checkEnd, _ = regexp.Compile("</pre>")

func main() {

	resp, err := http.Get("http://goo.gl/hbhlU")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	test, _ := os.Create("test")
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
	file, err := os.Create("hTest")
	defer file.Close()
	file.Write(hosts)

}
