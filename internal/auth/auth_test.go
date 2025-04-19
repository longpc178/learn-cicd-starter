package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey 1234")

	got, err := GetAPIKey(header)
	want := "1234"

	if err != nil {
		t.Fatalf("GetAPIKey %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("GetAPIKey fail expected: %v, got: %v", want, got)
	}
}
