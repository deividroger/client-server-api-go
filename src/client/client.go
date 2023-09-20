package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotation struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {

	cotation, err := GetCotation()

	if err != nil {
		log.Print(err)
		return
	}

	file, err := os.Create("cotacao.txt")

	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("Dólar: %v", cotation.Usdbrl.Bid))

}

func GetCotation() (*Cotation, error) {
	ctx, _cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer _cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var cotation Cotation
	err = json.NewDecoder(res.Body).Decode(&cotation)

	if err != nil {
		return nil, err
	}
	return &cotation, nil
}
