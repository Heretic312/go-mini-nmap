package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Heretic312/go-mini-nmap/config"
	"github.com/Heretic312/go-mini-nmap/scanner"
)

func main() {
	hosts := flag.String("hosts", "127.0.0.1", "Comma-separated hosts")
	startPort := flag.Int("start", config.DefaultStartPort, "Start port")
	endPort := flag.Int("end", config.DefaultEndPort, "End port")
	tcpOnly := flag.Bool("tcp", true, "Scan TCP")
	udpOnly := flag.Bool("udp", false, "Scan UDP")
	concurrency := flag.Int("concurrency", config.DefaultConcurrency, "Max concurrent scans")
	csvFile := flag.String("csv", config.DefaultCSVOutputFile, "CSV output file")
	jsonFile := flag.String("json", config.DefaultJSONOutputFile, "JSON output file")

	flag.Parse()

	hostList := strings.Split(*hosts, ",")

	fmt.Printf("Scanning hosts: %v, ports %d-%d\n", hostList, *startPort, *endPort)
	fmt.Printf("TCP: %v, UDP: %v, concurrency: %d\n", *tcpOnly, *udpOnly, *concurrency)

	var wg sync.WaitGroup
	resultsChan := make(chan scanner.ScanResult, (*endPort-*startPort+1)*len(hostList)*2)
	semaphore := make(chan bool, *concurrency)

	for _, host := range hostList {
		for port := *startPort; port <= *endPort; port++ {
			if *tcpOnly {
				wg.Add(1)
				go scanner.ScanTCP(host, port, 500*time.Millisecond, &wg, resultsChan, semaphore)
			}
			if *udpOnly {
				wg.Add(1)
				go scanner.ScanUDP(host, port, 500*time.Millisecond, &wg, resultsChan, semaphore)
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var allResults []scanner.ScanResult
	for res := range resultsChan {
		color := scanner.Green
		if res.Service == "Unknown" {
			color = scanner.Red
		}
		fmt.Printf("%s%s:%d [%s] - %s%s\n", color, res.Host, res.Port, res.Protocol, res.Service, scanner.Reset)
		allResults = append(allResults, res)
	}

	if err := scanner.SaveCSV(*csvFile, allResults); err != nil {
		log.Println("Error saving CSV:", err)
	}
	if err := scanner.SaveJSON(*jsonFile, allResults); err != nil {
		log.Println("Error saving JSON:", err)
	}

	fmt.Println(scanner.Yellow + "Scan complete. Results saved." + scanner.Reset)
}