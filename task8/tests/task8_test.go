package mywaitgroup_test

import (
	"testing"
	"time"

	"main/mywaitgroup"
)

func TestMyWaitGroup(t *testing.T) {
	wg := mywaitgroup.InitMyWaitGroup()

	wg.AddOne()
	done := make(chan bool)

	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		t.Fatalf("Wait unblock before Done")
	case <-time.After(time.Second):
	}

	wg.Done()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatalf("Wait unblock after Done")
	}
}

func TestMyWaitGroupDoubleDone(t *testing.T) {
	wg := mywaitgroup.InitMyWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic when Done more than AddOne")
		}
	}()

	wg.AddOne()
	wg.Done()
	wg.Done()
}

func TestMyWaitGroupNoAddNoWait(t *testing.T) {
	wg := mywaitgroup.InitMyWaitGroup()

	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		wg.Wait()
	}()

	select {
	case <-done:
		t.Fatalf("Wait not block if no AddOne is called")
	case <-time.After(time.Second):
		wg.AddOne()
		wg.Done()
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait is not unblock")
	}
}
