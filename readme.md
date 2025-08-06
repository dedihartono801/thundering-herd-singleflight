## 📦 Contoh Output

Request #1 mulai  
Mengambil dari DB...  
Request #4 mulai  
Request #2 mulai  
Request #3 mulai  
Request #5 mulai  
Request #1 dapat data: Data Produk 123  
Request #5 dapat data: Data Produk 123  
Request #3 dapat data: Data Produk 123  
Request #2 dapat data: Data Produk 123  
Request #4 dapat data: Data Produk 123  


## ✅ Kesimpulan

- `singleflight.Group` mencegah banyak goroutine memanggil proses yang sama secara bersamaan.
- Sangat berguna untuk mencegah **Thundering Herd** saat terjadi **cache miss**.
- Implementasi ini cocok untuk use-case seperti:
  - 🔁 **Cache + Database**
  - 🌐 **API Proxy**
  - 📁 **File System Access**