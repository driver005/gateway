package admin

import (
	"github.com/driver005/gateway/interfaces"
	"github.com/gofiber/fiber/v3"
)

// @oas:schema:AdminPostUploadsDownloadUrlReq
// type: object
// description: "The details of the file to retrieve its download URL."
// required:
//   - file_key
//
// properties:
//
//	file_key:
//	  description: "key of the file to obtain the download link for. This is obtained when you first uploaded the file, or by the file service if you used it directly."
//	  type: string
type AdminPostUploadsDownloadUrlReq struct {
	FileKey string `json:"file_key"`
}

// @oas:schema:AdminDeleteUploadsReq
// type: object
// description: "The details of the file to delete."
// required:
//   - file_key
//
// properties:
//
//	file_key:
//	  description: "key of the file to delete. This is obtained when you first uploaded the file, or by the file service if you used it directly."
//	  type: string
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

// @oas:path [post] /admin/uploads/download-url
// operationId: "PostUploadsDownloadUrl"
// summary: "Get a File's Download URL"
// description: "Create and retrieve a presigned or public download URL for a file. The URL creation is handled by the file service installed on the Medusa backend."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostUploadsDownloadUrlReq"
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.uploads.getPresignedDownloadUrl({
//     file_key
//     })
//     .then(({ download_url }) => {
//     console.log(download_url);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreatePresignedDownloadUrl } from "medusa-react"
//
//     const Image = () => {
//     const createPresignedUrl = useAdminCreatePresignedDownloadUrl()
//     // ...
//
//     const handlePresignedUrl = (fileKey: string) => {
//     createPresignedUrl.mutate({
//     file_key: fileKey
//     }, {
//     onSuccess: ({ download_url }) => {
//     console.log(download_url)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Image
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/uploads/download-url' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "file_key": "{file_key}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Uploads
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUploadsDownloadUrlRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
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

// @oas:path [post] /admin/uploads
// operationId: "PostUploads"
// summary: "Upload Files"
// description: "Upload at least one file to a public bucket or storage. The file upload is handled by the file service installed on the Medusa backend."
// x-authenticated: true
// requestBody:
//
//	content:
//	  multipart/form-data:
//	    schema:
//	      type: object
//	      properties:
//	        files:
//	          type: string
//	          format: binary
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.uploads.create(file)
//     .then(({ uploads }) => {
//     console.log(uploads.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUploadFile } from "medusa-react"
//
//     const UploadFile = () => {
//     const uploadFile = useAdminUploadFile()
//     // ...
//
//     const handleFileUpload = (file: File) => {
//     uploadFile.mutate(file, {
//     onSuccess: ({ uploads }) => {
//     console.log(uploads[0].key)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default UploadFile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/uploads' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: image/jpeg' \
//     --form 'files=@"<FILE_PATH_1>"' \
//     --form 'files=@"<FILE_PATH_1>"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Uploads
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUploadsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
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

// @oas:path [delete] /admin/uploads
// operationId: "DeleteUploads"
// summary: "Delete an Uploaded File"
// description: "Delete an uploaded file from storage. The file is deleted using the installed file service on the Medusa backend."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteUploadsReq"
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.uploads.delete({
//     file_key
//     })
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteFile } from "medusa-react"
//
//     const Image = () => {
//     const deleteFile = useAdminDeleteFile()
//     // ...
//
//     const handleDeleteFile = (fileKey: string) => {
//     deleteFile.mutate({
//     file_key: fileKey
//     }, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Image
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/uploads' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "file_key": "{file_key}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Uploads
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDeleteUploadsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
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

// @oas:path [post] /admin/uploads/protected
// operationId: "PostUploadsProtected"
// summary: "Protected File Upload"
// description: "Upload at least one file to an ACL or a non-public bucket. The file upload is handled by the file service installed on the Medusa backend."
// x-authenticated: true
// requestBody:
//
//	content:
//	  multipart/form-data:
//	    schema:
//	      type: object
//	      properties:
//	        files:
//	          type: string
//	          format: binary
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.uploads.createProtected(file)
//     .then(({ uploads }) => {
//     console.log(uploads.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUploadProtectedFile } from "medusa-react"
//
//     const UploadFile = () => {
//     const uploadFile = useAdminUploadProtectedFile()
//     // ...
//
//     const handleFileUpload = (file: File) => {
//     uploadFile.mutate(file, {
//     onSuccess: ({ uploads }) => {
//     console.log(uploads[0].key)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default UploadFile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/uploads/protected' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: image/jpeg' \
//     --form 'files=@"<FILE_PATH_1>"' \
//     --form 'files=@"<FILE_PATH_1>"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Uploads
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUploadsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
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
