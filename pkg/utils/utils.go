package utils

import (
	"github.com/aws/smithy-go/ptr"
)

// String ...
func String(v string) *string {
	return ptr.String(v)
}

// Int64 ...
func Int64(v int64) *int64 {
	return ptr.Int64(v)
}

// Float32 ...
func Float32(v float32) *float32 {
	return ptr.Float32(v)
}
