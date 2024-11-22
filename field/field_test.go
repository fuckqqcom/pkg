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

func TestField_ApplyTo1(t *testing.T) {
	type Person struct {
		Name        string
		Age         int
		Email       string
		PhoneNumber string
	}
	p := &Person{Age: 60}
	f := NewField()
	f.SetVal("name", "John", SkipFunc(func() bool {
		return false
	})).
		SetVal("email", "john@example.com",
			SkipFunc(func() bool {
				return false
			}),
			ValFunc(func() any {
				return "new-email@example.com"
			}),
		).SetVal("phoneNumber", "1190",
		SkipFunc(func() bool {
			return true
		}),
		ValFunc(func() any {
			return "1190"
		}),
	)
	f.Bind(p)
	//{ John 60 new-email@example.com}，
	fmt.Println(p)

}

func TestField_ApplyTo2(t *testing.T) {
	type Person struct {
		Name        string
		Age         int
		Email       string
		PhoneNumber string
	}
	p := &Person{Age: 60}
	check := NewField().SetVal("name", "John", SkipFunc(func() bool {
		return true
	})).
		SetVal("email", "john@example.com",
			SkipFunc(func() bool {
				return true
			}),
			ValFunc(func() any {
				return "new-email@example.com"
			}),
		).SetVal("phoneNumber", "1190",
		SkipFunc(func() bool {
			return true
		}),
		ValFunc(func() any {
			return "1190"
		}),
	).Bind(p).Check()
	//{ John 60 new-email@example.com}，
	fmt.Println(p)
	fmt.Println(check)

}
