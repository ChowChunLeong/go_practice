package test

import (
	"go_practice/helper"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Sync pool real time example

A programmer drinks bubble tea and needs a straw (the straw represents the buffer object in our code).
 After finishing the bubble tea, the straw is thrown away,
becoming plastic garbage (Garbage). Janitor Lao Li (the GC garbage collector) needs to follow closely to clean up.
Now, 1048576 programmers are simultaneously drinking bubble tea, each needing a new straw on the spot.

After drinking, there are immediately 1048576 plastic straw garbage on the ground.
Lao Li is estimated to be exhausted.
If, now, a recycling bin is placed in a secret corner (analogous to sync.Pool),
after the programmer finishes drinking bubble tea, the straw is thrown into the recycling bin.
The next programmer who needs a straw can reach into the bin to check if there are any straws.
If there are, they can use them. If not, they'll request a new straw.
This way, the number of new straws used greatly decreases, there is no garbage on the ground,
and Lao Li is also relieved, which is great.

Moreover, in extreme cases, if everyone drinks bubble tea quickly enough to ensure that
there is at least one used straw in the bin at all times,
then 1048576 programmers would probably have enough with just one straw...

*/

// 用来统计实例真正创建的次数
var numCalcsCreated int32

// 创建实例的函数
func createBuffer() interface{} {
	// 这里要注意下，非常重要的一点。这里必须使用原子加，不然有并发问题；
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func TestSyncPool(t *testing.T) {
	helper.Log("test.txt", "hello")
	// 创建实例
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	// 多 goroutine 并发测试
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// 申请一个 buffer 实例

			// * 1) create a few of instance only
			buffer := bufferPool.Get()
			// * 2) create 1048576 instance only
			// buffer := createBuffer()

			_ = buffer.(*[]byte)
			// 释放一个 buffer 实例
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()

	file, err := os.OpenFile("syncpool.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set the output of the logger to the file
	log.SetOutput(file)
	log.Printf("%d buffer objects were created.\n", numCalcsCreated)

}
