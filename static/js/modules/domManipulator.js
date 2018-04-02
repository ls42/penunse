import * as cfg from "./config.js"

// Insert a new tr in both tables to insert a new transaction into the database.
export function addInsertTRToTable(node) {
	// 1. Find the table in which the add-button has been clicked
	// 2. Generate new row and insert it into thattable
	let side = node.parentNode.className
	let body = document.getElementById(`body-${side}`)
	let insertRow = body.insertRow(0)
	cfg.config.headers.forEach(function (header) {
		let newTd = document.createElement("td")
		// TODO: input an td anhaengen dann td an tr haengen, fertig
		switch (header) {
			case "Date":
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
}

// Create a table based on `transactions` (which is an array of objects)
// Usually gets called from `reloadData()` in `modules/serverConnector.js`
export function constructTable(transactions) {
	let leftTable = document.getElementById("table-left")
	let rightTable = document.getElementById("table-right")

	// Cleanup existing table
	let tables = [leftTable, rightTable]
	tables.forEach(function (table) {
		while (table.firstChild) {
			table.removeChild(table.firstChild)
		}
	})

	// Generate table header
	let tHead = document.createElement("thead")
	let tr = document.createElement("tr")
	cfg.config.headers.forEach(function (e) {
		let cell = document.createElement("th")
		cell.appendChild(document.createTextNode(e))
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
	transactions.forEach(function (e) {
		let tr = document.createElement("tr")
		cfg.config.headers.forEach(function (header) {
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
					e.tags.forEach(function (tag) {
						let tagSpan = document.createElement("span")
						tagSpan.className = "tag"
						tagSpan.appendChild(document.createTextNode(tag))
						cell.appendChild(tagSpan)
					})
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
