package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	cache   = make(map[string]string) // Simulasi cache
	db      = map[string]string{"product:123": "Data Produk 123"}
	mu      sync.RWMutex       // Untuk akses aman ke cache
	sfGroup singleflight.Group // Singleflight instance
)

// Simulasi ambil data dari "database"
func fetchFromDB(key string) (string, error) {
	fmt.Println("Mengambil dari DB...")
	time.Sleep(2 * time.Second) // Simulasi delay
	return db[key], nil
}

// Fungsi ambil data dengan singleflight
func getData(key string) (string, error) {
	// Cek dulu di cache
	mu.RLock()
	if val, found := cache[key]; found {
		mu.RUnlock()
		fmt.Println("Dapat dari cache")
		return val, nil
	}
	mu.RUnlock()

	// Pakai singleflight untuk hindari request ganda
	val, err, _ := sfGroup.Do(key, func() (interface{}, error) {
		// Ambil dari DB
		result, err := fetchFromDB(key)
		if err != nil {
			return nil, err
		}

		// Simpan ke cache
		mu.Lock()
		cache[key] = result
		mu.Unlock()

		return result, nil
	})

	if err != nil {
		return "", err
	}
	return val.(string), nil
}

func main() {
	key := "product:123"

	var wg sync.WaitGroup

	// Simulasi 5 request bersamaan
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Request #%d mulai\n", i)
			data, err := getData(key)
			if err != nil {
				fmt.Printf("Request #%d error: %v\n", i, err)
				return
			}
			fmt.Printf("Request #%d dapat data: %s\n", i, data)
		}(i)
	}

	wg.Wait()
}
