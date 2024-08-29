package main

import "go-load-balancer/cmd/loadbalancer"

func main() {
	var loadbalancer loadbalancer.LoadBalancerServer = loadbalancer.LoadBalancerServer{
		Address: ":8080",
		TargetServers: []string{
			"",
			"",
			"",
		},
		Method: "rr",
	}

	loadbalancer.StartMainServerAndListen()
}
