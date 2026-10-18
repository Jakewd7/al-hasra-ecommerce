# 🛍️ Panduan Final Project: Bikin Toko Online "TKJ's Store"
> **Buku Panduan & Lembar Kerja Siswa Kelas 10**  
> Praktik Web Programming: HTML, CSS, dan JavaScript Dasar

---

## 👋 Halo Teman-Teman!

Selamat datang di **Final Project Web Programming**! 🎉  
Di proyek ini, kamu akan membuat website toko online bernama **TKJ's Store** yang terhubung langsung ke server backend milik tim kami 

Kamu akan mempraktikkan cara menampilkan katalog produk dari server, memasukkan barang ke keranjang belanja, melakukan checkout pesanan, mendapatkan nomor BCA Virtual Account, dan melakukan simulasi pembayaran hingga lunas.

---

## 🎒 1. Alat yang Digunakan

1. **Visual Studio Code (VS Code)** – Aplikasi untuk menulis kode.
2. **Ekstensi Live Server di VS Code (opsional)** – Untuk menjalankan website di browser secara otomatis.
3. **Web Browser (Google Chrome atau Microsoft Edge)** – Untuk melihat website dan mengecek error lewat menu Inspect/Console (`F12`).
4. **Struktur File Proyek di folder `frontend/`**:
   * `index.html` – Kerangka struktur halaman web.
   * `style.css` – Desain tampilan, warna, dan tata letak.
   * `app.js` – Logika interaktif, keranjang belanja, dan pemanggilan API backend.

---

## 🌐 2. Informasi Server Backend & Dokumentasi Swagger

### Apa itu Backend?
* Website yang kamu buat adalah **tampilan depan (Frontend)** tempat pengunjung melihat dan memilih barang.
* **Server Backend** bertugas sebagai penyedia data barang dari database dan memproses pesanan yang kamu kirimkan.

### Link Server Backend & Swagger:
* **Alamat Server API**:  
  `https://serverstaging-vps-5e6db201b3.asiahub.id/api`
* **Dokumentasi Swagger (Daftar API)**:  
  👉 [Buka Dokumentasi Swagger](https://serverstaging-vps-5e6db201b3.asiahub.id/swagger/index.html)

### 4 Endpoint API yang Digunakan:
1. **`GET /products`** – Mengambil seluruh katalog produk dari server.
2. **`POST /orders/checkout`** – Mengirim pesanan keranjang ke server untuk membuat tagihan dan nomor Virtual Account.
3. **`POST /sandbox/pay`** – Menyimulasikan pembayaran lunas Virtual Account.
4. **`GET /orders/{id}`** – Mengecek detail dan status pesanan.

---

## 📋 3. Langkah Demi Langkah Mengerjakan Proyek

Ikuti langkah-langkah di bawah ini secara berurutan:

### 🧩 Langkah 1: Siapkan Kerangka Halaman (`index.html`)
- [ ] Buat bagian atas (Navbar) yang berisi logo toko **TKJ's Store**, menu navigasi, dan tombol keranjang (`🛒 Cart`).
- [ ] Buat tombol menu hamburger untuk tampilan di layar HP.
- [ ] Buat bagian pembuka (Hero Banner) dengan judul toko dan tombol navigasi ke katalog.
- [ ] Siapkan kontainer untuk menampilkan produk: `<div id="product-grid"></div>`.
- [ ] Buat kontrol navigasi halaman produk (Tombol *Prev*, teks *Page Info*, dan tombol *Next*).
- [ ] Buat struktur pop-up (Modal) untuk Keranjang Belanja dan Formulir Pemesanan.
- [ ] Buat struktur pop-up (Modal) untuk Invoice Tagihan Pembayaran.

### 🎨 Langkah 2: Percantik Tampilan dengan CSS (`style.css`)
- [ ] Atur font sistem bawaan (`font-family: Arial, sans-serif`) agar ringan dan rapi.
- [ ] Gunakan **CSS Grid** pada `#product-grid` agar kartu produk tersusun rapi.
- [ ] Buat class `.modal-backdrop` dengan layar belakang transparan gelap untuk efek pop-up melayang di tengah layar.
- [ ] Buat class `.hidden { display: none !important; }` untuk menyembunyikan modal saat tidak aktif.
- [ ] Tambahkan aturan media query `@media (max-width: 768px)` agar menu navigasi HP berubah menjadi menu hamburger.

### 🚀 Langkah 3: Ambil Data Produk dari Server (`app.js`)
- [ ] Buat fungsi `async function muatProdukDariServer()`.
- [ ] Panggil API backend menggunakan `await fetch("https://.../api/products")`.
- [ ] Ubah data balasan server menjadi format JSON menggunakan `.json()`.
- [ ] Bungkus pemanggilan dengan blok `try...catch` untuk menangani jika terjadi kegagalan jaringan.

### 📄 Langkah 4: Buat Fitur Halaman Produk (Pagination)
- [ ] Tentukan jumlah barang per halaman: `const ITEMS_PER_PAGE = 6;`.
- [ ] Hitung total halaman menggunakan pembulatan ke atas:  
  `Math.ceil(daftarProduk.length / ITEMS_PER_PAGE)`.
- [ ] Potong daftar produk sesuai halaman yang sedang aktif dengan method `.slice(startIndex, endIndex)`.
- [ ] Pasang event listener pada tombol **Prev** dan **Next** untuk berpindah halaman.

### 🛒 Langkah 5: Kelola Keranjang Belanja & Simpan di LocalStorage
- [ ] Buat fungsi `tambahKeKeranjang(productId)`. Jika barang sudah ada di keranjang, tambahkan jumlahnya (`quantity += 1`).
- [ ] Simpan data keranjang ke memori browser menggunakan **LocalStorage**:  
  `localStorage.setItem("keranjang_tkj", JSON.stringify(keranjang));`
- [ ] Baca kembali data tersimpan saat website pertama kali dibuka:  
  `JSON.parse(localStorage.getItem("keranjang_tkj"));`
- [ ] Sediakan tombol kurangi kuantitas (`-`), tambah kuantitas (`+`), dan tombol hapus barang dari keranjang.

### 🪟 Langkah 6: Buat Modal Pop-up Interaktif
- [ ] Buka modal dengan menghapus class `.hidden`:  
  `cartModal.classList.remove("hidden");`
- [ ] Tutup modal dengan menambahkan class `.hidden`:  
  `cartModal.classList.add("hidden");`
- [ ] Pasang event listener pada tombol silang `[✕]` untuk menutup modal.
- [ ] Tambahkan fitur menutup modal saat area gelap di luar kotak modal diklik (`event.target === modal`).
- [ ] Tambahkan fitur menutup modal saat tombol keyboard **Escape (ESC)** ditekan.

### 📝 Langkah 7: Formulir Checkout & Kirim Pesanan
- [ ] Pasang event listener pada formulir: `checkoutForm.addEventListener("submit", ...)`.
- [ ] Tambahkan `event.preventDefault()` di baris awal fungsi listener agar halaman tidak me-reload sendiri.
- [ ] Ambil nilai input nama pembeli dan bersihkan spasi dengan `.value.trim()`.
- [ ] Kirim data barang belanjaan ke endpoint `/api/orders/checkout` menggunakan method `POST`.
- [ ] Setelah berhasil, tutup modal keranjang, bersihkan data LocalStorage, dan buka modal invoice.

### 💳 Langkah 8: Tampilkan Invoice & Pembayaran Sandbox
- [ ] Tampilkan Nomor Pesanan, Total Tagihan, dan Nomor **BCA Virtual Account** di layar invoice.
- [ ] Pasang event listener pada tombol **"Pay Now"**:
  - Kirim nomor Virtual Account ke endpoint `/api/sandbox/pay` dengan method `POST`.
  - Saat pembayaran berhasil, ubah badge status dari kuning (**PENDING**) menjadi hijau (**PAID**).
  - Tampilkan konfirmasi bahwa pembayaran telah berhasil diterima.

---

## 💡 4. Tips & Trik Ngoding

1. **Gunakan Developer Tools Browser (`F12`)**:
   - Tab **Console**: Cek baris error berwarna merah jika ada tombol atau fungsi yang tidak berjalan.
   - Tab **Network**: Lihat status dan data request setiap kali fungsi `fetch()` mengirim atau menerima data dari server.
   - Tab **Application ➔ Local Storage**: Periksa data keranjang `keranjang_tkj` yang tersimpan di browsermu.
2. **Mencegah Form Me-reload Halaman**:
   - Selalu pasang `event.preventDefault()` pada fungsi yang menangani event submit formulir.
3. **Menyalin Nomor Virtual Account Otomatis**:
   - Gunakan fungsi bawaan JavaScript modern:  
     `navigator.clipboard.writeText(nomorVA);`  
     untuk langsung menyalin teks nomor VA ke clipboard.

---

## 🔑 5. Kunci Jawaban & Referensi Solusi

Jika kamu mengalami kesulitan saat mengerjakan, kode lengkap yang sudah jadi dan berfungsi dapat kamu pelajari langsung pada file proyek ini:

* 📄 **[`index.html`](./index.html)** – Referensi struktur kerangka HTML.
* 🎨 **[`style.css`](./style.css)** – Referensi desain CSS, grid, modal, dan menu mobile.
* ⚙️ **[`app.js`](./app.js)** – Referensi logika JavaScript (Fetch API, Pagination, Modal, Checkout, dan LocalStorage).

---

*Selamat mengerjakan Final Project! Semoga sukses dan selamat belajar! 🚀✨*
