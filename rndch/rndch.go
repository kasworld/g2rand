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

package rndch

import (
	"math/rand"
	"time"
)

type ChSource struct {
	ch <-chan int64
}

func NewSource(seed int64, buffsize int) *ChSource {
	rnd := rand.New(rand.NewSource(seed))
	genCh := make(chan int64, buffsize)
	go func() {
		for {
			genCh <- rnd.Int63()
		}
	}()
	return &ChSource{
		ch: genCh,
	}
}

func (cs *ChSource) Int63() int64 {
	return <-cs.ch
}

func (cs *ChSource) Seed(seed int64) {
}

func New(buffsize int) *rand.Rand {
	seed := time.Now().UnixNano()
	cs := NewSource(seed, buffsize)
	return rand.New(cs)
}

var RndCh = New(1000)
