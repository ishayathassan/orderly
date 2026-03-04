package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id64, err := strconv.ParseUint(idStr, 10, 32) // ParseUint always returns uint64, bit size=32 for valid uint32 checking
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			c.Abort() // Stops the chain here
			return
		}
		// Store the parsed ID in the context for the next handler
		c.Set("orderID", uint(id64))
		c.Next()
	}
}