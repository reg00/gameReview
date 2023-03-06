package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/Reg00/gameReview/internal/infrastructure/appctx"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Server    http.Server
	Router    *gin.Engine
	Stop      chan bool
	IsStarted atomic.Bool
	Config    config.HTTP
	Cc        *appctx.CoreContext
}

func Register(
	ctx *appctx.CoreContext,
	cfg *config.Configuration,
	router *gin.Engine,
) *Server {
	webServer := new(Server)
	webServer.IsStarted.Store(false)
	webServer.Stop = make(chan bool)
	webServer.Config = cfg.Http
	webServer.Cc = ctx
	webServer.Router = router
	return webServer
}

func (w *Server) ServerStop() {
	if w.IsStarted.Load() {
		w.Stop <- true
	}
}

func (w *Server) IsServerStarted() bool {
	return w.IsStarted.Load()
}

func (w *Server) Serve() {
	gin.SetMode(gin.ReleaseMode)

	w.Server = http.Server{
		Addr:         ":" + w.Config.Port,
		Handler:      w.Router,
		ReadTimeout:  w.Config.GetTimeout(),
		WriteTimeout: w.Config.GetTimeout(),
	}

	go func() {
		w.IsStarted.Store(true)
		err := w.Server.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				fmt.Println("Can't start webserver: " + err.Error())
			} else {
				fmt.Println(err)
			}
		}
	}()

	go func() {
		<-w.Stop
		log.Println("Webserver stop signal received")
		shutDownCtx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		done := make(chan struct{})
		go func() {
			err := w.Server.Shutdown(shutDownCtx)
			if err != nil {
				log.Fatal("Webserver shutdown error: " + err.Error())
			}
			done <- struct{}{}
		}()
		select {
		case <-shutDownCtx.Done():
			log.Println("Webserver shutdown forced")
		case <-done:
			log.Println("Webserver shutdown completed")
		}
		cancel()
		close(w.Stop)
		w.IsStarted.Store(false)
	}()
}

func (w *Server) Run() {
	w.Serve()

	//graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	stopChannel := make(chan bool)

	go func(ch <-chan os.Signal, st chan<- bool) {
		<-ch
		log.Println("Stop signal received")
		w.ServerStop()
		log.Println("Stop signal sent to webserver")
		w.Cc.Cancel()
		log.Println("Waiting while webserver is stopping")
		for w.IsServerStarted() {
			time.Sleep(time.Microsecond)
		}
		log.Println("Webserver stopped")
		st <- true
	}(signalChannel, stopChannel)

	<-stopChannel
}
