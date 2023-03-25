package host

import (
	"context"
	"sync"
)

type EventHandler[T any] func(T)

// Event registration & callback
type Event[T any] struct {
	mutex         sync.RWMutex
	lastHandlerID int
	handlers      map[int]EventHandler[T]
}

// Register an event handler.
// To unregister, call the returned cancel function.
func (e *Event[T]) Register(handler EventHandler[T]) context.CancelFunc {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.lastHandlerID++
	id := e.lastHandlerID
	if e.handlers == nil {
		e.handlers = make(map[int]EventHandler[T])
	}
	e.handlers[id] = handler

	return func() {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		delete(e.handlers, id)
	}
}

// Invoke all handlers with given value
func (e *Event[T]) Invoke(value T) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	for _, h := range e.handlers {
		go h(value)
	}
}
