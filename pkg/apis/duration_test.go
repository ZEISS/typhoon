package apis

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringDuration(t *testing.T) {
	in := Duration(5*time.Minute + 10*time.Second)

	expect, got := "5m10s", in.String()
	assert.Equal(t, expect, got)
}

func TestMarshalDuration(t *testing.T) {
	const input = Duration(5*time.Minute + 10*time.Second)
	const expectOutput = `"5m10s"`

	b, err := json.Marshal(input)
	assert.NoError(t, err)
	assert.Equal(t, expectOutput, string(b))
}

func TestUnmarshalDuration(t *testing.T) {
	testCases := []struct {
		name              string
		input             string
		expectOutput      Duration
		expectErrContains string
	}{{
		name:         "Simple duration",
		input:        `"5m"`,
		expectOutput: Duration(5 * time.Minute),
	}, {
		name:         "Complex duration",
		input:        `"5m10s"`,
		expectOutput: Duration(5*time.Minute + 10*time.Second),
	}, {
		name:              "Missing unit",
		input:             `"5"`,
		expectErrContains: "missing unit in duration",
	}, {
		name:              "Invalid unit",
		input:             `"42kg"`,
		expectErrContains: "unknown unit",
	}, {
		name:              "Empty string",
		input:             `""`,
		expectErrContains: "invalid duration",
	}, {
		name:              "Not a JSON string",
		input:             "5m",
		expectErrContains: "invalid character",
	}}

	for _, tc := range testCases {
		//nolint:scopelint
		t.Run(tc.name, func(t *testing.T) {
			var d Duration
			err := json.Unmarshal([]byte(tc.input), &d)

			assert.Equal(t, tc.expectOutput, d)

			if errStr := tc.expectErrContains; errStr != "" {
				assert.Contains(t, err.Error(), errStr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
