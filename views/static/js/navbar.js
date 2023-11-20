const searchFieldBtn = document.getElementById("search_field_btn")
const searchField = document.getElementById("search_field")

const userName = document.getElementById("user_name")
const userImage = document.getElementById("user_image")

function search(value) {
  if (value.length == 0) {
    return window.location = "/products/list/all"
  }

  window.location = "/products/list/match?m=" + value
}

searchField.addEventListener("keypress", (event) => {
  if (event.key == "Enter") {
    let value = searchField.value
    search(value)
  }
})

searchFieldBtn.addEventListener("click", () => {
  let value = searchField.value
  search(value)
})

function getUserName() {
  fetch("/user/info/personal/my_name", {
    method: "GET"
  }).then((resp) => {
    if (resp.status == 200) {
      return resp.text()
    }
    return ""
  }).then((name) => {
    if (name.length > 0) {
      userName.innerHTML = name
      userName.setAttribute("href", "/user/info/profile/" + name)
    }
  })
}

function getUserImage() {
  fetch("/user/info/personal/my_profile_image", {
    method: "GET"
  }).then((resp) => {
    resp.blob().then((buffer) => {
      const image = URL.createObjectURL(buffer)
      userImage.childNodes[1].remove()
    })
  })
}

getUserName()
