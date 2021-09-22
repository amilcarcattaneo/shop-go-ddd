package usecase

import "shop-go-ddd/src/api/errors/entities"

func (usecases *UseCases) Check() (bool, *entities.ErrorType) {
	return usecases.checkRepository.Check()
}
