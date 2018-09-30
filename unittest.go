package unittest

import (
	"reflect"
	"runtime"
	"testing"
)

func D(lineCnt int, args ...interface{}) (result [][]interface{}) {
	line := []interface{}{}

	result = append(result)
	for _, v := range args {
		line = append(line, v)
		if len(line) == lineCnt {
			result = append(result, line)
			line = []interface{}{}
		}
	}
	return

}
func CallTestFunc(t *testing.T, fun interface{}, data [][]interface{}, haveDesc bool) {
	rt:=reflect.TypeOf(fun)
	inputLen:=rt.NumIn()
	outputLen:=rt.NumOut()
	descLen:=0
	if haveDesc {
		descLen=1
	}

	for _, d := range data {
		input := []reflect.Value{}
		for i := 0; i < inputLen; i++ {
			input = append(input, reflect.ValueOf(d[descLen+i]))
		}
		ot := reflect.ValueOf(fun).Call(input)

		out := []interface{}{}
		for _, x := range ot {
			out = append(out, x.Interface())
		}

		for i := 0; i < outputLen; i++ {
			//if out[i] != d[descLen+inputLen+i] {
			if !reflect.DeepEqual(out[i], d[descLen+inputLen+i]) {
				var descV interface{} = ""
				if descLen > 0 {
					descV = d[:descLen]
				}
				inputV := d[descLen : descLen+inputLen]
				outV := d[descLen+inputLen:]
				_, file, line, _ := runtime.Caller(1)

				t.Errorf("%s:%d\ndesc:%v\ninput:%v\n期待:%v\n实际:%v\n", file, line, descV, inputV, outV, out)
				break
			}

		}
	}

}
 func AssertEquals(t *testing.T,result interface{},active interface{},desc string){

	 if !reflect.DeepEqual(result,active) {
		 t.Errorf("desc:%v\n期待:%v\n实际:%v\n", desc, result, active)
	 }
 }