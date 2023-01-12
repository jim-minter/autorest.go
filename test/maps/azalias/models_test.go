package azalias

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolicyAssignmentProperties(t *testing.T) {
	const payload = `{"displayName":"Not allowed resource types - Virtual Machine","metadata":{"one":{"value":{"key":"value"}}},"parameters":{"effect":{"value":"Audit"},"listOfResourceTypesNotAllowed":{"value":["Microsoft.Compute/virtualMachines"]}}}`

	paprops := PolicyAssignmentProperties{}
	if err := json.Unmarshal([]byte(payload), &paprops); err != nil {
		t.Fatal(err)
	}
	var s string
	if err := json.Unmarshal(paprops.Parameters["effect"].Value, &s); err != nil {
		t.Fatal(err)
	}
	if s != "Audit" {
		t.Fatalf("got %s, want Audit", s)
	}
	sl := []string{}
	if err := json.Unmarshal(paprops.Parameters["listOfResourceTypesNotAllowed"].Value, &sl); err != nil {
		t.Fatal(err)
	}
	if len(sl) != 1 {
		t.Fatal("unexpected slice len")
	}
	if sl[0] != "Microsoft.Compute/virtualMachines" {
		t.Fatalf("got %s, want Microsoft.Compute/virtualMachines", sl[0])
	}
	m, ok := paprops.Metadata["one"]
	if !ok {
		t.Fatal("missing one")
	}
	mm := map[string]any{}
	if err := json.Unmarshal(m.Value, &mm); err != nil {
		t.Fatal(err)
	}
	if v := mm["key"]; v != "value" {
		t.Fatalf("got %s want value", v)
	}
	b, err := json.Marshal(paprops)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != payload {
		t.Fatalf("got %s", string(b))
	}
}

func TestUnmarshalFail(t *testing.T) {
	const data = `{"id": 123}`
	var geo GeoJSONFeature
	err := json.Unmarshal([]byte(data), &geo)
	require.Error(t, err)
	require.Equal(t, "unmarshalling type *azalias.GeoJSONFeature: struct field ID: json: cannot unmarshal number into Go value of type string", err.Error())
}
