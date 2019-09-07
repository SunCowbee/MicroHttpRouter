package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func getuser(w http.ResponseWriter ,r *http.Request,ps httprouter.Params )  {

	user:=ps.ByName("user")

	fmt.Fprintln(w,"获取到当前用户为",user)
	mod := r.Method
	fmt.Fprintln(w , "获取到当前用户的请求为",mod)
}




func main() {

	//创建一个路由的中间件
	rou := httprouter.New()

	//开始编写路由解析
	rou.GET("/getuser/:user",getuser)
	//rou.DELETE()
	//rou.POST()
	//rou.PUT()
	//:user是1个正则的操作

	//将中间件传到handle
	http.Handle("/",rou);
	http.ListenAndServe(":8080",nil)

}
