# gRPC test
之前有聽過gRPC，但是不知道是什麼，所以有空時就來玩玩看了。  
然後覺得也想寫寫看go的爬蟲，就寫了一個送氣象URL過去幫我解析日期、天氣與溫度的程式摟~

## cmd

產生proto  
ex: `protoc -I . helloworld.proto --go_out=plugins=grpc:.`  

執行  
`cd server`  
`go run main.go`  

```
2019/02/11 22:35:45 Received: test
2019/02/11 22:35:45 Received: test
2019/02/11 22:35:45 Received: https://www.cwb.gov.tw/V7/forecast/taiwan/inc/city/Taichung_City.htm
```

`cd client`  
`go run main.go`  
```
2019/02/11 22:38:58 HelloService: Hello test
2019/02/11 22:38:58 HiService: Hi test
2019/02/11 22:38:58 CrawlerService
2019/02/11 22:38:58      02/12
2019/02/11 22:38:58      晴時多雲
2019/02/11 22:38:58      17~29
2019/02/11 22:38:58      02/13
2019/02/11 22:38:58      晴時多雲
2019/02/11 22:38:58      19~28
2019/02/11 22:38:58      02/14
2019/02/11 22:38:58      多雲時陰
2019/02/11 22:38:58      19~27
2019/02/11 22:38:58      02/15
2019/02/11 22:38:58      晴時多雲
2019/02/11 22:38:58      19~29
2019/02/11 22:38:58      02/16
2019/02/11 22:38:58      多雲時陰
2019/02/11 22:38:58      19~27
2019/02/11 22:38:58      02/17
2019/02/11 22:38:58      多雲時陰
2019/02/11 22:38:58      19~27
2019/02/11 22:38:58      02/18
2019/02/11 22:38:58      多雲時晴
2019/02/11 22:38:58      18~27
```

## Reference
[grpc/grpc-go](https://github.com/grpc/grpc-go)  
[grpc-go使用教程之helloworld](https://blog.csdn.net/chenxun_2010/article/details/80015626)  
[anaskhan96/soup](https://github.com/anaskhan96/soup)  