package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jamal/api/helper"
	"jamal/api/models/domain"
	"jamal/api/models/web"
	"jamal/api/repository"
	"net/http"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	db                *gorm.DB
}

func NewProductService(productRepository repository.ProductRepository, db *gorm.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		db:                db,
	}
}

func (service ProductServiceImpl) Create(gin *gin.Context, createReq web.ProductCreate) web.WebResponse {
	var response web.WebResponse

	err := service.db.Transaction(func(tx *gorm.DB) error {
		createProduct := domain.Product{
			Name:          createReq.Name,
			PurchasePrice: createReq.PurchasePrice,
			SellingPrice:  createReq.SellingPrice,
			Stock:         createReq.Stock,
		}

		// if err
		product, err := service.ProductRepository.Create(tx, &createProduct)
		if err != nil {
			response = web.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
			return nil
		}

		// if success, membuat data yg boleh di terima client
		formatData := web.NewProductLimitData(product)

		response = web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   formatData,
		}
		return nil
	})

	helper.HandleErrorResponse(&response, err)

	return response
}

func (service ProductServiceImpl) Delete(ctx *gin.Context, productId int) web.WebResponse {
	var response web.WebResponse

	err := service.db.Transaction(func(tx *gorm.DB) error {

		// FIND DATA
		product, err := service.ProductRepository.FindById(tx, productId)
		if err != nil {
			response = web.WebResponse{Code: http.StatusNotFound, Status: "NOT FOUND", Data: nil}
			return nil
		}

		// DO DELETE
		err = service.ProductRepository.Delete(tx, product.Id)
		if err != nil {
			response = web.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: "has been deleted before"}
			return nil
		}

		// RESPONSE KETIKA BENAR
		response = web.WebResponse{Code: http.StatusOK, Status: "OK", Data: "SUCCESS DELETED DATA"}
		return nil
	})

	helper.HandleErrorResponse(&response, err)

	return response

}

func (service ProductServiceImpl) Update(ctx *gin.Context, updateReq web.ProductUpdate, productId int) web.WebResponse {
	var response web.WebResponse

	err := service.db.Transaction(func(tx *gorm.DB) error {

		//CHECK DATA FIRST
		product, err := service.ProductRepository.FindById(tx, productId)
		if err != nil {
			response = web.WebResponse{Code: http.StatusNotFound, Status: "NOT FOUND", Data: nil}
			return nil
		}

		//DO UPDATE DATA
		updatedProduct := domain.Product{ ///
			Name:          updateReq.Name,
			PurchasePrice: updateReq.PurchasePrice,
			SellingPrice:  updateReq.SellingPrice,
			Stock:         updateReq.Stock,
		}

		_, err = service.ProductRepository.Update(tx, &updatedProduct, product.Id)
		if err != nil {
			response = web.WebResponse{Code: http.StatusInternalServerError, Status: "INTERNAL SERVER ERROR", Data: nil}
			return err
		}

		// PREPARE RESPONSE,
		formatData := web.NewProductLimitData(product)

		response = web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   formatData,
		}
		return nil
	})

	// <-- SERVER ERROR
	helper.HandleErrorResponse(&response, err)

	return response
}

func (service ProductServiceImpl) FindById(ctx *gin.Context, productId int) web.WebResponse {
	var response web.WebResponse

	err := service.db.Transaction(func(tx *gorm.DB) error {
		product, err := service.ProductRepository.FindById(tx, productId)

		/// RESPONSE KETIKA SALAH
		if err != nil {
			response = web.WebResponse{Code: http.StatusNotFound, Status: "NOT FOUND", Data: nil}
			return nil
		}

		/// RESPONSE KETIKA BENAR, foramtin response
		formatData := web.NewProductLimitData(product)

		response = web.WebResponse{
			Code: http.StatusOK, Status: "OK", Data: formatData}
		return nil
	})

	// <-- SERVER ERROR
	helper.HandleErrorResponse(&response, err)

	///  MENGEMBALIKAN HASIL TRANSAKSI BERHASI;
	return response

}

func (service ProductServiceImpl) FindAll(ctx *gin.Context) web.WebResponse {
	var response web.WebResponse

	err := service.db.Transaction(func(tx *gorm.DB) error {
		products, err := service.ProductRepository.FindAll(tx)

		// OUTPUT DATA KETIKA SALAH
		if err != nil {
			// <-- SERVER ERROR
			helper.HandleErrorResponse(&response, err)
			return nil
		}

		// OUTPUT KETIKA BENAR

		response = web.WebResponse{
			Code: http.StatusOK, Status: "OK", Data: products}
		return nil
	})

	if err != nil { // <-- SERVER ERROR
		helper.HandleErrorResponse(&response, err)
	}

	return response
}
