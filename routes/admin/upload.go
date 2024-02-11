package admin

import (
	"github.com/driver005/gateway/interfaces"
	"github.com/gofiber/fiber/v3"
)

type AdminPostUploadsDownloadUrlReq struct {
	FileKey string `json:"file_key"`
}

type AdminDeleteUploadsReq struct {
	FileKey string `json:"file_key"`
}

type Upload struct {
	r Registry
}

func NewUpload(r Registry) *Upload {
	m := Upload{r: r}
	return &m
}

func (m *Upload) SetRoutes(router fiber.Router) {
	route := router.Group("/uploads")
	route.Post("", m.Create)
	route.Post("/protected", m.CreateProtectedUpload)
	route.Delete("", m.Delete)
	route.Post("/download-url", m.Get)
}

func (m *Upload) Get(context fiber.Ctx) error {
	var req AdminPostUploadsDownloadUrlReq
	if err := context.Bind().Query(&req); err != nil {
		return err
	}

	res, err := m.r.FileService().GetPresignedDownloadUrl(interfaces.GetUploadedFileType{FileKey: req.FileKey})
	if err != nil {
		return err
	}
	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"download_url": res,
	})
}

func (m *Upload) Create(context fiber.Ctx) error {
	var results []interfaces.FileServiceUploadResult
	if form, err := context.MultipartForm(); err == nil {
		files := form.File["files"]
		for _, file := range files {
			res, err := m.r.FileService().Upload(file)
			if err != nil {
				return err
			}

			results = append(results, *res)
		}
	}
	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"uploads": results,
	})
}

func (m *Upload) Delete(context fiber.Ctx) error {
	var req AdminDeleteUploadsReq
	if err := context.Bind().Query(&req); err != nil {
		return err
	}

	if err := m.r.FileService().Delete(interfaces.DeleteFileType{FileKey: req.FileKey}); err != nil {
		return err
	}
	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"id":      req.FileKey,
		"object":  "file",
		"deleted": true,
	})
}

func (m *Upload) CreateProtectedUpload(context fiber.Ctx) error {
	var results []interfaces.FileServiceUploadResult
	if form, err := context.MultipartForm(); err == nil {
		files := form.File["files"]
		for _, file := range files {
			res, err := m.r.FileService().UploadProtected(file)
			if err != nil {
				return err
			}

			results = append(results, *res)
		}
	}
	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"uploads": results,
	})
}
