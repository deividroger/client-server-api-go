package common

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func GetCotation[T interface{}](ctx context.Context, url string) (*T, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

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

	var c T

	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return &c, err
}
