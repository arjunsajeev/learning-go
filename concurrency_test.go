package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "test://test.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://arjun.xyz",
		"test://test.com",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://arjun.xyz":  true,
		"test://test.com":   false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
