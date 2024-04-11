package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Struct to hold IP and its count
type ipCount struct {
	IP    string
	Count int
}

// Function to read CIDR ranges from a file
func readCIDRRanges(filename string) ([]*net.IPNet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cidrRanges []*net.IPNet
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cidr := strings.TrimSpace(scanner.Text()) // Trim whitespace
		_, cidrNet, err := net.ParseCIDR(cidr)
		if err != nil {
			fmt.Printf("Error parsing CIDR '%s': %v\n", cidr, err)
			continue
		}
		cidrRanges = append(cidrRanges, cidrNet)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cidrRanges, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: contextifier <log_file> [cidr_file]")
		os.Exit(1)
	}

	logFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}
	defer logFile.Close()

	var cidrRanges []*net.IPNet
	if len(os.Args) > 2 {
		cidrRanges, err = readCIDRRanges(os.Args[2])
		if err != nil {
			fmt.Println("Error reading CIDR file:", err)
			os.Exit(1)
		}
	}

	ipv4Regex := regexp.MustCompile(`\b(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]|[0-9])\.(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]|[0-9])\.(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]|[0-9])\.(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]|[0-9])\b`)
	ipv6Regex := regexp.MustCompile(`\b([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\b`)

	ipMap := make(map[string]int)
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		line := scanner.Text()
		ips := ipv4Regex.FindAllString(line, -1)
		ips = append(ips, ipv6Regex.FindAllString(line, -1)...)

		for _, ip := range ips {
			ipAddr := net.ParseIP(ip)
			if ipAddr == nil {
				continue
			}

			excluded := false
			for _, cidr := range cidrRanges {
				if cidr.Contains(ipAddr) {
					excluded = true
					break
				}
			}

			if !excluded {
				ipMap[ip]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading log file:", err)
		os.Exit(1)
	}

	// Transfer data to slice and sort
	var counts []ipCount
	for ip, count := range ipMap {
		counts = append(counts, ipCount{IP: ip, Count: count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	// Output the sorted counts
	for _, count := range counts {
		fmt.Printf("%-20s %-10d https://wafris.org/ip-lookup/%s\n", count.IP, count.Count, count.IP)
	}

	// Print total unique IPs
	fmt.Printf("\nTotal unique IP addresses: %d\n", len(ipMap))
}
