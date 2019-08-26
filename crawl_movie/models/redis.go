package models

import (
	"github.com/monnand/goredis"
)

const (
	URL_QUEUE     = "url_queue"
	URL_VISIT_SET = "url_visit_set"
	MOVIE_NAME    = "movie_name"
)

var (
	client goredis.Client
)

func ConnectRedis(addr string) {
	client.Addr = addr
}

//queue
func PutinQueue(url string) {
	client.Lpush(URL_QUEUE, []byte(url))
}

//name
func Putiname(name string) {
	client.Lpush(MOVIE_NAME, []byte(name))
}

//name
func PopinMovieName() string {
	res, err := client.Rpop(MOVIE_NAME)
	if err != nil {
		panic(err)
	}
	return string(res)
}

//
func PopfromQueue() string {
	res, err := client.Rpop(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func AddToSet(url string) {
	client.Sadd(URL_VISIT_SET, []byte(url))
}
func IsVisit(url string) bool {
	bIsVisit, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil {
		return false
	} else {
		return bIsVisit
	}
}

func AddnameToSet(name string) {
	client.Sadd(MOVIE_NAME, []byte(name))
}

func NameIsVisit(name string) bool {
	bIsVisit, err := client.Sismember(MOVIE_NAME, []byte(name))
	if err != nil {
		return false
	} else {
		return bIsVisit
	}
}

func GetQueueLength() int {
	length, err := client.Llen(URL_QUEUE)
	if err != nil {
		return 0
	}
	return length
}
