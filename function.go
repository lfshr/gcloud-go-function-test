package function

import (
	"net/http"
)

func F(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; carset=utf8")
	w.Write([]byte("Hello World!"))
}
