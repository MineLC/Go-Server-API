package plugin

type Plugin interface {
	Enable()
	Disable()

	IsReloadeable() bool
	Name() string
}

type StructPlugin struct {
	Enable        func()
	Disable       func()
	IsReloadeable bool
	Name          string
}
