package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deividroger/client-server-api-go/dto"
	"github.com/deividroger/client-server-api-go/internal/common"
)

func ClientInit() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	cotation, err := common.GetCotation[dto.Cotation](ctx, "http://localhost:8080")

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
