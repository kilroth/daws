package interfaces

type IConfig interface {
	Save() error
	Load() (IDeployable, error)
}
