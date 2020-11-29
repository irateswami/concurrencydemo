package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

type Thing struct {
	Uuids []uuid.UUID `json:"uuids"`
	mut   *sync.Mutex
}

func (t *Thing) Gen(wg *sync.WaitGroup, c chan error) {
	defer wg.Done()

	newId, err := uuid.NewRandom()
	if err != nil {
		c <- err
		return
	}

	t.mut.Lock()
	t.Uuids = append(t.Uuids, newId)
	t.mut.Unlock()
}

func Demo(rw http.ResponseWriter, r *http.Request) {

	t := Thing{}
	l := 10

	c := make(chan error, l)
	w := new(sync.WaitGroup)

	for i := 0; i < l; i++ {
		go t.Gen(w, c)
	}
}

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("uuids", Demo)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Blocking channel
	<-c

	// Make a context
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	fmt.Println("shutting down")
	os.Exit(0)

}
