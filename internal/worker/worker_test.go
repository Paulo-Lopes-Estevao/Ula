package worker_test

import (
	"fmt"
	"sync"
	"testing"
)

var waitgroup sync.WaitGroup

func lengthEmail() int {
	email := []string{
		"test01@gmail.com",
		"test2@gmail.com",
		"test3@gmail.com",
		"test4@gmail.com",
	}

	return len(email)
}

func TestCreateWorkerQuantityWithExistingEmailQuantity(t *testing.T) {
	amountOfEmail := lengthEmail()
	if amountOfEmail != 4 {
		t.Error("Expected 4, got ", amountOfEmail)
	}

	waitgroup.Add(amountOfEmail)

	sum := 0
	for i := 0; i < amountOfEmail; i++ {
		go func() {
			defer waitgroup.Done()
			// do something
			sum += 1
			fmt.Println("do something", sum)
		}()
	}

}
