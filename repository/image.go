package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
)

type ImageRepo struct {
	sql.Repository[models.Image]
}

func ImageRepository(db *gorm.DB) *ImageRepo {
	return &ImageRepo{*sql.NewRepository[models.Image](db)}
}

func (r ImageRepo) InsertBulk(data []models.Image) ([]models.Image, *utils.ApplictaionError) {
	err := r.Db().Create(&data).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}

	return data, nil
}

func (r ImageRepo) FindByURLs(urls []string) ([]models.Image, *utils.ApplictaionError) {
	var images []models.Image
	err := r.Db().Where("url IN ?", urls).Find(&images).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}

	return images, nil
}

func (r ImageRepo) UpsertImages(images []models.Image) ([]models.Image, *utils.ApplictaionError) {
	var imageUrls []string
	for _, i := range images {
		imageUrls = append(imageUrls, i.Url)
	}
	existingImages, err := r.FindByURLs(imageUrls)
	if err != nil {
		return nil, err
	}

	existingImagesMap := make(map[string]models.Image)
	for _, img := range existingImages {
		existingImagesMap[img.Url] = img
	}

	var upsertedImgs []models.Image
	var imageToCreate []models.Image
	for _, url := range imageUrls {
		aImg, ok := existingImagesMap[url]
		if ok {
			upsertedImgs = append(upsertedImgs, aImg)
		} else {
			newImg := models.Image{Url: url}
			imageToCreate = append(imageToCreate, newImg)
		}
	}

	if len(imageToCreate) > 0 {
		newImgs, err := r.InsertBulk(imageToCreate)
		if err != nil {
			return nil, err
		}
		upsertedImgs = append(upsertedImgs, newImgs...)
	}

	return upsertedImgs, nil
}
