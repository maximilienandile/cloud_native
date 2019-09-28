package main

import (
  "net"
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
    Addr:    net.JoinHostPort("",port),
    Handler: r,
  }

  diagRouter := mux.NewRouter()
  diagRouter.HandleFunc("/health",func(w http.ResponseWriter, _ *http.Request){
    w.WriteHeader(http.StatusOK)
  })
  diag := http.Server{
    Addr: net.JoinHostPort("",diagPort),
    Handler:diagRouter,
  }

  go func() {
    logger.Info("Business logic server preparing...")
    server.ListenAndServe()
  }()
  logger.Info("Diagnostic server preparing...")

  diag.ListenAndServe()
}
