package useсase

type ContactUseCase interface {
	CreateContact() error
	ReadContact() error
	UpdateContact() error
	DeleteContact() error
}

type GroupUseCase interface {
	CreateGroup() error
	ReadGroup() error
	AddContactToGroup() error
}
