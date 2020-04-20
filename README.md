# RoundRobinWithWeight
带权重的round robin 算法：最大公约数&取模运算

# 结构体
type URL struct {
	Url       string //url连接
	Weight    int64  //权重
	CallTimes int64  //调用次数统计
}

# 运行1000次结果
两个url，权重为2:4, 最终调动次数为333:667
result : 
0 127.0.0.1:8080 2 333
1 127.0.0.1:8081 4 667
