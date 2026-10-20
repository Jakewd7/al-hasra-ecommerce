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
