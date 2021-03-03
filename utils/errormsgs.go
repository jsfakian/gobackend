package utils

//ParameterNotExistsErr is a function that returns an error if a parameter does not exist
func ParameterNotExistsErr(param string) string {
	return "the parameter " + param + " does not exist"
}

//ParameterNotUUIDErr is a function that returns an error if a parameter is uuid
func ParameterNotUUIDErr(param string) string {
	return "the parameter " + param + " is not Uuid"
}
