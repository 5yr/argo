package argo

import "strings"

type Argo map[string]string

func NewFromArgs(osArgs []string) *Argo {
	var (
		waiting bool // true - waiting for value
		k       string
		m       Argo
	)
	m = make(Argo)

	for _, arg := range osArgs[1:] { // skip file itself
		if strings.HasPrefix(arg, "-") {
			if waiting {
				m[k] = "" // last -k does not has pair value, save last key -k
				k = arg   // update k, keep waiting
			} else {
				//set k, start waiting
				k = arg
				waiting = true
			}
		} else {
			if waiting {
				// 等到了
				m[k] = arg
				waiting = false
			} else {
				// odd key without dash prefix
				m[arg] = ""
			}
		}
	}

	return &m
}

func (a Argo) Exist(key string) bool {
	_, exist := a[key]
	return exist
}

func (a Argo) Get(key string) string {
	return a[key]
}
