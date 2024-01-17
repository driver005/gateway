package admin

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
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
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Get("/:id", m.Get, m.r.Middleware().CanAccessBatchJob)
	route.Post("/:id/confirm", m.Confirm, m.r.Middleware().CanAccessBatchJob)
	route.Delete("/:id/cancel", m.Cancel, m.r.Middleware().CanAccessBatchJob)
}

func (m *Batch) Get(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)
	return context.Status(fiber.StatusOK).JSON(batch)
}

func (m *Batch) List(context fiber.Ctx) error {
	var req types.FilterableBatchJob
	if err := context.Bind().Query(&req); err != nil {
		return err
	}

	createdBy := utils.GetUser(context)
	config, err := sql.FromQuery(context)
	if err != nil {
		return err
	}

	req.CreatedBy = append(req.CreatedBy, createdBy)

	jobs, count, err := m.r.BatchJobService().ListAndCount(req, config)
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

func (m *Batch) Create(context fiber.Ctx) error {
	var req AdminPostBatchesReq
	if err := context.Bind().Body(&req); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}

	userId := utils.GetUser(context)

	m.r.BatchJobService().PrepareBatchJobForProcessing()
	return nil
}

func (m *Batch) Cancel(context fiber.Ctx) error {
	return nil
}

func (m *Batch) Confirm(context fiber.Ctx) error {
	return nil
}
