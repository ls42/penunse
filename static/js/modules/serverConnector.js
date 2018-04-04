import * as cfg from "./config.js"
import * as dm from "./domManipulator.js"

// Fetch all transactions from the API servers
export function reloadData() {
	let request = new Request(`${cfg.config.apiBase}/transaction/read/all`, {
		headers: new Headers({
			"X-Clacks-Overhead": "GNU Terry Pratchett"
		})
	})
	fetch(request).then(function (resp) {
		resp.json().then(function (transactions) {
			dm.constructTable(transactions)
		}).catch(function (err) {
			console.log(err)
			console.log("Couldn't convert API data to JSON")
		})
	}).catch(function (err) {
		console.log("Error calling API")
	})
}

export function sendNewTransaction(node) {
	// stuff from inputs to the server and call deleteInputAndAppendNewTransaction
	// also create toast if server says okay or not.
	console.log("nice to meet you")
	let side = node.parentNode.parentNode.className
	let insertRow = node.parentNode.parentNode.getElementById("insert-row")
	console.log(insertRow)
	// let newData = {
	// 	id: node.getElementById()
	// }
}
