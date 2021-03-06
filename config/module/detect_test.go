package module

import (
	"testing"
)

func TestDetect(t *testing.T) {
	cases := []struct {
		Input  string
		Pwd    string
		Output string
		Err    bool
	}{
		{"./foo", "/foo", "file:///foo/foo", false},
		{"git::./foo", "/foo", "git::file:///foo/foo", false},
		{
			"git::github.com/hashicorp/foo",
			"",
			"git::https://github.com/hashicorp/foo.git",
			false,
		},
		{
			"./foo//bar",
			"/foo",
			"file:///foo/foo//bar",
			false,
		},
		{
			"git::github.com/hashicorp/foo//bar",
			"",
			"git::https://github.com/hashicorp/foo.git//bar",
			false,
		},
	}

	for i, tc := range cases {
		output, err := Detect(tc.Input, tc.Pwd)
		if (err != nil) != tc.Err {
			t.Fatalf("%d: bad err: %s", i, err)
		}
		if output != tc.Output {
			t.Fatalf("%d: bad output: %s", i, output)
		}
	}
}
