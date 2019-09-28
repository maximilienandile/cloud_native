package main

import (
  "context"
  "github.com/maximilienandile/cloud_native/internal/version"
  "net"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gorilla/mux"
  "github.com/sirupsen/logrus"
)

func main() {

  logger := logrus.New().WithField("version", version.Version)
  logger.Infof("Application starting, Commit %v, Build Time %v", version.Commit, version.BuildTime)
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
  // health check
  diagRouter.HandleFunc("/health",func(w http.ResponseWriter, _ *http.Request){
    logger.Info("health received")
    w.WriteHeader(http.StatusOK)
  })
  // readiness
  diagRouter.HandleFunc("/ready",func(w http.ResponseWriter, _ *http.Request){
    logger.Infof("ready received")
    w.WriteHeader(http.StatusOK)
  })
  diag := http.Server{
    Addr: net.JoinHostPort("",diagPort),
    Handler:diagRouter,
  }

  go func() {
    logger.Infof("Business logic server preparing...")
    err := server.ListenAndServe()
    logger.Errorf("error %v",err)
  }()
  go func() {
    logger.Infof("Diagnostic server preparing...")
    err := diag.ListenAndServe()
    logger.Errorf("error %v",err)

  }()


  // graceful shutdown
  interrupt := make(chan os.Signal,1)
  signal.Notify(interrupt,os.Interrupt,syscall.SIGTERM)
  x:= <- interrupt
  logger.Infof("Received %v. Application stopped",x)
  timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancelFunc()
  err := diag.Shutdown(timeout)
  logger.Errorf("error %v",err)
}
