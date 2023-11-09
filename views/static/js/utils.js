function isInt(data) {
  for (let i = 0; i < data.length; i++) {
    let code = data.charCodeAt(i)
    if (code < '0'.charCodeAt() || code > '9'.charCodeAt()) {
      return false
    }
  }

  return true
}
