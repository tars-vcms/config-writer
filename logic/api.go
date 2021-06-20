package logic

type ConfigServer interface {
	CreateConfig()

	InsertItem()

	UpdateItem()
}

func NewConfigServer() ConfigServer {
	return newConfigServerImpl()
}
