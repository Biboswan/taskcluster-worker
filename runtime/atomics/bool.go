package atomics

import "sync/atomic"

// Bool is an atomic boolean, no need for locking which makes the
// code faster and simpler.
//
// This interface is really just to abstract away the 0 or 1 value of an int32
// modified using the sync/atomic package. Hopefully the go compiler will inline
// these methods so they'll be super fast.
type Bool struct {
	value int32
}

// NewBool returns an atomics.Bool initialized with value.
//
// Note it is perfectly safe to just declare an atomics.Bool; it defaults to
// false just like a normal boolean would do.
func NewBool(value bool) Bool {
	if value {
		return Bool{value: 1}
	}
	return Bool{value: 0}
}

// Set sets the value of the boolean to true or false
func (b *Bool) Set(value bool) {
	if value {
		atomic.StoreInt32(&b.value, 1)
	} else {
		atomic.StoreInt32(&b.value, 0)
	}
}

// Swap sets the value of the boolean to true or false and returns the old value
func (b *Bool) Swap(value bool) bool {
	if value {
		return atomic.SwapInt32(&b.value, 1) != 0
	}
	return atomic.SwapInt32(&b.value, 0) != 0
}

// Get returns the value of the boolean
func (b *Bool) Get() bool {
	return atomic.LoadInt32(&b.value) != 0
}
