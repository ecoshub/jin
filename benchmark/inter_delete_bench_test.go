package jin

import(
	"github.com/buger/jsonparser"
	"testing"
	"jin"
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
		jin.Delete(SmallFixture, "uuid")
		jin.Delete(SmallFixture, "tz")
		jin.Delete(SmallFixture, "ua")
		jin.Delete(SmallFixture, "st")
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
		jin.Delete(MediumFixture, "person", "name", "fullName")
		jin.Delete(MediumFixture, "person", "github", "followers")
		jin.Delete(MediumFixture, "company")
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
		jin.Delete(LargeFixture, "users", "0", "id")
		jin.Delete(LargeFixture, "users", "31", "id")
		jin.Delete(LargeFixture, "topics", "topics", "0", "id")
		jin.Delete(LargeFixture, "topics", "topics", "29", "id")
	}
}
