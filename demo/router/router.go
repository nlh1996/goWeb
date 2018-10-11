package router

import (
	"fmt"
	"net/http"

	"strings"
	"crypto/md5"
	"demo/controller"

	"github.com/gin-gonic/gin"
)

// Init .
func Init() {
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("/v1")
	{
		// v1.GET("/hello", handlers.HelloPage)
		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "hello %s", name)
		})
		// v1.GET("/welcome", func(c *gin.Context) {
		// 	firstname := c.DefaultQuery("username", "Guest")
		// 	lastname := c.Query("password")
		// 	c.String(http.StatusOK, "hello %s %s", firstname, lastname)
		// })

		v1.POST("/login", func(c *gin.Context) {
			username := c.PostForm("username")
			password := c.PostForm("password")
			h := md5.New()
			h.Write([]byte(password + username))
			password = fmt.Sprintf("%x", h.Sum(nil))
			user := &controller.User{UserName: username, PassWord: password}
			msg := user.Validator()
			c.String(http.StatusOK, "%s", msg)
		})

		v1.POST("/register", func(c *gin.Context) {
			username := c.PostForm("zhanghao")
			password := c.PostForm("password")
			realname := c.PostForm("realname")
			identityCard := c.PostForm("identity_card")
			phoneNumber := c.PostForm("phone_number")
			h := md5.New()
			h.Write([]byte(password + username))
			password = fmt.Sprintf("%x", h.Sum(nil))
			user := &controller.User{UserName: username, PassWord: password, RealName: realname, IdentityCard: identityCard, PhoneNumber: phoneNumber}
			if user.Find() == false {
				user.Save()
				c.String(http.StatusOK, "注册成功!")
			} else {
				c.String(http.StatusOK, "用户名存在!")
			}

		})

	}

	//r.LoadHTMLGlob("templates/*")
	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/index", func(c *gin.Context) {
	// 		c.HTML(http.StatusOK, "index.html", gin.H{
	// 			"title": "hello Gin.",
	// 		})
	// 	})

	// }

	r.Run(":8000")
	// 404 NotFound
	// 	r.NoRoute(func(c *gin.Context) {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"status": 404,
	// 			"error":  "404, page not exists!",
	// 		})
	// 	})
	// }
	}
	// Cors .
	func Cors() gin.HandlerFunc {
	    return func(c *gin.Context) {
	        method := c.Request.Method      //请求方法
	        origin := c.Request.Header.Get("Origin")        //请求头部
	        var headerKeys []string                             // 声明请求头keys
	        for k, _ := range c.Request.Header {
	            headerKeys = append(headerKeys, k)
	        }
	        headerStr := strings.Join(headerKeys, ", ")
	        if headerStr != "" {
	            headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
	        } else {
	            headerStr = "access-control-allow-origin, access-control-allow-headers"
	        }
	        if origin != "" {
	            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	            c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
	            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
	            //  header的类型
	            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
	            //              允许跨域设置                                                                                                      可以返回其他子段
	            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
	            c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
	            c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
	            c.Set("content-type", "application/json")       // 设置返回格式是json
	        }

	        //放行所有OPTIONS方法
	        if method == "OPTIONS" {
	            c.JSON(http.StatusOK, "Options Request!")
	        }
	        // 处理请求
	        c.Next()        //  处理请求
	    }
	}



