package test_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var (
	counter = 0
	lock sync.Mutex

	atomicCounter = AtomicInt{}
)

type AtomicInt struct {
	value int
	lock sync.Mutex
}

func (i *AtomicInt) Increase(){
	i.lock.Lock()
	defer i.lock.Unlock()

	i.value++
}

func (i *AtomicInt) Decrease(){
	i.lock.Lock()
	defer i.lock.Unlock()

	i.value--
}

func (i *AtomicInt) Value() int{
	return i.value
}

func TestGetMockedContext(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:123/something", nil)
	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockedContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("x-mock"))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}

func TestMutex(t *testing.T){
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++{
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("final counter %d", counter))
	fmt.Println(fmt.Sprintf("final atomic counter %d", atomicCounter.Value()))

}

func updateCounter(wg *sync.WaitGroup){
	lock.Lock()
	defer lock.Unlock()
	counter++

	atomicCounter.Increase()
	wg.Done()
}