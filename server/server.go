package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	handler "taskreminder/handler"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ServerHandlerMap struct {
	APIPath string
	Handler handler.APIHandler
}

type Server struct {
	Port        int
	ChiServer   *http.Server
	ChiApp      *chi.Mux
	APIRootPath string
	Handlers    []*ServerHandlerMap
}

func NewServerHandlerMap(apiPath string, handler handler.APIHandler) *ServerHandlerMap {
	return &ServerHandlerMap{APIPath: apiPath, Handler: handler}
}

func (s *Server) Setup() {
	if s.ChiApp == nil {
		log.Fatalln("Server setup incorrectly!")
	}
	AddMiddlewares(s.ChiApp)

	for _, eachHandlerMap := range s.Handlers {
		apiGroup := chi.NewRouter()
		s.ChiApp.Mount(s.APIRootPath+eachHandlerMap.APIPath, apiGroup)
		log.Println(s.APIRootPath+eachHandlerMap.APIPath, apiGroup)
		eachHandlerMap.Handler.RegisterRoutes(apiGroup)
	}
	s.ChiServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.Port),
		Handler: s.ChiApp,
	}
	log.Println(*s.ChiServer)
}

func (s *Server) StartServer() <-chan os.Signal {
	s.Setup()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ChiServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	return quit
}

func NewServer() *chi.Mux {
	router := chi.NewRouter()
	return router
}

func (s *Server) ShutdownGracefully() {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		// Release resources like Database connections
		cancel()

	}()

	shutdownChan := make(chan error, 1)
	go func() { shutdownChan <- s.ChiServer.Shutdown(context.Background()) }()

	select {
	case <-timeout.Done():
		log.Fatal("Server Shutdown Timed out before shutdown.")
	case err := <-shutdownChan:
		if err != nil {
			log.Fatal("Error while shutting down server", err)
		} else {
			log.Println("Server Shutdown Successful")
		}
	}

}

func AddMiddlewares(r chi.Router) {
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.AllowContentType("application/json"),
	)
}
