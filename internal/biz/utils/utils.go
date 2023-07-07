package utils

import (
	"github.com/star-table/go-common/pkg/encoding"
	"github.com/star-table/go-common/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func ProtoStructToModel(p *structpb.Struct, m interface{}) error {
	bts, err := encoding.GetJsonCodec().Marshal(p)
	if err != nil {
		return err
	}

	return encoding.GetJsonCodec().Unmarshal(bts, m)
}

// CopyProto 由于proto的struct复制会有问题，所以只能转换成pb再转回来
func CopyProto(src, dest proto.Message) error {
	bts, err := proto.Marshal(src)
	if err != nil {
		return errors.WithStack(err)
	}

	err = proto.Unmarshal(bts, dest)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func ModelToProtoStruct(m interface{}) (*structpb.Value, error) {
	bts, err := encoding.GetJsonCodec().Marshal(m)
	if err != nil {
		return nil, err
	}
	temp := map[string]interface{}{}
	err = encoding.GetJsonCodec().Unmarshal(bts, &temp)
	if err != nil {
		return nil, err
	}

	return structpb.NewValue(temp)
}
