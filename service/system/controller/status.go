package controller

import (
	"os"
	"runtime"
	"strconv"
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
	"github.com/feedlabs/elasticfeed/service/stream/model"
)

type StatusController struct {
	feedify.Controller
}

func (this *StatusController) Get() {
	var memstats runtime.MemStats;

	runtime.ReadMemStats(&memstats)
	hostname, _ := os.Hostname()

	this.Data["json"] = map[string]string{
		"hostname": hostname,
		"pid": strconv.Itoa(os.Getpid()),
		"cpus": strconv.Itoa(runtime.NumCPU()),
		"goroutines": strconv.Itoa(runtime.NumGoroutine()),
		"mem_alloc": strconv.Itoa(int(memstats.Alloc)),
		"mem_alloc_heap": strconv.Itoa(int(memstats.HeapAlloc)),
		"mem_alloc_total": strconv.Itoa(int(memstats.TotalAlloc)),
		"mem_sys": strconv.Itoa(int(memstats.Sys)),
		"subscribers": strconv.Itoa(room.Subscribers.Len()),
		"waitinglist": strconv.Itoa(room.WaitingList.Len()),
		"archived": strconv.Itoa(model.Archive.Len()),
	}

	this.Controller.ServeJson()
}
