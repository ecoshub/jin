package jin

// import (
// 	"jin"
// 	// "github.com/ecoshub/jin"
// 	"github.com/valyala/fastjson"
// 	"testing"
// )

// func BenchmarkFastjsonGetSmall(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		var p fastjson.Parser
// 		prs, _ := p.ParseBytes(smallfixture)
// 		prs.Get("uuid")
// 		prs.Get("tz")
// 		prs.Get("ua")
// 		prs.Get("st")
// 	}
// }

// func BenchmarkJinParseGetSmall(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		prs, _ := jin.Parse(smallfixture)
// 		prs.Get("uuid")
// 		prs.Get("tz")
// 		prs.Get("ua")
// 		prs.Get("st")
// 	}
// }

// func BenchmarkFastjsonGetMedium(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		var p fastjson.Parser
// 		prs, _ := p.ParseBytes(mediumfixture)
// 		prs.Get("person", "name", "fullName")
// 		prs.Get("person", "github", "followers")
// 		prs.Get("company")
// 	}
// }

// func BenchmarkJinParseGetMedium(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		prs, _ := jin.Parse(mediumfixture)
// 		prs.Get("person", "name", "fullName")
// 		prs.Get("person", "github", "followers")
// 		prs.Get("company")
// 	}
// }

// func BenchmarkFastjsonGetLarge(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		var p fastjson.Parser
// 		prs, _ := p.ParseBytes(largefixture)
// 		prs.Get("users", "0")
// 		prs.Get("users", "31")
// 		prs.Get("topics", "topics", "0")
// 		prs.Get("topics", "topics", "29")
// 	}
// }

// func BenchmarkJinParseGetLarge(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		prs, _ := jin.Parse(largefixture)
// 		prs.Get("users", "0")
// 		prs.Get("users", "31")
// 		prs.Get("topics", "topics", "0")
// 		prs.Get("topics", "topics", "29")
// 	}
// }
