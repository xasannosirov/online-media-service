package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/xasannosirov/online-media-service/internal/entity"
	"github.com/xasannosirov/online-media-service/internal/usecase"
	"github.com/xasannosirov/online-media-service/pkg/logger"
)

type filesRoutes struct {
	f usecase.FileRepo
	l logger.Interface
}

func newFileRoutes(handler *gin.RouterGroup, f usecase.File, l logger.Interface) {
	r := &filesRoutes{f, l}

	h := handler.Group("/file")
	{
		h.POST("/store", r.store)
		h.DELETE("/remove/:url", r.remove)
	}
}

type fileResponse struct {
	File entity.File `json:"file"`
}

func saveFile(file *multipart.FileHeader, c *gin.Context, r *filesRoutes) (url string, key string, err error) {
	uploadDir := "./uploads"

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		r.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusBadRequest, "invalid mkdir")

		return url, key, err
	}

	uid, err := uuid.NewV4()
	if err != nil {
		r.l.Error(err, "http - v1 - uuid")
		errorResponse(c, http.StatusBadRequest, "invalid uuid")

		return url, key, err
	}
	ext := filepath.Ext(file.Filename)

	filePath := filepath.Join(uploadDir, uid.String()+ext)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		r.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusBadRequest, "save file failed")

		return url, key, err
	}

	return filePath, uid.String() + ext, err
}

// @Summary     Store
// @Description Save file
// @Tags  	    file
// @Accept      multipart/form-data
// @Produce     json
// @Param		request formData file true "Send file"
// @Success     200 {object} entity.File
// @Failure     500 {object} response
// @Router      /file/store [post]
func (r *filesRoutes) store(c *gin.Context) {
	file, err := c.FormFile("request")
	if err != nil {
		r.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusBadRequest, "invalid request")

		return
	}

	filePath, key, err := saveFile(file, c, r)
	if err != nil {
		r.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusBadRequest, "save file failed")

		return
	}

	fileData, err := r.f.Store(c.Request.Context(), entity.File{
		Filename: file.Filename,
		FileURL:  filePath,
	})
	fileData.FileKey = key

	c.JSON(http.StatusOK, fileResponse{fileData})
}

type removeResponse struct {
	Message string `json:"message"       binding:"required"  example:"success"`
}

// @Summary     Remove
// @Description Remove a file
// @ID          url
// @Tags  	    file
// @Accept      json
// @Produce     json
// @Param       url path string true "File URL"
// @Success     200 {object} removeResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /file/remove/{url} [delete]
func (r *filesRoutes) remove(c *gin.Context) {
	url := c.Param("url")

	err := r.f.Remove(
		c.Request.Context(),
		"uploads/"+url,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - remove")
		errorResponse(c, http.StatusInternalServerError, "file service problems")

		return
	}

	if err := os.Remove("./uploads/" + url); err != nil {
		r.l.Error(err, "http - v1 - remove")
		errorResponse(c, http.StatusInternalServerError, "remove file failed")

		return
	}

	c.JSON(http.StatusOK, removeResponse{Message: "success"})
}
