package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("	dP    dP  88888888b              dP     dP   .d888888   a88888b. dP     dP ")
	fmt.Println("	Y8.  .8P  88                     88     88  d8'    88  d8'   `88 88   .d8' ")
	fmt.Println("	 Y8aa8P  a88aaaa                 88aaaaa88a 88aaaaa88a 88        88aaa8P'  ")
	fmt.Println("	   88     88                     88     88  88     88  88        88   `8b. ")
	fmt.Println("	   88     88                     88     88  88     88  Y8.   .88 88     88 ")
	fmt.Println("	   dP     88888888P______________dP     dP  88     88   Y88888P' dP     dP ")
	fmt.Println("")
	fmt.Println("Coded By ye_hack && Rewrite in Golang by MAJHCC")
	fmt.Println("")
	tmpips := []string{""}
	var input string
	tmpsites := []string{""}
	tmpoutputsites := []string{""}
	fmt.Printf("Your List >")
	fmt.Scanln(&input)
	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Println("File Not Found")
		os.Exit(1)
	}
	lines, err := readLines(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(lines) == 0 {
		fmt.Println("File Empty")
		os.Exit(1)
	}

	for _, i := range lines {
		if stringInSlice(i, tmpsites) == false {
			tmpips = rev(i, tmpips, tmpoutputsites)
			tmpsites = append(tmpsites, i)
		} else {
			fmt.Println("PASSED : " + i)
		}
	}
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func revSo(ip string) (*[]string, error) {
	var links []string
	api := "https://sonar.omnisint.io/reverse/"
	useragent := "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0"
	client := &http.Client{}
	req, err := http.NewRequest("GET", api+ip, nil)
	if err != nil {
		return nil, errors.New("Error")
	}
	req.Header.Set("User-Agent", useragent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error")
	}
	if string(body) == "null" {
		return nil, errors.New("Error")
	} else {
		r, _ := regexp.Compile(`(?:(?:https?|ftp):\/\/)?[\w/\-?=%.]+\.[\w/\-&?=%.]+`)
		domains := r.FindAllString(string(body), -1)
		for _, i := range domains {
			replace := strings.Replace(i, "www.", "", -1)
			replace2 := strings.Replace(replace, "cpanel.", "", -1)
			replace3 := strings.Replace(replace2, "webmail.", "", -1)
			replace4 := strings.Replace(replace3, "webdisk.", "", -1)
			replace5 := strings.Replace(replace4, "ftp.", "", -1)
			replace6 := strings.Replace(replace5, "cpcalendars.", "", -1)
			replace7 := strings.Replace(replace6, "cpcontacts.", "", -1)
			replace8 := strings.Replace(replace7, "mail.", "", -1)
			replace9 := strings.Replace(replace8, "ns1.", "", -1)
			replace10 := strings.Replace(replace9, "ns2.", "", -1)
			replace11 := strings.Replace(replace10, "autodiscover.", "", -1)
			links = append(links, replace11)
		}
		return &links, nil
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func rev(url string, tmpips []string, tmpoutputsites []string) []string {
	var resultSite []string
	outputfile, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	url2 := strings.Replace(url, "http://", "", -1)
	url3 := strings.Replace(url2, "https://", "", -1)
	url4 := strings.Replace(url3, "\n.", "", -1)
	url5 := strings.Replace(url4, "\r", "", -1)
	url6 := strings.Replace(url5, "/", "", -1)
	ips, err := net.LookupIP(url6)
	if err != nil {
		fmt.Println("Error")
	} else {
		for _, i := range ips {
			ipv4vil, _ := regexp.MatchString(`\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b`, i.String())
			if ipv4vil {
				if stringInSlice(i.String(), tmpips) {
				} else {
					tmpips = append(tmpips, i.String())
					res, err := revSo(i.String())
					if err != nil {
					} else {
						for _, k := range *res {
							if stringInSlice(k, tmpoutputsites) == false {
								resultSite = append(resultSite, k)
								outputfile.WriteString(k + "\n")
								fmt.Print(strconv.Itoa(len(resultSite)), " ", "SITES", " ", url, "\r")
							}
						}
					}
				}
			}
		}
	}
	fmt.Print(strconv.Itoa(len(resultSite)), " ", "SITES", " ", url, "\r")
	fmt.Println("")
	return tmpips
}
