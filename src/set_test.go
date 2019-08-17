package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 单个goroutine下的功能测试
func TestSet_Add(t *testing.T) {
	set := NewSet()
	s := fmt.Sprint(set)
	if s != "" {
		t.Error()
	}

	// 添加节点
	set.Add(2)
	set.Add(1)
	set.Add(3)
	s = fmt.Sprint(set)
	if s != "1\r\n2\r\n3\r\n" {
		t.Error()
	}

	// 添加重复节点
	set.Add(2)
	set.Add(3)
	s = fmt.Sprint(set)
	if s != "1\r\n2\r\n3\r\n" {
		t.Error()
	}
}

// 单个goroutine下的功能测试
func TestSet_Contain(t *testing.T) {
	set := NewSet()
	s := fmt.Sprint(set)
	if s != "" {
		t.Error()
	}

	set.Add(1)
	set.Add(3)
	set.Add(2)
	if !set.Contain(1) ||
		!set.Contain(2) ||
		!set.Contain(3) {
		t.Error()
	}
}

// 单个goroutine下的功能测试
func TestSet_Remove(t *testing.T) {
	set := NewSet()
	s := fmt.Sprint(set)
	if s != "" {
		t.Error()
	}

	// 添加节点
	set.Add(2)
	set.Add(1)
	set.Add(3)
	s = fmt.Sprint(set)
	if s != "1\r\n2\r\n3\r\n" {
		t.Error()
	}

	// 删除不存在的节点
	if set.Remove(4) {
		t.Error()
	}

	// 删除存在的节点
	set.Remove(1)
	s = fmt.Sprint(set)
	if s != "2\r\n3\r\n" {
		t.Error()
	}
}

func BenchmarkSet_Add(b *testing.B) {
	const n = 100000
	set := NewSet()
	b.ResetTimer()
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(n))
	}
}

func BenchmarkSet_Add2(b *testing.B) {
	set := NewSet()
	b.ResetTimer()
	const n = 100000
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			set.Add(rand.Intn(n))
		}()
	}
	wg.Wait()
}
