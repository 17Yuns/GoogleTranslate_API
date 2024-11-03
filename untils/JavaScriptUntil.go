package untils

import (
	"github.com/robertkrimen/otto"
	"os"
)

func JavaScript(filePath string, function string, data string) string {
	//先读入文件内容
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}
	value, err := vm.Call(function, nil, data)
	if err != nil {
		panic(err)
	}
	return value.String()
}
