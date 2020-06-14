package jin

import (
	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/ecoshub/jin"
	"testing"
)

func nop(_ ...interface{}) {}

func BenchmarkJsoniteratorGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Get(smallfixture, "uuid")
		jsoniter.Get(smallfixture, "tz")
		jsoniter.Get(smallfixture, "ua")
		jsoniter.Get(smallfixture, "st")
	}
}

func BenchmarkGjsonGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gjson.GetBytes(smallfixture, "uuid")
		gjson.GetBytes(smallfixture, "tz")
		gjson.GetBytes(smallfixture, "ua")
		gjson.GetBytes(smallfixture, "st")
	}
}

func BenchmarkJsonparserGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Get(smallfixture, "uuid")
		jsonparser.Get(smallfixture, "tz")
		jsonparser.Get(smallfixture, "ua")
		jsonparser.Get(smallfixture, "st")
	}
}

func BenchmarkJinGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Get(smallfixture, "uuid")
		jin.Get(smallfixture, "tz")
		jin.Get(smallfixture, "ua")
		jin.Get(smallfixture, "st")
	}
}

func BenchmarkGjsonGetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gjson.GetBytes(mediumfixture, "person.name.fullName")
		gjson.GetBytes(mediumfixture, "person.github.followers")
		gjson.GetBytes(mediumfixture, "company")

		result := gjson.GetBytes(mediumfixture, "person.gravatar.avatars")
		result.ForEach(func(key, value gjson.Result) bool {
			gjson.Get(value.String(), "url")
			nop()
			return true
		})
	}
}

func BenchmarkJsonparserGetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Get(mediumfixture, "person", "name", "fullName")
		jsonparser.Get(mediumfixture, "person", "github", "followers")
		jsonparser.Get(mediumfixture, "company")

		jsonparser.ArrayEach(mediumfixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "url")
			nop()
		}, "person", "gravatar", "avatars")
	}
}

func BenchmarkJinGetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.Get(mediumfixture, "person", "name", "fullName")
		jin.Get(mediumfixture, "person", "github", "followers")
		jin.Get(mediumfixture, "company")

		jin.IterateArray(mediumfixture, func(value []byte) bool {
			jin.Get(value, "url")
			nop()
			return true
		}, "person", "gravatar", "avatars")
	}
}

func BenchmarkGjsonrGetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := gjson.GetBytes(largefixture, "users")
		result.ForEach(func(key, value gjson.Result) bool {
			gjson.Get(value.String(), "username")
			nop()
			return true
		})

		result = gjson.GetBytes(largefixture, "topics.topics")
		result.ForEach(func(key, value gjson.Result) bool {
			gjson.Get(value.String(), "id")
			gjson.Get(value.String(), "slug")
			nop()
			return true
		})
	}
}

func BenchmarkJsonparserGetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.ArrayEach(largefixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "username")
			nop()
		}, "users")

		jsonparser.ArrayEach(largefixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "id")
			jsonparser.Get(value, "slug")
			nop()
		}, "topics", "topics")
	}
}

func BenchmarkJinGetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jin.IterateArray(largefixture, func(value []byte) bool {
			jin.Get(value, "username")
			return true
		}, "users")

		jin.IterateArray(largefixture, func(value []byte) bool {
			jin.Get(value, "id")
			jin.Get(value, "slug")
			return true
		}, "topics", "topics")
	}
}

func BenchmarkIterateArrayGetGjson(b *testing.B) {
	b.ReportAllocs()
	newData := jin.MakeEmptyJson()
	newData, _ = jin.AddKeyValue(newData, "array", fakearray)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := gjson.GetBytes(newData, "array")
		result.ForEach(func(key, value gjson.Result) bool {
			nop(key)
			nop(value)
			return true
		})
	}
}

func BenchmarkIterateArrayGetJsonparser(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsonparser.ArrayEach(fakearray, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nop(value)
		})
	}
}

func BenchmarkIterateArrayGetJin(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jin.IterateArray(fakearray, func(value []byte) bool {
			nop(value)
			return true
		})
	}
}

func BenchmarkIterateObjectGetGjson(b *testing.B) {
	b.ReportAllocs()
	newData := jin.MakeEmptyJson()
	newData, _ = jin.AddKeyValue(newData, "object", fakeobject)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := gjson.GetBytes(newData, "object")
		result.ForEach(func(key, value gjson.Result) bool {
			nop(key)
			nop(value)
			return true
		})
	}
}

func BenchmarkIterateObjectGetJsonparser(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsonparser.ObjectEach(fakeobject, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			nop(key, value)
			return nil
		})
	}
}

func BenchmarkIterateObjectGetJin(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jin.IterateKeyValue(fakeobject, func(key []byte, value []byte) bool {
			nop(key, value)
			return true
		})
	}
}
