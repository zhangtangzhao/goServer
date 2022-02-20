package main

import (
	"context"
	"fmt"
	"goServer/server"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	rootCtx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(rootCtx)
	srv := &http.Server{Addr: ":8080"}
	g.Go(func() error {
		return server.StartServer(srv)
	})

	g.Go(func() error {
		<- ctx.Done()
		fmt.Println("http server down...")
		return srv.Shutdown(ctx)
	})

	chanel := make(chan os.Signal)
	signal.Notify(chanel,syscall.SIGINT,syscall.SIGTERM)

	g.Go(func() error {
		for  {
			select {
			case <- ctx.Done():
				return ctx.Err()
			case s:= <- chanel:
				switch s{
				case syscall.SIGINT,syscall.SIGTERM:
					cancel()
				default:
					fmt.Println("unsignal syscall...")
				}
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil && err != context.Canceled {
		fmt.Printf("errgroup err : %+v\n",err.Error())
	}

	fmt.Println("httpserver shutdown .....")
}
