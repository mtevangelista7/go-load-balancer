package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type LoadBalancerServer struct {
	Address       string   // endereço e porta que recebem todas as requests
	TargetServers []string // lista com os servidores a receberem as requests
	Method        string   // apenas round-robin por enquanto
	Count         int
}

func (loadbalancer *LoadBalancerServer) StartMainServerAndListen() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ao entrar aqui já temos uma requisição
		targetServer := loadbalancer.roundRobin()

		fmt.Println(targetServer)

		if targetServer == "" {
			http.Error(w, "Servidor não encontrado", http.StatusNotFound)
			return
		}

		destinationURL, err := url.Parse(targetServer)
		if err != nil {
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
			return
		}

		// TODO: talvez implementar o próprio proxy reverso
		proxy := httputil.NewSingleHostReverseProxy(destinationURL)

		// Modifica a requisição para o proxy
		r.Host = destinationURL.Host

		// Encaminha a requisição para o servidor de destino
		proxy.ServeHTTP(w, r)
	})

	// Inicie o servidor após definir o handler
	if err := http.ListenAndServe(loadbalancer.Address, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %s", err)
	}

	// isso aqui parece estar no lugar errado
	loadbalancer.Count = -1
}

// roundRobin devolve em qual servidor a proxima requisição deve ser encaminhada
func (loadbalancer *LoadBalancerServer) roundRobin() string {
	loadbalancer.Count++
	if loadbalancer.Count > len(loadbalancer.TargetServers) {
		loadbalancer.Count = 0
	}
	return loadbalancer.TargetServers[loadbalancer.Count]
}
