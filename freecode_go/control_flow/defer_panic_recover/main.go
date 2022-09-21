package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred") // will run after main ends
	// defer goes last in first out order
	fmt.Println("end")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // be careful with defer here if looping
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots[:100]) // use [:100] to print part of text

	// panic
	/*defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something terrible happened")*/
	panicker()
	fmt.Println("end")

	/* Uncomment to start simple web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err2 := http.ListenAndServe(":8085", nil)
	if err2 != nil {
		panic(err2.Error())
	}
	*/
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something terrible happened")
	fmt.Println("done panicking")
}
