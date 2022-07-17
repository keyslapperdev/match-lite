package main

import (
	"github.com/gin-gonic/gin"
	"github.com/name-schema/match-lite/store"
	"github.com/name-schema/match-lite/types"
)

func AddStudent(s store.Storer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ret types.Student
		if err := c.Bind(&ret); err != nil {
			return
		}

		s.Add(types.StudentKey, types.NewStudent(ret.Name))
	}
}

func AddCompany(s store.Storer) gin.HandlerFunc {
	return func(c *gin.Context) {

		var ret types.Company
		if err := c.Bind(&ret); err != nil {
			return
		}

		s.Add(types.CompanyKey, types.NewCompany(ret.Name))
	}
}
