package copybook

import (
	"fmt"
	"sort"
	"strings"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = activity.Register(&Activity{})
}

// Utility functions
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func get_value(data map[string]interface{}, key []string) interface{} {
	var value interface{}
	if len(key) == 1 {
		value = data[key[0]]
	} else {
		newData, ok := data[key[0]].(map[string]interface{})
		if ok {
			value = get_value(newData, key[1:])
		}

	}
	return value
}

func get_type(key []string, copybook map[string]interface{}) string {
	if contains(key, "type") {
		val_type, _ := get_value(copybook, key).(string)
		return val_type
	}
	return ""
}

func get_mapping(key []string, copybook map[string]interface{}) string {
	if contains(key, "mapping") {
		mapping, _ := get_value(copybook, key).(string)
		return mapping
	}
	return ""
}

func get_length(key []string, copybook map[string]interface{}) int {
	if contains(key, "maxLength") {
		val_length := get_value(copybook, key).(float64)
		length := int(val_length)
		return length
	}
	return 0
}

func get_index(key []string, copybook map[string]interface{}) int {
	if contains(key, "index") {
		val_index := get_value(copybook, key).(float64)
		index := int(val_index)
		return index
	}
	return 0
}

func get_maxItems(key []string, copybook map[string]interface{}) int {
	if contains(key, "maxItems") {
		val_index := get_value(copybook, key).(float64)
		index := int(val_index)
		return index
	}
	return 0
}

func join_elements(data map[int]string) string {
	concated_elements := ""
	var indexes []int
	for k := range data {
		indexes = append(indexes, k)
	}
	sort.Ints(indexes)
	for _, index := range indexes {
		concated_elements += data[index]
	}
	return concated_elements
}

// Return ordered elements of a repetitive copybook object
func get_array_elements(root_key []string, element_mapping string, copybook map[string]interface{}, data map[string]interface{}) string {
	new_obj := get_value(copybook, append(root_key, "items")).(map[string]interface{})
	fnObjKeys := json.fnObjKeys
	f := &fnObjKeys{}
	new_keys, err := function.Eval(f, new_obj)
	//new_keys := json.fnObjKeys(new_obj, "")
	max_items := get_maxItems(append(root_key, "maxItems"), copybook)
	elements := make(map[int]string)
	element := get_value(data, strings.Split(element_mapping, ".")).([]interface{})
	n_elements_req := len(element)
	for i, e := range element {
		element_keys := obj_keys(e.(map[string]interface{}), "")
		for _, ek := range element_keys {
			sub_key := element_mapping + "." + ek
			for _, k := range new_keys {
				elem_key := strings.Split(k, ".")
				elem_root_key := elem_key[0 : len(elem_key)-1]
				mapps := get_mapping(elem_key, new_obj)
				if mapps == sub_key {
					length_key := append(elem_root_key, "maxLength")
					length := get_length(length_key, new_obj)
					index := get_index(append(elem_root_key, "index"), new_obj)
					mult := n_elements_req
					if n_elements_req > 2 {
						mult--
					}
					k := i*mult + index
					elem := padding.fnPaddingRight(get_value(e.(map[string]interface{}), []string{ek}).(string), length, "-")
					elements[k] = elem
				}
			}
		}
	}
	filler := ""
	if len(element) < max_items {
		i := 0
		for i < (max_items - len(element)) {
			for _, e := range new_obj {
				elem := padding("", int(e.(map[string]interface{})["maxLength"].(float64)), "-")
				filler += elem
			}
			i++
		}
	}
	val := join_elements(elements)
	result := val + filler
	return result
}

func get_element(key []string, copybook map[string]interface{}, data map[string]interface{}) (int, string) {
	element_type := get_type(key, copybook)
	root_key := key[0 : len(key)-1]
	mapping_key := append(root_key, "mapping")
	element_mapping := get_mapping(mapping_key, copybook)
	element_index := get_index(append(root_key, "index"), copybook)
	if (element_type != "array") && (element_type != "element") {
		element := get_value(data, strings.Split(element_mapping, ".")).(string)
		element_length := get_length(append(root_key, "maxLength"), copybook)
		padded_element := padding(element, element_length, "-")
		return element_index, padded_element
	} else if element_type == "array" {
		elements := get_array_elements(root_key, element_mapping, copybook, data)
		return element_index, elements
	}
	return 0, ""
}

func get_elements(keys []string, copybook map[string]interface{}, data map[string]interface{}) map[int]string {
	elements := make(map[int]string)
	for _, k := range keys {
		if strings.Contains(k, "type") {
			i, element := get_element(strings.Split(k, "."), copybook, data)
			if len(element) > 0 {
				elements[i] = element
			}
		}
	}
	return elements
}

type Input struct {
	Message       string `md:"message"`       // The message to log
	BookStructure string `md:"bookStructure"` // Copybook structure
	UsePrint      bool   `md:"usePrint"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message":       i.Message,
		"bookStructure": i.BookStructure,
		"usePrint":      i.UsePrint,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToObject(values["message"])
	if err != nil {
		return err
	}
	i.BookStructure, err = coerce.ToObject(values["bookStructure"])
	if err != nil {
		return err
	}

	i.UsePrint, err = coerce.ToBool(values["usePrint"])
	if err != nil {
		return err
	}

	return nil
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	msg := input.Message

	/*
		if input.AddDetails {
			msg = fmt.Sprintf("'%s' - HostID [%s], HostName [%s], Activity [%s]", msg,
				ctx.ActivityHost().ID(), ctx.ActivityHost().Name(), ctx.Name())
		}
	*/

	if input.UsePrint {
		fmt.Println(msg)
	} else {
		ctx.Logger().Info(msg)
	}

	return true, nil
}
