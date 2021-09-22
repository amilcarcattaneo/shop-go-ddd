package database

import (
	"net/http"

	"shop-go-ddd/src/api/errors/entities"
)

func (repository *CheckDatabaseRepository) Check() (bool, *entities.ErrorType) {
	var res int
	err := repository.database.DB().QueryRow("SELECT 1").Scan(&res)
	if err != nil {
		return false, &entities.ErrorType{
			Error:  "DB connection error",
			Status: http.StatusInternalServerError,
		}
	}

	if res != 1 {
		return false, &entities.ErrorType{
			Error:  "Unexpected query result",
			Status: http.StatusInternalServerError,
		}
	}

	return true, nil
}
