package library

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient() *APIClient {
	return &APIClient{
		baseURL: "https://library.kosgos.ru/api",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *APIClient) SearchBooks(query, author, year string) ([]ExternalBook, error) {
	params := url.Values{}
	if query != "" {
		params.Set("q", query)
	}
	if author != "" {
		params.Set("author", author)
	}
	if year != "" {
		params.Set("year", year)
	}

	url := fmt.Sprintf("%s/search?%s", c.baseURL, params.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []ExternalBook `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (c *APIClient) GetBookByID(id string) (*ExternalBook, error) {
	url := fmt.Sprintf("%s/books/%s", c.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var book ExternalBook
	if err := json.Unmarshal(body, &book); err != nil {
		return nil, err
	}

	return &book, nil
}
