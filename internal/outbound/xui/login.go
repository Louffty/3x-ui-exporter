package outboundxui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"xui-exporter/internal/domain"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (x *XUIClient) Login() error {
	body, err := json.Marshal(LoginRequest{
		Username: x.config.Username,
		Password: x.config.Password,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		x.config.PanelURL+"/login",
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := x.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("login failed with status %d", resp.StatusCode)
	}

	var apiResp domain.APIResponse[any]

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}

	if !apiResp.Success {
		return fmt.Errorf("login failed: %s", apiResp.Msg)
	}

	return nil
}
