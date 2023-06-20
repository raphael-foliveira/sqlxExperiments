package persistence

type ContextManager interface {
	Create(*interface{}) (*interface{}, error)
	GetById(*interface{}) (*interface{}, error)
}
