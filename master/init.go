package master

import (
	"database/sql"
	"gomart-api/master/controllers"
	"gomart-api/master/repositories"
	"gomart-api/master/usecases"

	"github.com/gorilla/mux"
)

func Init(db *sql.DB, r *mux.Router) {
	//product routers
	productRepo := repositories.InitProductRepositoryImpl(db)
	productUsecase := usecases.InitProductUsecase(productRepo)
	controllers.ProductController(r, productUsecase)

	//category routers
	categoryRepo := repositories.InitCategoryRepositoryImpl(db)
	categoryUsecase := usecases.InitCategoryUsecase(categoryRepo)
	controllers.CategoryController(r, categoryUsecase)

	//transactions routers
	transactionRepo := repositories.InitTransactionRepositoryImpl(db)
	transactionUsecase := usecases.InitTransactionUsecase(transactionRepo)
	controllers.TransactionController(r, transactionUsecase)

	//reports routers
	reportRepo := repositories.InitReportRepositoryImpl(db)
	reportUsecase := usecases.InitReportUsecase(reportRepo)
	controllers.ReportController(r, reportUsecase)

}
