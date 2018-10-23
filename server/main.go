package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, id+" is "+validUserName(id))
}

func validUserName(name string) string {
	if len(name) == 0 || len(name) > 32 {
		return "invalid"
	}
	if (name[0] < 'a' || name[0] > 'z') && name[0] != '_' {
		return "invalid"
	}

	for i, c := range name[1:] {
		if (c < 'a' || c > 'z') && (c != '_' && c != '.' && c != '-') {
			if i == len(name)-2 && c == '$' {
				break
			}
			return "invalid"
		}
	}
	return "valid"
}

func main() {
	http.Handle("/names", http.HandlerFunc(handler))
	http.ListenAndServe(":8080", nil)
}
