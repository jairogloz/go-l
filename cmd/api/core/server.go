package core

import (
	"github.com/gin-gonic/gin"
)

// Server represents the core server and holds the required components to expose
// the application services through http endpoints using the gin framework.
type Server struct {
	GinEngine *gin.Engine
}
