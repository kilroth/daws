package interfaces

type IDeployable interface {
	Deploy() error
	RunTests() error
}
