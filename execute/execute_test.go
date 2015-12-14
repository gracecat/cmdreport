package main

import "testing"

func Test_Execute_Ok(t *testing.T) {
	stdout, stderr, err := Execute("date")
	if err != nil {
		t.Error(err)
	}
	if len(stdout) != 1 {
		t.Error("stdout lenght not equal")
	}
	if len(stderr) != 0 {
		t.Error("stderr lenght not equal")
	}
}

func Test_Execute_Err(t *testing.T) {
	stdout, stderr, err := Execute("go", "build", "fixture/foo.src")
	if err != nil {
		t.Error(err)
	}
	if len(stdout) != 0 {
		t.Error("stdout lenght not equal")
	}
	if len(stderr) != 3 {
		t.Error("stderr lenght not equal")
	}
}
