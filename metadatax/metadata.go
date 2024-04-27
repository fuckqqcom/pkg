package metadatax

import (
	"context"
	"encoding/json"
	"github.com/fuckqqcom/pkg/constantx"
	"github.com/fuckqqcom/pkg/convertx"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ExtractedUserInfo(ctx context.Context) (*UserInfo, error) {
	val := GetValFromCtx(ctx, constantx.UserInfo)
	var user *UserInfo
	bytes, err := convertx.ToBytes(val)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bytes, &user); err != nil {
		return user, err
	}
	return user, nil
}

func GetMetaDataCtx(ctx context.Context, key string) (any, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	payloads, ok := md[key]
	if !ok {
		return 0, status.Errorf(412, "rpc metadata not found")
	}
	//payload := ToString(payloads[0])
	//if err != nil || payload == 0 {
	//	return 0, status.Errorf(412, err.Error())
	//}

	payload := payloads[0]
	if payload == "" {
		return 0, status.Errorf(412, "rpc metadata not found")
	}
	return payload, nil
}

func MetadataCtx(ctx context.Context, userId string) context.Context {
	md := metadata.New(map[string]string{
		constantx.UserId: userId,
	})
	return metadata.NewOutgoingContext(ctx, md)
}
