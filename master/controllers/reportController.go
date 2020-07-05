package controllers

import (
	"encoding/json"
	"fmt"
	"gomart-api/master/models"
	"gomart-api/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type reportHandler struct {
	reportUsecase usecases.ReportUsecase
}

var (
	// report         models.Report
	reportResponse models.Response
)

func GetMuxReport(value string, r *http.Request) string {
	parameter := mux.Vars(r)
	return parameter[value]
}

func ReportController(r *mux.Router, service usecases.ReportUsecase) {
	reportHandler := reportHandler{service}
	r.HandleFunc("/dailyreport/{year}/{month}/{day}", reportHandler.GetDailyReport).Methods(http.MethodGet)
	r.HandleFunc("/monthlyreport/{month}", reportHandler.GetMonthlyReport).Methods(http.MethodGet)
}

func (p *reportHandler) GetDailyReport(w http.ResponseWriter, r *http.Request) {
	year := GetMuxReport("year", r)
	month := GetMuxReport("month", r)
	day := GetMuxReport("day", r)
	report, err := p.reportUsecase.GetDailyReport(year, month, day)
	reportResponse.Messages = "Get The Daily Reports Data"
	reportResponse.Data = report
	reportResponse.Status = http.StatusOK
	byteOfreport, err := json.Marshal(reportResponse)
	if err != nil {
		w.Write([]byte("Something went wrong!"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfreport)
	fmt.Println(report)
}

func (p *reportHandler) GetMonthlyReport(w http.ResponseWriter, r *http.Request) {
	monthReport := GetMuxReport("month", r)
	report, err := p.reportUsecase.GetMonthlyReport(monthReport)
	reportResponse.Messages = "Get The Monthly Reports Data"
	reportResponse.Data = report
	reportResponse.Status = http.StatusOK
	byteOfreport, err := json.Marshal(reportResponse)
	if err != nil {
		w.Write([]byte("Something went wrong!"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfreport)
}
