package jin

import (
	"testing"

	"github.com/ecoshub/jin"
	"github.com/valyala/fastjson"
)

func BenchmarkFastjsonSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		prs, _ := p.ParseBytes(smallfixture)
		prs.Set("uuid", fastjson.MustParse(`"test-value"`))
		prs.Set("tz", fastjson.MustParse(`"test-value"`))
		prs.Set("ua", fastjson.MustParse(`"test-value"`))
		prs.Set("st", fastjson.MustParse(`"test-value"`))
	}
}

func BenchmarkJinParseSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs, _ := jin.Parse(smallfixture)
		prs.Set([]byte(`"test-value"`), "uuid")
		prs.Set([]byte(`"test-value"`), "tz")
		prs.Set([]byte(`"test-value"`), "ua")
		prs.Set([]byte(`"test-value"`), "st")
	}
}
