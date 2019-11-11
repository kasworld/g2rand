// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// extend stardard rand package
// can use instead of standard rand package
package g2rand

import (
	"math/rand"
	"sync"
	"time"
)

type G2Rand struct {
	mutex sync.Mutex
	mr    *rand.Rand
	src   rand.Source64
}

func New() *G2Rand {
	return NewWithSeed(time.Now().UnixNano())
}

func NewWithSeed(seed int64) *G2Rand {
	rndSource := rand.NewSource(seed)
	return &G2Rand{
		mr:  rand.New(rndSource),
		src: rndSource.(rand.Source64),
	}
}

// include s, not include e
func (rnd *G2Rand) IntRange(s, e int) int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.IntRangeNolock(s, e)
}

func (rnd *G2Rand) Read(dst []byte) (int, error) {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Read(dst)
}
func (rnd *G2Rand) ReadNolock(dst []byte) (int, error) {
	return rnd.mr.Read(dst)
}

func (rnd *G2Rand) IntRangeNolock(s, e int) int {
	r := e - s
	// if r < 1<<31 {
	// 	return int(rnd.Int31n(int32(r))) + s
	// }
	return rnd.mr.Intn(r) + s
}

func (rnd *G2Rand) NormIntRange(desiredMean, desiredStdDev int) int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return int(rnd.mr.NormFloat64()*float64(desiredStdDev)) + desiredMean
}

func (rnd *G2Rand) NormFloat64Range(desiredMean, desiredStdDev float64) float64 {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.NormFloat64()*desiredStdDev + desiredMean
}

func (rnd *G2Rand) Float64_ori() float64 {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Float64()
}

func (rnd *G2Rand) Perm(n int) []int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Perm(n)
}

func (rnd *G2Rand) Intn(n int) int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.IntnNolock(n)
}

func (rnd *G2Rand) Int63() int64 {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Int63()
}

func (rnd *G2Rand) IntnNolock(n int) int {
	// if n < 1<<31 {
	// 	return int(rnd.Int31n(int32(n)))
	// }
	return rnd.mr.Intn(n)
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *G2Rand) Float64() float64 {
	return float64(r.mr.Int63n(1<<53)) / (1 << 53)
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *G2Rand) Uint32() uint32 { return uint32(r.src.Int63() >> 31) }

// int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// n must be > 0, but int31n does not check this; the caller must ensure it.
// int31n exists because Int31n is inefficient, but Go 1 compatibility
// requires that the stream of values produced by math/rand remain unchanged.
// int31n can thus only be used internally, by newly introduced APIs.
//
// For implementation details, see:
// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction
// https://lemire.me/blog/2016/06/30/fast-random-shuffling
func (r *G2Rand) Int31n(n int32) int32 {
	v := r.Uint32()
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = r.Uint32()
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
		}
	}
	return int32(prod >> 32)
}

func (r *G2Rand) Shuffle(n int, swap func(i, j int)) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.mr.Shuffle(n, swap)
}
