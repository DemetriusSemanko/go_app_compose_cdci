package main

import (
   "log" 
   "net/http" 
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   w.Header().Set("Content-Type", "application/json")
   w.Write([]byte(`{"message": "I have successfully deployed CD/CI on 2025-04-04 at 10:24:00 EST"}`))
}

func main() {
  s := &Server{}
  http.Handle("/", s)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
