package copybook

import (
	"github.com/JCorpse96/contrib/activity/copybook/util"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{}, New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//var activityMd = activity.ToMetadata(&Input{}, &Output{})

// New optional factory method, should be used if one activity instance per configuration is desired

func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{}

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
	data, _ := coerce.ToObject(input.Request)
	//fmt.Println(data)

	copy, _ := coerce.ToObject(input.Copybook)
	//fmt.Println(copy)

	keys := util.ObjKeys(copy, "")
	//fmt.Println(keys)

	elements := util.GetElements(keys, copy, data)
	res := util.JoinElements(elements)

	ctx.Logger().Debugf("Input: %s", input.Request)

	output := &Output{Result: res}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
