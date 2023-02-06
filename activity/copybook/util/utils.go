package util

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/JCorpse96/contrib/function/padding/util"
	"github.com/project-flogo/core/data/coerce"
)

const request = `{
	"msg": {
		"id": "123456",
		"data": {
			"type": "2",
			"number": "0012111122223333444455556"
		},
		"object": {
			"entity": "11223",
			"reference": "123321123",
			"value": "3.45"
		},
		"destinations": [
			{
				"number": "123",
				"name": "ABC"
			},
			{
				"number": "456",
				"name": "DEF"
			},
			{
				"number": "789",
				"name": "GHI"
			},
			{
				"number": "101112",
				"name": "JKL"
			}
		],
		"location": "Rua da esquina",
		"source": "B"
	}
}`

const copybook = `{
	"header": {
		"MSG": "H001"
	},
	"v1": {
		"CPY-ID": {
			"mapping": "msg.id",
			"type": "number",
			"format": "integer",
			"index": 0,
			"maxLength": 10
		},
		"CPY-DATA": {
			"CPY-DATA-TYPE": {
				"mapping": "msg.data.type",
				"type": "number",
				"format": "integer",
				"index": 1,
				"maxLength": 4
			},
			"CPY-DATA-NUMBER": {
				"mapping": "msg.data.number",
				"type": "number",
				"format": "integer",
				"index": 2,
				"maxLength": 30
			}
		},
		"CPY-OBJECT": {
			"CPY-DATA-ENTITY": {
				"mapping": "msg.object.entity",
				"type": "number",
				"format": "integer",
				"index": 3,
				"maxLength": 8
			},
			"CPY-DATA-REFERENCE": {
				"mapping": "msg.object.reference",
				"type": "number",
				"format": "integer",
				"index": 4,
				"maxLength": 11
			},
			"CPY-DATA-VALUE": {
				"mapping": "msg.object.value",
				"type": "number",
				"format": "float",
				"index": 5,
				"maxLength": 7
			}
		},
		"CPY-DESTINATIONS": {
			"mapping": "msg.destinations",
			"type": "array",
			"index": 6,
			"maxItems": 4,
			"items": {
				"CPY-DESTINATIONS-ACC": {
					"mapping": "msg.destinations.number",
					"type": "element",
					"format": "integer",
					"index": 0,
					"maxLength": 6
				},
				"CPY-DESTINATIONS-NAME": {
					"mapping": "msg.destinations.name",
					"type": "element",
					"format": "string",
					"index": 1,
					"maxLength": 5
				}
			}
		},
		"CPY-LOCATION": {
			"mapping": "msg.location",
			"type": "string",
			"format": "string",
			"index": 7,
			"maxLength": 15
		},
		"CPY-SOURCE": {
			"mapping": "msg.source",
			"type": "string",
			"format": "string",
			"index": 8,
			"maxLength": 1
		}
	}
}`

/*
func padding(value string, length int, padCharacter string) string {
	var padCountInt = 1 + ((length - len(padCharacter)) / len(padCharacter))
	var retStr = value + strings.Repeat(padCharacter, padCountInt)
	return retStr[:length]
}*/

func padding(value string, length int, padCharacter string) string {
	retStr := util.PaddingRight(value, length, padCharacter)
	return retStr
}

func test(input string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce to object: %s", err.Error())
	}
	return data, nil
}

// Not used
/*
func objkey(data map[string]any, parent string) []string {
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		if len(parent) == 0 {
			keys[i] = k
		} else {
			keys[i] = parent + "." + k
		}

		newData, ok := data[k].(map[string]any)
		if ok {
			var newKeys []string = objkey(newData, keys[i])
			keys = append(keys, newKeys...)
		}
		i++
	}
	return keys
}
*/

// Used function
func ObjKeys(data map[string]interface{}, parent string) []string {
	//keys := make([]string, len(data))
	var keys []string
	i := 0
	for k := range data {
		key := ""
		if len(parent) == 0 {
			key += k
		} else {
			key += parent + "." + k
		}

		newData, ok := data[k].(map[string]interface{})
		if ok {
			var newKeys []string = ObjKeys(newData, key)
			keys = append(keys, newKeys...)
		} else {
			keys = append(keys, key)
		}
		i++
	}
	return keys
}

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

// NOT USED FUNCTION
/*
func get_element_alpha(key []string, copybook map[string]interface{}, data map[string]interface{}) (int, string) {
	element_type := get_type(key, copybook)
	mapping_key := append(key[0:len(key)-1], "mapping")
	if (element_type != "array") && (element_type != "element") {
		//mapping_key := append(key[0:len(key)-1], "mapping")
		element_mapping := get_mapping(mapping_key, copybook)
		element := get_value(data, strings.Split(element_mapping, ".")).(string)
		element_length := get_length(append(key[0:len(key)-1], "maxLength"), copybook)
		element_index := get_index(append(key[0:len(key)-1], "index"), copybook)
		//fmt.Println(element_length)
		padded_element := padding(element, element_length, "-")
		return element_index, padded_element
	} else if element_type == "array" {
		//mapping_key := append(key[0:len(key)-1], "mapping")
		element_mapping := get_mapping(mapping_key, copybook)
		element_index := get_index(append(key[0:len(key)-1], "index"), copybook)
		//fmt.Println("destinations:", key[0:len(key)-1])
		//fmt.Println("destiantion key:", element_mapping)
		new_obj := get_value(copybook, append(key[0:len(key)-1], "items")).(map[string]interface{})
		//fmt.Println("destination object:", new_obj)
		new_keys := obj_keys(new_obj, "")
		//for _, k := range new_keys {
		//	fmt.Println(k)
		//}
		arr := ""
		element := get_value(data, strings.Split(element_mapping, ".")).([]interface{})
		for _, e := range element {
			//fmt.Println(e)
			element_keys := obj_keys(e.(map[string]interface{}), "")
			for _, ek := range element_keys {
				sub_key := element_mapping + "." + ek
				//fmt.Println("sub key:", sub_key)
				for _, k := range new_keys {
					key := strings.Split(k, ".")
					mapps := get_mapping(key, new_obj)
					if mapps == sub_key {
						//fmt.Println("k: ", k)
						//fmt.Println("match:", mapps)
						length := get_length(append(key[0:len(key)-1], "maxLength"), new_obj)
						index := get_index(append(key[0:len(key)-1], "index"), new_obj)
						fmt.Println("element:", index)
						//fmt.Println("length:", length)
						elem := padding(get_value(e.(map[string]interface{}), []string{ek}).(string), length, "-")
						//fmt.Println("element:", elem)
						arr += elem
					}
				}
			}
		}
		//fmt.Println(arr)
		return element_index, arr
	}
	return 0, ""
}
*/

// USED FUNCTION
// Return copybook object
func get_element(key []string, copybook map[string]interface{}, data map[string]interface{}) (int, string) {
	element_type := get_type(key, copybook)
	root_key := key[0 : len(key)-1]
	mapping_key := append(root_key, "mapping")
	element_mapping := get_mapping(mapping_key, copybook)
	element_index := get_index(append(root_key, "index"), copybook)
	if (element_type != "array") && (element_type != "element") {
		element, _ := coerce.ToString(get_value(data, strings.Split(element_mapping, ".")))
		element_length := get_length(append(root_key, "maxLength"), copybook)
		padded_element := padding(element, element_length, "-")
		return element_index, padded_element
	} else if element_type == "array" {
		elements := get_array_elements(root_key, element_mapping, copybook, data)
		return element_index, elements
	}
	return 0, ""
}

// Return ordered elements of a repetitive copybook object
func get_array_elements(root_key []string, element_mapping string, copybook map[string]interface{}, data map[string]interface{}) string {
	new_obj := get_value(copybook, append(root_key, "items")).(map[string]interface{})
	new_keys := ObjKeys(new_obj, "")
	max_items := get_maxItems(append(root_key, "maxItems"), copybook)
	elements := make(map[int]string)
	element := get_value(data, strings.Split(element_mapping, ".")).([]interface{})
	n_elements_req := len(element)
	for i, e := range element {
		element_keys := ObjKeys(e.(map[string]interface{}), "")
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
					elem := padding(get_value(e.(map[string]interface{}), []string{ek}).(string), length, "-")
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
	val := JoinElements(elements)
	result := val + filler
	return result
}

func GetElements(keys []string, copybook map[string]interface{}, data map[string]interface{}) map[int]string {
	elements := make(map[int]string)
	for _, k := range keys {
		if strings.Contains(k, "type") {
			//fmt.Println(k)
			i, element := get_element(strings.Split(k, "."), copybook, data)
			if len(element) > 0 {
				elements[i] = element
			}
		}
	}
	return elements
}

func JoinElements(data map[int]string) string {
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
		val_length, _ := coerce.ToInt(get_value(copybook, key))
		length := int(val_length)
		return length
	}
	return 0
}

func get_index(key []string, copybook map[string]interface{}) int {
	if contains(key, "index") {
		val_index, _ := coerce.ToInt(get_value(copybook, key))
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

func main() {
	var message interface{}
	var copyB interface{}
	message, _ = test(request)
	copyB, _ = test(copybook)
	data := message.(map[string]interface{})
	dataCopy := copyB.(map[string]interface{})
	keys := ObjKeys(dataCopy, "")

	elements := GetElements(keys, dataCopy, data)

	fmt.Println()

	copy := JoinElements(elements)

	fmt.Println(copy)

	fmt.Println()

}
