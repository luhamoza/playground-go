package pointers

type BigData struct {
	// pretend 1 gb of data here
}

func DoSomething(b *BigData) {
	// some logic for Big Data
}

func Pointer2() {
	data := &BigData{}

	for i := 0; i < 10000; i++ {
		// the data will be copied 10000 times
		// thus need to make pointer to only copy 8 bytes of memory
		DoSomething(data)
	}
}
