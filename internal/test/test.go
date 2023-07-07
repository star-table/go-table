package main

import (
	"fmt"

	"github.com/star-table/go-common/pkg/encoding"
	pb "github.com/star-table/interface/golang/table/v1"
)

func main() {
	str := `{"name":"_field_1649316629903","label":"新建单选列","field":{"props":{"options":[{"color":"#377AFF","id":1,"value":"未开始"},{"color":"#4FBBEF","id":2,"value":"进行中"},{"color":"#4CBFAB","id":3,"value":"已完成"}],"select":{"options":[{"color":"#377AFF","id":1,"value":"未开始"},{"color":"#4FBBEF","id":2,"value":"进行中"},{"color":"#4CBFAB","id":3,"value":"已完成"}]}}}}`
	c := &pb.Column{}
	fmt.Println(encoding.GetJsonCodec().Unmarshal([]byte(str), c))
	fmt.Println(c.Field.Type.String())
}
