package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		h, err := os.Hostname()
		if err != nil {
			h = "unknown-host"
		}
		c.String(http.StatusOK, fmt.Sprintf("hello from %s", h))
	})

	// Intentionally insecure: simple CodeQL test route.
	r.GET("/run", func(c *gin.Context) {
		cmd := c.Query("cmd")
		out, _ := exec.Command("sh", "-c", cmd).CombinedOutput()
		c.String(http.StatusOK, string(out))
	})

	_ = r.Run(":8080")
}
