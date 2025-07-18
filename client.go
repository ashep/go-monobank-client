package monobank

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.monobank.ua"
	requestTimeout = 10 * time.Second
)

type Client struct {
	token string
	cli   *http.Client
}

func New(token string) *Client {
	return &Client{
		token: token,
		cli:   http.DefaultClient,
	}
}

func (c *Client) GetCurrencyRates(ctx context.Context) ([]CurrencyRate, error) {
	url := baseURL + "/bank/currency"
	b, err := c.request(ctx, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	res := make([]CurrencyRate, 0)
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return res, nil
}

func (c *Client) GetCustomerInfo(ctx context.Context) (*Customer, error) {
	url := baseURL + "/personal/client-info"
	b, err := c.request(ctx, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	res := &Customer{}
	if err := json.Unmarshal(b, res); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return res, nil
}

func (c *Client) GetStatement(ctx context.Context, accID string, from, to time.Time) (Statement, error) {
	url := baseURL + fmt.Sprintf("/personal/statement/%s/%d/%d", accID, from.Unix(), to.Unix())
	b, err := c.request(ctx, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	res := Statement{}
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return res, nil
}

func (c *Client) request(ctx context.Context, method string, url string) ([]byte, error) {
	ctx, ctxC := context.WithTimeout(ctx, requestTimeout)
	defer ctxC()

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("X-Token", c.token)

	res, err := c.cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	if res.StatusCode == http.StatusTooManyRequests {
		return nil, ErrTooManyRequests
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d; %s", res.StatusCode, body)
	}

	return body, nil
}
