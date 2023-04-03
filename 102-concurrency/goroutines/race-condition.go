// ** RACE-CONDITION
// ** Farklı go-routinelerinin aynı memory'ye erismeye calismasidir.
// ** burada bir goroutine o referansı okurken ; diger goroutine çoktan okumuş,çoktan editlemiş olabilir.
//
//	**  ama bir önceki go routine eski degerle islemi yapıp yeniden uzerine isleme yaparken diger go-routinein verisini ezmiş olabilir : Buna RACE-CONDITION deniyor.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// func main() {
// 	//raceExample()
// 	//reaceExampleFixed()
// }

func raceExample() {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(2) // ** 2 tane go-routineimiz var, bunları beklemesi icin waitGroup olusturduk.

	// ** bir shared value tanımlayalim. Bu 2 go-routine sonucu ben val = 200000000 olmasini bekliyorum ama cevap 101686682 çıktı.
	// ** Cunku 1. routine val degerini okurken, 2. go routine onu increment ediyor. daha sonrasında 1. go routine de eski degeri increment ediyor.
	val := 0

	// ** 1. GoRoutine
	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}
		waitGroup.Done()
	}()

	// ** 2. GoRoutine
	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
	fmt.Println(val)
}

// ** Bu sorunu cozebilmek icin 2 ana cözüm var.
// ** 1. Yol : MUTEX
// ** Mutex, bir memory adresine gercekten es zamanlı olarak 1 go-routineinin erisicegini garanti altına alır. (LOCK'lama ile)
func reaceExampleFixed() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	lock := sync.Mutex{}

	val := 0
	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
	fmt.Println(val)
}

// ** 2. Yol : Atomic
// ** sync paketinin altında atomic ile
func raceExampleFixedWithAtomic() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}
		waitGroup.Done()
	}()
}
