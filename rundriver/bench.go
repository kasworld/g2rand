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

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kasworld/g2rand"
	"github.com/kasworld/g2rand/rndch"
)

func main() {
	n := 10000000
	fmt.Printf("bench_g2rand %v %v\n", n, bench_g2rand(n))
	fmt.Printf("bench_grand %v %v\n", n, bench_grand(n))
	fmt.Printf("bench_rand %v %v\n", n, bench_rand(n))
	fmt.Printf("bench_ch 0 %v %v\n", n, bench_ch(0, n))
	fmt.Printf("bench_ch 100 %v %v\n", n, bench_ch(100, n))
}

func bench_grand(n int) time.Duration {
	sttime := time.Now()
	// rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sum int64
	for i := 0; i < n; i++ {
		sum += rand.Int63()
	}
	_ = sum
	return time.Now().Sub(sttime)
}

func bench_rand(n int) time.Duration {
	sttime := time.Now()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sum int64
	for i := 0; i < n; i++ {
		sum += rnd.Int63()
	}
	_ = sum
	return time.Now().Sub(sttime)
}

func bench_g2rand(n int) time.Duration {
	sttime := time.Now()
	rnd := g2rand.New()
	var sum int64
	for i := 0; i < n; i++ {
		sum += rnd.Int63()
	}
	_ = sum
	return time.Now().Sub(sttime)
}

func bench_ch(b, n int) time.Duration {
	sttime := time.Now()
	rnd := rndch.New(b)
	var sum int64
	for i := 0; i < n; i++ {
		sum += rnd.Int63()
	}
	_ = sum
	return time.Now().Sub(sttime)
}
