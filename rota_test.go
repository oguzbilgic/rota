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

func rotaMatchTester(r *Rota, path string, desiredMatch bool, t *testing.T) {
	if r.Match(path) != desiredMatch {
		t.Errorf("match fail %s", path)
	}
}

func TestRota(t *testing.T) {
	r := New("/articles/<id:int>")

	rotaMatchTester(r, "", false, t)
	rotaMatchTester(r, "/art", false, t)
	rotaMatchTester(r, "/articles", false, t)
	rotaMatchTester(r, "/articles/", false, t)
	rotaMatchTester(r, "/notMatch/99999", false, t)
	rotaMatchTester(r, "/articles/10/something", false, t)
	rotaMatchTester(r, "/articles/10/11234", false, t)

	rotaMatchTester(r, "/articles/0", true, t)
	rotaMatchTester(r, "/articles/99999", true, t)
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

// Benchmarks

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

func BenchmarkRotaMatchSingle(b *testing.B) {
	b.StopTimer()
	r := New("/articles/<id:int>")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = r.Match("/articles/234")
	}
}

func BenchmarkRotaMatchMultiple(b *testing.B) {
	b.StopTimer()
	r := New("/articles/<id:int>")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = r.Match("")
		_ = r.Match("/art")
		_ = r.Match("/articles")
		_ = r.Match("/articles/")
		_ = r.Match("/notMatch/99999")
		_ = r.Match("/articles/10/something")
		_ = r.Match("/articles/10/11234")
		_ = r.Match("/articles/0")
		_ = r.Match("/articles/99999")
	}
}
