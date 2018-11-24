import * as cfg from "./config.js"
import * as dm from "./domManipulator.js"

// Fetch all transactions from the API servers
export function reloadData() {
  let request = new Request(`${cfg.config.apiBase}/transaction/read`)
  fetch(request).then((resp) => {
    resp.json().then((transactions) => {
      dm.constructTable(transactions)
    }).catch((err) => {
      new Toast("API server sends garbage, contact support", Toast.TYPE_ERROR, 3000)
      console.log("Couldn't convert API data to JSON")
    })
  }).catch((err) => {
    new Toast("API unavailable", Toast.TYPE_ERROR, 3000)
    console.log("Error calling API")
    console.log(err)
  })
}

// Prepare form data and send it to API server
export function submitTransaction(node) {
  let side = node.parentNode.parentNode.className
  let insertRow = document.getElementById(`insert-row-${side}`)
  let tags = []
  insertRow.
    children[2].
    children['input_tags'].
    value.
    trim().
    toLowerCase().
    split(",").
    forEach(tag => {
      tags.push({ name: tag })
    })
  let newData = {
    user_id: side === "left" ? 0 : 1,
    amount: Number(insertRow.children[1].children['input_amount'].value.replace(",", ".")),
    tags: tags,
    note: insertRow.children[3].children['input_note'].value,
  }
  let newDataJSON = JSON.stringify(newData)
  let request = new Request(`${cfg.config.apiBase}/transaction/create`, {
    method: "POST",
    body: newDataJSON,
  })
  fetch(request).then((resp) => {
    if (resp.ok) {
      // Remove the `temporary`-flag from the newly inserted TR
      // or redraw table
      new Toast("Transaction submitted!", Toast.TYPE_DONE, 3000)
    } else {
      // Remove the newly inserted entry to the table (class `temporary`)
      new Toast("Server not satisfied with our request", Toast.TYPE_ERROR, 3000)
    }
  }).catch((err) => {
    new Toast("Could not send transaction to server", Toast.TYPE_ERROR, 3000)
    // Remove the newly inserted entry to the table (class `temporary`)
  }).then(() => {
    reloadData()
    node.value = "+"
    node.setAttribute("onclick", "dm.handlePlusButtonClick(this)")
  })
}

// send API request to delete a Transaction
export function sendDeleteTransaction(transaction_id) {
  console.log(`about to delete transaction ${transaction_id}`)
  let request = new Request(`${cfg.config.apiBase}/transaction/delete/${transaction_id}`, {
    method: "DELETE",
  })
  fetch(request).then((resp) => {
    if (resp.ok) {
      new Toast("Transaction removed!", Toast.TYPE_DONE, 3000)
    } else {
      // Remove the newly inserted entry to the table (class `temporary`)
      new Toast("Server not satisfied with our request", Toast.TYPE_ERROR, 3000)
    }
  }).catch((err) => {
    new Toast("Could not send delete request to server", Toast.TYPE_ERROR, 3000)
  }).then(() => {
    reloadData()
  })
}
