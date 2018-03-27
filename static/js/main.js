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
	for (; (l = s[i++]);) if (!~(n = hay.indexOf(l, n + 1))) return false
	return true
}

let apiBase = "https://penunse.ls42.de/api"

// Try a call to the API
function callAPI () {
	let request = new Request(`${apiBase}/transaction/read/all`, {
		headers: new Headers({
			"X-Clacks-Overhead": "GNU Terry Pratchett"
		})
	})
	fetch(request).then(function(resp) {
		resp.json().then(function(transactions) {
			constructTable()
		}).catch(function(err) {
			console.log(err)
			console.log("Couldn't convert API data to JSON")
		})
	}).catch(function(err) {
		console.log("Error calling API")
	})
}

function constructTable (transactions) {
	let target = document.getElementById("table-grid")
	let t = document.createElement("table")
	let thead = document.createElement("thead")
	let thead_tr = document.createElement("tr")
	["User", "Amount", "Tags", "Notes"].forEach(function(e) {
		let thead_tr_th = document.createElement("th")
		thead_tr_th.appendChild(document.createTextNode(e))
		thead_tr.appendChild(thead_tr_th)
	})
	thead.appendChild(thead_tr)
	t.appendChild(thead)
	let tbody = document.createElement("tbody")
	transactions.forEach(function(e) {
		let tbody_tr = document.createElement("tr")
		for (let prop in e) {
			let tbody_tr_td = document.createElement("td")
			tbody_tr_td.appendChild(document.createTextNode(e[prop]))
			tbody_tr.appendChild(tbody_tr_td)
		}
		tbody.appendChild(tbody_tr)
	})
	t.appendChild(tbody)
	target.appendChild(t)
}

callAPI()
