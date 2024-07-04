package storage

import (
	"context"

	"github.com/Duvewo/banquend/models"
)

type UsersRepository interface {
	Create(context.Context, models.UserModel) error
	Update(context.Context, models.UserModel) error
	Delete(context.Context, models.UserModel) error
	Search(context.Context, models.UserModel) (models.UserModel, error)
}

type PaymentsRepository interface{}

type AccountsRepository interface {
	Create(context.Context, models.AccountModel) error
}

type CurrenciesRepository interface {
	ByCode(context.Context, models.CurrencyModel) (models.CurrencyModel, error)
}
