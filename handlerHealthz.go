package main

import (
	"fmt"
	"net/http"
)

func handlerHealthz(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
	<html>
	
	<body>
	    <h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	
	</html>
	`, cfg.fileserverHits.Load())))

}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	cfg.fileserverHits.Store(0)
	x := cfg.fileserverHits.Load()
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	str := fmt.Sprintf("Hits: %d", x)
	w.Write([]byte(str))
}
