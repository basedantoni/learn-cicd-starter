package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	var val []string
	val = append(val, "ApiKey 12345")
	header["Authorization"] = val

	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	expect := "12345"
	if !reflect.DeepEqual(key, expect) {
		t.Fatalf("expected: %s, got: %s", expect, key)
	}
}
