package shared

type Validator interface {
	Validate(obj interface{}) error
}
