package rota

import (
	"testing"
)

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
