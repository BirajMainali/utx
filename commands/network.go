package commands

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

type NetworkCommand struct{}

func (n *NetworkCommand) Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: utx net <command>")
		return
	}

	switch args[0] {
	case "local-ip":
		getLocalIP()
	case "public-ip":
		getPublicIP()
	case "ping":
		if len(args) < 2 {
			fmt.Println("Usage: utx net ping <hostname/IP>")
			return
		}
		pingHost(args[1])
	case "dns":
		if len(args) < 2 {
			fmt.Println("Usage: utx net dns <domain>")
			return
		}
		dnsLookup(args[1])
	case "interfaces":
		listInterfaces()
	default:
		fmt.Println("Unknown network command")
	}
}

func getLocalIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting local IP:", err)
		return
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("Local IP:", ipNet.IP.String())
				return
			}
		}
	}
}

func getPublicIP() {
	cmd := exec.Command("curl", "-s", "https://api64.ipify.org")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error fetching public IP:", err)
		return
	}
	fmt.Println("Public IP:", strings.TrimSpace(string(output)))
}

func listOpenPorts() {
	fmt.Println("Feature not implemented yet")
}

func pingHost(host string) {
	cmd := exec.Command("ping", "-c", "4", host)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error pinging host:", err)
		return
	}
	fmt.Println(string(output))
}

func dnsLookup(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error resolving domain:", err)
		return
	}
	for _, ip := range ips {
		fmt.Println("IP:", ip.String())
	}
}

func listInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching network interfaces:", err)
		return
	}
	for _, iface := range interfaces {
		fmt.Println("Interface:", iface.Name)
	}
}
