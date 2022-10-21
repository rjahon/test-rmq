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
		// h.HandleBadRequest(c, err, "Id format should be uuid")
		log.Println("Id format should be int")
		return
	}

	resp, err := h.storage.Phone().Get(id)
	if err != nil {
		// h.HandleError(c, err, "Phone not found")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
