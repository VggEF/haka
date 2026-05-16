package schedule

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient() *APIClient {
	return &APIClient{
		baseURL: "https://eios.kosgos.ru",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *APIClient) GetSchedule(groupID int, date string) (*ExternalScheduleResponse, error) {
	url := fmt.Sprintf("%s/api/Rasp?idGroup=%d&sdate=%s", c.baseURL, groupID, date)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("client-version", "2026-04-10T12:06:09.286Z")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ExternalScheduleResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) GetGroupList() ([]Group, error) {
	url := fmt.Sprintf("%s/api/GetRaspGroups", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("client-version", "2026-04-10T12:06:09.286Z")

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
		Data  []Group `json:"data"`
		State int     `json:"state"`
		Msg   string  `json:"msg"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

type Group struct {
	Name      string `json:"name"`
	ID        int    `json:"id"`
	Kurs      int    `json:"kurs"`
	Facul     string `json:"facul"`
	FacultyID int    `json:"facultyID"`
}
