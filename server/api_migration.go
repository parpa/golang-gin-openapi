/*
 * Golang Gin Openapi API
 *
 * Here is golang-gin-openapi API project.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/parpa/golang-gin-openapi/migration"
)

// MigrationUp - migration up
var MigrationUpmu sync.Mutex

const MigrationUpSafeCode = "your code"

func MigrationUp(c *gin.Context) {
	code := c.PostForm("safe_code")
	if code != MigrationUpSafeCode {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	start := time.Now()

	MigrationUpmu.Lock()
	defer MigrationUpmu.Unlock()

	res, err := migration.Up()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &MigrationModel{
		Res:    res,
		During: time.Now().Sub(start).String(),
	})
}