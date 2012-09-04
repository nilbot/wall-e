// smarthosts

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

var platform string
var checkBegin, _ = regexp.Compile("#SmartHosts START")
var checkEnd, _ = regexp.Compile("#SmartHosts END")

func detectPlatform() string {
	return runtime.GOOS
}

//this function update the system hosts with smarthosts
func updateHosts(syshosts *os.File, content []byte) {
	var replace bool = false
	// var oldSmart []byte
	var iBegin int
	buffer := make([]byte, 1024*1024)

	bufferRead, readOldHostError := syshosts.Read(buffer)
	if bufferRead != 0 && readOldHostError != nil {
		log.Fatal(readOldHostError)
	}
	begin := checkBegin.FindIndex(buffer)
	end := checkEnd.FindIndex(buffer)

	//if hosts already contains smarthosts content
	if len(begin) != 0 && len(end) != 0 {
		iBegin = begin[0]
		// iEnd = end[1]
		replace = true
		// oldSmart = buffer[iBegin:iEnd]

	}

	if replace == false {
		fmt.Println("Your current hosts doesn't contain SmartHosts data, appending data now... \n 您现有的hosts还没有翻墙数据，程序接下来准备添加这些内容...")
		syshosts.Seek(0, 2) // append at end
		_, writeError := syshosts.Write(content)
		if writeError != nil {
			log.Fatal(writeError)
		}

	} else {
		fmt.Println("Your current already contains (old )SmartHosts data, replacing data now... \n 您现有的hosts已有（旧）翻墙数据，程序接下来准备替换这些内容...")
		syshosts.Seek(int64(iBegin), 0)
		_, writeError := syshosts.Write(content)
		if writeError != nil {
			log.Fatal(writeError)
		}

	}
}
func main() {
	//starting, prompt something
	fmt.Println("Getting hosts from smarthosts \n 程序运行中，正在下载最新版hosts文件...")
	resp, err := http.Get("http://smarthosts.googlecode.com/svn/trunk/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//doesn't perform regex error check because it has to contain #SmartHosts block
	body = body[checkBegin.FindIndex(body)[0]:checkEnd.FindIndex(body)[1]]

	switch detectPlatform() {
	case "windows":
		systemroot := os.Getenv("SYSTEMROOT")
		syshosts, err := os.OpenFile(systemroot+"/System32/drivers/etc/hosts", os.O_RDWR, 0640)
		if err != nil {
			log.Fatal(err)
		}
		defer syshosts.Close()
		updateHosts(syshosts, body)

	case "linux":
		syshosts, err := os.OpenFile("/etc/hosts", os.O_RDWR, 0640)
		if err != nil {
			log.Fatal(err)
		}
		defer syshosts.Close()
		updateHosts(syshosts, body)
	case "darwin":
		syshosts, err := os.OpenFile("/private/etc/hosts", os.O_RDWR, 0640)
		if err != nil {
			log.Fatal(err)
		}
		defer syshosts.Close()
		updateHosts(syshosts, body)
	}
	fmt.Println("Hosts has been updated. Restart your computer (Network Service) and have fun \n Hosts已被成功更新，重新启动你的计算机然后……开始开心吧")
}
