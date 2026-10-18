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
