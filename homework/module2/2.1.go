package main

import (
	"github.com/kataras/golog"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// 注册http handler
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// 将 request 中带的 header 写入 response header
		for k, v := range r.Header {
			w.Header()[k] = v
		}

		// 读取当前系统环境变量中的 VERSION 配置，并写入 response header
		w.Header()["VERSION"] = []string{os.Getenv("VERSION")}

		// 记录访问日志，日志内容为客户端IP地址，请求URI，请求协议，响应状态码，输出到标准输出
		golog.Infof("%s %s %s %d", r.RemoteAddr, r.RequestURI, r.Proto, http.StatusOK)

		// 向客户端返回状态码 200
		w.WriteHeader(http.StatusOK)

		// 在response body中写入字符串 ok
		_, err := w.Write([]byte("ok"))
		if err != nil {
			golog.Errorf("write response body failed, err: %v", err)
		}
	})

	// 启动http server，监听端口 8080
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			golog.Fatalf("start http server failed, err: %v", err)
		}
	}()

	// 监听操作系统信号，优雅关闭http server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	golog.Info("Shutdown Server ...")
}
