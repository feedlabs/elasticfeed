package template

func Error(err error) (entry map[string]string, code int) {
	return map[string]string{"result": err.Error(), "status": "error"}, HTTP_CODE_ENTITY_NOEXIST
}

func Success(msg string) (entry map[string]string, code int) {
	return map[string]string{"result": msg, "status": "ok"}, HTTP_CODE_VALID_REQUEST
}

func GetOK() int {
	return HTTP_CODE_VALID_REQUEST
}

func PostOK() int {
	return HTTP_CODE_ENTITY_CREATED
}

func PutOK() int {
	return HTTP_CODE_ENTITY_CREATED
}

func DeleteOK() int {
	return HTTP_CODE_VALID_REQUEST
}
