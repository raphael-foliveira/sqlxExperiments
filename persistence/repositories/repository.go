package repositories

type Repository interface {
	Create(*interface{}) (*interface{}, error)
	GetById(*interface{}) (*interface{}, error)
}
