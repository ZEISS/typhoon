package mturl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/zeiss/typhoon/pkg/mturl"
)

func TestURLPath(t *testing.T) {
	obj := &struct {
		metav1.ObjectMeta
	}{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "testns",
			Name:      "testname",
		},
	}

	assert.Equal(t, URLPath(obj), "/testns/testname")
}
