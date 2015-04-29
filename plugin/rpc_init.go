package plugin

import "encoding/gob"

func init() {
	gob.Register(make([]interface{}, 0))
	gob.Register(new(BasicError))
}
