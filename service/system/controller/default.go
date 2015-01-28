package controller

import (
	"os"
	"runtime"
	"strconv"
	"github.com/feedlabs/feedify"
)

type SystemController struct {
	feedify.Controller
}

func (this *SystemController) Get() {
	var memstats runtime.MemStats;

	runtime.ReadMemStats(&memstats)
	hostname, _ := os.Hostname()

	this.Data["json"] = map[string]string{
		"pid": strconv.Itoa(os.Getpid()),
		"hostname": hostname,
		"mem_alloc": strconv.Itoa(int(memstats.Alloc)),
		"mem_alloc_heap": strconv.Itoa(int(memstats.HeapAlloc)),
		"goroutines": strconv.Itoa(runtime.NumGoroutine()),
		"cpus": strconv.Itoa(runtime.NumCPU()),
	}

	this.Controller.ServeJson()
}
