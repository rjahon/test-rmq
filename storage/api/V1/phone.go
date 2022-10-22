package V1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) GetPhone(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println("Id format should be int")
		return
	}

	resp, err := h.storage.Phone().Get(id)
	if err != nil {
		log.Printf("%s: %s", "Could not get phone from db", err)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
