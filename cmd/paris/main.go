package main

import (
  "net/http"

  "github.com/gorilla/mux"
  "github.com/sirupsen/logrus"
)

func main() {

  logger := logrus.New()
  logger.Info("Application starting")

  r := mux.NewRouter()
  server := http.Server{
    Addr:    ":8080",
    Handler: r,
  }

  diag := http.Server{
    Addr: ":8081",
  }

  go func() {
    logger.Info("Business logic server preparing...")
    server.ListenAndServe()
  }()
  logger.Info("Diagnostic server preparing...")

  diag.ListenAndServe()
}
