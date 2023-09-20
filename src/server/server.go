package main

import (
	"client-server-api-go/dto"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	CreateDatabaseStructure()
	handler := func(w http.ResponseWriter, r *http.Request) {

		cotation, err := GetCotation()

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Falha ao obter a cotação"))
			return
		}

		err = StorageCotation(cotation)

		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)

			w.Write([]byte("Falha ao salvar a cotação"))
			log.Print(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cotation)

	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func GetCotation() (*dto.Cotation, error) {

	ctx, _cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer _cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var c dto.Cotation

	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return &c, err
}

func StorageCotation(cotation *dto.Cotation) error {
	ctx, _cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer _cancel()

	db, err := sql.Open("sqlite3", "./cotation.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	smt, err := db.Prepare("INSERT INTO cotation (bid) VALUES ($1)")

	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.ExecContext(ctx, cotation.Usdbrl.Bid)

	if err != nil {
		return err
	}

	return nil
}

func CreateDatabaseStructure() {
	db, err := sql.Open("sqlite3", "./cotation.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	smt, err := db.Prepare("CREATE TABLE IF NOT EXISTS cotation (bid TEXT)")

	if err != nil {
		panic(err)
	}
	defer smt.Close()

	_, err = smt.Exec()

	if err != nil {
		panic(err)
	}
}
