package tests

import (
	"encoding/json"
	"testing"
)

func marshal(t *testing.T, structure any) string {
	response, err := json.MarshalIndent(structure, "", "  ")
	if err != nil {
		t.Fatalf("%s JSON.MarshalIndent error = %s\n", t.Name(), err.Error())
	}
	return t.Name() + " response = " + string(response)
}
