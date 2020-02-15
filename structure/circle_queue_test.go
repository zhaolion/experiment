package structure

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestCirCleQueue(t *testing.T) {
	size := 3
	queue := NewCirCleQueue(size)
	assert.True(t, queue.IsEmpty())
	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Rear())

	for i := 0; i < size; i++ {
		assert.True(t, queue.Enqueue(1))
	}
	assert.False(t, queue.Enqueue(1))

	for i := 0; i < size; i++ {
		e, ok := queue.Dequeue()
		assert.NotEqual(t, -1, e)
		assert.True(t, ok)
	}
	e, ok := queue.Dequeue()
	assert.Equal(t, -1, e)
	assert.False(t, ok)

	queue = NewCirCleQueue(size)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Dequeue()
	queue.Enqueue(3)
	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 3, queue.Rear())
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(4)
	assert.Equal(t, 4, queue.Front())
	assert.Equal(t, 4, queue.Rear())
}

func TestCirCleQueue_Goconvey(t *testing.T) {
	Convey("CirCleQueue usages", t, func() {
		size := 3

		Convey("init a queue by default", func() {
			queue := NewCirCleQueue(size)
			So(queue.IsEmpty(), ShouldBeTrue)
			So(queue.Front(), ShouldEqual, -1)
			So(queue.Rear(), ShouldEqual, -1)
		})

		Convey("enqueue & dequeue many times", func() {
			queue := NewCirCleQueue(size)
			for i := 0; i < size; i++ {
				So(queue.Enqueue(1), ShouldBeTrue)
			}
			So(queue.Enqueue(1), ShouldBeFalse)

			for i := 0; i < size; i++ {
				e, ok := queue.Dequeue()
				So(e, ShouldNotEqual, -1)
				So(ok, ShouldBeTrue)
			}
			e, ok := queue.Dequeue()
			So(e, ShouldEqual, -1)
			So(ok, ShouldBeFalse)
		})

		Convey("enqueue and dequeue", func() {
			queue := NewCirCleQueue(size)
			queue.Enqueue(1)
			queue.Enqueue(2)
			queue.Dequeue()
			queue.Enqueue(3)
			So(queue.Front(), ShouldEqual, 2)
			So(queue.Rear(), ShouldEqual, 3)
			queue.Dequeue()
			queue.Dequeue()
			queue.Enqueue(4)
			So(queue.Front(), ShouldEqual, 4)
			So(queue.Rear(), ShouldEqual, 4)
			So(queue.IsEmpty(), ShouldBeFalse)
		})
	})
}
