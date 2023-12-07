package tests_assert

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EqualAsJSON(t *testing.T, expected, actual interface{}) bool {
	t.Helper()

	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actualJSON, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}

	return assert.Equal(t, string(expectedJSON), string(actualJSON))
}
