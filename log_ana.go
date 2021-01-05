package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// 读取Log

func Readlog(filename string) ([]string, []string, []string, []string, []string, []string, []string) {

	var IP_list []string
	var URI []string
	var Status_code []string
	var Byte_size []string
	var UserAgent []string
	var req_time []string
	var req_host []string
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buff_readfile := bufio.NewReader(file)
	for {
		ctx, _, err := buff_readfile.ReadLine()
		ctx_list := strings.Split(string(ctx), " ")

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		IP_list = append(IP_list, ctx_list[0])
		URI = append(URI, string(ctx_list[6]))
		Status_code = append(Status_code, string(ctx_list[8]))
		Byte_size = append(Byte_size, string(ctx_list[9]))
		UserAgent = append(UserAgent, string(ctx_list[11]))
		req_time = append(req_time, string(ctx_list[13]))
		req_host = append(req_host, string(ctx_list[14]))

		// fmt.Println(IP_list)

	}
	return IP_list, URI, Status_code, Byte_size, UserAgent, req_time, req_host
}

// 格式化表格
func Table_Format(Format_type string, count map[string]int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{Format_type, "Count"})
	for k, v := range count {
		table.Append([]string{k, strconv.Itoa(v)})
	}
	table.Render()
}

// 统计资源出现次数
func Count_resource_number(auto_list []string, auto_map map[string]int) {
	for _, v := range auto_list {
		if _, ok := auto_map[v]; !ok {
			auto_map[v] = 1
		} else {
			auto_map[v]++
		}
	}

}

// 解析IP地址
func resolver_IP(IP_list []string) {
	ip_count := make(map[string]int)
	Count_resource_number(IP_list, ip_count)
	Table_Format("IP", ip_count)
}

// 解析状态码
func resolver_Code(Status_code []string) {
	code_count := make(map[string]int)
	Count_resource_number(Status_code, code_count)
	Table_Format("Code", code_count)
}

func resolver_URI(URI_list []string) {
	uri_count := make(map[string]int)
	Count_resource_number(URI_list, uri_count)
	Table_Format("URI", uri_count)
}

func resolver_UserAgent(UserAgent []string) {
	useragent_count := make(map[string]int)
	Count_resource_number(UserAgent, useragent_count)
	Table_Format("UserAgent", useragent_count)
}

func resolver_Host(req_host []string) {
	req_host_count := make(map[string]int)
	Count_resource_number(req_host, req_host_count)
	Table_Format("Domain", req_host_count)
}

func P_format() {
	fmt.Println()
	fmt.Println("***********Menu info**********")
	fmt.Println()
	fmt.Println()
	fmt.Println(strings.Repeat("*", 60))
}

func main() {

	var (
		filename string
		opt      int
	)

	// filename := "lvscheck.xitong.qihoo.net-access.log"
	// IP_list, URI, Status_code, Byte_size, UserAgent, req_time, req_host := Readlog(filename)
	// IP_list, URI_list, Status_code, _, UserAgent, _, req_host := Readlog(filename)
	usage := `
		1. output Ip list count
		2. output statuscode count
		3. output URI count
		4. output useragent count
		5. output Host count
		6. output 1-5 all
		7. exit
	`
	flag.StringVar(&filename, "filename", "", "input filename")
	flag.IntVar(&opt, "opt", 0, usage)
	flag.Parse()

	flag.Usage = func() {
		flag.PrintDefaults()
		return
	}

	if filename == "" && opt == 0 {
		flag.Usage()
		os.Exit(1)
	}

	IP_list, URI_list, Status_code, _, UserAgent, _, req_host := Readlog(filename)

	switch opt {
	case 1:
		P_format()
		resolver_IP(IP_list)
	case 2:
		P_format()
		resolver_Code(Status_code)
	case 3:
		P_format()
		resolver_URI(URI_list)
	case 4:
		P_format()
		resolver_UserAgent(UserAgent)
	case 5:
		P_format()
		resolver_Host(req_host)
	case 6:
		P_format()
		resolver_IP(IP_list)
		resolver_Code(Status_code)
		resolver_URI(URI_list)
		resolver_UserAgent(UserAgent)
		resolver_Host(req_host)
	case 7:
		os.Exit(1)
	default:
		fmt.Println("Parameter not matched")
	}

}
