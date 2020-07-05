package repositories

import (
	"database/sql"
	"fmt"
	"gomart-api/master/models"
	"gomart-api/utils"
	"log"

	guuid "github.com/google/uuid"
)

type categoryRepoImpl struct {
	db *sql.DB
}

func (c *categoryRepoImpl) GetAllCategory() ([]*models.Category, error) {
	rows, err := c.db.Query(utils.GET_ALL_CATEGORY)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Categorys []*models.Category
	for rows.Next() {
		Category := models.Category{}
		err := rows.Scan(&Category.CategoryID, &Category.CategoryName, &Category.Status)

		if err != nil {
			return nil, err
		}

		Categorys = append(Categorys, &Category)
	}
	fmt.Println(Categorys)
	return Categorys, nil
}

func (c *categoryRepoImpl) GetCategoryById(id string) (*models.Category, error) {
	Category := models.Category{}
	err := c.db.QueryRow(utils.GET_CATEGORY_BY_ID, id).Scan(&Category.CategoryID, &Category.CategoryName, &Category.Status)
	if err != nil {
		return nil, err
	}
	return &Category, nil
}

func (c *categoryRepoImpl) AddCategory(CategoryName string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := c.db.Prepare(utils.ADD_CATEGORY)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(guuid.New(), CategoryName)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// UpdateCategory is a function to update Category data
func (c *categoryRepoImpl) UpdateCategory(categoryID, categoryName string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := c.db.Prepare(utils.UPDATE_CATEGORY)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(categoryName, categoryID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// DeleteCategory is a function to delete Category data
func (c *categoryRepoImpl) DeleteCategory(id string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := c.db.Prepare(utils.DELETE_CATEGORY)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}

	return tx.Commit()
}

func InitCategoryRepositoryImpl(db *sql.DB) CategoryRepository {
	return &categoryRepoImpl{db}
}
