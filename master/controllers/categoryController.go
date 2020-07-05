package controllers

import (
	"encoding/json"
	"gomart-api/master/models"
	"gomart-api/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type categoryHandler struct {
	categoryUsecase usecases.CategoryUsecase
}

var (
	category         models.Category
	categoryResponse models.Response
)

func GetMuxCategory(value string, r *http.Request) string {
	parameter := mux.Vars(r)
	return parameter[value]
}

func CategoryController(r *mux.Router, service usecases.CategoryUsecase) {
	categoryHandler := categoryHandler{service}
	r.HandleFunc("/categories", categoryHandler.ListCategory).Methods(http.MethodGet)
	r.HandleFunc("/category/{id}", categoryHandler.ListCategoryById).Methods(http.MethodGet)
	r.HandleFunc("/category", categoryHandler.AddCategoryPage).Methods(http.MethodPost)
	r.HandleFunc("/category", categoryHandler.UpdateCategoryPage).Methods(http.MethodPut)
	r.HandleFunc("/category/{id}", categoryHandler.DeleteCategoryPage).Methods(http.MethodDelete)
}

func (p *categoryHandler) ListCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := p.categoryUsecase.GetAllCategory()

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		categoryResponse.Messages = "List of Category Data"
		categoryResponse.Data = categories
		categoryResponse.Status = http.StatusOK
		byteOfcategory, err := json.Marshal(categoryResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfcategory)
	}
}

func (p *categoryHandler) ListCategoryById(w http.ResponseWriter, r *http.Request) {
	idcategory := GetMuxCategory("id", r)
	category, err := p.categoryUsecase.GetCategoryById(idcategory)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		categoryResponse.Messages = "data category berhasil diambil"
		categoryResponse.Data = category
		categoryResponse.Status = http.StatusOK
		byteOfcategory, err := json.Marshal(categoryResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfcategory)
	}
}

//AddCategoryPage is a controller of AddAcategory service (Insert / POST). It got the json data from Request Body RAW
func (p *categoryHandler) AddCategoryPage(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&category)

	err := p.categoryUsecase.AddCategory(category.CategoryName)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		categoryResponse.Messages = "Insert category Data Success!"
		categoryResponse.Status = http.StatusOK
		categoryResponse.Data = ""
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		byteOfcategory, err := json.Marshal(categoryResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfcategory)
	}
}

//UpdateCategoryPage is a controller of UpdateAcategory service (Update/PUT). It got the json data from Request Body RAW
func (p *categoryHandler) UpdateCategoryPage(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&category)
	err := p.categoryUsecase.UpdateCategory(category.CategoryID, category.CategoryName)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		categoryResponse.Messages = "Update category Success"
		categoryResponse.Status = http.StatusOK
		categoryResponse.Data = ""
		byteOfcategory, err := json.Marshal(categoryResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfcategory)
	}
}

// DeleteCategoryPage is a controller of  DeleteAcategory service  (DeleTE). It got the json data from Params
func (s *categoryHandler) DeleteCategoryPage(w http.ResponseWriter, r *http.Request) {
	idcategory := GetMuxCategory("id", r)

	err := s.categoryUsecase.DeleteCategory(idcategory)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		categoryResponse.Messages = "Delete category Success"
		categoryResponse.Status = http.StatusOK
		categoryResponse.Data = ""
		byteOfcategory, err := json.Marshal(categoryResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfcategory)
	}
}
