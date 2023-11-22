function addItem() {
  const productName = document.getElementById("product_name").innerHTML

  fetch("/cart/add_item", {
    method: "POST",
    body: productName
  }).then((resp) => {
    console.log("status: ", resp.status)
    if (resp.status == 200 || resp.status == 201) {
      alert("Produto adicionado ao carrinho")
    }
    if (resp.status == 208) {
      alert("Produto já adicionado ao carrino")
    }
  })
}
