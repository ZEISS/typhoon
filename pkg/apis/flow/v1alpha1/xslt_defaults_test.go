package v1alpha1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXSLTTransformationSetDefaults(t *testing.T) {
	testCases := map[string]struct {
		xslt      *XSLTTransformation
		defaulted *XSLTTransformation
	}{
		"XSLT with allow event xslt value set to false, needs no defaulting": {
			xslt: xsltTransform(
				xsltWithXSLT(valueFromField(vffWithValue(tValue))),
				xsltWithAllowEventXSLT(false),
			),
			defaulted: xsltTransform(
				xsltWithXSLT(valueFromField(vffWithValue(tValue))),
				xsltWithAllowEventXSLT(false),
			),
		},
		"XSLT with allow event xslt value set to true, needs no defaulting": {
			xslt: xsltTransform(
				xsltWithXSLT(valueFromField(vffWithValue(tValue))),
				xsltWithAllowEventXSLT(true),
			),
			defaulted: xsltTransform(
				xsltWithXSLT(valueFromField(vffWithValue(tValue))),
				xsltWithAllowEventXSLT(true),
			),
		},
		"XSLT without allow event xslt value, needs defaulting": {
			xslt: xsltTransform(xsltWithXSLT(valueFromField(vffWithValue(tValue)))),
			defaulted: xsltTransform(
				xsltWithXSLT(valueFromField(vffWithValue(tValue))),
				xsltWithAllowEventXSLT(false),
			),
		},
		"XSLT nil does not defaulting": {
			xslt:      nil,
			defaulted: nil,
		},
	}

	for name, tc := range testCases {
		//nolint:scopelint
		t.Run(name, func(t *testing.T) {
			tc.xslt.SetDefaults(context.Background())
			assert.Equal(t, tc.defaulted, tc.xslt)
		})
	}
}
