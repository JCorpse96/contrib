package padding

import (
	"github.com/JCorpse96/contrib/function/padding/util"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnPaddingRight struct {
}

func init() {
	function.Register(&fnPaddingRight{})
}

func (s *fnPaddingRight) Name() string {
	return "paddingRight"
}

func (s *fnPaddingRight) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeString}, false
}

func (s *fnPaddingRight) Eval(in ...interface{}) (interface{}, error) {

	//var output string

	value := in[0].(string)
	length := in[1].(int)
	padCharacter := in[2].(string)

	//fmt.Println("value ", value)
	//fmt.Println("length ", length)
	//fmt.Println("padCharacter ", padCharacter)

	//fmt.Println("output ", output)
	//return output, fmt.Errorf("fnPaddingLeft function must have three arguments")

	retStr := util.PaddingRight(value, length, padCharacter)

	return retStr, nil

}
