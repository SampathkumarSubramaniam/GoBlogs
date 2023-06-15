package modernWeb

import (
	"fmt"
	"net/http"
)

type  struct{}

func () Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func () Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func () Close() error {
	//TODO implement me
	panic("implement me")
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	_ = http.ListenAndServe(":8080", nil)
}
