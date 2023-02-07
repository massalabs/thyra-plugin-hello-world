package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

const ThyraRegisterEndpoint = "http://my.massa/plugin-manager/register"

type registerBody struct {
	ID          int64
	Name        string
	Author      string
	Description string
	Logo        string
	URL         string
	APISpec     string
	Home        string
}

func Register(
	pluginID int64,
	name string, author string,
	shortDescription string,
	socket net.Addr, spec string,
	logoURL string,
	home string,
) error {
	reg := registerBody{
		ID:          pluginID,
		Name:        name,
		Author:      author,
		Description: shortDescription,
		URL:         "http://" + socket.String(),
		APISpec:     spec,
		Logo:        logoURL,
		Home:        home,
	}

	body, err := json.Marshal(reg)
	if err != nil {
		return fmt.Errorf("while marshaling register body: %w", err)
	}

	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ThyraRegisterEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("while creating register request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("while doing register request: %w", err)
	}

	if response.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(response.Body)
		defer response.Body.Close()

		return fmt.Errorf("plugin registry failed: HTTP status: %d, HTTP body: %v", response.StatusCode, body)
	}

	return nil
}
