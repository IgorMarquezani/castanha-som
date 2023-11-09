//import('/static/js/utils.js')

const uploadBtn = document.getElementById("upload_btn")
const nameErr = document.getElementById("name_err")
const inCashErr = document.getElementById("in_cash_err")
const InstallmentsErr = document.getElementById("installments_err")

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
  InstallmentsErr.innerHTML = ""

  var errors = 0

  const files = uploadBtn.files

  if (files.length < 1) {
    alert("Por favor, selecione uma imagem para o produto")
    return
  }

  var file = files[0]

  var name = document.getElementById("name").value
  if (name.length < 1) {
    nameErr.innerHTML = "O nome produto não pode ser vazio"
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
    InstallmentsErr.innerHTML = "Preencha este campo"
  } else if (!isInt(installmentsValue)) {
    InstallmentsErr.innerHTML = "Este campo só aceita digitos númericos"
    errors++
  }

  if (errors > 0) {
    return
  }

  var descriptions = []
  for (let entrie of descriptionsTrack.values()) {
    descriptions.push(entrie)
  }

  var body = {
    name: "",
    in_cash_value: parseInt(inCashValue),
    installments_value: parseInt(installmentsValue),
    descriptions: descriptions
  }

  var form = new FormData()
  form.set("image", file)
  form.set("data", JSON.stringify(body))
}
