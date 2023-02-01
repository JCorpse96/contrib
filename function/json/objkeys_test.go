package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

const input = `{
    "content": {
    	"code": "123456",
        "data": {
            "type": "2",
            "number": "0012111122223333444455556"
        	},
        "object": {
            "entity": "11223",
            "reference": "123321123",
            "value": "3.45"
        },
        "location": "Location",
    	"source": "B"
    }
}`

const mapping = `{
	"v3": {
		"msg.id": 6,
		"msg.data.type": 1,
		"msg.data.number": 25,
		"msg.object.entity": 5,
		"msg.object.reference": 9,
		"msg.object.value": 5,
		"msg.location": 15,
		"msg.source": 1
	}
}`

func TestFnObjKeys(t *testing.T) {
	//data := make(map[string]interface{})
	var data interface{}
	err := json.Unmarshal([]byte(mapping), &data)
	assert.Nil(t, err)
	f := &fnObjKeys{}
	v, err := function.Eval(f, data)
	fmt.Println(v)
	assert.Nil(t, err)

}
