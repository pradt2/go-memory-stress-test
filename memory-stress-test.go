package main

import (
	"fmt"
	"runtime/debug"
    "time"
)

type memsize uint64

const byt memsize = 1
const kibibyte memsize = 1024 * byt
const mebibyte memsize = 1024 * kibibyte
const gibibyte memsize = 1024 * mebibyte

const kilobyte memsize = 1000 * byt
const megabyte memsize = 1000 * kilobyte
const gigabyte memsize = 1000 * megabyte

func alloc(size memsize) []byte {
	mem := make([]byte, size, size)
	var i uint64
	for i = 0; memsize(i) < size; i++ {
		mem[i] = 255
	}
	return mem
}

func allocAppend(size memsize, appender []uint8) []uint8 {
	mem := alloc(size)
	return append(appender, mem...)
}

func display(size memsize) {
	gbs := uint64(size / gigabyte)
	size -= memsize(gbs) * gigabyte
	mbs := uint64(size / megabyte)
	size -= memsize(mbs) * megabyte
	kbs := uint64(size / kilobyte)
	size -= memsize(kbs) * kilobyte
	bs := size

	fmt.Printf("%4d GBs | %4d MBs | %4d KBs | %4d Bs\n", gbs, mbs, kbs, bs)
}

const stepsize = 100 * megabyte

func main() {
	debug.SetGCPercent(-1)
	var size memsize = 0
	for {
		newSize := size + stepsize
		alloc(stepsize)
		size = newSize
		display(size)
        time.Sleep(1 * time.Second)
	}
}
