package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Gee\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geegee"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
