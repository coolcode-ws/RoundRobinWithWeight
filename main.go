package main

import (
	"fmt"
)

type URL struct {
	Url       string //url连接
	Weight    int64  //权重
	CallTimes int64  //调用次数统计
}

var weitht int64 = 0
var index int64 = -1

var urls []*URL = []*URL{
	&URL{
		Url:       "127.0.0.1:8080",
		Weight:    2,
		CallTimes: 0,
	},
	&URL{
		Url:       "127.0.0.1:8081",
		Weight:    4,
		CallTimes: 0,
	},
}

//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//计算多个数的最大公约数
func getGCD() int64 {
	var weights []int64

	for _, url := range urls {
		weights = append(weights, url.Weight)
	}

	g := weights[0]
	for i := 1; i < len(weights)-1; i++ {
		oldGcd := g
		g = gcd(oldGcd, weights[i])
	}
	return g
}

//获取当前轮次的url
func getUrl() *URL {
	for {
		index = (index + 1) % int64(len(urls))
		if index == 0 {
			weitht = weitht - getGCD()
			if weitht <= 0 {
				weitht = getMaxWeight()
				if weitht == 0 {
					return &URL{}
				}
			}
		}

		if urls[index].Weight >= weitht {
			return urls[index]
		}
	}
}

//获取最大权重
func getMaxWeight() int64 {
	var max int64 = 0
	for _, url := range urls {
		if url.Weight >= int64(max) {
			max = url.Weight
		}
	}

	return max
}

//测试函数
func main() {
	//模拟1000次调用，统计每个url的调用次数
	for i := 0; i < 1000; i++ {
		url := getUrl()
		fmt.Println("call: ", url.Url)
		if url.CallTimes != 0 {
			url.CallTimes++
		} else {
			url.CallTimes = 1
		}
	}
	fmt.Printf("result : \n")
	for i, url := range urls {
		fmt.Println(i, url.Url, url.Weight, url.CallTimes)
	}
}
