package main

import (
	"server/config"
	"testing"
)

func TestInit(t *testing.T) {
	err := config.Init()
	if err != nil {
		t.Fatal(err)
	}
}
