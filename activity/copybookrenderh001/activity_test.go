package copybookrenderh001

import (
	"testing"

	"github.com/JCorpse96/core/activity"
	"github.com/JCorpse96/core/support/test"
	"github.com/stretchr/testify/assert"
)

const data = `{
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

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	input := &Input{CPY_ID: "msg.id", CPY_DATA_TYPE: "msg.data.type", CCPY_DATA_NUMBER: "msg.data.number", CPY_OBJECT_ENTITY: "msg.object.entity", CPY_OBJECT_REFERENCE: "msg.object.reference", CPY_OBJECT_VALUE: "msg.object.value", CPY_DESTINATIONS: "msg.destinations", CPY_DESTINATIONS_ACC: "msg.destinations.number", CPY_DESTINATIONS_NAME: "msg.destinations.name", CPY_LOCATION: "msg.location", CPY_SOURCE: "msg.source", Request: data}
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)
	assert.Equal(t, "0000123456000200000001211112222333344445555600011223001233211230000345000123ABC--000456DEF--000789GHI--101112JKL--000000-----000000-----Rua da esquina---B--", output.Rendered)
}
