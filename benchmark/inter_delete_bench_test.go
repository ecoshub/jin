package jin

import (
	"github.com/buger/jsonparser"
	"github.com/ecoshub/jin"
	"testing"
)

func BenchmarkJsonParserDeleteSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Delete(smallfixture, "uuid")
		jsonparser.Delete(smallfixture, "tz")
		jsonparser.Delete(smallfixture, "ua")
		jsonparser.Delete(smallfixture, "st")
	}
}

func BenchmarkJinDeleteSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Delete(smallfixture, "uuid")
		jin.Delete(smallfixture, "tz")
		jin.Delete(smallfixture, "ua")
		jin.Delete(smallfixture, "st")
	}
}

func BenchmarkJsonParserDeleteMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Delete(mediumfixture, "person", "name", "fullName")
		jsonparser.Delete(mediumfixture, "person", "github", "followers")
		jsonparser.Delete(mediumfixture, "company")
	}
}

func BenchmarkJinDeleteMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Delete(mediumfixture, "person", "name", "fullName")
		jin.Delete(mediumfixture, "person", "github", "followers")
		jin.Delete(mediumfixture, "company")
	}
}

func BenchmarkJsonParserDeleteLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Delete(largefixture, "users", "[0]", "id")
		jsonparser.Delete(largefixture, "users", "[31]", "id")
		jsonparser.Delete(largefixture, "topics", "topics", "[0]", "id")
		jsonparser.Delete(largefixture, "topics", "topics", "[29]", "id")
	}
}

func BenchmarkJinDeleteLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Delete(largefixture, "users", "0", "id")
		jin.Delete(largefixture, "users", "31", "id")
		jin.Delete(largefixture, "topics", "topics", "0", "id")
		jin.Delete(largefixture, "topics", "topics", "29", "id")
	}
}
