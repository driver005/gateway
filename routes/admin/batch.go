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

func (m *Batch) Get(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)
	return context.Status(fiber.StatusOK).JSON(batch)
}

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

func (m *Batch) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateBatchJobInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	_ = utils.GetUser(context)

	m.r.BatchJobService().PrepareBatchJobForProcessing(model)
	return nil
}

func (m *Batch) Cancel(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)

	model, err := m.r.BatchJobService().Cancel(uuid.Nil, batch)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(model)
}

func (m *Batch) Confirm(context fiber.Ctx) error {
	batch := context.Locals("batch-job").(*models.BatchJob)

	model, err := m.r.BatchJobService().Confirm(uuid.Nil, batch)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(model)
}
