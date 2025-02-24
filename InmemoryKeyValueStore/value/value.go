package value

type Value struct {
	attributes map[string]IAttrValue
}

func getIAttrValueFromValue(input any) IAttrValue {
	switch input.(type) {
	case int:
		return &AttrValueInt{value: input.(int)}
	case string:
		return &AttrValueString{value: input.(string)}
	}
	return nil
}

func NewValue(attr map[string]interface{}) *Value {
	attributes := map[string]IAttrValue{}
	for k, v := range attr {
		attributes[k] = getIAttrValueFromValue(v)
	}
	return &Value{attributes: attributes}
}

func (v *Value) setAttribute(key string, val interface{}) {
	_, ok := v.attributes[key]
	if !ok {
		v.attributes[key] = getIAttrValueFromValue(v)
		return
	}
	v.attributes[key].Set(v)
}

func (v *Value) SetAttributes(attrMap map[string]interface{}) error {
	for attr, attrVal := range attrMap {
		err := v.attributes[attr].Set(attrVal)
		if err != nil {
			return err // TODO: What should be the handling here ?
		}
	}
	return nil
}

func (v *Value) GetValueAsMap() (map[string]interface{}, error) {
	valueMap := make(map[string]interface{}, 0)
	var err error
	for attr, attrVal := range v.attributes {
		if valueMap[attr], err = attrVal.Get(); err != nil {
			return nil, err
		}
	}
	return valueMap, nil
}
