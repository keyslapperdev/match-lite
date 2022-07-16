package main

import (
	"github.com/gin-gonic/gin"
	"github.com/name-schema/match-lite/store"
	"github.com/name-schema/match-lite/types"
)

func main() {
	store := store.NewMemoryStore()

	r := gin.Default()
	r.POST("/student", AddStudent(store))
	r.GET("/", func(ctx *gin.Context) {
		store.Get(types.StudentKey)
	})
	r.Run()
}
