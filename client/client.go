package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deividroger/client-server-api-go/dto"
)

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

func GetCotation() (*dto.Cotation, error) {
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

	var cotation dto.Cotation
	err = json.NewDecoder(res.Body).Decode(&cotation)

	if err != nil {
		return nil, err
	}
	return &cotation, nil
}
