function length_errors(name, nameErr, passwd, passwdErr) {
  var errors = 0

  if (name.length == 0) {
    nameErr.innerHTML = "Por favor digite seu nome"
    errors++
  }

  if (passwd.length < 8) {
    passwdErr.style.display = ""
    passwdErr.innerHTML = "A senha deve ter pelo menos 8 digitos"
    errors++
  } 
  
  return errors
}

function passwdErrors (passwd) {
  var noUppercase = true 
  var noNumeric = true
  var noEspecial = true
  var errors = 0

  for (i = 0; i < passwd.length; i++) {
    let charCode = passwd.charCodeAt(i)

    if (charCode >= "A".charCodeAt() && charCode <= "Z".charCodeAt()) {
      noUppercase = false
    } else if (charCode >= "1".charCodeAt() && charCode <= "9".charCodeAt()) {
      noNumeric = false
    } else if (charCode >= "!".charCodeAt() && charCode <= "/") {
      noEspecial = false
    } else if (charCode >= ":".charCodeAt() && charCode <= "@".charCodeAt()) {
      noEspecial = false
    } else if (charCode >= "[".charCodeAt() && charCode <= "`".charCodeAt()) {
      noEspecial = false
    } else if (charCode >= "{".charCodeAt()) {
      noEspecial = true
    }
  }

  if (noUppercase) {
    let uppercaseErr = document.getElementById("uppercase_err")
    uppercaseErr.innerHTML = "Deve conter um caracter maiúsculo"
    uppercaseErr.style.display = ""
    errors++
  }

  if (noNumeric) {
    let numericErr = document.getElementById("numeric_err")
    numericErr.innerHTML = "Deve conter um caracter númerico"
    numericErr.style.display = ""
    errors++
  }

  if (noEspecial) {
    let especialErr = document.getElementById("especial_err")
    especialErr.innerHTML = "Deve conter um caracter especial"
    especialErr.style.display = ""
    errors++
  }

  return errors
}

function register() {
  const nameErr = document.getElementById("name_err")
  const emailErr = document.getElementById("email_err")
  const passwdLengthErr = document.getElementById("length_err")
  const passwdUppercaseErr = document.getElementById("uppercase_err")
  const passwdNumericErr = document.getElementById("numeric_err")
  const passwdEspecialErr = document.getElementById("especial_err")
  const confirmationErr = document.getElementById("confirmation_err")

  nameErr.innerHTML = ""
  emailErr.innerHTML = ""

  passwdLengthErr.innerHTML = ""
  passwdLengthErr.style.display = "none"

  passwdUppercaseErr.innerHTML = ""
  passwdUppercaseErr.style.display 

  passwdNumericErr.innerHTML = ""
  passwdNumericErr.style.display 

  passwdEspecialErr.innerHTML = ""
  passwdEspecialErr.style.display 

  confirmationErr.innerHTML = ""

  const name = document.getElementById("name").value
  const email = document.getElementById("email").value
  const passwd = document.getElementById("password").value
  const confirmation = document.getElementById("confirmation").value

  var errors = 0

  errors += length_errors(name, nameErr, passwd, passwdLengthErr)

  const re = /^[\w]+@[A-z]+\.com/

  if (!re.test(email)) {
    emailErr.innerHTML = "E-mail inválido"
    errors++
  }

  if (passwd != confirmation) {
    confirmationErr.innerHTML = "Senhas incompátiveis"
    errors++
  }

  errors += passwdErrors(passwd)
  console.log(errors)

  if (errors > 0 ) {
    return
  }

  var form = new FormData()
  
  form.set("name", name)
  form.set("email", email)
  form.set("password", passwd)
  form.set("confirmation", confirmation)

  fetch("/user/register", {
    method: "POST",
    body: form,
  }).then((resp) => {
    if (resp.status == 200 || resp.status == 201) {
      console.log("ok")
      alert("Usuário cadastrado com sucesso. Obrigado por usar nosso sistema")
      window.location.href = "/static/html/login.html"
    } 

    if (resp.status == 208) {
      emailErr.innerHTML = "Email está sendo usado"
      return
    }
  })
}
