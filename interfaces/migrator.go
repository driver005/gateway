package interfaces

type IMigrator interface {
	Up() error
	Down() error
	GetName() string
}
