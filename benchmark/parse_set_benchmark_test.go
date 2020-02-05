package jin

import (
	"github.com/valyala/fastjson"
	"jin"
	"testing"
)

func BenchmarkFastjsonSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		prs, _ := p.ParseBytes(SmallFixture)
		prs.Set("uuid", fastjson.MustParse(`"test-value"`))
		prs.Set("tz", fastjson.MustParse(`"test-value"`))
		prs.Set("ua", fastjson.MustParse(`"test-value"`))
		prs.Set("st", fastjson.MustParse(`"test-value"`))
	}
}

func BenchmarkJintParseSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs, _ := jin.Parse(SmallFixture)
		prs.Set([]byte(`"test-value"`), "uuid")
		prs.Set([]byte(`"test-value"`), "tz")
		prs.Set([]byte(`"test-value"`), "ua")
		prs.Set([]byte(`"test-value"`), "st")
	}
}
