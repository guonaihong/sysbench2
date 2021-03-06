package main

import (
	"github.com/guonaihong/flag"
	"github.com/guonaihong/sysbench2/cpu"
	"github.com/guonaihong/sysbench2/memory"
	"github.com/guonaihong/sysbench2/threads"
	"os"
)

func main() {
	parent := flag.NewParentCommand(os.Args[0])

	parent.SubCommand("cpu", "CPU performance test", func() {
		cpu.Main(os.Args[0], parent.Args())
	})

	parent.SubCommand("memory", "Memory functions speed test", func() {
		memory.Main(os.Args[0], parent.Args())
	})

	parent.SubCommand("threads", "Threads subsystem performance test", func() {
		threads.Main(os.Args[0], parent.Args())
	})

	parent.Parse(os.Args[1:])
}
