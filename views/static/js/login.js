function login() {
  var email = document.getElementById("email").value
  var passwd = document.getElementById("password").value
  var err = document.getElementById("err")

  var form = new FormData()

  form.set("email", email)
  form.set("password", passwd)

  fetch("/user/login", {
    method: "POST",
    body: form,
  }).then((resp) => {
    if (resp.status == 200) {
      window.location.href = "/static/html/home.html"
    } else if (resp.status == 401) {
      err.innerHTML = "Credências invalidas"
    } else {
      alert("Erro inesperado =(, pedimos desculpa")
    }
  })
}
