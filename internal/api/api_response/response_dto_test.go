package api_response_test

import (
	"encoding/json"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"github.com/stretchr/testify/assert"
)

func TestResponseDTO_ToJSON(t *testing.T) {
	t.Run("Marshal valid JSON", func(t *testing.T) {
		data := map[string]interface{}{
			"name": "John",
			"age":  30,
		}
		response := api_response.Success(data)
		expected, _ := json.Marshal(response)

		result, err := response.ToJSON()

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Marshal invalid JSON", func(t *testing.T) {
		data := make(chan int)
		_, err := api_response.Success(data).ToJSON()
		assert.Error(t, err)
	})
}
