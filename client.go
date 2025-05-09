package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:9090/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro na requisição ao servidor: %v", err)
		return
	}
	defer resp.Body.Close()
	// var v interface{}
	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Printf("Erro: %v", err)
	// 	return
	// }
	// err = json.Unmarshal(respBody, v)
	// if err != nil {
	// 	log.Printf("error decoding sakura response: %v", err)
	// 	if e, ok := err.(*json.SyntaxError); ok {
	// 		log.Printf("syntax error at byte offset %d", e.Offset)
	// 	}
	// 	log.Printf("sakura response: %q", respBody)
	// 	return
	// }
	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		log.Printf("Erro ao decodificar resposta: %v", err)
		return
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Printf("Erro ao criar arquivo: %v", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Dólar: " + cotacao.Bid)
	if err != nil {
		log.Printf("Erro ao escrever no arquivo: %v", err)
	}

	log.Printf("Cotação salva com sucesso: Dólar: %s", cotacao.Bid)
}
