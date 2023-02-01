package json

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnObjKeys{})
}

type fnObjKeys struct {
}

// Name returns the name of the function
func (fnObjKeys) Name() string {
	return "objKeys"
}

// Sig returns the function signature
func (fnObjKeys) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeObject}, false
}

// Utility recursive function to return all keys and sub-keys in a JSON tring
// Iterates over all structures inside the paren structure
func obj_keys(data map[string]interface{}, parent string) []string {
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		if len(parent) == 0 {
			keys[i] = k
		} else {
			keys[i] = parent + "." + k
		}

		newData, ok := data[k].(map[string]interface{})
		if ok {
			var newKeys []string = obj_keys(newData, keys[i])
			keys = append(keys, newKeys...)
		}
		i++
	}
	return keys
}

// *
// Eval executes the function
func (fnObjKeys) Eval(params ...interface{}) (interface{}, error) {
	switch params[0].(type) {
	case []interface{}:
		return nil, fmt.Errorf("Cannot list keys for array type object")
	}
	input, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	keys := obj_keys(input, "")

	return keys, nil
}

//*/

/*
// Eval executes the function
func (fnObjKeys) Eval(params ...interface{}) (interface{}, error) {
	switch params[0].(type) {
	case []interface{}:
		return nil, fmt.Errorf("Cannot list keys for array type object")
	}
	input, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	keys := make([]string, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	return keys, nil
}
*/
