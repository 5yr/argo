package argo

import (
	"fmt"
	"testing"
)

func TestNewFromArgs(t *testing.T) {
	v := []string{"User/x/main.go", "-n", "3", "--color", "yellow", "-d", "--dumb-long1", "--size", "small", "jojo"}
	args := NewFromArgs(v)

	ts := []bool{
		args.Exist("-n"),
		args.Get("-n") == "3",
		args.Exist("--color"),
		args.Get("--color") == "yellow",
		args.Exist("-d"),
		args.Get("-d") == "",
		args.Exist("--dumb-long1"),
		args.Get("--dumb-long1") == "",
		args.Exist("--size"),
		args.Get("--size") == "small",
		args.Exist("jojo"),
		args.Get("jojo") == "",
	}

	fmt.Println(args)

	for _, o := range ts {
		if !o {
			t.Fail()
		}
	}
}
