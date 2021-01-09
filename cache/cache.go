package cache

import (
	"fmt"
	"github.com/allegro/bigcache"
	"time"
)

var (
	BigCache *bigcache.BigCache
)

//创建 bigcache
func InitCache() {
	config := bigcache.Config{
		Shards:             2,               // 存储的条目数量，值必须是2的幂
		LifeWindow:         1 * time.Minute, // 超时后条目被处理
		CleanWindow:        2 * time.Second, //处理超时条目的时间范围
		MaxEntriesInWindow: 0,               // 在 Life Window 中的最大数量，
		MaxEntrySize:       0,               // 条目最大尺寸，以字节为单位
		HardMaxCacheSize:   0,               // 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}
	var err error
	BigCache, err = bigcache.NewBigCache(config)
	if err != nil {
		fmt.Println(err)
	}
}

// 设置缓存
func SetCache(key string, value string, cc *bigcache.BigCache) {
	err := cc.Set(key, []byte(value))
	if err != nil {
		fmt.Println(err)
	}
}

// 获取缓存
func GetCache(key string, cc *bigcache.BigCache) string {
	data, err := cc.Get(key)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
