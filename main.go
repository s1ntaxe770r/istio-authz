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

func (ap *AuthorizationPolicy) GetNSAuthPolicy() (unstructured.Unstructured, error) {
	policy := unstructured.Unstructured{}
	policy.SetAPIVersion("networking.istio.io/v1alpha3")
	policy.SetKind("AuthorizationPolicy")
	policy.SetName(ap.Name)
	policy.SetNamespace(ap.Namespace)

	// Create the AuthorizationPolicy spec
	authPolicySpec := map[string]interface{}{
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
	}

	policy.Object = map[string]interface{}{
		"spec": authPolicySpec,
	}

	return policy, nil
}

func main() {
	authorizationPolicy := AuthorizationPolicy{
		Name:      "allow-ingress-namespace",
		Namespace: "your-namespace",
	}

	// print out full policy
	policy, _ := authorizationPolicy.GetNSAuthPolicy()

	// print out policy spec
	jsonPolicy, _ := json.Marshal(policy.Object["spec"])
	fmt.Println(string(jsonPolicy))

	fmt.Println(policy)
}
