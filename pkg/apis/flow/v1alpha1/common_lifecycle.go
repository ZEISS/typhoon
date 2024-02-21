package v1alpha1

import corev1 "k8s.io/api/core/v1"

// IsInformed returns if the value is informed in any of the available choices.
func (v *ValueFromField) IsInformed() bool {
	if v != nil &&
		(v.Value != "" ||
			v.ValueFromSecret != nil && v.ValueFromSecret.Name != "" && v.ValueFromSecret.Key != "" ||
			v.ValueFromConfigMap != nil && v.ValueFromConfigMap.Name != "" && v.ValueFromConfigMap.Key != "") {
		return true
	}

	return false
}

// ToEnvironmentVariable returns a kubernetes environment variable from
// a ValueFromField.
func (v *ValueFromField) ToEnvironmentVariable(name string) *corev1.EnvVar {
	env := &corev1.EnvVar{
		Name: name,
	}

	switch {
	case v == nil:

	case v.Value != "":
		env.Value = v.Value

	case v.ValueFromSecret != nil && v.ValueFromSecret.Name != "" && v.ValueFromSecret.Key != "":
		env.ValueFrom = &corev1.EnvVarSource{
			SecretKeyRef: v.ValueFromSecret,
		}

	case v.ValueFromConfigMap != nil && v.ValueFromConfigMap.Name != "" && v.ValueFromConfigMap.Key != "":
		env.ValueFrom = &corev1.EnvVarSource{
			ConfigMapKeyRef: v.ValueFromConfigMap,
		}
	}

	return env
}
