1.修改demo/database/conn.go中mongo的连接地址（Session,err = mgo.Dial("mongodb://admin:admin@115.159.77.155:11000?maxPoolSize=100")），默认是我云服务器的一个测试数据库地址，有auth认证
2.将docker构建文件夹放到服务器上，进入docker构建文件目录下 执行"docker build -t goweb ."完成自动构建docker镜像（该镜像有go环境，自行编译demo，没有使用预先编译好demo直接容器运行二进制文件的方法） 。
3.构建完成 执行“docker run -d --name=goweb -p 8080:8000 goweb”就能启动容器将服务器8080端口映射到容器内部的8000端口，8080端口可以自定义
4.web服务的端口是8000，可以在demo/router/router.go中自行修改，然后在dockerfile中开放修改后的端口。
5.测试 在服务器终端使用 curl http://localhost:8080/v1/hello/aa 进行简单测试
6.如有疑问qq:1217286494 联系我。