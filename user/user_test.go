package main

import "testing"

func TestUser(t *testing.T) {
	got := User("test")
	want := "test"

	if got != want {

		t.Errorf("got: %q want: %q", got, want)
	}
}
