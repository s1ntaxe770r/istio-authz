package main

import (
	"encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type AuthorizationPolicy struct {
	Name      string
	Namespace string
}

// GetIngressPolicy returns an AuthorizationPolicy for the ingress gateway
func (ap *AuthorizationPolicy) GetNSPolicy() unstructured.Unstructured {
	policy := unstructured.Unstructured{}
	policy.SetAPIVersion("networking.istio.io/v1alpha3")
	policy.SetKind("AuthorizationPolicy")
	policy.SetName(ap.Name)
	policy.SetNamespace(ap.Namespace)

	// Create the AuthorizationPolicy spec
	policy.Object = map[string]interface{}{
		"spec": map[string]interface{}{},
	}
	return policy
}

// GetNsPolicy returns an AuthorizationPolicy that allows all traffic within the namespace and from the ingress gateway
func (ap *AuthorizationPolicy) GetNsIngressPolicy() unstructured.Unstructured {
	policy := unstructured.Unstructured{}
	policy.SetAPIVersion("networking.istio.io/v1alpha3")
	policy.SetKind("AuthorizationPolicy")
	policy.SetName(ap.Name)
	policy.SetNamespace(ap.Namespace)
	policy.SetUnstructuredContent(map[string]interface{}{
		"selector": map[string]interface{}{
			"matchLabels": map[string]interface{}{
				"istio": "ingressgateway",
			},
		},
		"action": "ALLOW",
		"rules": []map[string]interface{}{
			{
				"to": []map[string]interface{}{
					{
						"operation": map[string]interface{}{
							"methods": []string{"*"},
						},
					},
				},
			},
		},
	})

	return policy
}

func main() {
	ap := AuthorizationPolicy{
		Name:      "allow-nothing",
		Namespace: "default",
	}

	authorizationPolicy := ap.GetNSPolicy()

	// Convert to JSON for printing
	data, err := json.MarshalIndent(authorizationPolicy, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))

}
