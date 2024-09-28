package convert

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"testing"
)

func TestAnyToInt64(t *testing.T) {

}

func TestStructToMap(t *testing.T) {
	node, err := snowflake.NewNode(1)
	fmt.Println(node.Generate().String(), err)
}
