package model

type Pipeline interface {

	Prepare(...interface{}) ([]string, error)

	Run(cache Cache) (Artifact, error)

	Cancel()
}
