package main

import "testing"

func BenchmarkFibonacchi(b *testing.B) {
	for i := 0; i < b.N; i++ { //Go는 N값을 적절히 증가시키면서 충분한 테스트로 함수 측정한다.
		fibonacchi(20)
	}
}
