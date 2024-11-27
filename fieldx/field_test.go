package fieldx

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
	f.SetVal("name", "John", WithSkipFunc(func() bool {
		return false
	})).
		SetVal("email", "john@example.com",
			WithSkipFunc(func() bool {
				return false
			}),
			WithValFunc(func() any {
				return "new-email@example.com"
			}),
		).SetVal("phoneNumber", "1190",
		WithSkipFunc(func() bool {
			return true
		}),
		WithValFunc(func() any {
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
	check := NewField().SetVal("name", "John", WithSkipFunc(func() bool {
		return false
	})).
		SetVal("email", "john@example.com",
			WithSkipFunc(func() bool {
				return true
			}),
			WithValFunc(func() any {
				return "new-email@example.com"
			}),
		).SetVal("phoneNumber", "1190",
		WithSkipFunc(func() bool {
			return true
		}),
		WithValFunc(func() any {
			return "1190"
		}),
	).SetIgnoreKey([]string{"name"}).Bind(p).Check()
	//{ John 60 new-email@example.com}，
	fmt.Println(p)
	fmt.Println(check)

}
