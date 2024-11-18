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
