package erasure

import (
	"bytes"
	"math/rand"
	"runtime"
	"sort"
	"testing"
)

func corrupt(source, errList []byte, shardLength int) []byte {
	corrupted := make([]byte, len(source))
	copy(corrupted, source)
	for _, err := range errList {
		for i := 0; i < shardLength; i++ {
			corrupted[int(err)*shardLength+i] = 0x00
		}
	}
	return corrupted
}

func randomErrorList(m, numberOfErrs int) []byte {
	set := make(map[int]bool, m)
	errListInts := make([]int, numberOfErrs)
	for i := 0; i < numberOfErrs; i++ {
		err := rand.Intn(m)
		for set[err] {
			err = rand.Intn(m)
		}
		set[err] = true
		errListInts[i] = err
	}

	sort.Ints(errListInts)

	errList := make([]byte, numberOfErrs)
	for i, err := range errListInts {
		errList[i] = byte(err)
	}

	return errList
}

func TestBasicErasure_12_8(t *testing.T) {
	m := 12
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 2, 3, 4}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestBasicErasure_16_8(t *testing.T) {
	m := 16
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 8 errors")
	}
}

func TestBasicErasure_20_8(t *testing.T) {
	m := 20
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 16, 17}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestBasicErasure_9_5(t *testing.T) {
	m := 9
	k := 5
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 2, 3, 4}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestRandomErasure_12_8(t *testing.T) {
	m := 12
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestRandomErasure_16_8(t *testing.T) {
	m := 16
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 8 errors")
	}
}

func TestRandomErasure_20_8(t *testing.T) {
	m := 20
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestRandomErasure_9_5(t *testing.T) {
	m := 9
	k := 5
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, false)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestBasicCacheErasure_12_8(t *testing.T) {
	m := 12
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 2, 3, 4}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestBasicCacheErasure_16_8(t *testing.T) {
	m := 16
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 8 errors")
	}
}

func TestBasicCacheErasure_20_8(t *testing.T) {
	m := 20
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 16, 17}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestBasicCacheErasure_9_5(t *testing.T) {
	m := 9
	k := 5
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := []byte{0, 2, 3, 4}

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestRandomCacheErasure_12_8(t *testing.T) {
	m := 12
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestRandomCacheErasure_16_8(t *testing.T) {
	m := 16
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 8 errors")
	}
}

func TestRandomCacheErasure_20_8(t *testing.T) {
	m := 20
	k := 8
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func TestRandomCacheErasure_9_5(t *testing.T) {
	m := 9
	k := 5
	shardLength := 16
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	errList := randomErrorList(m, rand.Intn(m-k))

	corrupted := corrupt(append(source, encoded...), errList, shardLength)

	recovered := code.Decode(corrupted, errList, true)

	if !bytes.Equal(source, recovered) {
		t.Error("Source was not sucessfully recovered with 4 errors")
	}
}

func BenchmarkBasicEncode_14_10(b *testing.B) {
	m := 24
	k := 14
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			code.Encode(source)
		}
	})
}

func BenchmarkBasicDecode_14_10(b *testing.B) {
	m := 24
	k := 14
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		errList := []byte{0, 2, 3, 4}

		corrupted := corrupt(append(source, encoded...), errList, shardLength)

		for pb.Next() {
			recovered := code.Decode(corrupted, errList, false)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}

func BenchmarkBasicCacheEncode_14_10(b *testing.B) {
	m := 24
	k := 14
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			code.Encode(source)
		}
	})
}

func BenchmarkBasicCacheDecode_14_10(b *testing.B) {
	m := 24
	k := 14
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		errList := []byte{0, 2, 3, 4}

		corrupted := corrupt(append(source, encoded...), errList, shardLength)

		for pb.Next() {
			recovered := code.Decode(corrupted, errList, true)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}

func BenchmarkBasicEncode_28_4(b *testing.B) {
	m := 32
	k := 28
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			code.Encode(source)
		}
	})
}

func BenchmarkBasicDecode_28_4(b *testing.B) {
	m := 32
	k := 28
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		errList := []byte{0, 2, 3, 4}

		corrupted := corrupt(append(source, encoded...), errList, shardLength)

		for pb.Next() {
			recovered := code.Decode(corrupted, errList, false)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}

func BenchmarkBasicCacheEncode_28_4(b *testing.B) {
	m := 32
	k := 28
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			code.Encode(source)
		}
	})
}

func BenchmarkBasicCacheDecode_28_4(b *testing.B) {
	m := 32
	k := 28
	shardLength := 16776168
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		errList := []byte{0, 2, 3, 4}

		corrupted := corrupt(append(source, encoded...), errList, shardLength)

		for pb.Next() {
			recovered := code.Decode(corrupted, errList, true)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}

func BenchmarkBasicEncode_12_8(b *testing.B) {
	m := 12
	k := 8
	shardLength := 8192
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			code.Encode(source)
		}
	})
}

func BenchmarkBasicDecode_12_8(b *testing.B) {
	m := 12
	k := 8
	shardLength := 8192
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		errList := []byte{0, 2, 3, 4}

		corrupted := corrupt(append(source, encoded...), errList, shardLength)

		for pb.Next() {
			recovered := code.Decode(corrupted, errList, false)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}

func BenchmarkBasicCacheEncode_12_8(b *testing.B) {
	m := 12
	k := 8
	shardLength := 8192
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			code.Encode(source)
		}
	})
}

func BenchmarkBasicCacheDecode_12_8(b *testing.B) {
	m := 12
	k := 8
	shardLength := 8192
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		errList := []byte{0, 2, 3, 4}

		corrupted := corrupt(append(source, encoded...), errList, shardLength)

		for pb.Next() {
			recovered := code.Decode(corrupted, errList, true)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}

func BenchmarkRandomDecode_12_8(b *testing.B) {
	m := 12
	k := 8
	shardLength := 8192
	size := k * shardLength

	code := NewCode(m, k, size)

	source := make([]byte, size)
	for i := range source {
		source[i] = byte(rand.Int63() & 0xff) //0x62
	}

	encoded := code.Encode(source)

	b.SetBytes(int64(size))
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			errList := randomErrorList(m, rand.Intn(m-k)+1)

			corrupted := corrupt(append(source, encoded...), errList, shardLength)

			recovered := code.Decode(corrupted, errList, rand.Float32() > 0.5)

			if !bytes.Equal(source, recovered) {
				b.Error("Source was not sucessfully recovered with 4 errors")
			}
		}
	})
}
