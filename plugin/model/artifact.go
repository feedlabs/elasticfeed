package model

type Artifact interface {
	BuilderId() string

	Files() []string

	Id() string

	String() string

	State(name string) interface{}

	Destroy() error
}
