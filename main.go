package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/mdns"
)

func main() {

	addr := flag.String("addr", "127.0.0.1", "addr")
	mac := flag.String("mac", "30:d5:3e:4f:a7:dc", "mac")
	port := flag.Int("port", 62078, "port")
	hostname := flag.String("hostname", "iphone.local.", "hostname")
	domain := flag.String("domain", "local.", "domain")
	servicename := flag.String("service", "_apple-mobdev2._tcp", "setvice")
	flag.Parse()

	// Setup our service export
	instance := fmt.Sprintf("%s@supportsRP", *mac)
	info := []string{"user define service"}

	fmt.Printf("instance:[%s]\nservicename:[%s]\ndomain:[%s]\nhostname:[%s]\nport:[%d]\nip:[%s]\n", instance, *servicename, *domain, *hostname, *port, *addr)

	ip := net.ParseIP(*addr)
	svc, _ := mdns.NewMDNSService(instance, *servicename, *domain, *hostname, *port, []net.IP{ip}, info)

	fmt.Println(svc)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: svc})

	defer server.Shutdown()
	time.Sleep(time.Second * 10)
}
