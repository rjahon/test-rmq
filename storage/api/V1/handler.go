package V1

import (
	"github.com/rjahon/labs-rmq/storage/config"
	"github.com/rjahon/labs-rmq/storage/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

// func (h *handlerV1) HandleError(c *gin.Context, err error, message string) {
// 	if err == sql.ErrNoRows {
// 		h.HandleNotFoundError(c, err, message)
// 	} else {
// 		h.HandleInternalServerError(c, err, message)
// 	}
// }

// func (h *handlerV1) HandleBadRequest(c *gin.Context, err error, message string) {
// 	h.log.Error(message, logger.Error(err))
// 	c.JSON(400, models.Response{
// 		Error: models.Error{
// 			Code:    helper.ErrorCodeBadRequest,
// 			Message: message,
// 		},
// 	})
// }

// func ParseQueryParam(c *gin.Context, key string, defaultValue string) (int32, error) {
// 	valueStr := c.DefaultQuery(key, defaultValue)

// 	value, err := strconv.Atoi(valueStr)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		c.Abort()
// 		return 0, err
// 	}

// 	return int32(value), nil
// }

// func ParseBoolQueryParam(c *gin.Context, name string) (bool, error) {
// 	a, err := strconv.ParseBool(c.DefaultQuery(name, "false"))
// 	if err != nil {
// 		return false, err
// 	}
// 	return a, nil
// }

// func (h *handlerV1) HandleInternalServerError(c *gin.Context, err error, message string) {
// 	h.log.Error(message, logger.Error(err))
// 	c.JSON(http.StatusInternalServerError, models.Response{
// 		Error: models.Error{
// 			Code:    helper.ErrorCodeInternal,
// 			Message: message,
// 		},
// 	})
// }

// func (h *handlerV1) HandleNotFoundError(c *gin.Context, err error, message string) {
// 	h.log.Warn(message, err)
// 	c.JSON(http.StatusNotFound, models.Response{
// 		Error: models.Error{
// 			Code:    helper.ErrorCodeNotFound,
// 			Message: message,
// 		},
// 	})
// }

// func GetLanguage(c *gin.Context) string {
// 	value := c.GetHeader("Accept-Language")
// 	if value == "" || value != "uz" && value != "ru" && value != "en" {
// 		value = "ru"
// 	}

// 	return value
// }

// func (h *handlerV1) GetCompanyID(c *gin.Context) string {
// 	if h.cfg.Environment == "develop" {
// 		return "d5f7b467-aa93-4741-b59f-278bd1c6e4de"
// 		// return "8b5abc0e-112a-4b94-81bd-b04aace025d1"
// 	}

// 	res, ok := c.Get("auth")
// 	if !ok {
// 		return ""
// 	}

// 	v := reflect.ValueOf(res)
// 	f := v.FieldByName("CompanyID")
// 	return f.String()
// }
