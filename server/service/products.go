package service

import (
	"image"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/chai2010/webp"
	"github.com/gosimple/slug"
	"github.com/nfnt/resize"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	product.Slug = slug.Make(product.Name)
	err := s.repo.InsertProduct(product)
	return err
}
func (s *ProductService) UpdateProduct(product *models.Product) error {
	err := s.repo.UpdateProduct(product)
	return err
}
func (s *ProductService) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	return err
}
func (s *ProductService) GetProduct(filter dto.ProductFiler) (dto.ProductsResponce, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 20
	}
	products, err := s.repo.SelectProducts(filter)
	return products, err

}
func (s *ProductService) GetProductSlug(slug string) (models.Product, error) {
	product, err := s.repo.SelectProductSlug(slug)
	return product, err
}

func PaginationParse(page, limit string) (int, int, error) {
	if page == "" || limit == "" {
		return 1, 10, nil
	} else {
		p, err := strconv.Atoi(page)
		if err != nil {
			return 1, 10, err
		}
		l, err := strconv.Atoi(limit)
		if err != nil {
			return 1, 10, err
		}
		return p, l, nil
	}
}

func (s *ProductService) GetProductImage(slug string) {

}

func (s *ProductService) SetProductImages(slug string, files *[]*multipart.FileHeader) error {
	err := s.repo.AddProductImageCount(slug, len(*files))
	if err != nil {
		return err
	}
	for i, file := range *files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		img, _, err := image.Decode(src)
		if err != nil {
			return err
		}
		pathSmall := "static/images/products/" + slug + "/small/"
		pathBig := "static/images/products/" + slug + "/big/"
		_ = os.MkdirAll(pathSmall, os.ModePerm)
		_ = os.MkdirAll(pathBig, os.ModePerm)
		out, _ := os.Create(pathSmall + strconv.Itoa(i+1) + ".webp")
		resized := resize.Resize(400, 0, img, resize.Lanczos3)
		webp.Encode(out, resized, &webp.Options{Quality: 80})
		out, _ = os.Create(pathBig + strconv.Itoa(i+1) + ".webp")
		webp.Encode(out, img, &webp.Options{Quality: 80})
	}
	return nil

}
