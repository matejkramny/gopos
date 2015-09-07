package gopos

import (
	"fmt"
	// "encoding/json"
	"github.com/garyburd/redigo/redis"
)

var Redis *redis.Conn

type printData struct {
	Print string `json:"print"`
}

func ConnectRedis() error {
	redis_c, err := redis.Dial("tcp", "127.0.0.1:6379")
	Redis = &redis_c

	return err
}

func RedisListen(printer *ESCPOS) {
	sub := redis.PubSubConn{Conn: *Redis}
	sub.Subscribe("oc_print.receipt")

	for {
		switch n := sub.Receive().(type) {
		case redis.Message:
			fmt.Printf("Got message: %s\n", n.Data)

			// var data printData
			// if err := json.Unmarshal(n.Data, &data); err != nil {
			// 	panic(err)
			// }

			printer.Connection.Write(n.Data)
			// printer.PrintTemplate(data.Print)
		case error:
			fmt.Printf("Redis error: %v\n", n)
			return
		}
	}
}