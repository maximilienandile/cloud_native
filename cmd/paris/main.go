package main

import (
  "net/http"
  "os"

  "github.com/gorilla/mux"
  "github.com/sirupsen/logrus"
)

func main() {

  logger := logrus.New()
  logger.Info("Application starting")
  port := os.Getenv("PORT")
  if port == "" {
    logger.Fatal("Business logic port is not set (PORT)")
  }
  diagPort := os.Getenv("DIAG_PORT")
  if diagPort == "" {
    logger.Fatal("Diagnostic port is not set (DIAG_PORT)")
  }
  r := mux.NewRouter()
  server := http.Server{
    Addr:    ":"+port,
    Handler: r,
  }

  diag := http.Server{
    Addr: ":"+diagPort,
  }

  go func() {
    logger.Info("Business logic server preparing...")
    server.ListenAndServe()
  }()
  logger.Info("Diagnostic server preparing...")

  diag.ListenAndServe()
}
