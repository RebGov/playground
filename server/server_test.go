package server

import (
	"context"
	"sync"
	"testing"
	"time"
)

// TestWaitUntilIsDoneOrCanceled tests the waitUntilIsDoneOrCanceled function.
func TestWaitUntilIsDoneOrCanceled(t *testing.T) {
	t.Run("TestWaitUntilIsDoneOrCanceled", func(t *testing.T) {
		// Create a cancelable context with a timeout of 5 seconds
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Create done channels
		done1 := make(chan struct{})
		done2 := make(chan struct{})

		// Use a WaitGroup to wait for the goroutine to finish
		var wg sync.WaitGroup
		wg.Add(1)

		// Run the waitUntilIsDoneOrCanceled function in a goroutine
		go func() {
			defer wg.Done()
			if err := waitUntilIsDoneOrCanceled(ctx, done1, done2); err != nil {
				t.Logf("Error: %v", err)
			}
		}()

		// Simulate closing one done channel
		close(done1)

		// Simulate closing the other done channel
		close(done2)

		// Wait for the goroutine to finish or timeout
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			// The goroutine has finished
		case <-ctx.Done():
			t.Error("Timeout waiting for the goroutine to finish")
		}
	})
}
