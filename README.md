#### 项目打包
```
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o go-blog  main.go
```

### 项目部署
```
1.替换/mnt/docker/go-blog-end下go-blog文件
2.chmod -R 777 go-blog
3.docker-compose up --build
```
