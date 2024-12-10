package benchmark

import (
	"fmt"
	"runtime"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Benchmark struct {
	startTime   time.Time
	endTime     time.Time
	Duration    time.Duration
	startMemory runtime.MemStats
	endMemory   runtime.MemStats
	Memory      int
	Tag         string
}

func (b *Benchmark) Start(str string) {
	b.Reset()

	b.Tag = str
	b.startTime = time.Now()
	runtime.ReadMemStats(&b.startMemory)
}

func (b *Benchmark) End() {
	b.endTime = time.Now()
	runtime.ReadMemStats(&b.endMemory)

	b.Duration = b.endTime.Sub(b.startTime)
	b.Memory = int(b.endMemory.Alloc - b.startMemory.Alloc)
}

func (b *Benchmark) Print() {
	p := message.NewPrinter(language.BrazilianPortuguese)

	fmt.Printf("\n == %s ==", b.Tag)
	fmt.Printf("\nTime: %s", b.Duration)
	fmt.Printf("\nMemory: %s bytes (%s kb)\n", p.Sprint(b.Memory), p.Sprint(b.Memory/1024))
}

func (b *Benchmark) EndPrint() {
	b.End()
	b.Print()
}

func (b *Benchmark) Reset() {
	b = &Benchmark{}
}
