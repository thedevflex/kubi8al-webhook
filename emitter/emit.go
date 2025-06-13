package emitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/thedevflex/kubi8al-webhook/model"
	logs "github.com/thedevflex/kubi8al-webhook/utils/logger"
)

type EmitterPayload struct {
	RepositoryName string `json:"repository_name"`
	FullName       string `json:"full_name"`
	Owner          string `json:"owner"`
	Private        bool   `json:"private"`
	CommitHash     string `json:"commit_hash"`
	EventType      string `json:"event_type"`
	Branch         string `json:"branch"`
}

func EmitWebhookPayload(payload model.ParsedWebHookPayload) error {
	numOfRetries := 5
	logs.Info("Emitting webhook payload to EMMITER_API_ADDRESS")

	emitPayload := EmitterPayload{
		RepositoryName: payload.Payload.Repository.Name,
		FullName:       payload.Payload.Repository.FullName,
		Owner:          payload.Payload.Repository.Owner.Login,
		Private:        payload.Payload.Repository.Private,
		CommitHash:     payload.Payload.HeadCommit.ID,
		EventType:      payload.Event,
		Branch:         payload.Payload.Ref,
	}

	jsonData, err := json.Marshal(emitPayload)
	if err != nil {
		logs.Error("Failed to marshal payload", err)
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	emitterAPI := os.Getenv("EMMITER_API_ADDRESS")
	if emitterAPI == "" {
		logs.Error("EMMITER_API_ADDRESS environment variable not set")
		return fmt.Errorf("EMMITER_API_ADDRESS environment variable not set")
	}
	for retryCount := numOfRetries; retryCount > 0; retryCount-- {
		resp, err := http.Post(emitterAPI, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			logs.Error("Failed to send payload to emitter API", err)
			if retryCount == 1 {
				return fmt.Errorf("failed to send payload to emitter API: %v", err)
			}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			logs.Info("Successfully emitted webhook payload")
			return nil
		}

		if resp.StatusCode >= 400 {
			logs.Errorf("Emitter API returned error status: %d", resp.StatusCode)
			return fmt.Errorf("emitter API returned error status: %d", resp.StatusCode)
		}
	}

	logs.Error("Failed to emit webhook payload")
	return fmt.Errorf("failed to emit webhook payload")
}
