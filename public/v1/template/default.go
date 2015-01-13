package template

func GetError(err error) (entry map[string]string, code int) {
	return map[string]string{"result": err.Error(), "status": "error"}, HTTP_CODE_ENTITY_NOEXIST
}
