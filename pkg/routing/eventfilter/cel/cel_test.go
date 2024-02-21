package cel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileExpression(t *testing.T) {
	cases := map[string]struct {
		expression string
		wantError  bool
	}{
		"Invalid type": {
			expression: `$var.(foo) == 0`,
			wantError:  true,
		},
		"Malformed var": {
			expression: `TigerMesh!`,
			wantError:  true,
		},
		"Missing type": {
			expression: `$hello == "world"`,
			wantError:  true,
		},
		"Type missmatch": {
			expression: `$var.(int64) == "foo"`,
			wantError:  true,
		},
		"Missing key": {
			expression: `var.(bool) == true`,
			wantError:  true,
		},
		"Invalid variable": {
			expression: `$var(int64) > -1`,
			wantError:  true,
		},
		"Extraneous bits": {
			expression: `$var.(int64) < 1 a`,
			wantError:  true,
		},
		"Non-bool result": {
			expression: `$var1.(int64) + $var2.(int64)`,
			wantError:  true,
		},
		"Malformed type": {
			expression: `$var(bool).(int64) == 3`,
			wantError:  true,
		},
		"Double type": {
			expression: `$var.(bool).(int64) == 3`,
			wantError:  true,
		},
		"Empty string": {
			expression: ` `,
			wantError:  true,
		},
		"Valid expression 1": {
			expression: `$id.first.(int64) + $id.second.(int64) >= -8`,
		},
		"Valid expression 2": {
			expression: `$0.list.(string) == "foo"`,
		},
		"Valid expression 3": {
			expression: `$var.(bool)`,
		},
		"Valid expression 4": {
			expression: `1 > 2`,
		},
		"Valid expression 5": {
			expression: `true`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := CompileExpression(tc.expression)
			if tc.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
