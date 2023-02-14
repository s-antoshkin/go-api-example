package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Record struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func getID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Writer.WriteHeader(400)
		return 0, false
	}
	return id, true
}

func getRecords(c *gin.Context) {
	var str string
	if len(c.Request.URL.RawQuery) > 0 {
		str = c.Request.URL.Query().Get("name")
		if str == "" {
			c.Writer.WriteHeader(400)
			return
		}
	}
	recs, err := readAll(str)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, recs)
}

func getRecord(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	rec, err := readOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "phone not found"})
	}
	c.IndentedJSON(http.StatusOK, rec)
}

func addRecord(c *gin.Context) {
	var rec Record
	err := c.BindJSON(&rec)
	if err != nil || rec.Name == "" || rec.Phone == "" {
		c.Writer.WriteHeader(400)
		return
	}
	if _, err := insert(rec.Name, rec.Phone); err != nil {
		c.Writer.WriteHeader(500)
		return
	}
	c.Writer.WriteHeader(201)
}

func updateRecord(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	var rec Record
	err := json.NewDecoder(c.Request.Body).Decode(&rec)
	if err != nil || rec.Name == "" || rec.Phone == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	res, err := update(id, rec.Name, rec.Phone)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	n := res.RowsAffected()
	if n == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "No Content"})
}

func deleteRecord(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	if _, err := remove(id); err != nil {
		c.Writer.WriteHeader(500)
	}
	c.Writer.WriteHeader(204)
}
