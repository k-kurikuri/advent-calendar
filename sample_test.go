package main_test

import (
	"fmt"
	"os"
	"testing"
)

func TestSample(t *testing.T) {
	for _, tt := range []struct {
		name   string
		expect bool
		skip   bool
	}{
		{
			name:   "success",
			expect: true,
		},
		{
			name: "skip",
			skip: true,
		},
		{
			name:   "fail",
			expect: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Env CI:", os.Getenv("CI"))
			if tt.expect {
				return
			}

			if tt.skip {
				// todo:
				t.Skip("skip test")
				return
			}

			if !tt.expect {
				t.Error("fail test case")
				return
			}
		})
	}
}
