// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
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
}

// include s, not include e
func (rnd *G2Rand) IntRange(s, e int) int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Intn(e-s) + s
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

func (rnd *G2Rand) Float64() float64 {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Float64()
}

func (rnd *G2Rand) Intn(n int) int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Intn(n)
}

func (rnd *G2Rand) Perm(n int) []int {
	rnd.mutex.Lock()
	defer rnd.mutex.Unlock()
	return rnd.mr.Perm(n)
}

func New() *G2Rand {
	rndSource := rand.NewSource(time.Now().UnixNano())
	return &G2Rand{
		mr: rand.New(rndSource),
	}
}
