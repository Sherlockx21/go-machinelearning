package Cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func Simple() {
	c := cache.New(1*time.Minute, 30*time.Second)
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey,value:%s\n", v)
	}
}
