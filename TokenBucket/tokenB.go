package main

import (
	"net/http"

	"github.com/SongLiangChen/RateLimiter"
)

func main() {
	rules := RateLimiter.NewRules()
	// 规定任何用户1s内只允许访问5次/test1
	rules.AddRule("/test1", &RateLimiter.Rule{
		Duration: 1,
		Limit:    5,
	})
	// 同时规定任何用户10s内只能访问10次/test2
	rules.AddRule("/test2", &RateLimiter.Rule{
		Duration: 10,
		Limit:    5,
	})
	// 并且规定对任何uri的访问，60s内只能访问10次
	rules.AddRule("", &RateLimiter.Rule{
		Duration: 60,
		Limit:    10,
	})

	RateLimiter.InitRateLimiter(rules)

	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		uid := r.FormValue("uid")
		if !RateLimiter.TakeAccess(uid, "/test1") || !RateLimiter.TakeAccess(uid, "") {
			w.Write([]byte("请求太频繁"))
			return
		}

		// ...do your work
	})

	http.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		uid := r.FormValue("uid")
		if !RateLimiter.TakeAccess(uid, "/test2") || !RateLimiter.TakeAccess(uid, "") {
			w.Write([]byte("请求太频繁"))
			return
		}

		// ...do your work
	})

	http.ListenAndServe(":8080", nil)
}
