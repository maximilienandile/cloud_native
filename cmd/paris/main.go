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
    logger.Infof("ready received, will wait 2 seconds before answer")
    time.Sleep(1 * time.Minute)
    w.WriteHeader(http.StatusOK)
  })
  diag := http.Server{
    Addr: net.JoinHostPort("",diagPort),
    Handler:diagRouter,
  }

  // Bufered channel of two,
  // because we have two services
  shutdown := make(chan error, 2)

  go func() {
    logger.Infof("Business logic server preparing...")
    err := server.ListenAndServe()
    // Note: here we check that we do not have the error
    // server closed, because it's not really an error.
    // this pattern might be remembered because you can have the same
    // pattern for database usage (ex: no rows)
    if err != nil && err != http.ErrServerClosed {
      shutdown<-err
    }
  }()
  go func() {
    logger.Infof("Diagnostic server preparing...")
    err := diag.ListenAndServe()
    if err != nil && err != http.ErrServerClosed{
      shutdown<-err
    }
  }()


  // graceful shutdown
  // listen for interrupt
  interrupt := make(chan os.Signal,1)
  signal.Notify(interrupt,os.Interrupt,syscall.SIGTERM)

  // listen for both channels
  select {
      case x := <- interrupt :
        logger.Infof("Received %v",x)
      case err := <- shutdown:
        // TODO : here we handle just the first error
        // how to make it listen to error 1 and 2
        logger.Infof("Received %v",err)

  }
  // NOTE : What happens if we have an INTERRUPT, and right after a shutdown ?
  // not handled.
  // maybe add a waitgroup ?
  // Be carefull to not have to complex shutdown mechanism
  timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancelFunc()

  // if we did not pass a timeout,
  // when we interrupt the process
  // it will wait for the requests to be processed
  // with the timeout we will wait for the timeout duration not more.
  err := diag.Shutdown(timeout)
  if err != nil {
    logger.Errorf("error %v",err)
  }
  err = server.Shutdown(timeout)
  if err != nil {
    logger.Errorf("error %v",err)

  }
}
