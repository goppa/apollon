package http

import (
	"net/http"
	"log"
	"os"
	"strconv"
)

var running = false        // サーバ起動フラグ

func Start(port int, mux http.Handler) error {
	// 起動済みの場合
	if running == true {
		log.Println("already server running")
		return nil
	}
	// サーバの起動
	addr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Println("ListenAndServe error")
		return err
	}
	log.Printf("start pid %d\n", os.Getpid())
	running = true
	// シグナルハンドリング
	// そのうちやる
	// http://reiki4040.hatenablog.com/entry/2014/10/03/002239
	return nil
}

func Stop() error {
	// サーバが起動していない場合
	if running == false {
		log.Println("server not running")
		return nil
	}
	// 終了処理実行
	log.Printf("shutdown pid %d\n", os.Getpid())
	return nil
}
