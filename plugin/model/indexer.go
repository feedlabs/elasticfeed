package model

type Indexer interface {
	Prepare(...interface{}) ([]string, error)

	Run(cache Cache) (Artifact, error)

	Cancel()
}
