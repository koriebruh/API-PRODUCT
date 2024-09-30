package helper

import (
	"jamal/api/models/web"
	"net/http"
)

// product create ama update filed sama makanya buat 1 saja

// Fungsi untuk memvalidasi data input produk
func validateProductRequest(productReq web.ProductCreate) *web.WebResponse {
	if productReq.Name == "" {
		return &web.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
	}
	if productReq.PurchasePrice < 0 {
		return &web.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
	}
	if productReq.SellingPrice < 0 {
		return &web.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
	}

	return nil // Tidak ada kesalahan, validasi berhasil
}
