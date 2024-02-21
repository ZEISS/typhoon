package awslambdatarget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustParseARN(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expectPanic bool
	}{
		"valid input": {
			input:       "arn:aws:lambda:us-west-2:testproject:function:lambdadumper",
			expectPanic: false,
		},
		"invalid input": {
			input:       "arn:",
			expectPanic: true,
		},
	}

	for name, tc := range testCases {
		//nolint:scopelint
		t.Run(name, func(t *testing.T) {
			var testFn assert.PanicTestFunc = func() {
				// we do not test the output because the
				// parsing logic belongs to the AWS SDK
				_ = MustParseARN(tc.input)
			}

			if tc.expectPanic {
				assert.PanicsWithValue(t, `failed to parse "`+tc.input+`": arn: not enough sections`, testFn)
			} else {
				assert.NotPanics(t, testFn)
			}
		})
	}
}
