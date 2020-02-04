package jin

import(
	"github.com/buger/jsonparser"
	"testing"
	"jin"
)

func BenchmarkJsonParserSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Set(SmallFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "uuid")
		jsonparser.Set(SmallFixture, []byte("-3"), "tz")
		jsonparser.Set(SmallFixture, []byte(`"server_agent"`), "ua")
		jsonparser.Set(SmallFixture, []byte("3"), "st")
	}
}

func BenchmarkJintSetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Set(SmallFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "uuid")
		jin.Set(SmallFixture, []byte("-3"), "tz")
		jin.Set(SmallFixture, []byte(`"server_agent"`), "ua")
		jin.Set(SmallFixture, []byte("3"), "st")
	}
}

func BenchmarkJsonParserSetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Set(MediumFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "person", "name", "fullName")
		jsonparser.Set(MediumFixture, []byte("-3"), "person", "github", "followers")
		jsonparser.Set(MediumFixture, []byte(`"server_agent"`), "company")
	}
}

func BenchmarkJintSetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Set(MediumFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "person", "name", "fullName")
		jin.Set(MediumFixture, []byte("-3"), "person", "github", "followers")
		jin.Set(MediumFixture, []byte(`"server_agent"`), "company")
	}
}

func BenchmarkJsonParserSetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Set(LargeFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "[31]", "id")
		jsonparser.Set(LargeFixture, []byte("-3"), "topics", "topics", "[29]", "id")
	}
}

func BenchmarkJintSetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Set(LargeFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "31", "id")
		jin.Set(LargeFixture, []byte("-3"), "topics", "topics", "29", "id")
	}
}
