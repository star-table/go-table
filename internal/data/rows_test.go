package data

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/spf13/cast"

	"github.com/star-table/go-common/utils/unsafe"

	"github.com/star-table/go-common/pkg/encoding"

	"github.com/star-table/go-table/internal/data/po"

	"github.com/star-table/go-table/internal/data/consts"
)

var flagconf string

func init() {
	encoding.Init()
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf config.yaml")
}

func TestPgData(t *testing.T) {
	fmt.Println(cast.ToFloat64(math.MaxInt64))
	fmt.Println(fmt.Sprintf("%c", consts.ColumnIdRandomKey[10]))
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		"pgm-uf6nfhul11220z6w.pg.rds.aliyuncs.com", "polarisgray", "runx@123", "postgres", 5432)
	fmt.Println(dsn)
	db := openDB("postgres", dsn).Debug()

	data := po.NewRow()
	//fmt.Println(db.Exec(`insert into "lc_data"(id,"orgId","appId","projectId","tableId","creator","issueId","updator") values('123','47964','124','124','124','124','124','123')`).Error)
	db.Raw(`select "id","issueId","data" :: jsonb -> 'auditStatus' "auditStatus","parentId","data" :: jsonb -> 'title' "title","data" :: jsonb -> 'isParentDesc' "isParentDesc","code","data" :: jsonb -> 'ownerId' "ownerId","data" :: jsonb -> 'issueStatus' "issueStatus","data" :: jsonb -> 'planStartTime' "planStartTime","data" :: jsonb -> 'planEndTime' "planEndTime","data" :: jsonb -> 'remark' "remark","data" :: jsonb -> 'followerIds' "followerIds","data" :: jsonb -> '_field_priority' "_field_priority","creator","updator","data" :: jsonb -> 'baRelating' "baRelating","data" :: jsonb -> 'G6cm' "G6cm","data" :: jsonb -> 'eApJ' "eApJ" from "lc_data" where (("recycleFlag" = '2'  and "tableId" = '1601424466779312128'  and "orgId" = '47964' )) limit 8000 offset 0`).Find(data)

	//db.Table("lc_data").Delete(&data, jsonb.NewQuery(&tablePb.Condition{
	//	Type: tablePb.ConditionType_and,
	//	Conditions: []*tablePb.Condition{
	//		{
	//			Column: `"orgId"`,
	//			Type:   tablePb.ConditionType_equal,
	//			Value:  `[47964]`,
	//		},
	//	},
	//}))
	//fmt.Println(db.Table("_form_2373_1499356343193141249").Select([]string{`data->'tableId'`, `data->'projectId'`}).Find(&data, jsonb.NewQuery(&tablePb.Condition{
	//	Type:  "and",
	//	Value: ``,
	//	Conditions: []*tablePb.Condition{
	//		{Type: "in", Value: `["1595345131723034624"]`, Column: `data->'tableId'`},
	//		{Type: "equal", Value: `[61902]`, Column: `data->'projectId'`},
	//	},
	//})).Error)
	fmt.Println(data.Buf.String())
}

func TestList(t *testing.T) {
	userIds, _ := encoding.GetJsonCodec().Marshal([]int64{1, 2, 3})
	deptIds, _ := encoding.GetJsonCodec().Marshal([]int64{1, 2, 3})
	replyStr := fmt.Sprintf(`{"userIds":%s, "deptIds":%s, "lastUpdateTime":"%s", "count":%d}`,
		unsafe.BytesString(userIds), unsafe.BytesString(deptIds), "123123", 10)

	buffer := &bytes.Buffer{}
	buffer.Grow(200 + len(replyStr) + 2)
	buffer.WriteString(replyStr)
	buffer.WriteString("#")
	buffer.Write([]byte(`[123,457]`))

	fmt.Println(string(buffer.Bytes()))
	respStr := string(buffer.Bytes())
	idx := strings.Index(respStr, "#")
	fmt.Println(respStr[:idx])
	fmt.Println(respStr[idx+1:])
}

type Scanner interface {
	Scan(values []interface{}, columns []string)
}

type ScannerSelf struct {
}

func (s *ScannerSelf) Scan(values []interface{}, columns []string) {

}

func TestQuerySql(t *testing.T) {
	s := &ScannerSelf{}
	checkType(s)
}

func checkType(v interface{}) {
	switch dest := v.(type) {
	case Scanner:
		fmt.Println(dest)
	}
}
