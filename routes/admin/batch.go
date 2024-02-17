package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AdminPostBatchesReq struct {
	Type    string     `json:"type"`
	Context core.JSONB `json:"context"`
	DryRun  bool       `json:"dry_run,omitempty"`
}

type Batch struct {
	r Registry
}

func NewBatch(r Registry) *Batch {
	m := Batch{r: r}
	return &m
}

func (m *Batch) SetRoutes(router fiber.Router) {
	route := router.Group("/batch-jobs")
	route.Get("/:id", m.Get, m.r.Middleware().CanAccessBatchJob)
	route.Get("", m.List)
	route.Post("", m.Create)

	route.Post("/confirm", m.Confirm, m.r.Middleware().CanAccessBatchJob)
	route.Post("/cancel", m.Cancel, m.r.Middleware().CanAccessBatchJob)
}

// @oas:path [get] /admin/batch-jobs/{id}
// operationId: "GetBatchJobsBatchJob"
// summary: "Get a Batch Job"
// description: "Retrieve the details of a batch job."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Batch Job
//
// x-codegen:
//
//	method: retrieve
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.batchJobs.retrieve(batchJobId)
//     .then(({ batch_job }) => {
//     console.log(batch_job.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminBatchJob } from "medusa-react"
//
//     type Props = {
//     batchJobId: string
//     }
//
//     const BatchJob = ({ batchJobId }: Props) => {
//     const { batch_job, isLoading } = useAdminBatchJob(batchJobId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {batch_job && <span>{batch_job.created_by}</span>}
//     </div>
//     )
//     }
//
//     export default BatchJob
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/batch-jobs/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Batch Jobs
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminBatchJobRes"
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
func (m *Batch) Get(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)
	return context.Status(fiber.StatusOK).JSON(batch)
}

// @oas:path [get] /admin/batch-jobs
// operationId: "GetBatchJobs"
// summary: "List Batch Jobs"
// description: "Retrieve a list of Batch Jobs. The batch jobs can be filtered by fields such as `type` or `confirmed_at`. The batch jobs can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) limit=10 {integer} Limit the number of batch jobs returned.
//   - (query) offset=0 {integer} The number of batch jobs to skip when retrieving the batch jobs.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by the batch ID
//     schema:
//     oneOf:
//   - type: string
//     description: batch job ID
//   - type: array
//     description: multiple batch job IDs
//     items:
//     type: string
//   - in: query
//     name: type
//     style: form
//     explode: false
//     description: Filter by the batch type
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: confirmed_at
//     style: form
//     explode: false
//     description: Filter by a confirmation date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: pre_processed_at
//     style: form
//     explode: false
//     description: Filter by a pre-processing date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: completed_at
//     style: form
//     explode: false
//     description: Filter by a completion date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: failed_at
//     style: form
//     explode: false
//     description: Filter by a failure date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: canceled_at
//     style: form
//     explode: false
//     description: Filter by a cancelation date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - (query) order {string} A batch-job field to sort-order the retrieved batch jobs by.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned batch jobs.
//   - (query) fields {string} Comma-separated fields that should be included in the returned batch jobs.
//   - in: query
//     name: created_at
//     style: form
//     explode: false
//     description: Filter by a creation date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: updated_at
//     style: form
//     explode: false
//     description: Filter by an update date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetBatchParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.batchJobs.list()
//     .then(({ batch_jobs, limit, offset, count }) => {
//     console.log(batch_jobs.length)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminBatchJobs } from "medusa-react"
//
//     const BatchJobs = () => {
//     const {
//     batch_jobs,
//     limit,
//     offset,
//     count,
//     isLoading
//     } = useAdminBatchJobs()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {batch_jobs?.length && (
//     <ul>
//     {batch_jobs.map((batchJob) => (
//     <li key={batchJob.id}>
//     {batchJob.id}
//     </li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default BatchJobs
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/batch-jobs' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Batch Jobs
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminBatchJobListRes"
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
func (m *Batch) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableBatchJob](context)
	if err != nil {
		return err
	}

	createdBy := utils.GetUser(context)

	model.CreatedBy = append(model.CreatedBy, createdBy)

	jobs, count, err := m.r.BatchJobService().ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"batch_jobs": jobs,
		"count":      count,
		"offset":     config.Skip,
		"limit":      config.Take,
	})
}

// @oas:path [post] /admin/batch-jobs
// operationId: "PostBatchJobs"
// summary: "Create a Batch Job"
// description: "Create a Batch Job to be executed asynchronously in the Medusa backend. If `dry_run` is set to `true`, the batch job will not be executed until the it is confirmed,
//
//	which can be done using the Confirm Batch Job API Route."
//
// externalDocs:
//
//	description: "How to create a batch job"
//	url: "https://docs.medusajs.com/development/batch-jobs/create#create-batch-job"
//
// x-authenticated: true
// requestBody:
//
//	content:
//	 application/json:
//	   schema:
//	     $ref: "#/components/schemas/AdminPostBatchesReq"
//
// x-codegen:
//
//	method: create
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.batchJobs.create({
//     type: 'product-export',
//     context: {},
//     dry_run: false
//     }).then((({ batch_job }) => {
//     console.log(batch_job.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateBatchJob } from "medusa-react"
//
//     const CreateBatchJob = () => {
//     const createBatchJob = useAdminCreateBatchJob()
//     // ...
//
//     const handleCreateBatchJob = () => {
//     createBatchJob.mutate({
//     type: "publish-products",
//     context: {},
//     dry_run: true
//     }, {
//     onSuccess: ({ batch_job }) => {
//     console.log(batch_job)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateBatchJob
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/batch-jobs' \
//     -H 'Content-Type: application/json' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     --data-raw '{
//     "type": "product-export",
//     "context": { }
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Batch Jobs
//
// responses:
//
//	201:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminBatchJobRes"
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
func (m *Batch) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateBatchJobInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	_ = utils.GetUser(context)

	m.r.BatchJobService().PrepareBatchJobForProcessing(model)
	return nil
}

// @oas:path [post] /admin/batch-jobs/{id}/cancel
// operationId: "PostBatchJobsBatchJobCancel"
// summary: "Cancel a Batch Job"
// description: "Mark a batch job as canceled. When a batch job is canceled, the processing of the batch job doesn’t automatically stop."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the batch job.
//
// x-codegen:
//
//	method: cancel
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.batchJobs.cancel(batchJobId)
//     .then(({ batch_job }) => {
//     console.log(batch_job.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelBatchJob } from "medusa-react"
//
//     type Props = {
//     batchJobId: string
//     }
//
//     const BatchJob = ({ batchJobId }: Props) => {
//     const cancelBatchJob = useAdminCancelBatchJob(batchJobId)
//     // ...
//
//     const handleCancel = () => {
//     cancelBatchJob.mutate(undefined, {
//     onSuccess: ({ batch_job }) => {
//     console.log(batch_job)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default BatchJob
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/batch-jobs/{id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Batch Jobs
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminBatchJobRes"
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
func (m *Batch) Cancel(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)

	model, err := m.r.BatchJobService().Cancel(uuid.Nil, batch)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(model)
}

// @oas:path [post] /admin/batch-jobs/{id}/confirm
// operationId: "PostBatchJobsBatchJobConfirmProcessing"
// summary: "Confirm a Batch Job"
// description: "When a batch job is created, it is not executed automatically if `dry_run` is set to `true`. This API Route confirms that the batch job should be executed."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the batch job.
//
// x-codegen:
//
//	method: confirm
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.batchJobs.confirm(batchJobId)
//     .then(({ batch_job }) => {
//     console.log(batch_job.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminConfirmBatchJob } from "medusa-react"
//
//     type Props = {
//     batchJobId: string
//     }
//
//     const BatchJob = ({ batchJobId }: Props) => {
//     const confirmBatchJob = useAdminConfirmBatchJob(batchJobId)
//     // ...
//
//     const handleConfirm = () => {
//     confirmBatchJob.mutate(undefined, {
//     onSuccess: ({ batch_job }) => {
//     console.log(batch_job)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default BatchJob
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/batch-jobs/{id}/confirm' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Batch Jobs
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminBatchJobRes"
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
func (m *Batch) Confirm(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)

	model, err := m.r.BatchJobService().Confirm(uuid.Nil, batch)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(model)
}
