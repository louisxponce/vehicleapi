package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func getIp(r *http.Request) string {
	// Try to get from forwared first
	log.Println("Checking X-FORWARDED-FOR")
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.SplitSeq(ips, ",")
	for ip := range splitIps {
		log.Println("Parsing")
		netIp := net.ParseIP(ip)
		if netIp != nil {
			log.Println("Found. Returning")
			return ip
		}
	}

	// try x-real-ip
	log.Print("X-REAL-IP")
	ip := r.Header.Get("X-REAL-IP")
	netIp := net.ParseIP(ip)
	if netIp != nil {
		log.Println("Found. Returning")
		return ip
	}

	// If nothing found above, get from remote address
	log.Print("Trying Remote Address")
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("Not found...")
		return ""
	}

	log.Printf("Parsing ...")

	netIp = net.ParseIP(ip)
	if netIp != nil {
		log.Println("Found. Returning")
		return ip
	}

	return ""
}

func MyIpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking requester ip")
	ip := getIp(r)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"ip": "%s"}`, ip)
}
