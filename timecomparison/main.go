package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Thing struct {
	Uuids []string `json:"uuids"`
	mut   sync.Mutex
}

func (t *Thing) ConGen(wg *sync.WaitGroup, c chan error) {
	defer wg.Done()
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)
	b := make([]byte, 16)
	_, err := gen.Read(b)
	if err != nil {
		c <- err
		return
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	t.mut.Lock()
	t.Uuids = append(t.Uuids, uuid)
	t.mut.Unlock()
}

func (t *Thing) NonConGen() error {
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)
	b := make([]byte, 16)
	_, err := gen.Read(b)
	if err != nil {
		return err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	t.Uuids = append(t.Uuids, uuid)
	return nil
}

func (t *Thing) GooglePkg(wg *sync.WaitGroup, c chan error) {
	defer wg.Done()

	newId, err := uuid.NewRandom()
	if err != nil {
		c <- err
		return
	}

	idString := fmt.Sprintf("%s", newId)

	t.mut.Lock()
	t.Uuids = append(t.Uuids, idString)
	t.mut.Unlock()
}

func ConGooglePkg() {

	t := new(Thing)
	l := 1000

	c := make(chan error, l)
	w := new(sync.WaitGroup)
	w.Add(l)

	for i := 0; i < l; i++ {
		go t.GooglePkg(w, c)
	}
	w.Wait()

	if len(c) != 0 {
		for i := range c {
			log.Println(i)
		}
	}
}

func Con() {

	t := new(Thing)
	l := 1000

	c := make(chan error, l)
	w := new(sync.WaitGroup)
	w.Add(l)

	for i := 0; i < l; i++ {
		go t.ConGen(w, c)
	}
	w.Wait()

	if len(c) != 0 {
		for i := range c {
			log.Println(i)
		}
	}
}

func NonCon() {

	ncpu := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpu)
	t := Thing{}
	l := 1000

	for i := 0; i < l; i++ {
		t.NonConGen()
	}
}
