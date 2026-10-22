# TKJ's Store

**Final Project Web Programming — Kelas 10**

Membangun website toko online yang terhubung ke backend API, mulai dari menampilkan produk sampai simulasi pembayaran menggunakan BCA Virtual Account.

![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![REST API](https://img.shields.io/badge/REST_API-0056b3?style=for-the-badge&logo=swagger&logoColor=white)

---

## Tentang Proyek Ini

Di final project ini, kamu akan membuat website toko online bernama **TKJ's Store**. Website ini tidak hanya menampilkan produk, tetapi juga terhubung langsung ke server untuk mengambil data produk, memproses pesanan, dan memeriksa status pembayaran.

Jadi, kamu nggak cuma belajar bikin tampilan website. Kamu juga akan belajar bagaimana frontend berkomunikasi dengan backend menggunakan API.

### Apa saja yang akan kamu buat?

| Fitur | Yang akan dipelajari |
|---|---|
| Katalog produk | `fetch()`, `async/await`, JSON |
| Pagination | `slice()`, `Math.ceil()` |
| Keranjang belanja | Array, event, `localStorage` |
| Modal keranjang dan invoice | DOM manipulation, `classList` |
| Checkout pesanan | Form, `POST`, `preventDefault()` |
| Pembayaran Virtual Account | REST API, query string |
| Simulator pembayaran | Mengambil dan mengirim data ke server |

**Sebelum mulai:** baca panduan ini sambil praktik langsung di VS Code. Jangan buru-buru copy semua kode sekaligus. Coba pahami fungsi setiap bagian, lalu jalankan website untuk melihat hasilnya.

Kalau kamu tertinggal saat mengerjakan project, tenang. Ada panduan langkah demi langkah, troubleshooting, dan kode referensi di bagian akhir README ini.

## Daftar Isi

1. [Alat dan Struktur File](#1-alat-dan-struktur-file)
2. [Backend API dan Swagger](#2-backend-api-dan-swagger)
3. [Cara Menjalankan Proyek](#3-cara-menjalankan-proyek)
4. [Langkah Pengerjaan](#4-langkah-pengerjaan)
5. [Kalau Kamu Tertinggal](#5-kalau-kamu-tertinggal)
6. [Panduan Kustomisasi](#6-panduan-kustomisasi)
7. [Error Umum dan Solusinya](#7-error-umum-dan-solusinya)
8. [Tips Ngoding](#8-tips-ngoding)
9. [Kode Referensi Lengkap](#9-kode-referensi-lengkap)
10. [Alur Kerja Website](#10-alur-kerja-website)

---

## 1. Alat dan Struktur File

### Tools yang dibutuhkan

| Tools | Kegunaan |
|---|---|
| Visual Studio Code | Menulis dan mengedit kode |
| Google Chrome atau Microsoft Edge | Menjalankan website dan memeriksa error |
| Live Server (ekstensi VS Code) | Menjalankan website melalui server lokal dan auto-refresh saat file disimpan |
| Koneksi internet | Mengakses backend API dan mengambil data produk |

### Struktur folder

Pastikan semua file berikut berada di dalam satu folder yang sama.

```text
frontend/
├── index.html     # Halaman utama toko
├── sandbox.html   # Simulator pembayaran
├── style.css      # Tampilan kedua halaman
├── app.js         # Logika website dan komunikasi dengan API
└── README.md      # Panduan pengerjaan
```

**Fungsi setiap file:**

- `index.html`: struktur halaman toko, katalog, keranjang, dan invoice.
- `sandbox.html`: halaman simulator pembayaran.
- `style.css`: mengatur warna, layout, tombol, modal, dan tampilan responsif.
- `app.js`: mengatur pengambilan produk, keranjang, checkout, dan pembayaran.
- `README.md`: panduan pengerjaan dan referensi kode.

Perhatikan bahwa `index.html` dan `sandbox.html` menggunakan file `style.css` dan `app.js` yang sama.

---

## 2. Backend API dan Swagger

Website yang kamu buat adalah **frontend**, yaitu bagian yang berinteraksi langsung dengan pengguna. Sementara itu, **backend** bertugas memproses data dan menyediakan layanan yang dibutuhkan website.

Dalam project ini, data produk, stok, pesanan, dan status pembayaran diperoleh melalui backend API.

| Resource | Link |
|---|---|
| Base URL API | [Buka API Server](https://serverstaging-vps-5e6db201b3.asiahub.id/api) |
| Dokumentasi Swagger | [Buka Swagger UI](https://serverstaging-vps-5e6db201b3.asiahub.id/swagger/index.html) |

### Endpoint yang digunakan

Endpoint adalah alamat yang digunakan website untuk meminta atau mengirim data ke server.

| Method | Endpoint | Fungsi |
|---|---|---|
| `GET` | `/products` | Mengambil daftar produk dan stok |
| `POST` | `/orders/checkout` | Mengirim pesanan dan membuat tagihan |
| `GET` | `/sandbox/va/{va_number}` | Memeriksa detail dan status tagihan |
| `POST` | `/sandbox/pay` | Menjalankan simulasi pembayaran |

Kamu bisa melihat detail setiap endpoint melalui Swagger. Coba perhatikan method, parameter, request body, dan response yang tersedia.

> Catatan: pembayaran dalam project ini menggunakan simulator. Ini bukan transaksi perbankan sungguhan dan tidak terhubung ke rekening BCA untuk memindahkan uang.

Alamat API disimpan dalam variabel `API_BASE` di bagian awal `app.js`. Jika alamat server berubah, sesuaikan variabel tersebut.

---

## 3. Cara Menjalankan Proyek

### Cara A — Menggunakan Live Server (disarankan)

1. Buka folder `frontend/` di VS Code.
2. Buka menu Extensions.
3. Cari ekstensi **Live Server**, lalu install.
4. Klik kanan `index.html`.
5. Pilih **Open with Live Server**.

Website akan terbuka di browser, biasanya melalui alamat `http://127.0.0.1:5500`.

### Cara B — Membuka HTML langsung

Kamu juga bisa membuka `index.html` langsung melalui browser. Namun, jika muncul masalah saat mengakses API, coba jalankan menggunakan Live Server terlebih dahulu.

### Checklist setelah website dibuka

- [ ] Halaman toko tampil dengan layout yang benar.
- [ ] Produk berhasil dimuat dari server.
- [ ] Informasi harga dan stok muncul.
- [ ] Tombol pagination bisa digunakan.
- [ ] Tombol keranjang bisa diklik.

Kalau produk tidak muncul, jangan langsung mengubah semua kode. Periksa koneksi internet, alamat API, dan Console browser terlebih dahulu.

---

## 4. Langkah Pengerjaan

Kerjakan project secara bertahap. Setelah satu bagian selesai, jalankan website dan periksa hasilnya sebelum melanjutkan.

### Langkah 1 — Membuat struktur halaman

File: `index.html`

Buat struktur utama website yang mencakup:

- Navbar dengan logo, menu, dan tombol keranjang.
- Hero section yang berisi judul dan deskripsi toko.
- Katalog untuk menampilkan produk.
- Tombol pagination.
- Modal keranjang dan form checkout.
- Modal invoice untuk menampilkan detail pesanan dan pembayaran.
- Footer.

Pastikan setiap elemen yang akan diakses melalui JavaScript memiliki `id` yang sesuai.

**Sudah selesai?** Buka halaman di browser dan pastikan seluruh bagian utama muncul.

### Langkah 2 — Mengatur tampilan dengan CSS

File: `style.css`

Gunakan CSS untuk membuat tampilan website lebih rapi dan nyaman digunakan.

Beberapa hal yang perlu dibuat:

- Reset dasar CSS dan pengaturan font.
- Navbar dan hero section.
- Grid untuk kartu produk.
- Style tombol dan status stok.
- Modal untuk keranjang dan invoice.
- Tampilan responsif untuk perangkat mobile.
- Layout khusus halaman simulator pembayaran.

Kamu bisa menggunakan CSS Grid dan Flexbox untuk membantu menyusun elemen.

**Sudah selesai?** Coba ubah warna tombol atau background dan lihat apakah tampilannya ikut berubah.

### Langkah 3 — Mengambil produk dari server

File: `app.js`

Buat fungsi `muatProduk()` untuk mengambil daftar produk dari backend.

Hal yang perlu dilakukan:

1. Gunakan `fetch()` untuk mengakses endpoint `/products`.
2. Gunakan `await res.json()` untuk membaca response.
3. Simpan hasilnya ke variabel `daftarProduk`.
4. Panggil fungsi untuk menampilkan katalog.
5. Tambahkan penanganan error menggunakan `try/catch`.

Setiap produk harus menampilkan informasi seperti nama, harga, gambar, dan stok.

Jika stok produk habis, tombol pembelian harus dinonaktifkan.

**Sudah selesai?** Bandingkan jumlah dan informasi produk di website dengan response dari API.

### Langkah 4 — Membuat pagination

Pagination digunakan untuk membagi katalog menjadi beberapa halaman.

Dalam project ini, setiap halaman menampilkan enam produk.

Pelajari penggunaan:

- `ITEMS_PER_PAGE` untuk menentukan jumlah produk per halaman.
- `Math.ceil()` untuk menghitung total halaman.
- `slice()` untuk mengambil produk yang akan ditampilkan.
- `currentPage` untuk menyimpan halaman aktif.

Jangan lupa menonaktifkan tombol Prev pada halaman pertama dan Next pada halaman terakhir.

**Sudah selesai?** Coba berpindah halaman dan pastikan produk yang muncul berubah.

### Langkah 5 — Membuat keranjang belanja

Keranjang memungkinkan pengguna memilih produk sebelum checkout.

Buat fungsi untuk:

- Menambahkan produk ke keranjang.
- Menambah atau mengurangi jumlah item.
- Menghapus item.
- Menghitung total harga dan jumlah barang.
- Menyimpan data menggunakan `localStorage`.

Pastikan jumlah pembelian tidak melebihi stok produk yang tersedia.

Data keranjang disimpan menggunakan key `keranjang_tkj`, sehingga isi keranjang dapat dipertahankan setelah halaman di-refresh.

**Sudah selesai?** Tambahkan produk, ubah jumlahnya, lalu refresh browser. Periksa apakah isi keranjang masih tersedia.

### Langkah 6 — Membuat modal

Modal adalah pop-up yang muncul di atas halaman utama.

Dalam project ini, modal digunakan untuk keranjang belanja dan invoice.

Gunakan `classList.remove("hidden")` untuk menampilkan modal dan `classList.add("hidden")` untuk menyembunyikannya.

Pasang event pada tombol buka, tutup, dan Continue Shopping.

**Sudah selesai?** Pastikan modal bisa dibuka dan ditutup tanpa membuat halaman reload.

### Langkah 7 — Membuat proses checkout

File: `app.js`

Saat pengguna mengirim form checkout, website harus mengirim data pesanan ke backend.

Urutannya:

1. Tangani event `submit` menggunakan `addEventListener()`.
2. Gunakan `e.preventDefault()` agar halaman tidak reload.
3. Validasi nama pembeli dan isi keranjang.
4. Susun data produk menjadi array `items`, yang berisi `product_id` dan `quantity`.
5. Kirim request `POST` ke `/orders/checkout`.
6. Baca response server dan tampilkan invoice jika pesanan berhasil dibuat.

Setelah checkout berhasil, keranjang dikosongkan dan katalog diperbarui.

**Sudah selesai?** Lakukan checkout dan pastikan invoice menampilkan informasi pesanan serta nomor Virtual Account dari server.

### Langkah 8 — Membuat invoice dan simulator pembayaran

File: `index.html`, `sandbox.html`, dan `app.js`

Bagian ini menghubungkan proses checkout dengan simulasi pembayaran.

Pada halaman invoice, tampilkan nomor pesanan, total tagihan, status pembayaran, dan nomor Virtual Account.

Selanjutnya:

1. Buat tombol untuk menyalin nomor VA.
2. Arahkan tombol simulator ke `sandbox.html?va=...`.
3. Baca parameter `va` menggunakan `URLSearchParams`.
4. Ambil detail tagihan melalui endpoint `/sandbox/va/{va_number}`.
5. Kirim request pembayaran ke `/sandbox/pay`.
6. Kembali ke halaman toko dan tekan tombol Cek Status Pembayaran.

Jika server mengembalikan status `PAID`, invoice harus menampilkan status lunas.

**Penting:** status pembayaran diperiksa saat pengguna menekan tombol cek status. Website tidak melakukan polling otomatis.

**Sudah selesai?** Uji alur dari memilih produk sampai invoice menunjukkan pembayaran lunas.

---

## 5. Kalau Kamu Tertinggal

Ketinggalan saat mengerjakan project? Nggak masalah. Jangan buru-buru mengganti seluruh kode dengan jawaban akhir.

Coba ikuti langkah berikut:

1. Cari bagian yang sedang kamu kerjakan di Bagian 4.
2. Periksa file yang berkaitan dengan bagian tersebut.
3. Bandingkan kode milikmu dengan referensi di Bagian 9.
4. Jika masih error, buka Developer Tools dengan menekan `F12`.
5. Baca pesan error di tab Console dan periksa request di tab Network.

### File mana yang harus diperiksa?

| Masalah | File |
|---|---|
| Struktur halaman atau elemen tidak muncul | `index.html` atau `sandbox.html` |
| Warna, ukuran, atau layout berantakan | `style.css` |
| Produk tidak muncul atau tombol tidak berfungsi | `app.js` |
| Checkout atau pembayaran gagal | `app.js` dan response API |

### Checklist sebelum meminta bantuan

- [ ] Semua file berada di folder yang sama.
- [ ] File CSS dan JavaScript sudah dihubungkan ke HTML.
- [ ] Semua perubahan kode sudah disimpan.
- [ ] Koneksi internet berfungsi.
- [ ] Nama `id` di HTML sesuai dengan yang dipanggil di JavaScript.
- [ ] Pesan error di Console sudah diperiksa.

Kalau sudah mencoba semuanya tetapi masih mentok, catat pesan error dan bagian kode yang bermasalah. Itu akan membantu proses debugging jadi lebih cepat.

**Gunakan kode referensi untuk memahami dan memperbaiki pekerjaanmu, bukan sekadar copy-paste tanpa tahu fungsinya.**

---

## 6. Panduan Kustomisasi

Website ini boleh kamu modifikasi supaya punya identitas sendiri. Kamu bisa mengubah warna, font, nama toko, layout, dan jumlah produk per halaman tanpa harus membuat website dari awal.

### Mengubah warna utama

Warna utama bawaan website adalah `#0056b3`.

Untuk menggantinya:

1. Buka `style.css`.
2. Tekan `Ctrl + H` atau `Cmd + H`.
3. Cari kode warna `#0056b3`.
4. Ganti dengan warna pilihanmu.
5. Periksa kembali hasilnya di browser.

Contoh warna alternatif:

| Warna | Hex |
|---|---|
| Ungu | `#7c3aed` |
| Merah | `#e11d48` |
| Hijau | `#15803d` |
| Teal | `#0f766e` |

Perhatikan juga warna hover, background, dan teks agar kontrasnya tetap nyaman dibaca.

### Mengubah font

Buka bagian `body` di `style.css`, lalu ubah properti `font-family`.

```css
body {
  font-family: Arial, Helvetica, sans-serif;
}
```

Kamu juga bisa mencoba font Google Fonts seperti Plus Jakarta Sans atau Poppins, asalkan font tersebut sudah dimuat melalui HTML.

### Mengubah nama toko dan teks

Buka `index.html`, lalu cari teks yang ingin diubah.

Beberapa bagian yang bisa dikustomisasi:

- Judul tab browser.
- Logo toko.
- Judul dan deskripsi hero.
- Judul katalog.
- Teks tombol.
- Footer.

### Mengubah jumlah produk per halaman

Buka `app.js`, kemudian cari:

```javascript
const ITEMS_PER_PAGE = 6;
```

Ganti angka `6` menjadi jumlah produk yang kamu inginkan, misalnya `3` atau `9`.

### Mengubah format mata uang

Website menggunakan format Rupiah dengan locale Indonesia:

```javascript
Number(price).toLocaleString("id-ID");
```

Format tersebut menghasilkan pemisah ribuan yang sesuai dengan kebiasaan penulisan angka di Indonesia.

Jika ingin menggunakan mata uang lain, kamu perlu menyesuaikan simbol mata uang dan locale yang digunakan.

> Tips: setelah mengubah kode, simpan file terlebih dahulu. Jika perubahan belum terlihat, coba hard refresh dengan `Ctrl + Shift + R` atau `Cmd + Shift + R`.

---

## 7. Error Umum dan Solusinya

| Gejala | Kemungkinan penyebab | Solusi |
|---|---|---|
| Produk terus loading | API belum merespons atau JavaScript gagal dijalankan | Periksa Console dan Network |
| Muncul pesan gagal memuat produk | Server tidak dapat diakses atau alamat API salah | Periksa koneksi dan `API_BASE` |
| `Failed to fetch` | Koneksi bermasalah, server tidak tersedia, atau request diblokir | Periksa URL API dan detail error di browser |
| Website tampil tanpa styling | File CSS tidak ditemukan | Periksa `href="style.css"` dan lokasi file |
| Tombol tidak berfungsi | JavaScript tidak berjalan atau `id` tidak sesuai | Periksa Console dan event listener |
| `Cannot read properties of null` | Elemen yang dicari tidak ditemukan | Cocokkan `id` pada HTML dengan JavaScript |
| Halaman reload saat checkout | `preventDefault()` belum digunakan | Periksa event `submit` |
| Keranjang kosong setelah refresh | Data LocalStorage tidak tersimpan atau gagal dibaca | Periksa Application → Local Storage |
| Harga tampil `NaN` | Data harga bukan angka yang valid | Periksa nilai `price` dari API |
| Nomor VA tidak muncul | Response checkout tidak sesuai harapan | Periksa response endpoint checkout |
| Tagihan tidak ditemukan | Nomor VA salah atau pesanan belum dibuat | Periksa nomor VA dan response API |
| Invoice masih PENDING | Status belum diperiksa kembali | Klik Cek Status Pembayaran setelah simulasi pembayaran |

### Cara membaca error dengan F12

- **Console:** menampilkan error JavaScript beserta lokasi masalahnya.
- **Network:** menampilkan request dan response dari server. Klik request yang bermasalah untuk memeriksa status dan response-nya.
- **Application → Local Storage:** menampilkan data keranjang yang disimpan di browser.

Status HTTP `200` umumnya menunjukkan request berhasil. Status `4xx` berkaitan dengan request atau kondisi yang tidak dapat dipenuhi, sedangkan `5xx` menunjukkan masalah di sisi server. Periksa response untuk mengetahui penyebab yang lebih spesifik.

---

## 8. Tips Ngoding

Beberapa kebiasaan sederhana yang bakal membantu selama mengerjakan project:

1. **Biasakan membaca error.** Jangan langsung menghapus atau mengganti kode yang bermasalah.
2. **Pahami `fetch()` dan `async/await`.** Keduanya penting untuk mengambil data dari API.
3. **Periksa response server.** Jangan menebak struktur JSON; lihat data yang benar-benar dikirim backend.
4. **Simpan pekerjaan secara berkala.** Gunakan `Ctrl + S` atau `Cmd + S`.
5. **Uji satu fitur pada satu waktu.** Pastikan fitur yang sedang dikerjakan berfungsi sebelum lanjut.
6. **Gunakan Developer Tools.** Console, Network, dan Local Storage adalah temanmu saat debugging.
7. **Jangan takut bereksperimen.** Setelah fitur utama berfungsi, coba ubah tampilan atau tambahkan fitur kecil.

Yang penting bukan sekadar websitenya berhasil dibuka, tetapi kamu juga tahu bagaimana kode di dalamnya bekerja.

---

## 9. Kode Referensi Lengkap

Bagian ini berisi implementasi referensi untuk membantu kamu membandingkan kode sendiri dengan versi yang disediakan.

Gunakan referensi sesuai kebutuhan. Kalau hanya ada masalah pada CSS, kamu tidak perlu mengganti seluruh file JavaScript.

**Isi bagian ini dengan kode lengkap yang disediakan bersama project:**

- `index.html` — struktur halaman toko.
- `sandbox.html` — struktur simulator pembayaran.
- `style.css` — seluruh styling website.
- `app.js` — logika frontend dan komunikasi dengan backend API.

```html
<details>
  <summary><strong>index.html</strong></summary>
    <!-- <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>TKJ's Store</title>
      <link rel="preconnect" href="https://fonts.googleapis.com">
      <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
      <link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&family=JetBrains+Mono:wght@600;700&display=swap" rel="stylesheet">
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
            <a href="sandbox.html" target="_blank" class="nav-link nav-sandbox-link">💳 Simulator Bayar</a>
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
              <p class="inst-title">Petunjuk Pembayaran BCA Virtual Account:</p>
              <ol>
                <li>Salin <strong>Nomor Virtual Account</strong> di atas.</li>
                <li>Buka <strong>Simulator Pembayaran (m-Banking BCA)</strong> di tab terpisah.</li>
                <li>Periksa detail tagihan Anda dan klik <strong>Bayar Sekarang</strong>.</li>
                <li>Kembali ke halaman ini, lalu klik <strong>Cek Status Pembayaran</strong>.</li>
              </ol>
            </div>

            <div id="sandbox-action-box" class="simulator-action-box">
              <a id="btn-open-simulator" href="sandbox.html" target="_blank" class="btn btn-simulator btn-block">
                🏦 Buka Simulator Pembayaran (BCA) ↗
              </a>
              <button id="btn-check-payment" type="button" class="btn btn-secondary btn-block" style="margin-top: 8px;">
                🔄 Cek Status Pembayaran
              </button>
            </div>

            <div id="payment-success-box" class="success-box hidden">
              <div class="success-icon">🎉</div>
              <h4>Pembayaran Lunas & Dikonfirmasi!</h4>
              <p>Terima kasih! Pembayaran Virtual Account Anda telah terverifikasi oleh server. Pesanan kini sedang diproses.</p>
              <div id="payment-confirmed-time" class="confirmed-time"></div>
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
    </html> -->

</details>
```

<details>
  <summary><strong>sandbox.html</strong></summary>

  <!-- 
      <!DOCTYPE html>
    <html lang="id">
    <head>
      <meta charset="UTF-8" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Simulator Pembayaran BCA</title>
      <link rel="stylesheet" href="style.css" />
    </head>
    <body class="sandbox-page">
      <div class="sandbox-card">
        <h2 class="text-primary" style="margin-bottom: 14px;">Simulator BCA</h2>
        
        <div class="form-group">
          <label for="va-input">Nomor Virtual Account:</label>
          <input type="text" id="va-input" placeholder="88001xxxx" />
        </div>
        <button class="btn btn-primary btn-block" onclick="cekTagihan()">Cek Tagihan</button>

        <div id="bill-card" class="invoice-details hidden" style="margin-top: 16px;">
          <div class="detail-row"><span>Order:</span><strong id="info-order">-</strong></div>
          <div class="detail-row"><span>Total:</span><strong id="info-total" class="text-primary">Rp 0</strong></div>
          <div class="detail-row"><span>Status:</span><strong id="info-status">PENDING</strong></div>
          <button id="btn-pay" class="btn btn-simulator btn-block" style="margin-top: 10px;" onclick="bayarTagihan()">Bayar Sekarang</button>
        </div>

        <p style="text-align: center; margin-top: 16px;">
          <a href="index.html" class="nav-link">← Kembali ke Toko Online</a>
        </p>
      </div>

      <script src="app.js"></script>
    </body>
    </html>
  -->

</details>

<details>
  <summary><strong>app.js</strong></summary>

  <!--
    const API_BASE = "https://serverstaging-vps-5e6db201b3.asiahub.id/api";
    const ITEMS_PER_PAGE = 6;

    let currentPage = 1;
    let daftarProduk = [];
    let keranjang = JSON.parse(localStorage.getItem("keranjang_tkj")) || [];
    let pesananAktif = null;

    async function muatProduk() {
      try {
        const res = await fetch(API_BASE + "/products");
        daftarProduk = await res.json();
        tampilkanKatalog();
      } catch (err) {
        document.getElementById("product-grid").innerHTML = "<p>Gagal memuat produk dari server.</p>";
      }
    }

    function tampilkanKatalog() {
      const grid = document.getElementById("product-grid");
      const totalPages = Math.ceil(daftarProduk.length / ITEMS_PER_PAGE) || 1;
      const start = (currentPage - 1) * ITEMS_PER_PAGE;
      const produkTampil = daftarProduk.slice(start, start + ITEMS_PER_PAGE);

      grid.innerHTML = produkTampil.map(p => {
        const isHabis = p.stock <= 0;
        return `
          <div class="product-card">
            <div class="product-img-box">
              <img src="${p.image_url}" alt="${p.name}" onerror="this.src='https://via.placeholder.com/240x160?text=Produk'" />
              <span class="stock-badge ${isHabis ? 'out-of-stock' : 'in-stock'}">
                ${isHabis ? 'Stok Habis' : 'Stok: ' + p.stock}
              </span>
            </div>
            <div class="product-details">
              <h3 class="product-title">${p.name}</h3>
              <p class="product-desc">${p.description || ""}</p>
              <div class="product-meta">
                <div class="product-price">Rp ${Number(p.price).toLocaleString("id-ID")}</div>
                <div class="product-stock-text">Stok: <strong>${p.stock}</strong></div>
              </div>
              <button class="btn ${isHabis ? 'btn-disabled' : 'btn-primary'} btn-block" 
                      ${isHabis ? 'disabled' : ''} 
                      onclick="tambahKeKeranjang('${p.id}')">
                ${isHabis ? 'Stok Habis' : '+ Add to Cart'}
              </button>
            </div>
          </div>
        `;
      }).join("");

      document.getElementById("page-info").textContent = `Page ${currentPage} of ${totalPages}`;
      document.getElementById("btn-prev-page").disabled = (currentPage === 1);
      document.getElementById("btn-next-page").disabled = (currentPage === totalPages);
    }

    function tambahKeKeranjang(id) {
      const produk = daftarProduk.find(p => p.id === id);
      if (!produk || produk.stock <= 0) return alert("Stok produk ini sedang habis!");

      const item = keranjang.find(i => i.product_id === id);
      const qtySekarang = item ? item.quantity : 0;

      if (qtySekarang + 1 > produk.stock) {
        return alert(`Stok tidak cukup! Hanya tersisa ${produk.stock} pcs.`);
      }

      if (item) {
        item.quantity++;
      } else {
        keranjang.push({
          product_id: produk.id,
          name: produk.name,
          price: produk.price,
          image_url: produk.image_url,
          quantity: 1
        });
      }

      simpanKeranjang();
      alert(`"${produk.name}" dimasukkan ke keranjang!`);
    }

    function ubahJumlahItem(id, perubahan) {
      const item = keranjang.find(i => i.product_id === id);
      if (!item) return;

      if (perubahan > 0) {
        const produk = daftarProduk.find(p => p.id === id);
        if (produk && item.quantity + perubahan > produk.stock) {
          return alert(`Maksimal pembelian adalah ${produk.stock} pcs.`);
        }
      }

      item.quantity += perubahan;
      if (item.quantity <= 0) {
        keranjang = keranjang.filter(i => i.product_id !== id);
      }

      simpanKeranjang();
    }

    function hapusItemDariKeranjang(id) {
      keranjang = keranjang.filter(i => i.product_id !== id);
      simpanKeranjang();
    }

    function simpanKeranjang() {
      localStorage.setItem("keranjang_tkj", JSON.stringify(keranjang));
      perbaruiTampilanKeranjang();
    }

    function perbaruiTampilanKeranjang() {
      const cartList = document.getElementById("cart-items-list");
      if (!cartList) return;

      let totalQty = 0;
      let totalHarga = 0;

      keranjang.forEach(item => {
        totalQty += item.quantity;
        totalHarga += item.price * item.quantity;
      });

      document.getElementById("cart-counter").textContent = totalQty;
      document.getElementById("cart-grand-total").textContent = "Rp " + totalHarga.toLocaleString("id-ID");

      if (keranjang.length === 0) {
        cartList.innerHTML = "<p style='text-align: center; color: #888;'>Keranjang masih kosong.</p>";
        return;
      }

      cartList.innerHTML = keranjang.map(item => `
        <div class="cart-item-row">
          <img class="cart-item-thumb" src="${item.image_url || 'https://via.placeholder.com/48'}" alt="${item.name}" />
          <div class="cart-item-info">
            <div class="cart-item-name">${item.name}</div>
            <div class="cart-item-price">Rp ${Number(item.price).toLocaleString("id-ID")}</div>
          </div>
          <div class="cart-item-qty">
            <button class="btn-qty" onclick="ubahJumlahItem('${item.product_id}', -1)">-</button>
            <span>${item.quantity}</span>
            <button class="btn-qty" onclick="ubahJumlahItem('${item.product_id}', 1)">+</button>
            <button class="btn-delete-item" onclick="hapusItemDariKeranjang('${item.product_id}')">Hapus</button>
          </div>
        </div>
      `).join("");
    }

    function tampilkanInvoice(order) {
      document.getElementById("invoice-order-number").textContent = order.order_number;
      document.getElementById("invoice-buyer-name").textContent = order.buyer_name || "Pelanggan";
      document.getElementById("invoice-va-number").textContent = order.payment.va_number;
      document.getElementById("invoice-total-amount").textContent = "Rp " + Number(order.total_amount).toLocaleString("id-ID");

      const badge = document.getElementById("invoice-status");
      badge.className = "badge badge-pending";
      badge.textContent = "PENDING";

      document.getElementById("sandbox-action-box").style.display = "block";
      document.getElementById("payment-success-box").classList.add("hidden");

      document.getElementById("btn-open-simulator").href = "sandbox.html?va=" + order.payment.va_number;

      document.getElementById("invoice-modal").classList.remove("hidden");
    }

    async function cekStatusPembayaran() {
      if (!pesananAktif) return;
      const va = pesananAktif.payment.va_number;

      try {
        const res = await fetch(`${API_BASE}/sandbox/va/${va}`);
        const data = await res.json();

        if (data.status === "PAID") {
          const badge = document.getElementById("invoice-status");
          badge.className = "badge badge-paid";
          badge.textContent = "PAID (LUNAS)";

          document.getElementById("sandbox-action-box").style.display = "none";
          document.getElementById("payment-success-box").classList.remove("hidden");

          muatProduk();
          alert("Pembayaran Terverifikasi Lunas!");
        } else {
          alert("Status masih PENDING. Silakan bayar di Simulator BCA.");
        }
      } catch (err) {
        alert("Gagal memeriksa status: " + err.message);
      }
    }

    function bukaModalKeranjang() {
      perbaruiTampilanKeranjang();
      document.getElementById("cart-modal").classList.remove("hidden");
    }
    function tutupModalKeranjang() {
      document.getElementById("cart-modal").classList.add("hidden");
    }
    function tutupModalInvoice() {
      document.getElementById("invoice-modal").classList.add("hidden");
      muatProduk();
    }

    const checkoutForm = document.getElementById("checkout-form");
    if (checkoutForm) {
      checkoutForm.addEventListener("submit", async function(e) {
        e.preventDefault();
        const nama = document.getElementById("buyer-name").value.trim();
        if (!nama) return alert("Silakan isi nama Anda!");
        if (keranjang.length === 0) return alert("Keranjang belanja kosong!");

        const items = keranjang.map(i => ({ product_id: i.product_id, quantity: i.quantity }));

        try {
          const res = await fetch(API_BASE + "/orders/checkout", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ items })
          });
          const data = await res.json();

          pesananAktif = data;
          pesananAktif.buyer_name = nama;

          document.getElementById("buyer-name").value = "";
          keranjang = [];
          simpanKeranjang();
          tutupModalKeranjang();
          muatProduk();

          tampilkanInvoice(data);
        } catch (err) {
          alert("Checkout gagal: " + err.message);
        }
      });

      document.getElementById("btn-open-cart").onclick = bukaModalKeranjang;
      document.getElementById("btn-close-cart").onclick = tutupModalKeranjang;
      document.getElementById("btn-continue-shopping").onclick = tutupModalKeranjang;
      document.getElementById("btn-close-invoice").onclick = tutupModalInvoice;
      document.getElementById("btn-finish-invoice").onclick = tutupModalInvoice;
      document.getElementById("btn-check-payment").onclick = cekStatusPembayaran;

      document.getElementById("btn-copy-va").onclick = function() {
        const va = document.getElementById("invoice-va-number").textContent;
        if (va && va !== "-") {
          navigator.clipboard.writeText(va);
          alert("Nomor VA disalin: " + va);
        }
      };

      document.getElementById("btn-prev-page").onclick = function() {
        if (currentPage > 1) {
          currentPage--;
          tampilkanKatalog();
        }
      };
      document.getElementById("btn-next-page").onclick = function() {
        const totalPages = Math.ceil(daftarProduk.length / ITEMS_PER_PAGE);
        if (currentPage < totalPages) {
          currentPage++;
          tampilkanKatalog();
        }
      };

      const hamburgerBtn = document.getElementById("hamburger-btn");
      const navMenu = document.getElementById("nav-menu");
      if (hamburgerBtn) {
        hamburgerBtn.onclick = function() {
          navMenu.classList.toggle("active");
        };
      }

      muatProduk();
      perbaruiTampilanKeranjang();
    }

    async function cekTagihan() {
      const va = document.getElementById("va-input").value.trim();
      const res = await fetch(`${API_BASE}/sandbox/va/${va}`);
      if (!res.ok) return alert("Nomor VA tidak ditemukan!");
      const data = await res.json();
      document.getElementById("info-order").textContent = data.order_number;
      document.getElementById("info-total").textContent = "Rp " + Number(data.amount).toLocaleString("id-ID");
      document.getElementById("info-status").textContent = data.status;
      document.getElementById("btn-pay").style.display = data.status === "PAID" ? "none" : "block";
      document.getElementById("bill-card").classList.remove("hidden");
    }

    async function bayarTagihan() {
      const va = document.getElementById("va-input").value.trim();
      const res = await fetch(`${API_BASE}/sandbox/pay`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ va_number: va })
      });
      if (!res.ok) return alert("Gagal memproses pembayaran!");
      alert("Pembayaran Berhasil!");
      cekTagihan();
    }

    const vaInput = document.getElementById("va-input");
    if (vaInput) {
      const params = new URLSearchParams(window.location.search);
      if (params.get("va")) {
        vaInput.value = params.get("va");
        cekTagihan();
      }
    }

    window.muatProdukDariServer = muatProduk;
    window.tampilkanKatalogDenganPagination = tampilkanKatalog;

  -->

</details>
<details>
  <summary><strong>style.css</strong></summary>

  <!--
  * {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  }

  body {
    font-family: Arial, Helvetica, sans-serif;
    background-color: #f8fafc;
    color: #1e293b;
    line-height: 1.6;
  }

  .container {
    max-width: 1050px;
    margin: 0 auto;
    padding: 0 20px;
  }

  .hidden {
    display: none !important;
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
    color: #475569;
  }

  .btn-simulator {
    background: #003893;
    color: #ffffff;
  }

  .btn-block {
    width: 100%;
  }

  .btn-disabled, .btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .navbar {
    background-color: #ffffff;
    border-bottom: 1px solid #e2e8f0;
    position: sticky;
    top: 0;
    z-index: 100;
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
    color: #475569;
    font-weight: bold;
    font-size: 0.95rem;
  }

  .nav-sandbox-link {
    background: #eff6ff;
    color: #1d4ed8;
    padding: 6px 12px;
    border-radius: 20px;
    font-size: 0.85rem;
    border: 1px solid #bfdbfe;
  }

  .btn-cart {
    background-color: #ebf5ff;
    border: 1px solid #bce0fd;
    color: #0056b3;
    padding: 8px 16px;
    border-radius: 20px;
    font-weight: bold;
    cursor: pointer;
  }

  .hamburger-btn {
    display: none;
    flex-direction: column;
    gap: 4px;
    background: none;
    border: none;
    cursor: pointer;
  }

  .hamburger-btn span {
    width: 24px;
    height: 3px;
    background-color: #1e293b;
  }

  .hero-section {
    padding: 60px 0;
    background: linear-gradient(180deg, #ffffff 0%, #f0f6ff 100%);
    border-bottom: 1px solid #e2e8f0;
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
    margin-bottom: 12px;
  }

  .text-primary {
    color: #0056b3;
  }

  .hero-text p {
    color: #64748b;
    max-width: 600px;
    margin: 0 auto 20px;
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
    margin-bottom: 6px;
  }

  .section-header p {
    color: #64748b;
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
    display: flex;
    flex-direction: column;
  }

  .product-img-box {
    width: 100%;
    height: 160px;
    background-color: #edf2f7;
    position: relative;
  }

  .product-img-box img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .stock-badge {
    position: absolute;
    top: 8px;
    right: 8px;
    padding: 3px 8px;
    border-radius: 12px;
    font-size: 0.75rem;
    font-weight: bold;
  }

  .stock-badge.in-stock {
    background: #dcfce7;
    color: #166534;
  }

  .stock-badge.out-of-stock {
    background: #fee2e2;
    color: #991b1b;
  }

  .product-details {
    padding: 16px;
    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  .product-title {
    font-size: 1rem;
    margin-bottom: 6px;
  }

  .product-desc {
    font-size: 0.82rem;
    color: #64748b;
    margin-bottom: 12px;
    flex-grow: 1;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .product-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
  }

  .product-price {
    font-size: 1.1rem;
    font-weight: bold;
    color: #0056b3;
  }

  .product-stock-text {
    font-size: 0.8rem;
    color: #64748b;
  }

  .pagination-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
  }

  .page-info {
    font-weight: bold;
    color: #64748b;
  }

  .modal-backdrop {
    position: fixed;
    inset: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    padding: 16px;
  }

  .modal-box {
    background-color: #ffffff;
    width: 100%;
    max-width: 500px;
    border-radius: 8px;
    max-height: 90vh;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
  }

  .modal-header {
    padding: 16px 20px;
    border-bottom: 1px solid #e2e8f0;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .btn-close {
    background: none;
    border: none;
    font-size: 1.25rem;
    cursor: pointer;
    color: #64748b;
  }

  .modal-body {
    padding: 20px;
  }

  .modal-footer {
    padding: 12px 20px;
    border-top: 1px solid #e2e8f0;
    display: flex;
    justify-content: flex-end;
  }

  .cart-items-list {
    display: flex;
    flex-direction: column;
    gap: 10px;
    margin-bottom: 16px;
    max-height: 200px;
    overflow-y: auto;
  }

  .cart-item-row {
    display: flex;
    align-items: center;
    gap: 10px;
    border-bottom: 1px solid #edf2f7;
    padding-bottom: 8px;
  }

  .cart-item-thumb {
    width: 44px;
    height: 44px;
    object-fit: cover;
    border-radius: 4px;
  }

  .cart-item-info {
    flex-grow: 1;
    font-size: 0.85rem;
  }

  .cart-item-name {
    font-weight: bold;
  }

  .cart-item-price {
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
    background: white;
    cursor: pointer;
    border-radius: 3px;
    font-weight: bold;
  }

  .btn-delete-item {
    background: none;
    border: none;
    color: #dc2626;
    font-size: 0.8rem;
    cursor: pointer;
    margin-left: 6px;
  }

  .cart-total-box {
    display: flex;
    justify-content: space-between;
    padding: 12px;
    background-color: #f8fafc;
    border: 1px solid #e2e8f0;
    border-radius: 6px;
    margin-bottom: 16px;
    font-weight: bold;
  }

  .total-price {
    color: #0056b3;
    font-size: 1.15rem;
  }

  .checkout-form-container {
    border-top: 1px solid #e2e8f0;
    padding-top: 14px;
  }

  .form-group {
    margin-bottom: 12px;
  }

  .form-group label {
    display: block;
    font-size: 0.85rem;
    font-weight: bold;
    margin-bottom: 4px;
  }

  .form-group input {
    width: 100%;
    padding: 8px 12px;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
  }

  .status-banner {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 12px;
    background-color: #fef3c7;
    border: 1px solid #fde68a;
    border-radius: 6px;
    margin-bottom: 14px;
  }

  .badge {
    padding: 3px 8px;
    border-radius: 12px;
    font-size: 0.75rem;
    font-weight: bold;
  }

  .badge-pending {
    background-color: #d97706;
    color: white;
  }

  .badge-paid {
    background-color: #16a34a;
    color: white;
  }

  .invoice-details {
    background-color: #f8fafc;
    border: 1px solid #e2e8f0;
    border-radius: 6px;
    padding: 12px;
    margin-bottom: 14px;
    font-size: 0.88rem;
  }

  .detail-row {
    display: flex;
    justify-content: space-between;
    margin-bottom: 8px;
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
    font-family: monospace;
    font-size: 1rem;
  }

  .btn-copy {
    background-color: white;
    border: 1px solid #cbd5e1;
    padding: 2px 8px;
    border-radius: 4px;
    font-size: 0.75rem;
    cursor: pointer;
  }

  .payment-instructions {
    background-color: #f8fafc;
    border: 1px solid #e2e8f0;
    border-radius: 6px;
    padding: 12px;
    font-size: 0.82rem;
    margin-bottom: 14px;
  }

  .payment-instructions ol {
    padding-left: 20px;
  }

  .simulator-action-box {
    background-color: #f8fafc;
    border: 1px solid #e2e8f0;
    border-radius: 8px;
    padding: 14px;
    margin-bottom: 14px;
  }

  .success-box {
    background-color: #ecfdf5;
    border: 1px solid #a7f3d0;
    padding: 16px;
    border-radius: 8px;
    text-align: center;
  }

  .success-icon {
    font-size: 2rem;
    margin-bottom: 4px;
  }

  .footer {
    text-align: center;
    padding: 24px 0;
    background-color: #ffffff;
    border-top: 1px solid #e2e8f0;
    color: #64748b;
    font-size: 0.85rem;
  }

  .empty-state {
    text-align: center;
    padding: 30px;
    color: #64748b;
    grid-column: 1 / -1;
  }

  @media (max-width: 768px) {
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
      flex-direction: column;
      padding: 16px 20px;
      gap: 12px;
    }

    .nav-menu.active {
      display: flex;
    }
  }

  .sandbox-page {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 20px;
  }

  .sandbox-card {
    background: #ffffff;
    width: 100%;
    max-width: 400px;
    padding: 24px;
    border-radius: 8px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.06);
  }
  -->

</details>

---

## 10. Alur Kerja Website

Berikut gambaran singkat proses yang terjadi dari awal sampai pembayaran selesai.

```text
Pembeli membuka website
        |
        v
Frontend meminta daftar produk
(GET /products)
        |
        v
Produk ditampilkan di katalog
        |
        v
Pembeli memilih produk
        |
        v
Keranjang disimpan di LocalStorage
        |
        v
Pembeli melakukan checkout
(POST /orders/checkout)
        |
        v
Server membuat pesanan dan tagihan VA
        |
        v
Invoice ditampilkan
        |
        v
Pembeli membuka simulator pembayaran
        |
        v
Simulator mengambil detail tagihan
(GET /sandbox/va/{va_number})
        |
        v
Pembeli menjalankan simulasi pembayaran
(POST /sandbox/pay)
        |
        v
Pembeli kembali ke halaman toko
        |
        v
Klik Cek Status Pembayaran
        |
        v
Server mengembalikan status PAID
        |
        v
Invoice menampilkan status lunas
```

Dari alur ini, kamu bisa melihat bahwa frontend tidak memproses seluruh transaksi sendiri. Frontend mengirim request, menerima response, dan menampilkan hasilnya kepada pengguna. Data pesanan dan status pembayaran diproses oleh backend.

---

<div align="center">

**Selamat mengerjakan final project!**

Kerjakan pelan-pelan, pahami setiap langkah, dan jangan takut mencoba.

*Happy coding!*

</div>
