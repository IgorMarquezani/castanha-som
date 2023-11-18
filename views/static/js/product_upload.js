//import('/static/js/utils.js')

const uploadBtn = document.getElementById("upload_btn")
const nameErr = document.getElementById("name_err")
const inCashErr = document.getElementById("in_cash_err")
const installmentsErr = document.getElementById("installments_err")
const typeErr = document.getElementById("type_err")

// this map is meant to be use to keep track of the descriptions provided by the user
// making it possible to avoid repeated fields being created
const descriptionsTrack = new Map();
// this variable is used to create unique id values for each description added
var descriptionsSerial = 0

function triggerUploadBtn() {
  uploadBtn.click()
}

function previewImage() {
  const file = uploadBtn.files[0]
  const fileUrl = URL.createObjectURL(file)

  const preview = document.getElementById("image_preview")
  preview.src = fileUrl
}

function addDescription() {
  const description = document.getElementById("description_input").value
  const descriptionList = document.getElementById("descriptions")

  if (description.length < 1) {
    return
  }

  if (descriptionsTrack.get(description)) {
    return
  }

  const descriptionTemplate = `
                <div class="col col-sm text-start">
                  <p>${description}</p>
                </div>
                <div class="col col-ms text-end">
                  <input type="button" class="btn btn-danger" onclick="deleteDescription(${descriptionsSerial})" value="Deletar">
                </div>
`

  const newDescriptionElement = document.createElement("div")
  newDescriptionElement.className = "row justify-content-md-center"
  newDescriptionElement.id = descriptionsSerial
  newDescriptionElement.innerHTML = descriptionTemplate

  descriptionList.append(newDescriptionElement)

  descriptionsTrack.set(description, description)
  descriptionsSerial++

  const form = document.getElementById("form_2")
  const heigh = form.style.height
  form.style.height = parseInt(heigh) + 40 + "px"
}

function deleteDescription(descriptionId) {
  const element = document.getElementById(descriptionId)
  element.remove()
  descriptionsTrack.delete(element.innerHTML)

  const form = document.getElementById("form_2")
  const heigh = form.style.height
  form.style.height = parseInt(heigh) - 40 + "px"
}

function register() {
  nameErr.innerHTML = ""
  inCashErr.innerHTML = ""
  installmentsErr.innerHTML = ""
  typeErr.innerHTML = ""

  var errors = 0

  const files = uploadBtn.files

  if (files.length < 1) {
    alert("Por favor, selecione uma imagem para o produto")
    errors++
  }

  var file = files[0]

  var name = document.getElementById("name").value
  if (name.length < 1) {
    nameErr.innerHTML = "O nome do produto não pode ser nulo"
    errors++
  }

  var type = document.getElementById("type").value
  if (type.length < 1) {
    typeErr.innerHTML = "O tipo do produto não pode ser nulo"
    errors++
  }

  var inCashValue = document.getElementById("in_cash_value").value
  if (inCashValue.length < 1) {
    inCashErr.innerHTML = "Preencha este campo"
  } else if (!isInt(inCashValue)) {
    inCashErr.innerHTML = "Este campo só aceita digitos númericos"
    errors++
  }

  var installmentsValue = document.getElementById("installments_value").value
  if (installmentsValue.length < 1) {
    installmentsErr.innerHTML = "Preencha este campo"
  } else if (!isInt(installmentsValue)) {
    installmentsErr.innerHTML = "Este campo só aceita digitos númericos"
    errors++
  }

  if (errors > 0) {
    return
  }

  var descriptions = []
  for (let entrie of descriptionsTrack.values()) {
    let description = {
      value: entrie
    }
    descriptions.push(description)
  }

  var body = {
    name: name,
    type: type,
    in_cash_value: parseFloat(inCashValue),
    installment_value: parseFloat(installmentsValue),
    descriptions: descriptions
  }

  var form = new FormData()
  form.set("image", file)
  form.set("data", JSON.stringify(body))

  fetch("/product/register", {
    method: "POST",
    body: form,
  }).then((resp) => {
    if (resp.status == 208) {
      alert("Produto já cadastrado")
    } else if (resp.status == 200 || resp.status == 201) {
      alert("Produto cadastrado com sucesso")
    } else {
      alert("Erro inesperado, pedimos desculpa")
    }
  })
}
