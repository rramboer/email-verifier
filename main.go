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
		if strings.Contains(domain, "www.") {
			domain = strings.Replace(domain, "www.", "", -1)
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
	fmt.Println("\nIP Addresses:")
	for _, ip := range ips {
		fmt.Println(ip)
	}

	MXs, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nMX Records:")
	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}

	TXTs, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nSPF Records:")
	for _, TXT := range TXTs {
		if strings.Contains(TXT, "v=spf1") {
			fmt.Println(TXT)
			break
		}
	}

	NSs, err := net.LookupNS(domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nNS Records:")
	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}

	dmarc, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nDMARC Record:")
	for _, dmarc := range dmarc {
		if strings.Contains(dmarc, "v=DMARC1") {
			fmt.Println(dmarc)
			break
		}
	}
}