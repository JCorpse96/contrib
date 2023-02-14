package copybookrenderh001

import (
	"encoding/json"
	"fmt"

	"github.com/JCorpse96/core/activity"
	"github.com/JCorpse96/core/data/copybook"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{}, New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

var cpy_id = copybook.Element{ElementIndex: 0, ElementName: "CPY_ID", ElementMaxLength: 10, ElementFormat: copybook.Integer{TypeOf: "number"}}

var cpy_data_type = copybook.Element{ElementIndex: 1, ElementName: "CPY_DATA_TYPE", ElementMaxLength: 4, ElementFormat: copybook.Integer{TypeOf: "number"}}

var cpy_data_number = copybook.Element{ElementIndex: 2, ElementName: "CPY_DATA_NUMBER", ElementMaxLength: 30, ElementFormat: copybook.Integer{TypeOf: "number"}}

var cpy_object_entity = copybook.Element{ElementIndex: 3, ElementName: "CPY_OBJECT_ENTITY", ElementMaxLength: 8, ElementFormat: copybook.Integer{TypeOf: "number"}}

var cpy_object_reference = copybook.Element{ElementIndex: 4, ElementName: "CPY_OBJECT_REFERENCE", ElementMaxLength: 11, ElementFormat: copybook.Integer{TypeOf: "number"}}

var cpy_object_value = copybook.Element{ElementIndex: 5, ElementName: "CPY_OBJECT_VALUE", ElementMaxLength: 7, ElementFormat: copybook.Float{Decimals: 2, TypeOf: "number"}}

//array

var cpy_destinations_acc = copybook.Element{ElementIndex: 0, ElementName: "CPY_DESTINATIONS_ACC", ElementMaxLength: 6, ElementFormat: copybook.Integer{TypeOf: "number"}}

var cpy_destinations_name = copybook.Element{ElementIndex: 1, ElementName: "CPY_DESTINATIONS_NAME", ElementMaxLength: 5, ElementFormat: copybook.String{TypeOf: "string"}}

var cpy_destinations = copybook.Element{ElementIndex: 6, ElementName: "CPY_DESTINATIONS", ElementMaxLength: 6, ElementFormat: copybook.Array{TypeOf: "array", Elements: []copybook.Element{cpy_destinations_acc, cpy_destinations_name}}}

//end array

var cpy_location = copybook.Element{ElementIndex: 7, ElementName: "CPY_LOCATION", ElementMaxLength: 17, ElementFormat: copybook.String{TypeOf: "string"}}

var cpy_source = copybook.Element{ElementIndex: 8, ElementName: "CPY_SOURCE", ElementMaxLength: 3, ElementFormat: copybook.String{TypeOf: "string"}}

type H001 struct {
	copybook.HXXX `json:"H001"`
	jsonString    string
	mapStruct     map[string]interface{}
}

func (h H001) setJsonString() {
	h.jsonString = h.getHJson()
}

func (h H001) setMapStruct() {
	h.mapStruct, _ = copybook.MapJson(h.jsonString)
}

func initialize() H001 {
	h001 := H001{copybook.HXXX{Elements: []copybook.Element{cpy_id, cpy_data_type, cpy_data_number, cpy_object_entity, cpy_object_reference, cpy_object_value, cpy_destinations, cpy_location, cpy_source}}, "", nil}
	h001.setJsonString()
	h001.setMapStruct()
	return h001
}

func (h H001) getHJson() string {
	h_ind, _ := json.MarshalIndent(h, "", "  ")
	return string(h_ind)
}

// New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	h001 := initialize()

	var message interface{}
	message, _ = coerce.ToObject(input.Request)
	req := message.(map[string]interface{})

	var inInterface map[string]string
	inrec, _ := json.Marshal(input)
	json.Unmarshal(inrec, &inInterface)

	rendered := h001.HXXX.RenderCopybook(inInterface, req)
	fmt.Println(rendered)

	ctx.Logger().Debugf("Input: %s", input.CPY_ID)

	output := &Output{Rendered: rendered}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
