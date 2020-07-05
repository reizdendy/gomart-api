package controllers

import (
	"encoding/json"
	"gomart-api/master/models"
	"gomart-api/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type productHandler struct {
	productUsecase usecases.ProductUsecase
}

var (
	product         models.Product
	productResponse models.Response
)

func GetMuxProduct(value string, r *http.Request) string {
	parameter := mux.Vars(r)
	return parameter[value]
}

func ProductController(r *mux.Router, service usecases.ProductUsecase) {
	productHandler := productHandler{service}
	r.HandleFunc("/products", productHandler.ListProduct).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", productHandler.ListProductByID).Methods(http.MethodGet)
	r.HandleFunc("/product", productHandler.AddProduct).Methods(http.MethodPost)
	r.HandleFunc("/product", productHandler.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", productHandler.DeleteProduct).Methods(http.MethodDelete)
}

func (p *productHandler) ListProduct(w http.ResponseWriter, r *http.Request) {
	products, err := p.productUsecase.GetAllProduct()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		productResponse.Messages = "List of Product Data"
		productResponse.Data = products
		productResponse.Status = http.StatusOK
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.Write([]byte("Data Not Found"))
		}
		byteOfproduct, err := json.Marshal(productResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfproduct)
	}
}

func (p *productHandler) ListProductByID(w http.ResponseWriter, r *http.Request) {
	idproduct := GetMuxProduct("id", r)
	product, err := p.productUsecase.GetProductById(idproduct)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		productResponse.Messages = "List Of Product Data"
		productResponse.Data = product
		productResponse.Status = http.StatusOK
		if err != nil {
			w.Write([]byte("Data Not Found"))
		}
		byteOfproduct, err := json.Marshal(productResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfproduct)
	}
}

//AddProduct is a controller of AddAproduct service (Insert / POST). It got the json data from Request Body RAW
func (p *productHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&product)
	err := p.productUsecase.AddProduct(product.ProductCode, product.ProductName, product.ProductCategory, product.ProductPrice)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		productResponse.Messages = "Insert product Data Success!"
		productResponse.Status = http.StatusOK
		productResponse.Data = ""
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		byteOfproduct, err := json.Marshal(productResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfproduct)
	}
}

//UpdateProduct is a controller of UpdateAproduct service (Update/PUT). It got the json data from Request Body RAW
func (p *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&product)
	err := p.productUsecase.UpdateProduct(product.ProductID, product.ProductCode, product.ProductName, product.ProductCategory, product.ProductPrice)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		productResponse.Messages = "Update product Success"
		productResponse.Status = http.StatusOK
		productResponse.Data = ""
		byteOfproduct, err := json.Marshal(productResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfproduct)
	}
}

// DeleteProduct is a controller of  DeleteAproduct service  (DeleTE). It got the json data from Params
func (s *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idproduct := GetMuxProduct("id", r)
	err := s.productUsecase.DeleteProduct(idproduct)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		productResponse.Messages = "Delete product Success"
		productResponse.Status = http.StatusOK
		productResponse.Data = ""
		byteOfproduct, err := json.Marshal(productResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfproduct)
	}
}
