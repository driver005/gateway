package database

// func (h *Handler) GetBatchJob(ctx context.Context, config types.FilterableBatchJobProps, model models.BatchJob) (*models.BatchJob, error) {
// 	var m models.BatchJob

// 	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
// 		return nil, helper.ParseError(err)
// 	}

// 	return &m, nil
// }

// func (h *Handler) GetBatchJobs(ctx context.Context, filters models.Filter, config types.FilterableBatchJobProps, model models.BatchJob) ([]models.BatchJob, *int64, error) {
// 	var m = make([]models.BatchJob, 0)

// 	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
// 		return nil, nil, helper.ParseError(err)
// 	}

// 	n := h.Query(ctx, config, model).Find(&models.BatchJob{})
// 	if n.Error != nil {
// 		return nil, nil, helper.ParseError(n.Error)
// 	}

// 	if m == nil {
// 		m = []models.BatchJob{}
// 	}

// 	return m, &n.RowsAffected, nil
// }

// func (h *Handler) GetBatchJobCount(ctx context.Context, config types.FilterableBatchJobProps, model models.BatchJob) ([]models.BatchJob, *int64, error) {
// 	var m = make([]models.BatchJob, 0)
// 	var count int64

// 	if err := h.Query(ctx, config, model).Find(&m).Count(&count).Error; err != nil {
// 		return nil, nil, helper.ParseError(err)
// 	}

// 	return m, &count, nil
// }

// func (h *Handler) CreateBatchJob(ctx context.Context, m *models.BatchJob) (*models.BatchJob, error) {
// 	m.CreatedAt = time.Now().UTC().Round(time.Second)
// 	m.UpdatedAt = m.CreatedAt
// 	if err := h.r.Manager(ctx).Session(&gorm.Session{DryRun: m.DryRun}).Create(&m).Error; err != nil {
// 		return nil, helper.ParseError(err)
// 	}

// 	return m, nil
// }

// func (h *Handler) UpdateBatchJob(ctx context.Context, m *models.BatchJob) (*models.BatchJob, error) {
// 	q := models.BatchJob{}
// 	q.Id = m.Id

// 	o, err := h.GetBatchJob(ctx, types.FilterableBatchJobProps{}, q)
// 	if err != nil {
// 		return nil, helper.ParseError(err)
// 	}

// 	m.Id = o.Id
// 	m.UpdatedAt = time.Now().UTC().Round(time.Second)
// 	if err := h.r.Manager(ctx).Session(&gorm.Session{DryRun: m.DryRun}).Model(&o).Updates(m).Error; err != nil {
// 		return nil, helper.ParseError(err)
// 	}

// 	return m, nil
// }

// func (h *Handler) DeleteBatchJob(ctx context.Context, id uuid.UUID) error {
// 	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.BatchJob{}).Error; err != nil {
// 		return helper.ParseError(err)
// 	}

// 	return nil
// }
