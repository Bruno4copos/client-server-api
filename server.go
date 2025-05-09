package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CotacaoAPI struct {
	USD struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY AUTOINCREMENT, bid TEXT)")
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctxAPI, cancelAPI := context.WithTimeout(r.Context(), 500*time.Millisecond)
		defer cancelAPI()

		req, err := http.NewRequestWithContext(ctxAPI, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			http.Error(w, "Erro criando requisição externa", http.StatusInternalServerError)
			log.Printf("Erro criando requisição externa: %v", err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Erro ao buscar cotação externa. "+err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao buscar cotação externa: %v", err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Erro lendo resposta externa", http.StatusInternalServerError)
			return
		}

		var cotacaoAPI CotacaoAPI
		if err := json.Unmarshal(body, &cotacaoAPI); err != nil {
			http.Error(w, "Erro ao decodificar cotação", http.StatusInternalServerError)
			return
		}

		bid := cotacaoAPI.USD.Bid

		ctxDB, cancelDB := context.WithTimeout(r.Context(), 10*time.Millisecond)
		defer cancelDB()

		err = salvarCotacao(ctxDB, db, bid)
		if err != nil {
			log.Printf("Erro ao salvar no banco: %v", err)
		}

		json.NewEncoder(w).Encode(Cotacao{Bid: bid})
	})

	log.Println("Servidor iniciado na porta 9090")
	http.ListenAndServe(":9090", nil)
}

func salvarCotacao(ctx context.Context, db *sql.DB, bid string) error {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes(bid) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, bid)
	return err
}
