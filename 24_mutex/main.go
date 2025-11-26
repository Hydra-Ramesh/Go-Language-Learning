package main

import (
	"fmt"
	"sync"
)

// Post represents a blog post with a view count.
type Post struct {
	views int
	mu    sync.Mutex
}

// inc increments the view count of the post safely using a mutex.
func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		// Unlock the mutex and signal that the goroutine is done.
		p.mu.Unlock()
		wg.Done()
	}()
	// Lock the mutex before incrementing the view count.
	p.mu.Lock()
	p.views++
}
func main() {
	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup
	p := Post{views: 0}
	for i := 0; i < 100; i++ {
		// Increment the WaitGroup counter and start a new goroutine.
		wg.Add(1)
		// Call the inc method in a goroutine.
		go p.inc(&wg)
	}
	// Wait for all goroutines to finish.
	wg.Wait()
	// Print the final view count.
	fmt.Println(p.views)
}
