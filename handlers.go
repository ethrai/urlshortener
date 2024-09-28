package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func createURLHandler(c *gin.Context) {
	s, _ := c.MustGet(ctxStoreKey).(*Store)
	var form struct {
		URL string `form:"url" binding:"required"`
	}

	if err := c.Bind(&form); err != nil {
		c.Error(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	record, err := s.FindByURL(form.URL)
	if err != nil {
		c.Error(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	if record != nil {
		c.HTML(http.StatusCreated, "result.html", gin.H{"alias": record.Alias})
		return
	}

	alias := generateAlias()

	if err := s.SaveRecord(Record{Alias: alias, URL: form.URL}); err != nil {
		c.Error(err)
		c.HTML(http.StatusInternalServerError, "error.html", struct{}{})
		return
	}
	c.HTML(http.StatusCreated, "result.html", gin.H{"alias": alias})
}

func redirectHandler(c *gin.Context) {
	s, _ := c.MustGet(ctxStoreKey).(*Store)

	alias := c.Param("alias")

	record, err := s.FindByAlias(alias)
	if err != nil {
		c.Error(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusPermanentRedirect, record.URL)
}

func generateAlias() string {
	abc := "qwertyuioasdfghjklzxcvnmQWERTYUIOASDFGHJKLZXCVBNM"
	var sb strings.Builder
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(abc))
		sb.WriteByte(abc[idx])
	}
	return sb.String()
}
