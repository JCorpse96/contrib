package padding

import (
	"fmt"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnPaddingLeft struct {
}

func init() {
	function.Register(&fnPaddingLeft{})
}

func (s *fnPaddingLeft) Name() string {
	return "paddingLeft"
}

func (s *fnPaddingLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeString}, false
}

func (s *fnPaddingLeft) Eval(in ...interface{}) (interface{}, error) {

	//var output string

	value := in[0].(string)
	length := in[1].(int)
	padCharacter := in[2].(string)

	fmt.Println("value ", value)
	fmt.Println("length ", length)
	fmt.Println("padCharacter ", padCharacter)

	//for i := len(value); i <= length; i++ {
	//	output = padCharacter + value
	//}

	//fmt.Println("output ", output)
	//return output, fmt.Errorf("fnPaddingLeft function must have three arguments")

	var padCountInt = 1 + ((length - len(padCharacter)) / len(padCharacter))
	var retStr = strings.Repeat(padCharacter, padCountInt) + value
	fmt.Println("retStr ", retStr[(len(retStr)-length):])
	return retStr[(len(retStr) - length):], nil

}
