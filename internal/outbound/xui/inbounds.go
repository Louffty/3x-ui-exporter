package outboundxui

import (
	"encoding/json"
	"fmt"
	"net/http"

	"xui-exporter/internal/domain"
)

func (x *XUIClient) GetInbounds() ([]domain.Inbound, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		x.config.PanelURL+"/xui/API/inbounds/",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		if err = x.Login(); err != nil {
			return nil, err
		}

		req, err = http.NewRequest(
			http.MethodGet,
			x.config.PanelURL+"/xui/API/inbounds/",
			nil,
		)
		if err != nil {
			return nil, err
		}

		resp, err = x.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("get inbounds failed with status %d", resp.StatusCode)
	}

	var apiResp domain.APIResponse[[]domain.Inbound]

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("get inbounds failed: %s", apiResp.Msg)
	}

	return apiResp.Obj, nil
}
