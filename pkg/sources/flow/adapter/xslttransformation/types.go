//go:build !noclibs

package xslttransformation

// XSLTTransformationStructuredRequest contains an opinionated structure
// that informs both the XML and XSLT to transform.
type XSLTTransformationStructuredRequest struct {
	XML  string `json:"xml"`
	XSLT string `json:"xslt,omitempty"`
}
