// main.go
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cert/:websiteURL/:port", func(c *gin.Context) {
		websiteURL := c.Param("websiteURL")
		port := c.Param("port")

		portInt, err := strconv.Atoi(port)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid port"})
			return
		}

		certificate, err := getCertificate(websiteURL, portInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting certificate: %s", err)})
			return
		}

		certJSON, err := convertCertificateToJSON(certificate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error converting certificate to JSON: %s", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"certificate": certJSON})
	})

	router.Run(":8080")
}
