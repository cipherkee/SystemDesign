package value

import "errors"

type IAttrValue interface {
	Get() (interface{}, error)
	Set(interface{}) error
}

var (
	IncompatibleValueUpdate = errors.New("Incompatible value update")
	ValueIsNull             = errors.New("Value object is null")
)

//--------------------------------------------------------------------------------------------------------------------------------------------

type AttrValueString struct {
	value string
}

func (a *AttrValueString) Get() (interface{}, error) {
	if a == nil {
		return nil, ValueIsNull
	}
	return a.value, nil
}

func (a *AttrValueString) Set(newValue interface{}) error {
	if a == nil {
		return ValueIsNull
	}
	switch newValue.(type) {
	case string:
		a.value = newValue.(string)
		return nil
	default:
		return IncompatibleValueUpdate
	}
}

//--------------------------------------------------------------------------------------------------------------------------------------------

type AttrValueInt struct {
	value int
}

func (a *AttrValueInt) Get() (interface{}, error) {
	if a == nil {
		return nil, ValueIsNull
	}
	return a.value, nil
}

func (a *AttrValueInt) Set(newValue interface{}) error {
	if a == nil {
		return ValueIsNull
	}
	switch newValue.(type) {
	case int:
		a.value = newValue.(int)
		return nil
	default:
		return IncompatibleValueUpdate
	}
}
