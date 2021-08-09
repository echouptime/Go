package main

import "fmt"

//定义空接口
var (
	userInfo []interface{}
	t_string []string
	t_bool []bool
	tint []int
	t_float []float64
)


//接口类型断言
func type_filter (info interface{}) interface{} {
	switch obj := info.(type) {
	case int:
		tint = append(tint, obj)
	case string:
		t_string = append(t_string, obj)
	case bool:
		t_bool = append(t_bool,obj)
	case float64:
		t_float = append(t_float,obj)
	}
	return nil
}



func main() {

	//空接口写数据
	userInfo = append(userInfo,  2,false,0.81313,6,5,10,11,12,13,true,"abcd","ok",0.123)
	//fmt.Print("%",userInfo)

	//循环接口数据
	for _, vaule := range userInfo {
		type_filter(vaule)
		//fmt.Printf("%#v",vaule)

	}
	fmt.Printf("string_list :%#v\n",t_string)
	fmt.Printf("bool_list :%#v\n",t_bool)
	fmt.Printf("int_list :%#v\n",tint)
	fmt.Printf("float_list :%#v\n",t_float)

}
