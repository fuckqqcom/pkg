package fieldx

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"testing"
)

// 假设这是一个简单的Proto Message定义
type MyProtoMessage struct {
	ID    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age   int32  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Email string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *MyProtoMessage) Reset() { *m = MyProtoMessage{} }
func (m *MyProtoMessage) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Age: %d, Email: %s", m.ID, m.Name, m.Age, m.Email)
}
func (m *MyProtoMessage) ProtoReflect() protoreflect.Message {
	// 你可以返回一个带有字段描述的实现
	// 这里只是一个简化的示例
	return nil
}

// 测试 CropProtoFields 函数
func TestFilterProtoFields(t *testing.T) {
	tests := []struct {
		name           string
		includeFields  []string
		excludeFields  []string
		merge          bool
		expectedResult map[string]interface{}
	}{
		{
			name:          "Test with includeFields only",
			includeFields: []string{"ID", "Name"},
			excludeFields: []string{"Email"},
			merge:         false,
			expectedResult: map[string]interface{}{
				"ID":   "123",
				"Name": "Alice",
			},
		},
		{
			name:          "Test with includeFields and excludeFields, merge = false",
			includeFields: []string{"ID", "Name", "Age"},
			excludeFields: []string{"Email"},
			merge:         false,
			expectedResult: map[string]interface{}{
				"ID":   "123",
				"Name": "Alice",
				"Age":  25,
			},
		},
		{
			name:          "Test with includeFields and excludeFields, merge = true",
			includeFields: []string{"ID", "Name", "Age"},
			excludeFields: []string{"Email"},
			merge:         true,
			expectedResult: map[string]interface{}{
				"ID":   "123",
				"Name": "Alice",
				"Age":  25,
			},
		},
		{
			name:          "Test with includeFields and excludeFields, merge = true and fieldx in both",
			includeFields: []string{"ID", "Name", "Age"},
			excludeFields: []string{"Age", "Email"},
			merge:         true,
			expectedResult: map[string]interface{}{
				"ID":   "123",
				"Name": "Alice",
			},
		},
		{
			name:          "Test with all fields excluded in merge mode",
			includeFields: []string{"ID", "Name", "Age", "Email"},
			excludeFields: []string{"ID", "Name", "Age"},
			merge:         true,
			expectedResult: map[string]interface{}{
				"Email": "alice@example.com",
			},
		},
	}

	// 创建测试 Proto Message 实例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建一个示例 MyProtoMessage
			protoMessage := &MyProtoMessage{
				ID:    "123",
				Name:  "Alice",
				Age:   25,
				Email: "alice@example.com",
			}

			// 调用 CropProtoFields 函数
			result := FilterProtoField(protoMessage, tt.includeFields, tt.excludeFields, tt.merge)

			// 验证返回结果是否符合预期
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
