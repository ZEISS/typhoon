package aws

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreclientv1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Credentials returns the AWS security credentials referenced in a source's
// spec, using the provided Secrets client if necessary.
func Credentials(cli coreclientv1.SecretInterface, creds *v1alpha1.AWSSecurityCredentials) (*credentials.Value, error) {
	accessKeyID := creds.AccessKeyID.Value
	secretAccessKey := creds.SecretAccessKey.Value
	sessionToken := creds.SessionToken.Value

	// cache a Secret object by name to avoid GET-ing the same Secret
	// object multiple times
	var secretCache map[string]*corev1.Secret

	if vfs := creds.AccessKeyID.ValueFromSecret; vfs != nil {
		secr, err := cli.Get(context.Background(), vfs.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("getting Secret from cluster: %w", err)
		}

		// cache Secret containing the access key ID so it can be reused
		// below in case the same Secret contains the secret access key
		secretCache = map[string]*corev1.Secret{
			vfs.Name: secr,
		}

		accessKeyID = string(secr.Data[vfs.Key])
	}

	if vfs := creds.SecretAccessKey.ValueFromSecret; vfs != nil {
		var secr *corev1.Secret
		var err error

		if secretCache != nil && secretCache[vfs.Name] != nil {
			secr = secretCache[vfs.Name]
		} else {
			secr, err = cli.Get(context.Background(), vfs.Name, metav1.GetOptions{})
			if err != nil {
				return nil, fmt.Errorf("getting Secret from cluster: %w", err)
			}
		}

		secretAccessKey = string(secr.Data[vfs.Key])
	}

	if vfs := creds.SessionToken.ValueFromSecret; vfs != nil {
		var secr *corev1.Secret
		var err error

		if secretCache != nil && secretCache[vfs.Name] != nil {
			secr = secretCache[vfs.Name]
		} else {
			secr, err = cli.Get(context.Background(), vfs.Name, metav1.GetOptions{})
			if err != nil {
				return nil, fmt.Errorf("getting Secret from cluster: %w", err)
			}
		}

		sessionToken = string(secr.Data[vfs.Key])
	}

	return &credentials.Value{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		SessionToken:    sessionToken,
	}, nil
}
