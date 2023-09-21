package main

import (
	"fmt"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	conf   *config.AppConfig
	engine *gin.Engine
}

func NewServer(engine *gin.Engine, conf *config.AppConfig) *Server {
	return &Server{
		engine: engine,
		conf:   conf,
	}
}

func (s *Server) Run() {
	s.engine.Run(fmt.Sprintf(":%s", s.conf.GetString("app.port")))
}
