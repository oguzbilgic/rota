package rota

import (
	"testing"
)

func captureFnTester(cf CaptureFn, path string, desiredMatch bool, desiredRest string, t *testing.T) {
	if match, rest := cf(path); match != desiredMatch {
		t.Errorf("match fail %s", path)
	} else if rest != desiredRest {
		t.Errorf("rest fail  '%s' != '%s'", desiredRest, rest)
	}
}

func TestConstCaptureFn(t *testing.T) {
	cf := ConstCaptureFn("/articles")

	// Shorter
	captureFnTester(cf, "", false, "", t)
	captureFnTester(cf, "/art", false, "/art", t)
	captureFnTester(cf, "/not", false, "/not", t)

	// Same lenght
	captureFnTester(cf, "/articles", true, "", t)
	captureFnTester(cf, "/notMatch", false, "/notMatch", t)

	// Longer
	captureFnTester(cf, "/articles/comment", true, "/comment", t)
	captureFnTester(cf, "/notMatch/comment", false, "/notMatch/comment", t)
}

func TestIntVarCaptureFn(t *testing.T) {
	cf := IntVarCaptureFn()

	captureFnTester(cf, "", false, "", t)
	captureFnTester(cf, "12", true, "", t)
	captureFnTester(cf, "deneme", false, "deneme", t)
	captureFnTester(cf, "12/1234", true, "/1234", t)
	captureFnTester(cf, "12/comment", true, "/comment", t)
	captureFnTester(cf, "12comment", true, "comment", t)
	captureFnTester(cf, "12.jpeg", true, ".jpeg", t)
}

func TestStrVarCaptureFn(t *testing.T) {
	cf := StrVarCaptureFn()

	captureFnTester(cf, "", false, "", t)
	captureFnTester(cf, "12", false, "12", t)
	captureFnTester(cf, "articles", true, "", t)
	captureFnTester(cf, "articles/comments", true, "/comments", t)
	captureFnTester(cf, "articles.jpeg", true, ".jpeg", t)
}

func BenchmarkConstCaptureFn(b *testing.B) {
	b.StopTimer()
	cf := ConstCaptureFn("/articles")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = cf("/art")
		_, _ = cf("/not")
		_, _ = cf("/articles")
		_, _ = cf("/notMatch")
		_, _ = cf("/articles/comment")
		_, _ = cf("/notMatch/comment")
	}
}

func BenchmarkIntVarCaptureFn(b *testing.B) {
	b.StopTimer()
	cf := IntVarCaptureFn()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = cf("lala")
		_, _ = cf("9876567898576878768/lala")
		_, _ = cf("99999999999999999999/lala")
	}
}
