// Package suid enables the generation of unique, non-sequential, and URL-friendly IDs.
//
//   4 bytes of Time (10 millisecond)
// + 3 bytes Random
// ------------------
//   7 bytes
//
package suid

import (
	"math/rand"
	"sync"
	"time"
)

const (
	encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	TimeUnit  = 1e7 // 10ms
)

type weeproof struct {
	mutex sync.Mutex
	seq   uint64
	rand  []byte
	short []byte
}

var once sync.Once
var wp *weeproof

func init() {
	once.Do(func() {
		rand.Seed(time.Now().UTC().UnixNano())
		wp = &weeproof{
			rand:  make([]byte, 3),
			short: make([]byte, 7),
		}
	})
}

// NewUID generates and returns a new unique ID
func NewUID() []byte {
	wp.mutex.Lock()
	defer wp.mutex.Unlock()
	now := time.Now().UnixNano() / TimeUnit
	rand.Read(wp.rand)
	val := uint(wp.rand[0])<<16 | uint(wp.rand[1])<<8 | uint(wp.rand[2])
	wp.short[1] = encodeStd[now>>18&0x3D]
	wp.short[3] = encodeStd[now>>12&0x3D]
	wp.short[5] = encodeStd[now>>6&0x3D]
	wp.short[6] = encodeStd[now&0x3D]
	wp.short[0] = encodeStd[val>>12&0x3D]
	wp.short[2] = encodeStd[val>>6&0x3D]
	wp.short[4] = encodeStd[val&0x3D]
	return wp.short
}
