package service

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

var Cache *bigcache.BigCache

func InitCache() {
	config := bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   1024, //1G in RAM
		OnRemove:           onRemove,
		OnRemoveWithReason: onRemoveWithReason,
	}
	var err error
	Cache, err = bigcache.NewBigCache(config)
	if err != nil {
		log.Fatal(err)
	}

}

func onRemove(key string, entry []byte) {
	log.Println(time.Now().Local().String(), "entry removed:", key)
}

func onRemoveWithReason(key string, entry []byte, reason bigcache.RemoveReason) {
	log.Println("entry removed:", key, "----reason:", reason)
}
