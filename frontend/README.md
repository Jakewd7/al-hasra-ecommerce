# 🛍️ Panduan Final Project: Bikin Toko Online "TKJ's Store"
> **Buku Panduan & Lembar Kerja Siswa Kelas 10**  
> Praktik Web Programming: HTML, CSS, dan JavaScript Dasar

---

## 👋 Halo Teman-Teman!

Selamat datang di **Final Project Web Programming**! 🎉  
Di proyek ini, kamu akan membuat website toko online bernama **TKJ's Store** yang terhubung langsung ke server backend milik tim kami.

Kamu akan mempraktikkan cara menampilkan katalog produk dari server, memasukkan barang ke keranjang belanja, melakukan checkout pesanan, mendapatkan nomor BCA Virtual Account, dan melakukan simulasi pembayaran hingga lunas.

> 💡 **Ketinggalan langkah? Tenang.** README ini bukan cuma langkah-langkah. Di dalamnya sudah ada:
> - 🔧 Cara menjalankan proyek (Bagian 3)
> - ✅ Checklist per langkah (Bagian 4)
> - 🆘 **Kalau kamu ketinggalan / ragu, baca "Bagian 5: Jangan Panik — Panduan Kalau Ketinggalan"**
> - 🎨 **Di mana ubah warna tombol, nama toko, font, dll → "Bagian 6: Panduan Kustomisasi"**
> - 🚨 Daftar error umum & solusinya (Bagian 7)
> - 📜 **Kode lengkap semua script (Bagian 9)** — tinggal copy kalau kamu tertinggal.

---

## 🎒 1. Alat yang Digunakan

1. **Visual Studio Code (VS Code)** – Aplikasi untuk menulis kode.
2. **Ekstensi Live Server di VS Code (opsional)** – Untuk menjalankan website di browser secara otomatis.
3. **Web Browser (Google Chrome atau Microsoft Edge, etc.)** – Untuk melihat website dan mengecek error lewat menu Inspect/Console (`F12`).
4. **Struktur File Proyek di folder `frontend/`**:

| File | Isinya | Untuk apa |
|------|--------|-----------|
| `index.html` | Kerangka struktur halaman web | Menyusun navbar, katalog, modal, invoice |
| `style.css` | Desain tampilan, warna, dan tata letak | Mengatur warna, ukuran, posisi, font |
| `app.js` | Logika interaktif & pemanggilan API | Ambil produk, keranjang, checkout, bayar |

> ⚠️ **Penting:** ketiga file ini harus ada di **satu folder yang sama**. Kalau `style.css` atau `app.js` dipindah ke folder lain, website akan tampil polos tanpa desain.

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

> 📌 Alamat server ini disimpan di **`app.js` baris 1** pada variabel `API_BASE`. Kalau alamat server berubah, cukup ubah **satu baris itu** saja.

---

## ▶️ 3. Cara Menjalankan Proyek

Pilih salah satu cara di bawah ini:

### Cara A — Langsung buka di browser (paling gampang)
1. Buka folder `frontend/` di VS Code.
2. Klik-kanan file **`index.html`** → **Open with Live Server** *(kalau punya ekstensi Live Server)*, **atau** klik dua kali file `index.html` sampai terbuka di Chrome/Edge.
3. Halaman TKJ's Store akan terbuka. Produk otomatis muncul dari server.

### Cara B — Memakai Live Server (disarankan)
1. Install ekstensi **Live Server** di VS Code (Extensions → cari "Live Server").
2. Klik-kanan `index.html` → **Open with Live Server**.
3. Browser terbuka otomatis di alamat `http://127.0.0.1:5500`.

### Cara Cek Berhasil atau Belum
* Kalau muncul daftar produk + gambar → ✅ **berhasil**.
* Kalau muncul tulisan *"Loading products from catalog..."* terus → ❌ koneksi ke server bermasalah, lihat **Bagian 7**.

---

## 📋 4. Langkah Demi Langkah Mengerjakan Proyek

Ikuti langkah-langkah di bawah ini secara berurutan. Centang `[x]` kalau sudah selesai.

### 🧩 Langkah 1: Siapkan Kerangka Halaman (`index.html`)
- [ ] Buat bagian atas (Navbar) yang berisi logo toko **TKJ's Store**, menu navigasi, dan tombol keranjang (`🛒 Cart`).
- [ ] Buat tombol menu hamburger untuk tampilan di layar HP.
- [ ] Buat bagian pembuka (Hero Banner) dengan judul toko dan tombol navigasi ke katalog.
- [ ] Siapkan kontainer untuk menampilkan produk: `<div id="product-grid"></div>`.
- [ ] Buat kontrol navigasi halaman produk (Tombol *Prev*, teks *Page Info*, dan tombol *Next*).
- [ ] Buat struktur pop-up (Modal) untuk Keranjang Belanja dan Formulir Pemesanan.
- [ ] Buat struktur pop-up (Modal) untuk Invoice Tagihan Pembayaran.

> 🔎 **Sudah jadi?** Lihat `index.html` baris 11–110 (navbar + hero + katalog) dan baris 70–178 (modal keranjang + invoice).

### 🎨 Langkah 2: Percantik Tampilan dengan CSS (`style.css`)
- [ ] Atur font sistem bawaan (`font-family: Arial, sans-serif`) agar ringan dan rapi.
- [ ] Gunakan **CSS Grid** pada `#product-grid` agar kartu produk tersusun rapi.
- [ ] Buat class `.modal-backdrop` dengan layar belakang transparan gelap untuk efek pop-up melayang di tengah layar.
- [ ] Buat class `.hidden { display: none !important; }` untuk menyembunyikan modal saat tidak aktif.
- [ ] Tambahkan aturan media query `@media (max-width: 768px)` agar menu navigasi HP berubah menjadi menu hamburger.

### 🚀 Langkah 3: Ambil Data Produk dari Server (`app.js`)
- [ ] Buat fungsi `async function muatProdukDariServer()`.
- [ ] Panggil API backend menggunakan `await fetch(API_BASE + "/products")`.
- [ ] Ubah data balasan server menjadi format JSON menggunakan `.json()`.
- [ ] Bungkus pemanggilan dengan blok `try...catch` untuk menangani jika terjadi kegagalan jaringan.

### 📄 Langkah 4: Buat Fitur Halaman Produk (Pagination)
- [ ] Tentukan jumlah barang per halaman: `const ITEMS_PER_PAGE = 6;`.
- [ ] Hitung total halaman menggunakan pembulatan ke atas: `Math.ceil(daftarProduk.length / ITEMS_PER_PAGE)`.
- [ ] Potong daftar produk sesuai halaman yang sedang aktif dengan method `.slice(startIndex, endIndex)`.
- [ ] Pasang event listener pada tombol **Prev** dan **Next** untuk berpindah halaman.

### 🛒 Langkah 5: Kelola Keranjang Belanja & Simpan di LocalStorage
- [ ] Buat fungsi `tambahKeKeranjang(productId)`. Jika barang sudah ada di keranjang, tambahkan jumlahnya (`quantity += 1`).
- [ ] Simpan data keranjang ke memori browser menggunakan **LocalStorage**: `localStorage.setItem("keranjang_tkj", JSON.stringify(keranjang));`
- [ ] Baca kembali data tersimpan saat website pertama kali dibuka: `JSON.parse(localStorage.getItem("keranjang_tkj"));`
- [ ] Sediakan tombol kurangi kuantitas (`-`), tambah kuantitas (`+`), dan tombol hapus barang dari keranjang.

### 🪟 Langkah 6: Buat Modal Pop-up Interaktif
- [ ] Buka modal dengan menghapus class `.hidden`: `cartModal.classList.remove("hidden");`
- [ ] Tutup modal dengan menambahkan class `.hidden`: `cartModal.classList.add("hidden");`
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

## 🆘 5. Jangan Panik — Panduan Kalau Ketinggalan

Lewat dari langkah yang diinstruksikan? Jangan bingung. Ikuti alur ini:

1. **Lihat langkahmu yang mana.** Buka **Bagian 4** dan cari langkah yang belum dicentang.
2. **Bandingkan kode kamu dengan kode lengkap.** Kode yang sudah 100% jalan ada di **Bagian 9**:
   - Kerangka halaman → `index.html`
   - Desain & warna → `style.css`
   - Logika & API → `app.js`
3. **Cari bagian yang salah.** Gunakan `Ctrl+F` (Windows) / `Cmd+F` (Mac) di editor untuk mencari nama fungsi/class. Contoh: cari `muatProdukDariServer` untuk memastikan fungsinya ada.
4. **Copy-paste kalau mentok.** Kalau kode kamu berantakan, **ganti seluruh isi file** kamu dengan kode di Bagian 9 yang sesuai. Tidak masalah — yang penting kamu paham cara kerjanya.
5. **Cek error di Console.** Tekan `F12` → tab **Console**. Tulisan merah = petunjuk di mana masalahnya. Cocokkan dengan **Bagian 7**.
6. **Masih error?** Cek urutan hal ini dulu:
   - [ ] Tiga file (`index.html`, `style.css`, `app.js`) ada di **satu folder**.
   - [ ] Di `index.html` sudah ada `<link rel="stylesheet" href="style.css">` dan `<script src="app.js"></script>`.
   - [ ] Nama `id` di HTML **sama persis** dengan yang dipanggil di `app.js` (contoh `id="product-grid"`).
   - [ ] Laptop terhubung internet (data produk diambil dari server).

> ✅ **Aturan emas:** kalau ragu, **README ini adalah kunci jawaban**. Kode di Bagian 9 adalah versi final yang sudah pasti jalan.

---

## 🎨 6. Panduan Kustomisasi (Ubah Warna, Tombol, Nama Toko, dll)

Mau bikin tampilan kamu beda dari teman? Semua tempat yang bisa diubah ada di sini.

### 🎯 Warna Utama (Yang Paling Sering Diubah)
Warna biru `#0056b3` adalah **warna utama (primary)** toko. Ubah di `style.css`:

| Mau ubah jadi | Cari di `style.css` | Ganti nilai |
|---------------|---------------------|-------------|
| Warna biru utama | semua kemunculan `#0056b3` | `#0056b3` → warna baru, mis. `#7c3aed` (ungu) |

> 💡 **Tips cepat:** tekan `Ctrl+H` / `Cmd+H` (Find & Replace) di VS Code, cari `#0056b3`, ganti semua dengan warna pilihanmu. Semua teks, tombol, dan harga ikut berubah otomatis.

### 🎨 Warna Tombol (Per Tombol)

| Yang mau diubah | Class di CSS | Baris | Warna contoh |
|-----------------|--------------|-------|--------------|
| Tombol **Add to Cart** / tombol biru utama | `.btn-primary` | ± 242 | `background-color: #0056b3;` |
| Tombol **Prev / Next / Continue Shopping** (abu) | `.btn-secondary` | ± 251 | `background-color: #e2e8f0;` |
| Tombol **Pay Now** (hijau) | `.btn-success` | ± 265 | `background-color: #16a34a;` |
| Tombol **🛒 Cart** di navbar | `.btn-cart` | ± 58 | `background-color: #ebf5ff;` dan `color: #0056b3;` |
| Tombol **Copy** nomor VA | `.btn-copy` | ± 539 | `color: #0056b3;` |
| Tombol **Remove** item keranjang (merah) | `.btn-delete-item` | ± 406 | `color: #e53e3e;` |

### 🌈 Warna Latar & Teks Halaman

| Yang mau diubah | Class / selector | Baris | Nilai |
|-----------------|------------------|-------|-------|
| Warna latar seluruh halaman | `body` → `background-color` | ± 9 | `#f7f9fc` |
| Warna teks seluruh halaman | `body` → `color` | ± 10 | `#333333` |
| Warna latar Navbar | `.navbar` | ± 21 | `#ffffff` |
| Warna teks logo "TKJ's Store" | `.logo h2` → `color` | ± 38 | `#0056b3` |
| Warna latar Hero (gradasi) | `.hero-section` → `background` | ± 79 | `linear-gradient(...)` |
| Warna badge di Hero | `.hero-badge` | ± 96 | `#dbeafe` / `#1d4ed8` |
| Warna kartu produk | `.product-card` | ± 150 | `#ffffff` |
| Warna latar belakang pop-up (modal) | `.modal-box` | ± 297 | `#ffffff` |
| Warna lapisan gelap di belakang modal | `.modal-backdrop` | ± 284 | `rgba(0, 0, 0, 0.55)` |
| Warna badge status **PENDING** (kuning) | `.badge-pending` | ± 490 | `#d97706` |
| Warna badge status **PAID** (hijau) | `.badge-paid` | ± 495 | `#16a34a` |
| Warna footer | `.footer` | ± 606 | `#ffffff` |

### 🔤 Font / Tulisan

| Yang mau diubah | Lokasi | Baris | Contoh ganti |
|-----------------|--------|-------|--------------|
| Jenis font seluruh website | `body` → `font-family` | ± 8 | `'Poppins', Arial, sans-serif` |
| Ukuran judul Hero | `.hero-text h1` → `font-size` | ± 106 | `2.2rem` → `2.8rem` |
| Ukuran judul katalog | `.section-header h2` | ± 132 | `1.8rem` |

> ⚠️ Kalau pakai font dari internet (mis. Poppins), tambahkan dulu di `index.html` di dalam `<head>`:
> ```html
> <link href="https://fonts.googleapis.com/css2?family=Poppins&display=swap" rel="stylesheet">
> ```

### 🏪 Nama Toko & Teks

| Yang mau diubah | File | Lokasi |
|-----------------|------|--------|
| Nama toko "TKJ's Store" | `index.html` | baris 6 (judul tab), 14 (logo), 35 (hero) |
| Judul katalog "Featured Catalog" | `index.html` | baris 48 |
| Teks footer `© 2026 TKJ's Store` | `index.html` | baris 182 |

### 🧩 Tata Letak & Jumlah Produk

| Yang mau diubah | Lokasi | Baris | Catatan |
|-----------------|--------|-------|---------|
| Jumlah kolom kartu produk | `.product-grid` → `grid-template-columns` | ± 144 | `minmax(230px, 1fr)` — angka lebih kecil = kartu lebih kecil |
| Jumlah produk per halaman | `app.js` → `const ITEMS_PER_PAGE` | 3 | ganti `6` jadi angka lain, mis. `3` atau `9` |
| Sudut melengkung (radius) | cari `border-radius` | — | angka lebih besar = sudut lebih bulat |
| Efek bayangan kartu | `box-shadow` di `.product-card` | ± 154 | — |

### 💰 Format Harga & Mata Uang

| Yang mau diubah | Lokasi | Baris | Contoh |
|-----------------|--------|-------|--------|
| Simbol mata uang (Rp) | `app.js` cari `"Rp "` | 95, 168, 185, 353 | ganti `"Rp "` jadi `"$ "` |
| Format angka (id-ID) | `app.js` cari `toLocaleString("id-ID")` | 95, 168, 185, 353 | `"id-ID"` → `"en-US"` |

> 📌 **Ingat:** perubahan tampilan (warna, ukuran, font) → di `style.css`. Perubahan isi/teks → di `index.html`. Perubahan logika/harga/jumlah → di `app.js`.

---

## 🚨 7. Error Umum & Solusinya

| Gejala | Kemungkinan penyebab | Solusi |
|--------|----------------------|--------|
| Muncul tulisan *"Loading products..."* terus | Tidak ada koneksi internet / server mati | Cek internet. Buka Console (`F12`) → tab Network → lihat request `products` |
| Muncul *"Unable to load products"* | Server backend tidak bisa diakses | Cek kembali alamat `API_BASE` di `app.js` baris 1 |
| Console: `Failed to fetch` | Koneksi diblokir / salah alamat | Pastikan internet jalan & alamat server benar |
| Tampilan polos tanpa warna sama sekali | `style.css` tidak terpasang | Cek `<link rel="stylesheet" href="style.css">` di `index.html`; pastikan file ada di folder yang sama |
| Tombol tidak berfungsi (klik tidak ada reaksi) | `app.js` tidak ter-load, atau `id` tidak cocok | Cek `<script src="app.js"></script>` di `index.html` baris 186; cocokkan `id` HTML dengan `getElementById` di `app.js` |
| Console: `Cannot read properties of null` | `id` di HTML salah beda dengan `app.js` | Pastikan mis. `id="product-grid"` benar-benar ada di HTML |
| Halaman me-reload saat klik "Proceed to Order" | Lupa `event.preventDefault()` | Cek `checkoutForm` listener di `app.js` baris 290 |
| Keranjang kosong lagi setelah refresh | LocalStorage gagal dibaca/ditulis | Cek `F12 → Application → Local Storage`, cari key `keranjang_tkj` |
| Harga tampil `Rp NaN` | `price` dari server bukan angka | Pastikan pakai `Number(p.price)` sebelum dihitung |
| Nomor VA tidak muncul | Data `payment.va_number` kosong | Cek balikan server di tab **Network** saat checkout |
| Gambar produk tidak muncul | `image_url` kosong/rusak | Gambar akan pakai placeholder otomatis, ini normal |
| Menu HP tidak muncul | Layar belum dikecilkan / lebar > 768px | Perkecil jendela browser sampai di bawah 768px, atau pakai mode device (`F12` → ikon HP) |
| Modal tidak mau tertutup | Class `.hidden` tidak ada di CSS | Pastikan `.hidden { display: none !important; }` ada di `style.css` baris 622 |

### 🔍 Cara Debug Cepat dengan `F12`
1. **Tab Console** → baca tulisan **merah**. Itu letak error-nya (beserta nama file + baris).
2. **Tab Network** → klik tombol yang bermasalah, lihat request `products` / `checkout` / `pay`. Kalau status `200` = sukses, `4xx/5xx` = error dari server.
3. **Tab Application → Local Storage** → lihat isi `keranjang_tkj` untuk memastikan keranjang tersimpan.

---

## 💡 8. Tips & Trik Ngoding

1. **Gunakan Developer Tools Browser (`F12`)**:
   - Tab **Console**: Cek baris error berwarna merah jika ada tombol atau fungsi yang tidak berjalan.
   - Tab **Network**: Lihat status dan data request setiap kali fungsi `fetch()` mengirim atau menerima data dari server.
   - Tab **Application ➔ Local Storage**: Periksa data keranjang `keranjang_tkj` yang tersimpan di browsermu.
2. **Mencegah Form Me-reload Halaman**:
   - Selalu pasang `event.preventDefault()` pada fungsi yang menangani event submit formulir.
3. **Menyalin Nomor Virtual Account Otomatis**:
   - Gunakan fungsi bawaan JavaScript modern `navigator.clipboard.writeText(nomorVA);` untuk langsung menyalin teks nomor VA ke clipboard.
4. **Selalu simpan (Save) file sebelum refresh browser.** Perubahan kode tidak akan muncul kalau belum di-save (`Ctrl+S` / `Cmd+S`).

---

## 📜 9. Kode Lengkap (Kunci Jawaban)

Kode di bawah ini adalah versi **lengkap dan sudah jalan**. Kalau kamu ketinggalan atau file kamu berantakan, tinggal cocokkan / copy-paste.

### 📄 `index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>TKJ's Store</title>
  <link rel="stylesheet" href="style.css" />
</head>
<body>

  <header class="navbar">
    <div class="container navbar-container">
      <div class="logo">
        <h2>TKJ's Store</h2>
      </div>
      <button id="hamburger-btn" class="hamburger-btn" aria-label="Toggle Navigation">
        <span></span>
        <span></span>
        <span></span>
      </button>
      <nav id="nav-menu" class="nav-menu">
        <a href="#landing" class="nav-link">Home</a>
        <a href="#katalog" class="nav-link">Catalog</a>
        <button id="btn-open-cart" class="btn-cart">
          🛒 Cart (<span id="cart-counter">0</span>)
        </button>
      </nav>
    </div>
  </header>

  <section id="landing" class="hero-section">
    <div class="container hero-container">
      <div class="hero-text">
        <span class="hero-badge">Special Collection 2026</span>
        <h1>Discover Quality Gear at <br><span class="text-primary">TKJ's Store</span></h1>
        <p>
          Explore premium tech essentials, accessories, and daily lifestyle products 
          with fast online ordering and seamless instant checkout.
        </p>
        <a href="#katalog" class="btn btn-primary">Shop Collection ↓</a>
      </div>
    </div>
  </section>

  <section id="katalog" class="catalog-section">
    <div class="container">
      <div class="section-header">
        <h2>Featured Catalog</h2>
        <p>Handpicked essentials crafted for everyday productivity and modern living.</p>
      </div>

      <div id="product-grid" class="product-grid">
        <div class="loading-state">Loading products from catalog...</div>
      </div>

      <div class="pagination-wrapper">
        <button id="btn-prev-page" class="btn btn-secondary btn-page">
          ◀ Prev
        </button>
        <span id="page-info" class="page-info">
          Page 1 of 1
        </span>
        <button id="btn-next-page" class="btn btn-secondary btn-page">
          Next ▶
        </button>
      </div>
    </div>
  </section>

  <div id="cart-modal" class="modal-backdrop hidden">
    <div class="modal-box">
      <div class="modal-header">
        <h3>🛒 Shopping Cart</h3>
        <button id="btn-close-cart" class="btn-close" title="Close">✕</button>
      </div>

      <div class="modal-body">
        <div id="cart-items-list" class="cart-items-list"></div>

        <div class="cart-total-box">
          <span>Total Amount:</span>
          <strong id="cart-grand-total" class="total-price">Rp 0</strong>
        </div>

        <div class="checkout-form-container">
          <h4>Checkout Details:</h4>
          <form id="checkout-form">
            <div class="form-group">
              <label for="buyer-name">Customer Name:</label>
              <input 
                type="text" 
                id="buyer-name" 
                placeholder="Enter your full name..." 
                autocomplete="off"
              />
            </div>
            <button type="submit" id="btn-submit-order" class="btn btn-primary btn-block">
              Proceed to Order ➔
            </button>
          </form>
        </div>
      </div>

      <div class="modal-footer">
        <button id="btn-continue-shopping" class="btn btn-secondary">
          Continue Shopping
        </button>
      </div>
    </div>
  </div>

  <div id="invoice-modal" class="modal-backdrop hidden">
    <div class="modal-box invoice-box">
      <div class="modal-header">
        <h3>🧾 Order Invoice</h3>
        <button id="btn-close-invoice" class="btn-close" title="Close">✕</button>
      </div>

      <div class="modal-body">
        <div class="status-banner">
          <span>Order Status:</span>
          <span id="invoice-status" class="badge badge-pending">PENDING</span>
        </div>

        <div class="invoice-details">
          <div class="detail-row">
            <span>Order ID:</span>
            <strong id="invoice-order-number">-</strong>
          </div>
          <div class="detail-row">
            <span>Customer:</span>
            <strong id="invoice-buyer-name">-</strong>
          </div>
          <div class="detail-row">
            <span>Payment Method:</span>
            <strong>BCA Virtual Account</strong>
          </div>
          <div class="detail-row va-highlight">
            <span>VA Number:</span>
            <div class="va-wrapper">
              <strong id="invoice-va-number" class="va-code">-</strong>
              <button id="btn-copy-va" class="btn-copy" type="button">Copy</button>
            </div>
          </div>
          <div class="detail-row">
            <span>Total Due:</span>
            <strong id="invoice-total-amount" class="text-primary">-</strong>
          </div>
        </div>

        <div id="payment-instructions" class="payment-instructions">
          <p class="inst-title">Virtual Account Payment Steps:</p>
          <ol>
            <li>Open your Mobile Banking or ATM.</li>
            <li>Select <strong>Transfer ➔ Virtual Account</strong>.</li>
            <li>Input the Virtual Account number above and confirm the transaction.</li>
          </ol>
        </div>

        <div id="sandbox-action-box" class="sandbox-box">
          <button id="btn-pay-sandbox" class="btn btn-success btn-block">
            Pay Now
          </button>
        </div>

        <div id="payment-success-box" class="success-box hidden">
          <h4>Payment Confirmed</h4>
          <p>Your payment has been received and your order is now being processed.</p>
        </div>
      </div>

      <div class="modal-footer">
        <button id="btn-finish-invoice" class="btn btn-secondary">
          Close & Back to Store
        </button>
      </div>
    </div>
  </div>

  <footer class="footer">
    <div class="container">
      <p>© 2026 TKJ's Store. All rights reserved.</p>
    </div>
  </footer>

  <script src="app.js"></script>
</body>
</html>
```

### 🎨 `style.css`

```css
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: Arial, Helvetica, sans-serif;
  background-color: #f7f9fc;
  color: #333333;
  line-height: 1.6;
}

.container {
  max-width: 1050px;
  margin: 0 auto;
  padding: 0 20px;
}

.navbar {
  background-color: #ffffff;
  border-bottom: 1px solid #e2e8f0;
  position: sticky;
  top: 0;
  z-index: 100;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.04);
}

.navbar-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 65px;
}

.logo h2 {
  font-size: 1.3rem;
  color: #0056b3;
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-link {
  text-decoration: none;
  color: #4a5568;
  font-weight: bold;
  font-size: 0.95rem;
}

.nav-link:hover {
  color: #0056b3;
}

.btn-cart {
  background-color: #ebf5ff;
  border: 1px solid #bce0fd;
  color: #0056b3;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: bold;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: background-color 0.2s;
}

.btn-cart:hover {
  background-color: #d8edfe;
}

.hero-section {
  padding: 60px 0;
  background: linear-gradient(180deg, #ffffff 0%, #f0f6ff 100%);
  border-bottom: 1px solid #e2e8f0;
}

.hero-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.hero-text {
  max-width: 650px;
  text-align: center;
}

.hero-badge {
  display: inline-block;
  background-color: #dbeafe;
  color: #1d4ed8;
  font-size: 0.8rem;
  font-weight: bold;
  padding: 4px 12px;
  border-radius: 12px;
  margin-bottom: 12px;
}

.hero-text h1 {
  font-size: 2.2rem;
  line-height: 1.25;
  color: #1a202c;
  margin-bottom: 14px;
}

.text-primary {
  color: #0056b3;
}

.hero-text p {
  font-size: 1rem;
  color: #4a5568;
  margin-bottom: 24px;
}

.catalog-section {
  padding: 50px 0 70px;
}

.section-header {
  text-align: center;
  margin-bottom: 30px;
}

.section-header h2 {
  font-size: 1.8rem;
  color: #2d3748;
  margin-bottom: 6px;
}

.section-header p {
  color: #718096;
  font-size: 0.95rem;
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(230px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.product-card {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  transition: transform 0.2s, box-shadow 0.2s;
}

.product-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 14px rgba(0, 0, 0, 0.08);
}

.product-img-box {
  width: 100%;
  height: 160px;
  background-color: #edf2f7;
  overflow: hidden;
}

.product-img-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-details {
  padding: 16px;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}

.product-title {
  font-size: 1rem;
  color: #2d3748;
  margin-bottom: 8px;
  line-height: 1.3;
}

.product-desc {
  font-size: 0.82rem;
  color: #718096;
  margin-bottom: 12px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  flex-grow: 1;
}

.product-price {
  font-size: 1.1rem;
  font-weight: bold;
  color: #0056b3;
  margin-bottom: 12px;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-top: 20px;
}

.btn-page {
  padding: 8px 16px;
  font-size: 0.9rem;
}

.page-info {
  font-size: 0.95rem;
  font-weight: bold;
  color: #4a5568;
}

.btn {
  display: inline-block;
  padding: 9px 16px;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: bold;
  text-decoration: none;
  cursor: pointer;
  border: none;
  text-align: center;
  transition: all 0.2s;
}

.btn-primary {
  background-color: #0056b3;
  color: #ffffff;
}

.btn-primary:hover {
  background-color: #004494;
}

.btn-secondary {
  background-color: #e2e8f0;
  color: #4a5568;
}

.btn-secondary:hover:not(:disabled) {
  background-color: #cbd5e1;
}

.btn-secondary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-success {
  background-color: #16a34a;
  color: #ffffff;
}

.btn-success:hover {
  background-color: #15803d;
}

.btn-block {
  width: 100%;
}

.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.55);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 16px;
}

.modal-backdrop.hidden {
  display: none !important;
}

.modal-box {
  background-color: #ffffff;
  width: 100%;
  max-width: 520px;
  border-radius: 8px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  max-height: 90vh;
  animation: modalIn 0.2s ease-out;
}

@keyframes modalIn {
  from { transform: scale(0.95); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}

.modal-header {
  padding: 16px 20px;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h3 {
  font-size: 1.15rem;
  color: #1a202c;
}

.btn-close {
  background: transparent;
  border: none;
  font-size: 1.25rem;
  cursor: pointer;
  color: #718096;
}

.btn-close:hover {
  color: #1a202c;
}

.modal-body {
  padding: 20px;
  overflow-y: auto;
}

.modal-footer {
  padding: 14px 20px;
  border-top: 1px solid #e2e8f0;
  display: flex;
  justify-content: flex-end;
}

.cart-items-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 16px;
  max-height: 220px;
  overflow-y: auto;
}

.cart-item-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 10px;
  border-bottom: 1px solid #edf2f7;
}

.cart-item-thumb {
  width: 48px;
  height: 48px;
  object-fit: cover;
  border-radius: 4px;
  background-color: #edf2f7;
}

.cart-item-info {
  flex-grow: 1;
}

.cart-item-name {
  font-size: 0.9rem;
  font-weight: bold;
  color: #2d3748;
}

.cart-item-price {
  font-size: 0.85rem;
  color: #0056b3;
}

.cart-item-qty {
  display: flex;
  align-items: center;
  gap: 6px;
}

.btn-qty {
  width: 24px;
  height: 24px;
  border: 1px solid #cbd5e1;
  background-color: #ffffff;
  border-radius: 3px;
  cursor: pointer;
  font-weight: bold;
}

.btn-delete-item {
  background: none;
  border: none;
  color: #e53e3e;
  font-size: 0.8rem;
  cursor: pointer;
  margin-left: 6px;
}

.btn-delete-item:hover {
  text-decoration: underline;
}

.cart-total-box {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  margin-bottom: 16px;
}

.total-price {
  font-size: 1.15rem;
  color: #0056b3;
}

.checkout-form-container {
  border-top: 1px solid #e2e8f0;
  padding-top: 16px;
}

.checkout-form-container h4 {
  font-size: 0.95rem;
  margin-bottom: 10px;
  color: #2d3748;
}

.form-group {
  margin-bottom: 12px;
}

.form-group label {
  display: block;
  font-size: 0.85rem;
  font-weight: bold;
  margin-bottom: 4px;
  color: #4a5568;
}

.form-group input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #cbd5e1;
  border-radius: 4px;
  font-size: 0.9rem;
}

.form-group input:focus {
  outline: none;
  border-color: #0056b3;
}

.status-banner {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 14px;
  background-color: #fef3c7;
  border: 1px solid #fde68a;
  border-radius: 6px;
  margin-bottom: 16px;
  font-size: 0.9rem;
}

.badge {
  padding: 3px 10px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: bold;
}

.badge-pending {
  background-color: #d97706;
  color: #ffffff;
}

.badge-paid {
  background-color: #16a34a;
  color: #ffffff;
}

.invoice-details {
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  padding: 14px;
  margin-bottom: 14px;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 0.88rem;
}

.detail-row:last-child {
  margin-bottom: 0;
  border-top: 1px dashed #cbd5e1;
  padding-top: 8px;
}

.va-highlight {
  background-color: #eff6ff;
  padding: 6px 8px;
  border-radius: 4px;
}

.va-wrapper {
  display: flex;
  align-items: center;
  gap: 8px;
}

.va-code {
  color: #0056b3;
  letter-spacing: 1px;
}

.btn-copy {
  background-color: #ffffff;
  border: 1px solid #bce0fd;
  color: #0056b3;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: bold;
  cursor: pointer;
}

.btn-copy:hover {
  background-color: #e8f4fd;
}

.payment-instructions {
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  padding: 12px 14px;
  font-size: 0.82rem;
  color: #4a5568;
  margin-bottom: 14px;
}

.payment-instructions .inst-title {
  font-weight: bold;
  margin-bottom: 4px;
  color: #2d3748;
}

.payment-instructions ol {
  padding-left: 18px;
  margin: 0;
}

.payment-instructions li {
  margin-bottom: 2px;
}

.sandbox-box {
  margin-bottom: 12px;
}

.success-box {
  background-color: #ecfdf5;
  border: 1px solid #a7f3d0;
  padding: 16px;
  border-radius: 6px;
  text-align: center;
}

.success-icon {
  font-size: 2rem;
  margin-bottom: 4px;
}

.success-box h4 {
  color: #15803d;
  margin-bottom: 4px;
}

.success-box p {
  font-size: 0.85rem;
  color: #4a5568;
}

.footer {
  text-align: center;
  padding: 24px 0;
  background-color: #ffffff;
  border-top: 1px solid #e2e8f0;
  color: #718096;
  font-size: 0.85rem;
}

.loading-state, .empty-state {
  text-align: center;
  padding: 30px;
  color: #718096;
  grid-column: 1 / -1;
}

.hidden {
  display: none !important;
}

.hamburger-btn {
  display: none;
  flex-direction: column;
  justify-content: space-between;
  width: 26px;
  height: 19px;
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 0;
}

.hamburger-btn span {
  width: 100%;
  height: 3px;
  background-color: #1a202c;
  border-radius: 2px;
  transition: all 0.3s ease;
}

.hamburger-btn.active span:nth-child(1) {
  transform: translateY(8px) rotate(45deg);
}

.hamburger-btn.active span:nth-child(2) {
  opacity: 0;
}

.hamburger-btn.active span:nth-child(3) {
  transform: translateY(-8px) rotate(-45deg);
}

@media (max-width: 768px) {
  .hero-container {
    flex-direction: column;
    text-align: center;
  }

  .hamburger-btn {
    display: flex;
  }

  .nav-menu {
    display: none;
    position: absolute;
    top: 65px;
    left: 0;
    width: 100%;
    background-color: #ffffff;
    border-bottom: 1px solid #e2e8f0;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    flex-direction: column;
    align-items: flex-start;
    padding: 16px 20px;
    gap: 12px;
  }

  .nav-menu.active {
    display: flex;
  }

  .nav-menu .nav-link {
    width: 100%;
    padding: 8px 0;
    border-bottom: 1px solid #f1f5f9;
  }

  .nav-menu .btn-cart {
    width: 100%;
    justify-content: center;
    margin-top: 6px;
  }
}
```

### ⚙️ `app.js`

```javascript
const API_BASE = "https://serverstaging-vps-5e6db201b3.asiahub.id/api";

const ITEMS_PER_PAGE = 6;
let currentPage = 1;

let daftarProduk = [];
let keranjang = [];
let pesananAktif = null;

function simpanKeranjangKeLocalStorage() {
  const jsonString = JSON.stringify(keranjang);
  localStorage.setItem("keranjang_tkj", jsonString);
}

function ambilKeranjangDariLocalStorage() {
  const dataTersimpan = localStorage.getItem("keranjang_tkj");

  if (dataTersimpan !== null) {
    try {
      keranjang = JSON.parse(dataTersimpan);
    } catch (err) {
      console.error(err);
      keranjang = [];
    }
  }
}

function bersihkanKeranjang() {
  localStorage.removeItem("keranjang_tkj");
  keranjang = [];
  perbaruiTampilanKeranjang();
}

async function muatProdukDariServer() {
  const container = document.getElementById("product-grid");

  try {
    const response = await fetch(API_BASE + "/products");

    if (!response.ok) {
      throw new Error("Failed to fetch product catalog. Status: " + response.status);
    }

    daftarProduk = await response.json();

    currentPage = 1;
    tampilkanKatalogDenganPagination();

  } catch (error) {
    console.error("Error loading products:", error);
    container.innerHTML = `
      <div class="empty-state" style="color: #dc3545;">
        ⚠️ Unable to load products: ${error.message}
      </div>
    `;
  }
}

function tampilkanKatalogDenganPagination() {
  const gridContainer = document.getElementById("product-grid");
  const pageInfo = document.getElementById("page-info");
  const btnPrev = document.getElementById("btn-prev-page");
  const btnNext = document.getElementById("btn-next-page");

  if (!daftarProduk || daftarProduk.length === 0) {
    gridContainer.innerHTML = "<div class='empty-state'>No products available right now.</div>";
    pageInfo.textContent = "Page 0 of 0";
    btnPrev.disabled = true;
    btnNext.disabled = true;
    return;
  }

  const totalPages = Math.ceil(daftarProduk.length / ITEMS_PER_PAGE);

  if (currentPage > totalPages) currentPage = totalPages;
  if (currentPage < 1) currentPage = 1;

  const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
  const endIndex = startIndex + ITEMS_PER_PAGE;
  const produkHalamanIni = daftarProduk.slice(startIndex, endIndex);

  let cardsHtml = "";
  for (let i = 0; i < produkHalamanIni.length; i++) {
    const p = produkHalamanIni[i];
    const gambar = p.image_url || "https://via.placeholder.com/240x160?text=Product";

    cardsHtml += `
      <div class="product-card">
        <div class="product-img-box">
          <img src="${gambar}" alt="${p.name}" onerror="this.src='https://via.placeholder.com/240x160?text=No+Image'" />
        </div>
        <div class="product-details">
          <h3 class="product-title">${p.name}</h3>
          <p class="product-desc">${p.description || "No description available."}</p>
          <div class="product-price">Rp ${Number(p.price).toLocaleString("id-ID")}</div>
          <button class="btn btn-primary btn-block" onclick="tambahKeKeranjang('${p.id}')">
            + Add to Cart
          </button>
        </div>
      </div>
    `;
  }
  gridContainer.innerHTML = cardsHtml;

  pageInfo.textContent = `Page ${currentPage} of ${totalPages}`;

  btnPrev.disabled = (currentPage === 1);
  btnNext.disabled = (currentPage === totalPages);
}

function tambahKeKeranjang(productId) {
  const produk = daftarProduk.find(item => item.id === productId);
  if (!produk) return;

  const itemAda = keranjang.find(item => item.product_id === productId);

  if (itemAda) {
    itemAda.quantity += 1;
  } else {
    keranjang.push({
      product_id: produk.id,
      name: produk.name,
      price: produk.price,
      image_url: produk.image_url,
      quantity: 1
    });
  }

  simpanKeranjangKeLocalStorage();
  perbaruiTampilanKeranjang();
  alert(`"${produk.name}" has been added to your cart.`);
}

function ubahJumlahItem(productId, perubahan) {
  const item = keranjang.find(i => i.product_id === productId);
  if (!item) return;

  item.quantity += perubahan;

  if (item.quantity <= 0) {
    keranjang = keranjang.filter(i => i.product_id !== productId);
  }

  simpanKeranjangKeLocalStorage();
  perbaruiTampilanKeranjang();
}

function hapusItemDariKeranjang(productId) {
  keranjang = keranjang.filter(i => i.product_id !== productId);
  simpanKeranjangKeLocalStorage();
  perbaruiTampilanKeranjang();
}

function perbaruiTampilanKeranjang() {
  const cartCounter = document.getElementById("cart-counter");
  const cartList = document.getElementById("cart-items-list");
  const grandTotalEl = document.getElementById("cart-grand-total");

  let totalJumlah = 0;
  let totalBiaya = 0;

  for (let i = 0; i < keranjang.length; i++) {
    totalJumlah += keranjang[i].quantity;
    totalBiaya += (keranjang[i].price * keranjang[i].quantity);
  }

  cartCounter.textContent = totalJumlah;
  grandTotalEl.textContent = "Rp " + totalBiaya.toLocaleString("id-ID");

  if (keranjang.length === 0) {
    cartList.innerHTML = "<p style='text-align: center; color: #718096; padding: 20px 0;'>Your shopping cart is currently empty.</p>";
    return;
  }

  let itemsHtml = "";
  for (let i = 0; i < keranjang.length; i++) {
    const item = keranjang[i];
    const gambar = item.image_url || "https://via.placeholder.com/48";

    itemsHtml += `
      <div class="cart-item-row">
        <img class="cart-item-thumb" src="${gambar}" alt="${item.name}" />
        <div class="cart-item-info">
          <div class="cart-item-name">${item.name}</div>
          <div class="cart-item-price">Rp ${Number(item.price).toLocaleString("id-ID")}</div>
        </div>
        <div class="cart-item-qty">
          <button class="btn-qty" onclick="ubahJumlahItem('${item.product_id}', -1)">-</button>
          <span>${item.quantity}</span>
          <button class="btn-qty" onclick="ubahJumlahItem('${item.product_id}', 1)">+</button>
          <button class="btn-delete-item" onclick="hapusItemDariKeranjang('${item.product_id}')">Remove</button>
        </div>
      </div>
    `;
  }

  cartList.innerHTML = itemsHtml;
}

const cartModal = document.getElementById("cart-modal");
const invoiceModal = document.getElementById("invoice-modal");

const hamburgerBtn = document.getElementById("hamburger-btn");
const navMenu = document.getElementById("nav-menu");

if (hamburgerBtn && navMenu) {
  hamburgerBtn.addEventListener("click", function() {
    hamburgerBtn.classList.toggle("active");
    navMenu.classList.toggle("active");
  });

  document.querySelectorAll(".nav-link").forEach(link => {
    link.addEventListener("click", function() {
      hamburgerBtn.classList.remove("active");
      navMenu.classList.remove("active");
    });
  });
}

function bukaModalKeranjang() {
  if (hamburgerBtn && navMenu) {
    hamburgerBtn.classList.remove("active");
    navMenu.classList.remove("active");
  }
  perbaruiTampilanKeranjang();
  cartModal.classList.remove("hidden");
}

function tutupModalKeranjang() {
  cartModal.classList.add("hidden");
}

function tutupModalInvoice() {
  invoiceModal.classList.add("hidden");
}

document.getElementById("btn-open-cart").addEventListener("click", bukaModalKeranjang);
document.getElementById("btn-close-cart").addEventListener("click", tutupModalKeranjang);
document.getElementById("btn-continue-shopping").addEventListener("click", tutupModalKeranjang);

document.getElementById("btn-close-invoice").addEventListener("click", tutupModalInvoice);
document.getElementById("btn-finish-invoice").addEventListener("click", tutupModalInvoice);

document.getElementById("btn-copy-va").addEventListener("click", function() {
  const vaText = document.getElementById("invoice-va-number").textContent;
  if (vaText && vaText !== "-") {
    navigator.clipboard.writeText(vaText);
    alert("Virtual Account number copied: " + vaText);
  }
});

window.addEventListener("click", function(event) {
  if (event.target === cartModal) {
    tutupModalKeranjang();
  }
  if (event.target === invoiceModal) {
    tutupModalInvoice();
  }
  if (hamburgerBtn && navMenu && !navMenu.contains(event.target) && !hamburgerBtn.contains(event.target)) {
    hamburgerBtn.classList.remove("active");
    navMenu.classList.remove("active");
  }
});

window.addEventListener("keydown", function(event) {
  if (event.key === "Escape") {
    tutupModalKeranjang();
    tutupModalInvoice();
  }
});

document.getElementById("btn-prev-page").addEventListener("click", function() {
  if (currentPage > 1) {
    currentPage--;
    tampilkanKatalogDenganPagination();
  }
});

document.getElementById("btn-next-page").addEventListener("click", function() {
  const totalPages = Math.ceil(daftarProduk.length / ITEMS_PER_PAGE);
  if (currentPage < totalPages) {
    currentPage++;
    tampilkanKatalogDenganPagination();
  }
});

const checkoutForm = document.getElementById("checkout-form");

checkoutForm.addEventListener("submit", async function(event) {
  event.preventDefault();

  const buyerInput = document.getElementById("buyer-name");
  const namaPembeli = buyerInput.value.trim();

  if (namaPembeli === "") {
    alert("Please enter your name!");
    buyerInput.focus();
    return;
  }

  if (keranjang.length === 0) {
    alert("Your shopping cart is empty!");
    return;
  }

  const btnSubmit = document.getElementById("btn-submit-order");
  btnSubmit.disabled = true;
  btnSubmit.textContent = "Processing Order...";

  try {
    const itemsData = keranjang.map(item => ({
      product_id: item.product_id,
      quantity: item.quantity
    }));

    const response = await fetch(API_BASE + "/orders/checkout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ items: itemsData })
    });

    if (!response.ok) {
      const errJson = await response.json();
      throw new Error(errJson.error || "Failed to place order.");
    }

    const orderData = await response.json();

    pesananAktif = orderData;
    pesananAktif.buyer_name = namaPembeli;

    tutupModalKeranjang();
    buyerInput.value = "";
    bersihkanKeranjang();

    tampilkanInvoice(pesananAktif);

  } catch (error) {
    console.error("Checkout failed:", error);
    alert("Checkout failed: " + error.message);
  } finally {
    btnSubmit.disabled = false;
    btnSubmit.textContent = "Proceed to Order ➔";
  }
});

function tampilkanInvoice(order) {
  document.getElementById("invoice-order-number").textContent = order.order_number;
  document.getElementById("invoice-buyer-name").textContent = order.buyer_name || "Valued Customer";
  document.getElementById("invoice-va-number").textContent = order.payment ? order.payment.va_number : "-";
  document.getElementById("invoice-total-amount").textContent = "Rp " + Number(order.total_amount).toLocaleString("id-ID");

  const statusBadge = document.getElementById("invoice-status");
  statusBadge.className = "badge badge-pending";
  statusBadge.textContent = "PENDING";

  document.getElementById("payment-instructions").style.display = "block";
  document.getElementById("sandbox-action-box").style.display = "block";
  document.getElementById("payment-success-box").classList.add("hidden");

  invoiceModal.classList.remove("hidden");
}

const btnPaySandbox = document.getElementById("btn-pay-sandbox");

btnPaySandbox.addEventListener("click", async function() {
  if (!pesananAktif || !pesananAktif.payment || !pesananAktif.payment.va_number) {
    alert("Virtual Account number is invalid!");
    return;
  }

  btnPaySandbox.disabled = true;
  btnPaySandbox.textContent = "Processing Payment...";

  try {
    const response = await fetch(API_BASE + "/sandbox/pay", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        va_number: pesananAktif.payment.va_number
      })
    });

    if (!response.ok) {
      const errData = await response.json();
      throw new Error(errData.error || "Sandbox payment failed.");
    }

    await response.json();

    const statusBadge = document.getElementById("invoice-status");
    statusBadge.className = "badge badge-paid";
    statusBadge.textContent = "PAID";

    document.getElementById("sandbox-action-box").style.display = "none";
    document.getElementById("payment-instructions").style.display = "none";
    document.getElementById("payment-success-box").classList.remove("hidden");

  } catch (error) {
    console.error("Payment failed:", error);
    alert("Payment failed: " + error.message);
  } finally {
    btnPaySandbox.disabled = false;
    btnPaySandbox.textContent = "Pay Now";
  }
});

window.addEventListener("DOMContentLoaded", function() {
  ambilKeranjangDariLocalStorage();
  perbaruiTampilanKeranjang();
  muatProdukDariServer();
});
```

---

## 📖 10. Ringkasan Alur Website

```
Buka index.html
      │
      ▼
app.js jalan → muatProdukDariServer() → GET /products
      │
      ▼
Produk tampil (pagination 6 per halaman)
      │
      ▼  klik "+ Add to Cart"
Keranjang tersimpan di LocalStorage
      │
      ▼  klik "🛒 Cart" → isi nama → "Proceed to Order"
POST /orders/checkout → dapat nomor BCA Virtual Account
      │
      ▼  klik "Pay Now"
POST /sandbox/pay → status berubah PENDING ➜ PAID ✅
```

---

*Selamat mengerjakan Final Project! Kalau tersesat, buka lagi README ini. Semoga sukses dan selamat belajar! 🚀✨*
