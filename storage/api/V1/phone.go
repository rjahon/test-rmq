package V1

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) GetPhone(c *gin.Context) {
	if i := rand.Intn(4); i == 1 {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Id format should be int")
		return
	}

	resp, err := h.storage.Phone().Get(id)
	if err != nil {
		log.Printf("Could not get phone from db: %s", err)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	c.JSON(http.StatusFound, resp)
}
