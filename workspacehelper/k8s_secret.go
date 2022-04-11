package workspacehelper

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

// GetSecretData retrieves the data from a secret in a given namespace
func (r *WorkspaceHelper) GetSecretData(namespace string, name string) (map[string][]byte, error) {
	// If no secretName defined, return empty map
	if name == "" {
		return make(map[string][]byte), nil
	}

	r.reqLogger.Info("Getting Secret", "Namespace", namespace, "Name", name)

	secret := &corev1.Secret{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, secret)

	if err != nil {
		r.reqLogger.Error(err, "Failed to get Secret", "Namespace", namespace, "Name", name)

		return nil, err
	}

	return secret.Data, nil
}
