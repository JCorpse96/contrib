package copybookrenderh001

import "github.com/JCorpse96/core/data/coerce"

type Input struct {
	CPY_ID                string `md:"CPY-ID"`
	CPY_DATA_TYPE         string `md:"CPY-DATA-TYPE"`
	CCPY_DATA_NUMBER      string `md:"CCPY-DATA-NUMBER"`
	CPY_OBJECT_ENTITY     string `md:"CPY-OBJECT-ENTITY"`
	CPY_OBJECT_REFERENCE  string `md:"CPY-OBJECT-REFERENCE"`
	CPY_OBJECT_VALUE      string `md:"CPY-OBJECT-VALUE"`
	CPY_DESTINATIONS      string `md:"CPY-DESTINATIONS"`
	CPY_DESTINATIONS_ACC  string `md:"CPY-DESTINATIONS-ACC"`
	CPY_DESTINATIONS_NAME string `md:"CPY-DESTINATIONS-NAME"`
	CPY_LOCATION          string `md:"CPY-LOCATION"`
	CPY_SOURCE            string `md:"CPY-SOURCE"`
	Request               string `md:"request"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	i.CPY_ID, _ = coerce.ToString(values["CPY-ID"])
	i.CPY_DATA_TYPE, _ = coerce.ToString(values["CPY-DATA-TYPE"])
	i.CCPY_DATA_NUMBER, _ = coerce.ToString(values["CCPY-DATA-NUMBER"])
	i.CPY_OBJECT_ENTITY, _ = coerce.ToString(values["CPY-OBJECT-ENTITY"])
	i.CPY_OBJECT_REFERENCE, _ = coerce.ToString(values["CPY-OBJECT-REFERENCE"])
	i.CPY_OBJECT_VALUE, _ = coerce.ToString(values["CPY-OBJECT-VALUE"])
	i.CPY_DESTINATIONS, _ = coerce.ToString(values["CPY-DESTINATIONS"])
	i.CPY_DESTINATIONS_ACC, _ = coerce.ToString(values["CPY-DESTINATIONS-ACC"])
	i.CPY_DESTINATIONS_NAME, _ = coerce.ToString(values["CPY-DESTINATIONS-NAME"])
	i.CPY_LOCATION, _ = coerce.ToString(values["CPY-LOCATION"])
	i.CPY_SOURCE, _ = coerce.ToString(values["CPY-SOURCE"])
	i.Request, _ = coerce.ToString(values["request"])
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"CPY-ID":                i.CPY_ID,
		"CPY-DATA-TYPE":         i.CPY_DATA_TYPE,
		"CCPY-DATA-NUMBER":      i.CCPY_DATA_NUMBER,
		"CPY-OBJECT-ENTITY":     i.CPY_OBJECT_ENTITY,
		"CPY-OBJECT-REFERENCE":  i.CPY_OBJECT_REFERENCE,
		"CPY-OBJECT-VALUE":      i.CPY_OBJECT_VALUE,
		"CPY-DESTINATIONS":      i.CPY_DESTINATIONS,
		"CPY-DESTINATIONS-ACC":  i.CPY_DESTINATIONS_ACC,
		"CPY-DESTINATIONS-NAME": i.CPY_DESTINATIONS_NAME,
		"CPY-LOCATION":          i.CPY_LOCATION,
		"CPY-SOURCE":            i.CPY_SOURCE,
		"request":               i.Request,
	}
}

type Output struct {
	Rendered string `md:"rendered"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.Rendered, _ = coerce.ToString(values["rendered"])
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"rendered": o.Rendered,
	}
}
