package queue

import "testing"

const queueSize = 3

func BenchmarkQueue1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := make([]string, 0, queueSize)
		for i := 0; i < 100; i++ {
			if len(q) == queueSize {
				_ = q[0] //dequeue
				q = q[1:]
			}
			q = append(q, "conn")
		}
	}
}

func BenchmarkQueue2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := make([]string, 0, queueSize)
		for i := 0; i < 100; i++ {
			if len(q) == queueSize {
				_ = q[0] //dequeue
				copy(q, q[1:])
				q = q[:len(q)-1]
			}
			q = append(q, "conn")
		}
	}
}
