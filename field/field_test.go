package field

import (
	"fmt"
	"testing"
)

func TestCropModel(t *testing.T) {

	type A struct {
		Name        string `protobuf:"varint,6,opt,name=name,json=name,proto3,oneof" json:"name,omitempty"`
		Age         int    `protobuf:"varint,6,opt,name=age,json=age,proto3,oneof" json:"age,omitempty"`
		PhoneNumber string `protobuf:"varint,6,opt,name=phone_number,json=phone_number,proto3,oneof" json:"phone_number,omitempty"`
	}
	a := A{Age: 100, Name: "xiaohan", PhoneNumber: "phoneNumber"}
	//CropObjFields(&a, []string{"name"})
	//CropObjFields
	//marshal, err := json.Marshal(a)
	//fmt.Println("a", a, string(marshal), err)

	CropObjFields(&a, []string{"phoneNumber"})
	fmt.Println(a)
}

func TestField_ApplyTo(t *testing.T) {
	type Person struct {
		Name  string
		Age   int
		Email string
	}
	p := &Person{}

	f := NewField()

	f.SetVal("Name", "John").
		Next().
		SetVal("Age", 30).
		Next().
		SetVal("Email", "john@example.com",
			f.SkipFunc(func() bool {
				// 模拟条件判断，决定是否跳过字段
				return false
			}),
			f.ValFunc(func() any {
				// 动态计算值
				return "new-email@example.com"
			}),
		)

	// 应用配置到对象
	errs := f.ApplyTo(p)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Updated Person:", p)
	}
}
