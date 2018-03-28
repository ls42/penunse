//
// Initial stuff
//

// Used for text search in table headers
String.prototype.fuzzy = function (s) {
	var hay = this.toLowerCase(),
		i = 0,
		n = -1,
		l
	s = s.toLowerCase()
	for (;
		(l = s[i++]);)
		if (!~(n = hay.indexOf(l, n + 1))) return false
	return true
}

// let apiBase = "http://localhost:4202/api"
let apiBase = "https://penunse.ls42.de/api"

// Get the transaction data from the server
function callAPI() {
	let request = new Request(`${apiBase}/transaction/read/all`, {
		headers: new Headers({
			"X-Clacks-Overhead": "GNU Terry Pratchett"
		})
	})
	fetch(request).then(function (resp) {
		resp.json().then(function (transactions) {
			constructTable(transactions)
		}).catch(function (err) {
			console.log(err)
			console.log("Couldn't convert API data to JSON")
		})
	}).catch(function (err) {
		console.log("Error calling API")
	})
}

// Create a table based on the data in `transactions`
function constructTable(transactions) {

	// Generate table header
	let headers = ["Date", "Amount", "Tags", "Note"]
	let leftTable = document.getElementById("table-left")
	let rightTable = document.getElementById("table-right")
	let tHead = document.createElement("thead")
	let tr = document.createElement("tr")
	headers.forEach(function(e) {
		let cell = document.createElement("th")
		cell.appendChild(document.createTextNode(e))
		tr.appendChild(cell)
	})
	tHead.appendChild(tr)
	leftTable.appendChild(tHead)
	rightTable.appendChild(tHead.cloneNode(true))

	// Generate table body
	let leftBody = document.createElement("tbody")
	let rightBody = document.createElement("tbody")
	leftBody.id = "body-left"
	rightBody.id = "body-right"
	transactions.forEach(function(e) {
		let tr = document.createElement("tr")
		headers.forEach(function(f) {
			let cell = document.createElement("td")
			switch (f) {
				case "Date":
					let createdDate = new Date(e.created)
					let dateCell = document.createTextNode(createdDate.toLocaleString())
					cell.className = "date"
					cell.appendChild(dateCell)
					break
				case "Amount":
					cell.appendChild(document.createTextNode(e.amount))
					break
				case "Tags":
					cell.appendChild(document.createTextNode(e.tags.join(", ")))
					break
				case "Note":
					cell.appendChild(document.createTextNode(e.note))
					break
			}
			tr.appendChild(cell)
		})
		if (e.user_id === 0) {
			leftBody.appendChild(tr)
		} else {
			rightBody.appendChild(tr)
		}
	})
	leftTable.appendChild(leftBody)
	rightTable.appendChild(rightBody)
}

callAPI()
