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

// 测试在单个goroutine下，调用n次Add()，n次Contain()，n次Remove()的效率
func BenchmarkSet_Add_Contain_Remove(b *testing.B) {
	const n = 100000
	set := NewSet()
	b.ResetTimer()
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(n))
	}
	for i := 0; i < n/2; i++ {
		set.Contain(rand.Intn(n))
	}
	for i := 0; i < n; i++ {
		set.Remove(rand.Intn(n))
	}
	for i := 0; i < n/2; i++ {
		set.Contain(rand.Intn(n))
	}
}

// 测试在多个goroutine下，调用n次Add()，n次Contain()，n次Remove()的效率
func BenchmarkSet_Add_Contain_Remove2(b *testing.B) {
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

	wg.Add(n/2)
	for i := 0; i < n/2; i++ {
		go func() {
			defer wg.Done()
			set.Contain(rand.Intn(n))
		}()
	}
	wg.Wait()

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			set.Remove(rand.Intn(n))
		}()
	}
	wg.Wait()

	wg.Add(n/2)
	for i := 0; i < n/2; i++ {
		go func() {
			defer wg.Done()
			set.Contain(rand.Intn(n))
		}()
	}
	wg.Wait()
}
