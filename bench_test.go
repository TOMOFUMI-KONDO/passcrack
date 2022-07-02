package passcrack

import "testing"

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(1, 6)
	}
}

func BenchmarkRunConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunConcurrent(1, 6)
	}
}
