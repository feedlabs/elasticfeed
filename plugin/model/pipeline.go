package model

type Pipeline interface {

	Prepare(...interface{}) ([]string, error)

	Run(data interface {}) (interface {}, error)

	Cancel()
}
