package copybook

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
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

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)
	fmt.Println(ref)
	fmt.Println(act)
	f := activity.GetFactory(ref)
	fmt.Println(f)
	assert.NotNil(t, act)
	assert.NotNil(t, f)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	input := &Input{Request: request, Copybook: copybook}
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	err = tc.GetOutputObject(output)
	fmt.Println(output.Result)
	assert.Nil(t, err)
	//assert.Equal(t, request+copybook, output.Result)
}
