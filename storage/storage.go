package storage

import (
	"context"

	"github.com/Duvewo/banquend/models"
)

type UsersRepository interface {
	Create(ctx context.Context, u models.UserModel) error
	Update(ctx context.Context, u models.UserModel) error
	Delete(ctx context.Context, u models.UserModel) error

	ByPublicID(ctx context.Context, PID string) (models.UserModel, error)
	ByEmail(ctx context.Context, email string) (models.UserModel, error)
	ByPhoneNumber(ctx context.Context, phoneNumber string) (models.UserModel, error)
}

type AuthRepository interface {
	Restore() //Restore Password
	//Logout() // idk maybe for sessions
}

type CardsRepository interface {
	Create(ctx context.Context, card models.CardModel) error
	Freeze(ctx context.Context, card models.CardModel) error
	Block(ctx context.Context, card models.CardModel) error
	Delete(ctx context.Context, card models.CardModel) error

	ByAccountID() ([]models.CardModel, error)
	ByPublicNumber() (models.CardModel, error)
}

type PaymentsRepository interface{}

type AccountsRepository interface {
	Create(ctx context.Context, account models.AccountModel) error
	ByOwnerID(ctx context.Context, ownerID int) ([]models.AccountModel, error)
	ByID(ctx context.Context, ID int) (models.AccountModel, error)
}

type CurrenciesRepository interface {
	ByCode(ctx context.Context, code string) (models.CurrencyModel, error)
	ListAll(ctx context.Context) ([]models.CurrencyModel, error)
}
