package copybook

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Request  string `md:"request"`
	Copybook string `md:"copybook"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//strVal, _ := coerce.ToString(values["request"])
	//i.Request = strVal
	var err error
	i.Request, err = coerce.ToString(values["request"])
	if err != nil {
		return err
	}
	i.Copybook, err = coerce.ToString(values["copybook"])
	if err != nil {
		return err
	}
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"request":  i.Request,
		"copybook": i.Copybook,
	}
}

type Output struct {
	Result string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	//strVal, _ := coerce.ToString(values["result"])
	//o.Result = strVal
	var err error
	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
