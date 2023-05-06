package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n\nEnter domain: ")
	for scanner.Scan() {
		domain := scanner.Text()
		if strings.Contains(domain, "http://") {
			domain = strings.Replace(domain, "http://", "", -1)
		}
		if strings.Contains(domain, "https://") {
			domain = strings.Replace(domain, "https://", "", -1)
		}
		checkDomain(domain)
		fmt.Print("\n\nEnter domain: ")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkDomain(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

	MXs, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, MX := range MXs {
		fmt.Println("MX Records:\n")
		fmt.Println(MX.Host)
	}

	TXTs, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, TXT := range TXTs {
		fmt.Println("SPF Records:\n")
		if strings.Contains(TXT, "v=spf1") {
			fmt.Println(TXT)
		}
	}

	NSs, err := net.LookupNS(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, NS := range NSs {
		fmt.Println("NS Records:\n")
		fmt.Println(NS.Host)
	}

	

}