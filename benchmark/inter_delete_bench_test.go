package jint

import(
	"github.com/buger/jsonparser"
	"testing"
	"jint"
)

func BenchmarkJsonParserDeleteSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Delete(SmallFixture, "uuid")
		jsonparser.Delete(SmallFixture, "tz")
		jsonparser.Delete(SmallFixture, "ua")
		jsonparser.Delete(SmallFixture, "st")
	}
}

func BenchmarkJintDeleteSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jint.Delete(SmallFixture, "uuid")
		jint.Delete(SmallFixture, "tz")
		jint.Delete(SmallFixture, "ua")
		jint.Delete(SmallFixture, "st")
	}
}

func BenchmarkJsonParserDeleteMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Delete(MediumFixture, "person", "name", "fullName")
		jsonparser.Delete(MediumFixture, "person", "github", "followers")
		jsonparser.Delete(MediumFixture, "company")
	}
}

func BenchmarkJintDeleteMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jint.Delete(MediumFixture, "person", "name", "fullName")
		jint.Delete(MediumFixture, "person", "github", "followers")
		jint.Delete(MediumFixture, "company")
	}
}

func BenchmarkJsonParserDeleteLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Delete(LargeFixture, "users", "[0]", "id")
		jsonparser.Delete(LargeFixture, "users", "[31]", "id")
		jsonparser.Delete(LargeFixture, "topics", "topics", "[0]", "id")
		jsonparser.Delete(LargeFixture, "topics", "topics", "[29]", "id")
	}
}

func BenchmarkJintDeleteLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jint.Delete(LargeFixture, "users", "0", "id")
		jint.Delete(LargeFixture, "users", "31", "id")
		jint.Delete(LargeFixture, "topics", "topics", "0", "id")
		jint.Delete(LargeFixture, "topics", "topics", "29", "id")
	}
}
