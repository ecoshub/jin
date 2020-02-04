package jint

import(
	"github.com/buger/jsonparser"
	"testing"
	"jint"
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
		jint.Set(SmallFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "uuid")
		jint.Set(SmallFixture, []byte("-3"), "tz")
		jint.Set(SmallFixture, []byte(`"server_agent"`), "ua")
		jint.Set(SmallFixture, []byte("3"), "st")
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
		jint.Set(MediumFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "person", "name", "fullName")
		jint.Set(MediumFixture, []byte("-3"), "person", "github", "followers")
		jint.Set(MediumFixture, []byte(`"server_agent"`), "company")
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
		jint.Set(LargeFixture, []byte(`"c90927dd-1588-4fe7-a14f-8a8950cfcbd8"`), "users", "31", "id")
		jint.Set(LargeFixture, []byte("-3"), "topics", "topics", "29", "id")
	}
}
