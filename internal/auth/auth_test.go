package auth

import (
	"errors"
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

func TestGetEmptyApiKey(t *testing.T) {
	header := http.Header{}
	var val []string
	val = append(val, "ApiKey")
	header["Authorization"] = val

	want := errors.New("malformed authorization header")

	_, got := GetAPIKey(header)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %s", want, got)
	}
}

func TestGetNoApiKey(t *testing.T) {
	header := http.Header{}
	var val []string
	val = append(val, "")
	header["Authorization"] = val

	want := errors.New("no authorization header included")

	_, got := GetAPIKey(header)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %s", want, got)
	}
}
