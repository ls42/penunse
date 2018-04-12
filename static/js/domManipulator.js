import * as cfg from "./config.js"
import * as ev from "./events.js"

// Insert a new tr in both tables to insert a new transaction into the database.
export function handlePlusButtonClick(node) {
	// Add input row to table
	let side = node.parentNode.parentNode.className
	let body = document.getElementById(`body-${side}`)
	let insertRow = body.insertRow(0)
	insertRow.id = `insert-row-${side}`
	cfg.config.headers.forEach((header) => {
		let newTd = document.createElement("td")
		switch (header) {
			case "Date":
				newTd.appendChild(document.createTextNode(""))
				insertRow.appendChild(newTd)
				break;
			case "Action":
				newTd.appendChild(document.createTextNode(""))
				insertRow.appendChild(newTd)
				break;
			default:
				let inputId = `input_${header.toLowerCase()}`
				let newInput = document.createElement("input")
				newInput.id = inputId
				newInput.type = "text"
				newInput.className = "transaction-input"
				newTd.appendChild(newInput)
				insertRow.appendChild(newTd)
				break;
		}
	})

	// Change `+`-button to submit
	node.value = "save"
	node.setAttribute("onclick", "sc.sendNewTransaction(this)")
}

// Create a table based on `transactions` (which is an array of objects)
// Usually gets called from `reloadData()` in `modules/serverConnector.js`
export function constructTable(transactions) {
	let leftTable = document.getElementById("table-left")
	let rightTable = document.getElementById("table-right")

	// Cleanup existing table
	let tables = [leftTable, rightTable]
	tables.forEach((table) => {
		while (table.firstChild) {
			table.removeChild(table.firstChild)
		}
	})

	// Generate table header
	let tHead = document.createElement("thead")
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
	tHead.appendChild(tr)
	leftTable.appendChild(tHead)
	rightTable.appendChild(tHead.cloneNode(true))

	// Generate table body
	let leftTotal = 0
	let rightTotal = 0
	let leftBody = document.createElement("tbody")
	let rightBody = document.createElement("tbody")
	leftBody.id = "body-left"
	rightBody.id = "body-right"
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
							tagSpan.appendChild(document.createTextNode(tag))
							cell.appendChild(tagSpan)
						})
					}
					break
				case "Note":
					let noteCell
					if (e.note.length >= 20) {
						noteCell = document.createTextNode(e.note.substr(0, 20) + " [...]")
					} else {
						noteCell = document.createTextNode(e.note)
					}
					cell.title = e.note
					cell.appendChild(noteCell)
					break
				case "Action":
					// This is the target for the action buttons
					let editButton = document.createElement("input")
					editButton.type = "button"
					editButton.setAttribute("data-transaction-id", e.id)
					editButton.className = "edit-button action-button"
					editButton.value = "edit"
					editButton.style.visibility = "hidden";
					editButton.addEventListener("onclick", ev.clickEditButton)
					cell.className = "action-field"
					let deleteButton = document.createElement("input")
					deleteButton.type = "button"
					deleteButton.setAttribute("data-transaction-id", e.id)
					deleteButton.className = "delete-button action-button"
					deleteButton.value = "delete"
					deleteButton.addEventListener("onclick", ev.clickDeleteButton)
					deleteButton.style.visibility = "hidden";
					cell.appendChild(editButton)
					cell.appendChild(deleteButton)
			}
			tr.appendChild(cell)
		})
		if (e.user_id === 0) {
			leftBody.appendChild(tr)
			leftTotal += e.amount
		} else {
			rightBody.appendChild(tr)
			rightTotal += e.amount
		}

	})
	let totalTrLeft = document.createElement("tr")
	totalTrLeft.className = "total"
	let totalTrRight = document.createElement("tr")
	totalTrRight.className = "total"
	let totalCellText
	for (let i = 0; i < cfg.config.headers.length; i++) {
		let totalCell = document.createElement("td")
		if (cfg.config.headers[i] == "Amount") {
			totalCellText = document.createTextNode(leftTotal.toFixed(2))
		} else {
			totalCellText = document.createTextNode("")
		}
		totalCell.appendChild(totalCellText)
		totalTrLeft.appendChild(totalCell)
	}
	for (let i = 0; i < cfg.config.headers.length; i++) {
		let totalCell = document.createElement("td")
		if (cfg.config.headers[i] == "Amount") {
			totalCellText = document.createTextNode(rightTotal.toFixed(2))
		} else {
			totalCellText = document.createTextNode("")
		}
		totalCell.appendChild(totalCellText)
		totalTrRight.appendChild(totalCell)
	}
	leftBody.appendChild(totalTrLeft)
	rightBody.appendChild(totalTrRight)
	leftTable.appendChild(leftBody)
	rightTable.appendChild(rightBody)
}

// This function gets called on mouseover of a TR and inserts the edit
// buttons into the 'Action' column of the current row
export function inserEditButtonsToTR(node) {
	let actionTD = node.querySelectorAll('.action-field')[0]
	for (let i = 0; i < actionTD.children.length; i++) {
		actionTD.children[i].style.visibility = "visible"
	}
}

// This function gets called on mouseout of a TR and cleans the 'Action' column
// of the current row
export function removeEditButtonsFromTR(node) {
	let actionTD = node.querySelectorAll('.action-field')[0]
	for (let i = 0; i < actionTD.children.length; i++) {
		actionTD.children[i].style.visibility = "hidden"
	}
}
