package convert

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToMap(t *testing.T) {
	testCases := []struct {
		path   string
		value  string
		result map[string]interface{}
	}{
		{
			path:  "foo.bar",
			value: "",
			result: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "",
				},
			},
		},
		{
			path:  "foo.[0].bar",
			value: "",
			result: map[string]interface{}{
				"foo": map[string]interface{}{
					"": []interface{}{
						map[string]interface{}{
							"bar": "",
						},
					},
				},
			},
		},
		{
			path:  "[1].foo",
			value: "bar",
			result: map[string]interface{}{
				"": []interface{}{
					nil,
					map[string]interface{}{
						"foo": "bar",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			assert.Equal(t, tc.result, SliceToMap(strings.Split(tc.path, "."), tc.value))
		})
	}
}

func TestMergeJSONWithMap(t *testing.T) {
	testCases := []struct {
		source   string
		appendix string
		result   interface{}
	}{
		{
			source:   `{"old":"value"}`,
			appendix: "foo.bar",
			result: map[string]interface{}{
				"old": "value",
				"foo": "bar",
			},
		}, {
			source:   `{"old":"value"}`,
			appendix: "foo.bar[1].baz",
			result: map[string]interface{}{
				"old": "value",
				"foo": map[string]interface{}{
					"bar": []interface{}{
						nil,
						"baz",
					},
				},
			},
		}, {
			source:   `{"old":"value"}`,
			appendix: "[1].foo.bar",
			result: []interface{}{
				nil,
				map[string]interface{}{
					"foo": "bar",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.appendix, func(t *testing.T) {
			var data interface{}
			assert.NoError(t, json.Unmarshal([]byte(tc.source), &data))
			s := strings.Split(tc.appendix, ".")
			appendix := SliceToMap(s[:len(s)-1], s[len(s)-1])
			assert.Equal(t, MergeJSONWithMap(data, appendix), tc.result)
		})
	}
}
