package test

import (
	"sync/atomic"
	"testing"
)

/*

"Atomic operation" is a term in computer science used to describe an operation or a group of operations
that cannot be divided. This means that during the execution of an atomic operation,
the operation either fully executes or does not execute at all,
 and there are no intermediate states or partial executions.

Atomic operations are commonly used in multi-threaded programming and concurrency control to ensure data consistency and reliability. In a multi-threaded environment,
 multiple threads may simultaneously access and modify shared data.
 Without proper synchronization and control of these operations,
  data races and unpredictable behavior can occur.
Atomic operations provide an effective way to handle such concurrent access and ensure data correctness.
*/

// scenario 1(normal test)
// 1) access go_practice/test
// 2) go test -run TestDataRaceCondition

// scenario 2(race condition test)
// 1) access go_practice/test
// 2)  go test -run TestDataRaceCondition -race
func TestDataRaceCondition(t *testing.T) {
	var counter int32
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter += int32(i)
		}(i)
	}
}

// Atomic scenario 1(normal test)
// 1) access go_practice/test
// 2) go test -run TestAtomicDataRaceCondition

// Atomic scenario 2(race condition test)
// 1) access go_practice/test
// 2)  go test -run TestAtomicDataRaceCondition -race
func TestAtomicDataRaceCondition(t *testing.T) {
	var counter int32
	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&counter, int32(i))
		}(i)
	}
}
