package snowflakex

import "github.com/bwmarrin/snowflake"

func GetId() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}

func GetNonce() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	return node.Generate().Base64(), nil
}

func GetIds(n int) (arr []int64, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	for i := 0; i < n; i++ {
		arr = append(arr, node.Generate().Int64())
	}
	return
}
