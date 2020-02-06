package jin

import (
	"github.com/buger/jsonparser"
	"jin"
	"testing"
)

func BenchmarkJsonParserSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Set(smallfixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "uuid")
		jsonparser.Set(smallfixture, []byte("-3"), "tz")
		jsonparser.Set(smallfixture, []byte(`"server_agent"`), "ua")
		jsonparser.Set(smallfixture, []byte("3"), "st")
	}
}

func BenchmarkJinSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Set(smallfixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "uuid")
		jin.Set(smallfixture, []byte("-3"), "tz")
		jin.Set(smallfixture, []byte(`"server_agent"`), "ua")
		jin.Set(smallfixture, []byte("3"), "st")
	}
}

func BenchmarkJsonParserSetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Set(mediumfixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "person", "name", "fullName")
		jsonparser.Set(mediumfixture, []byte("-3"), "person", "github", "followers")
		jsonparser.Set(mediumfixture, []byte(`"server_agent"`), "company")
	}
}

func BenchmarkJinSetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Set(mediumfixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "person", "name", "fullName")
		jin.Set(mediumfixture, []byte("-3"), "person", "github", "followers")
		jin.Set(mediumfixture, []byte(`"server_agent"`), "company")
	}
}

func BenchmarkJsonParserSetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Set(largefixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "[0]", "id")
		jsonparser.Set(largefixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "[31]", "id")
		jsonparser.Set(largefixture, []byte("-3"), "topics", "topics", "[0]", "id")
		jsonparser.Set(largefixture, []byte("-3"), "topics", "topics", "[29]", "id")
	}
}

func BenchmarkJinSetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Set(largefixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "0", "id")
		jin.Set(largefixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "31", "id")
		jin.Set(largefixture, []byte("-3"), "topics", "topics", "0", "id")
		jin.Set(largefixture, []byte("-3"), "topics", "topics", "29", "id")
	}
}
