package jint

import (
	"github.com/json-iterator/go"
	"github.com/valyala/fastjson"
	"jint"
	"testing"
)

func BenchmarkJsoniteratorGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs := jsoniter.Get(SmallFixture)
		prs.Get("uuid").ToString()
		prs.Get("tz").ToString()
		prs.Get("ua").ToString()
		prs.Get("st").ToString()
	}
}

func BenchmarkFastjsonGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		prs, _ := p.ParseBytes(SmallFixture)
		prs.Get("uuid")
		prs.Get("tz")
		prs.Get("ua")
		prs.Get("st")
	}
}

func BenchmarkJintGetSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs, _ := jint.Parse(SmallFixture)
		prs.Get("uuid")
		prs.Get("tz")
		prs.Get("ua")
		prs.Get("st")
	}
}

func BenchmarkJsoniteratorGetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs := jsoniter.Get(MediumFixture)
		prs.Get("person", "name", "fullName")
		prs.Get("person", "github", "followers")
		prs.Get("company")
	}
}

func BenchmarkFastjsonGetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		prs, _ := p.ParseBytes(MediumFixture)
		prs.Get("person", "name", "fullName")
		prs.Get("person", "github", "followers")
		prs.Get("company")
	}
}

func BenchmarkJintGetMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs, _ := jint.Parse(MediumFixture)
		prs.Get("person", "name", "fullName")
		prs.Get("person", "github", "followers")
		prs.Get("company")
	}
}

func BenchmarkJsoniteratorGetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs := jsoniter.Get(LargeFixture)
		prs.Get("users", 0, "id")
		prs.Get("users", 31, "id")
		prs.Get("topics", "topics", 0, "id")
		prs.Get("topics", "topics", 29, "id")
	}
}

func BenchmarkFastjsonGetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		prs, _ := p.ParseBytes(LargeFixture)
		prs.Get("users", "0", "id")
		prs.Get("users", "31", "id")
		prs.Get("topics", "topics", "0", "id")
		prs.Get("topics", "topics", "29", "id")
	}
}

func BenchmarkJintGetLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prs, _ := jint.Parse(LargeFixture)
		prs.Get("users", "0", "id")
		prs.Get("users", "31", "id")
		prs.Get("topics", "topics", "0", "id")
		prs.Get("topics", "topics", "29", "id")
	}
}