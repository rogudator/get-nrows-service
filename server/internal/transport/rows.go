package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RowsGetter interface {
	getRows(c *gin.Context)
}

// @Summary Get n Rows
// @Tags rows
// @Description get n rows
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getRowsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rows [get]
func (t *Transport) getRows(c *gin.Context) {
	// log incoming request
	log.Printf("server: got GET request: %s", c.Request.URL.String())
	// get n parameter from query
	parameters := c.Request.URL.Query()
	// check if parameter exists
	if len(parameters["n"]) == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "No \"n\" parameter")
		return
	}
	// Query() function returns parameter as string. We need to convert it to int.
	n, err := strconv.Atoi(parameters["n"][0])
	if err != nil {
		log.Fatal("server: error while converting to int of n paramater")
		newErrorResponse(c, http.StatusInternalServerError, "error while converting to int of n paramater")
		return
	}
	// if n is bigger then the number of rows, we just return every row
	rows := t.serv.GetRows()
	if n > len(rows) {
		n = len(rows)
	}
	// successfully return rows
	c.JSON(http.StatusOK, getRowsResponse{
		Names: rows[:n],
	})
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getRowsResponse struct {
	Names []string `json:"names"`
}
