package main

import (
	"time"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func main() {
	coll, err := ebpf.LoadCollection("bpf/hello")
	if err != nil {
		panic(err)
	}
	defer coll.Close()

	probe, err := link.Kprobe("sys_clone", coll.Programs["kprobe_sys_clone"], nil)
	if err != nil {
		panic(err)
	}
	defer probe.Close()

	time.Sleep(time.Minute)
} 
