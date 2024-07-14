package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"sort"
	"sync"
)

// 定义点赞数据结构
type LikeData struct {
	ID    string
	Likes int
}

// 定义全局变量
var (
	localCache   = make(map[string]int)
	localCacheMu sync.Mutex
	redisClient  *redis.Client
	ctx          = context.Background()
)

// 初始化 Redis 连接
func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// 加载数据并缓存
func loadData() {
	// 示例数据
	data := []LikeData{
		{"item1", 100},
		{"item2", 200},
		{"item3", 300},
	}

	// 更新本地缓存和 Redis
	localCacheMu.Lock()
	defer localCacheMu.Unlock()
	for _, item := range data {
		localCache[item.ID] = item.Likes
		redisClient.ZAdd(ctx, "likes", &redis.Z{
			Score:  float64(item.Likes),
			Member: item.ID,
		})
	}
}

// 获取前 N 的点赞数据
func getTopN(n int) []LikeData {
	// 尝试从本地缓存中获取
	localCacheMu.Lock()
	items := make([]LikeData, 0, len(localCache))
	for id, likes := range localCache {
		items = append(items, LikeData{ID: id, Likes: likes})
	}
	localCacheMu.Unlock()

	// 按点赞数量排序并返回前 N 个
	sort.Slice(items, func(i, j int) bool {
		return items[i].Likes > items[j].Likes
	})

	if len(items) > n {
		items = items[:n]
	}
	return items
}

// 更新缓存
func updateLikes(id string, likes int) {
	// 更新本地缓存
	localCacheMu.Lock()
	localCache[id] = likes
	localCacheMu.Unlock()

	// 更新 Redis
	redisClient.ZAdd(ctx, "likes", &redis.Z{
		Score:  float64(likes),
		Member: id,
	})
}

func main() {
	// 初始化 Redis 连接
	initRedis()

	// 加载数据并缓存
	loadData()

	// 获取前 3 的点赞数据
	top3 := getTopN(3)
	fmt.Println("Top 3 likes:", top3)

	// 更新点赞数据
	updateLikes("item1", 400)

	// 再次获取前 3 的点赞数据
	top3 = getTopN(3)
	fmt.Println("Top 3 likes after update:", top3)
}
