package main

import (
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../../assets"))))
}

// package main

// import (
// 	"log"
// 	"net/http"
// )

// const dir = "../../assets"

// func main() {
// 	fs := http.FileServer(http.Dir(dir))
// 	log.Print("Serving " + dir + " on http://localhost:8080")
// 	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
// 		// resp.Header().Add("Cache-Control", "no-cache")
// 		resp.Header().Set("Content-Type", "text/html; charset=ascii")
// 		resp.Header().Set("Access-Control-Allow-Origin", "*")
// 		resp.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
// 		resp.Write([]byte("Hello, World!"))
// 		// if strings.HasSuffix(req.URL.Path, ".wasm") {
// 		// 	resp.Header().Set("content-type", "application/wasm")
// 		// }
// 		fs.ServeHTTP(resp, req)
// 	}))
// }

// package main

// import (
// 	"log"
// 	"net/http"
// )

// func Cors(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=ascii")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
// 	w.Write([]byte("Hello, World!"))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/plm/cors", Cors)
// 	log.Fatal(http.ListenAndServe(":9000", http.ListenAndServe(":8081", mux)))
// 	// http.ListenAndServe(":8081", mux)
// }

// package main

// import (
// 	"net/http"

// 	"github.com/rs/cors"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte("{\"hello\": \"world\"}"))
// 	})

// 	// cors.Default() setup the middleware with default options being
// 	// all origins accepted with simple methods (GET, POST). See
// 	// documentation below for more options.
// 	handler := cors.Default().Handler(mux)
// 	http.ListenAndServe(":8080", handler)
// }
