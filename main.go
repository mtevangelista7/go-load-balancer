package main

import "go-load-balancer/cmd/loadbalancer"

func main() {
	var loadbalancer loadbalancer.LoadBalancerServer = loadbalancer.LoadBalancerServer{
		Address: ":8080",
		TargetServers: []string{
			"https://www.google.com/",
			"https://copilot.microsoft.com/",
			"https://www.youtube.com/",
		},
		Method: "rr",
	}

	loadbalancer.StartMainServerAndListen()
}
