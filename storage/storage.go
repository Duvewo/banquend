package storage

type UsersRepository interface {
	Create()
	Update()
	Delete()
	Search()
}

type PaymentsRepository interface{}
