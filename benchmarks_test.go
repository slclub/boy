package boy

import (
	"github.com/slclub/gcore"
	"github.com/slclub/gnet"
	"github.com/slclub/link"
	"net/http"
	"testing"
)

func BenchmarkServe(B *testing.B) {
	Install()
	Initialize()

	R.GET("/hello/world", func(ctx gnet.Contexter) {
	})

	run_request(B, App, http.MethodGet, "/hello/world")
}

func run_request(B *testing.B, en *gcore.Engine, method, path string) {
	// create fake request
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	link.DEBUG_PRINT("GCORE: we start benchmarks", "\n")
	w := new_mock_writer()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		en.Core().ServeHTTP(w, req)
	}
}

type header_writer struct {
	header http.Header
}

func (m *header_writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *header_writer) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *header_writer) Header() http.Header {
	return m.header
}

func (m *header_writer) WriteHeader(int) {}

func new_mock_writer() *header_writer {
	return &header_writer{
		http.Header{},
	}
}
