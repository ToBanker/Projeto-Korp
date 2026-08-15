package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var (
	totalRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisições recebidas no endpoint /projeto-korp",
		},
	)
)

func init() {
	prometheus.MustRegister(totalRequests)
}

type RespostaKorp struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func korpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	totalRequests.Inc()

	resposta := RespostaKorp{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	http.HandleFunc("/projeto-korp", korpHandler)
	
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("🚀 Servidor Golang com Monitoramento rodando na porta 8080...")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}