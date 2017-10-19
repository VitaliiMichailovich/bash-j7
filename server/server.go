package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/VitaliiMichailovich/bash-j7/bash"
	"net/http"
	"html/template"
)

var Router *gin.Engine

func Server() {
	// Set gin mode
	gin.SetMode(gin.ReleaseMode)
	// Set the router as the default one provided by Gin
	Router = gin.Default()
	// Set up a static server.
	Router.Use(static.Serve("/client/", static.LocalFile("./client", true)))
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	Router.LoadHTMLGlob("templates/*")
	// Handlers
	Router.GET("/", MainGetHandler)
	Router.POST("/", MainPostHandler)
	// Start serving the application
	Router.Run()
}

func MainGetHandler(c *gin.Context) {
	quote := bash.BashQuote()
	c.HTML(http.StatusOK, "index.html", gin.H {"quote" : template.HTML(quote)})
}

func MainPostHandler(c *gin.Context) {
	quote := bash.BashQuote()
	c.HTML(http.StatusOK, "index.html", gin.H {"quote" : template.HTML(quote)})
}