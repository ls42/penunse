import * as cfg from "./config.js"
import * as ev from "./events.js"

// Insert a new tr in both tables to insert a new transaction into the database.
export function handleAddButtonClick(node) {
  // Add input row to table
  let inputId, newInput, newNameSelect
  let body = document.getElementById(`body`)
  let insertRow = body.insertRow(0)
  insertRow.id = `insert-row`
  cfg.config.headers.forEach((header) => {
    let newTd = document.createElement("td")
    switch (header) {
      case "Date":
        newTd.appendChild(document.createTextNode(""))
        insertRow.appendChild(newTd)
        break
      case "Action":
        newTd.appendChild(document.createTextNode(""))
        insertRow.appendChild(newTd)
        break
      case "User":
        inputId = `input_${header.toLowerCase()}`
        newNameSelect = document.createElement("select")
        newNameSelect.id = inputId
        newNameSelect.className = "transaction-input"
        cfg.config.users.forEach((user) => {
          console.log(user.name)
          var option = document.createElement("option")
          option.value = user.id
          option.text = user.name
          newNameSelect.appendChild(option)
        })
        newTd.appendChild(newNameSelect)
        insertRow.appendChild(newTd)
        break
      default:
        inputId = `input_${header.toLowerCase()}`
        newInput = document.createElement("input")
        newInput.id = inputId
        newInput.type = "text"
        newInput.className = "transaction-input"
        newTd.appendChild(newInput)
        insertRow.appendChild(newTd)
        break
    }
  })

  // Change `add`-button to submit
  node.value = "save"
  node.setAttribute("onclick", "sc.submitTransaction(this)")
}

// Create a table based on `transactions` (which is an array of objects)
// Usually gets called from `reloadData()` in `modules/serverConnector.js`
export function constructTable(transactions) {
  let table = document.getElementById("table")
  while (table.firstChild) {
    table.removeChild(table.firstChild)
  }

  // Generate table header
  let tableHead = document.createElement("thead")
  let tr = document.createElement("tr")
  cfg.config.headers.forEach((e) => {
    let cell = document.createElement("th")
    if (e !== "Action") {
      cell.appendChild(document.createTextNode(e))
    } else {
      cell.appendChild(document.createTextNode(""))
    }
    tr.appendChild(cell)
  })
  tableHead.appendChild(tr)
  table.appendChild(tableHead)

  // Generate table body
  let tableTotal = 0
  let tableBody = document.createElement("tbody")
  tableBody.id = "body"
  transactions.forEach((e) => {
    let tr = document.createElement("tr")
    tr.id = e.id
    tr.addEventListener("mouseenter", ev.mouseEnterTransactionTR)
    tr.addEventListener("mouseleave", ev.mouseOutTransactionTR)
    tr.className = "transaction-data-row"
    cfg.config.headers.forEach((header) => {
      let cell = document.createElement("td")
      switch (header) {
        case "Date":
          let createdDate = new Date(e.created).toLocaleString()
          let dateCell = document.createTextNode(createdDate)
          cell.className = "date"
          cell.appendChild(dateCell)
          break
        case "Amount":
          cell.appendChild(document.createTextNode(e.amount.toFixed(2)))
          break
        case "Tags":
          if (e.tags !== null) {
            e.tags.forEach(function (tag) {
              let tagSpan = document.createElement("span")
              tagSpan.className = "tag"
              tagSpan.appendChild(document.createTextNode(tag.name))
              cell.appendChild(tagSpan)
            })
          }
          break
        case "Note":
          let noteCell
          if (e.note.length > 30) {
            noteCell = document.createTextNode(e.note.substr(0, 30) + " [...]")
          } else {
            noteCell = document.createTextNode(e.note)
          }
          cell.className = "note-td"
          cell.title = e.note
          cell.appendChild(noteCell)
          break
        case "User":
          let userCell, username
          cfg.config.users.forEach((user) => {
            if (user.id === e.user_id) {
              userCell = document.createTextNode(user.name)
            }
          })
          cell.className = "user-td"
          cell.title = e.user_id
          cell.appendChild(userCell)
          break
        case "Action":
          // This is the target for the action buttons
          let editButton = document.createElement("input")
          editButton.type = "button"
          editButton.setAttribute("data-transaction-id", e.id)
          editButton.className = "edit-button action-button"
          editButton.value = "edit"
          editButton.style.visibility = "hidden";
          editButton.addEventListener("click", ev.clickEditButton)

          cell.className = "action-field"
          let deleteButton = document.createElement("input")
          deleteButton.type = "button"
          deleteButton.setAttribute("data-transaction-id", e.id)
          deleteButton.className = "delete-button action-button"
          deleteButton.value = "delete"
          deleteButton.addEventListener("click", ev.clickDeleteButton)
          deleteButton.style.visibility = "hidden";
          cell.appendChild(editButton)
          cell.appendChild(deleteButton)
      }
      tr.appendChild(cell)
    })
    tableBody.appendChild(tr)
    tableTotal += e.amount
  })
  let totalTr = document.createElement("tr")
  totalTr.className = "total"
  let totalCellText
  for (let i = 0; i < cfg.config.headers.length; i++) {
    let totalCell = document.createElement("td")
    if (cfg.config.headers[i] == "Amount") {
      totalCellText = document.createTextNode(tableTotal.toFixed(2))
    } else {
      totalCellText = document.createTextNode("")
    }
    totalCell.appendChild(totalCellText)
    totalTr.appendChild(totalCell)
  }
  tableBody.appendChild(totalTr)
  table.appendChild(tableBody)
}

// This function gets called on mouseover of a TR and makes
// the action buttons visible
export function inserEditButtonsToTR(node) {
  let actionTD = node.querySelectorAll('.action-field')[0]
  for (let i = 0; i < actionTD.children.length; i++) {
    actionTD.children[i].style.visibility = "visible"
  }
}

// This function gets called on mouseout of a TR and hides
// the action buttons
export function removeEditButtonsFromTR(node) {
  let actionTD = node.querySelectorAll('.action-field')[0]
  for (let i = 0; i < actionTD.children.length; i++) {
    actionTD.children[i].style.visibility = "hidden"
  }
}
