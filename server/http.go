package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	errp "github.com/cyansilver/go-lib/err"
	log "github.com/cyansilver/go-lib/log"
)

type HTTPApiServerItc interface {
	SetHttpSrv(srv *http.Server)
	Start(port string)
	ListenOSStopSignal()
	Stop()
	Healthcheck(w http.ResponseWriter, r *http.Request)
	HandleErrorResp(r *Result, err error, w http.ResponseWriter)
	HandleInvalidDataErrorResp(r *Result, err error, w http.ResponseWriter)
	HandleInvalidJsonErrorResp(r *Result, err error, w http.ResponseWriter)
	HandleMissingParamsErrorResp(r *Result, param string, w http.ResponseWriter)
}

type HTTPApiServer struct {
	httpSrv *http.Server
}

func (s *HTTPApiServer) SetHttpSrv(srv *http.Server) {
	s.httpSrv = srv
}

// Start register the grpc server and listen on the port
func (s *HTTPApiServer) Start(port string) {
	fmt.Println(fmt.Sprintf("API server listener port %v", port))

	go func() {
		listener, err := net.Listen("tcp", ":"+port)
		if err != nil {
			log.Logger.Fatalf("error creating the server %v", err)
		}
		if err := s.httpSrv.Serve(listener); err != nil {
			log.Logger.Fatalf("API server listener failed %v", err)
		}
	}()
}

// ListenOSStopSignal listen stop signal from os
func (s *HTTPApiServer) ListenOSStopSignal() {
	// Respect OS stop signals.
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a termination signal.
	<-c

	graceSeconds := 30

	// If a shutdown grace period is allowed, prepare a timer.
	var timer *time.Timer
	timerCh := make(<-chan time.Time, 1)
	if graceSeconds != 0 {
		timer = time.NewTimer(time.Duration(graceSeconds) * time.Second)
		timerCh = timer.C
		log.Logger.Info("Shutdown started - use CTRL^C to force stop server")
	} else {
		// No grace period.
		log.Logger.Info("Shutdown started")
	}

	// Stop any running authoritative matches and do not accept any new ones.
	select {
	case <-timerCh:
		// Timer has expired, terminate matches immediately.
		log.Logger.Info("Shutdown grace period expired")
	case <-c:
		// A second interrupt has been received.
		log.Logger.Info("Skipping graceful shutdown")
	}
	if timer != nil {
		timer.Stop()
	}
}

// Stop the server and call graceful stop
func (s *HTTPApiServer) Stop() {
	ctx := context.Background()
	s.httpSrv.Shutdown(ctx)
}

// Healthcheck to know the server is up
func (s *HTTPApiServer) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// HandleErrorResp returns the error response
func (s *HTTPApiServer) HandleErrorResp(r *Result, err error, w http.ResponseWriter) {
	var e *errp.Error
	if errors.As(err, &e) {
		r.SetError(err.(*errp.Error))
	} else {
		r.SetError(errp.ErrInternal)
	}
	log.Logger.WithError(err).Error(r)
	w.WriteHeader(e.HttpStatus)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
	return
}

// HandleInvalidDataErrorResp returns the error response
func (s *HTTPApiServer) HandleInvalidDataErrorResp(r *Result, err error, w http.ResponseWriter) {
	log.Logger.WithError(err).Error(errp.ErrInvalidData.Error())
	r.SetError(errp.ErrInvalidData)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)

	return
}

// HandleInvalidJsonErrorResp returns the error response
func (s *HTTPApiServer) HandleInvalidJsonErrorResp(r *Result, err error, w http.ResponseWriter) {
	log.Logger.WithError(err).Error(errp.ErrInvalidJson.Error())
	r.SetError(errp.ErrInvalidJson)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)

	return
}

// HandleMissingParamsErrorResp returns the error response
func (s *HTTPApiServer) HandleMissingParamsErrorResp(r *Result, param string, w http.ResponseWriter) {
	log.Logger.WithField("param", param).Error(errp.ErrMissingParams.Error())

	r.SetError(errp.ErrMissingParams)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)

	return
}
