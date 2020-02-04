package jint

// import(
// 	"github.com/buger/jsonparser"
// 	"testing"
// 	"jint"
// )

// func nop(_ ...interface{}) {}

// func BenchmarkJsonparserGetSmall(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		jsonparser.Get(SmallFixture, "uuid")
// 		jsonparser.Get(SmallFixture, "tz")
// 		jsonparser.Get(SmallFixture, "ua")
// 		jsonparser.Get(SmallFixture, "st")
// 	}
// }

// func BenchmarkJintGetSmall(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		jint.Get(SmallFixture, "uuid")
// 		jint.Get(SmallFixture, "tz")
// 		jint.Get(SmallFixture, "ua")
// 		jint.Get(SmallFixture, "st")
// 	}
// }

// func BenchmarkJsonparserGetMedium(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		jsonparser.Get(MediumFixture, "person", "name", "fullName")
// 		jsonparser.Get(MediumFixture, "person", "github", "followers")
// 		jsonparser.Get(MediumFixture, "company")

// 		jsonparser.ArrayEach(MediumFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 			jsonparser.Get(value, "url")
// 			nop()
// 		}, "person", "gravatar", "avatars")
// 	}
// }

// func BenchmarkJintGetMedium(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		jint.Get(MediumFixture, "person", "name", "fullName")
// 		jint.Get(MediumFixture, "person", "github", "followers")
// 		jint.Get(MediumFixture, "company")

// 		jint.IterateArray(MediumFixture, func(value []byte) bool {
// 			jint.Get(value, "url")
// 			nop()
// 			return true
// 		}, "person", "gravatar", "avatars")
// 	}
// }

// func BenchmarkJsonparserGetLarge(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		jsonparser.ArrayEach(LargeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 			jsonparser.Get(value, "username")
// 			nop()
// 		}, "users")

// 		jsonparser.ArrayEach(LargeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 			jsonparser.Get(value, "id")
// 			jsonparser.Get(value, "slug")
// 			nop()
// 		}, "topics", "topics")
// 	}
// }

// func BenchmarkJintGetLarge(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		jint.IterateArray(LargeFixture, func(value []byte) bool {
// 			jint.Get(value, "username")
// 			return true
// 		}, "users")

// 		jint.IterateArray(LargeFixture, func(value []byte) bool {
// 			jint.Get(value, "id")
// 			jint.Get(value, "slug")
// 			return true
// 		}, "topics", "topics")
// 	}
// }

// func BenchmarkIterateArrayGetJsonparser(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		jsonparser.ArrayEach(FakeArray, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 			nop(value)
// 		})
// 	}
// }

// func BenchmarkIterateArrayGetJint(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		jint.IterateArray(FakeArray, func(value []byte) bool {
// 			nop(value)
// 			return true
// 		})
// 	}
// }

// func BenchmarkIterateObjectGetJsonparser(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		jsonparser.ObjectEach(FakeObject, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
// 			nop(key, value)
// 			return nil
// 		})
// 	}
// }

// func BenchmarkIterateObjectGetJint(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		jint.IterateKeyValue(FakeObject, func(key []byte, value []byte) bool {
// 			nop(key, value)
// 			return true
// 		})
// 	}
// }
