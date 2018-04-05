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
	let side = node.parentNode.parentNode.className
	let insertRow = document.getElementById(`insert-row-${side}`)
	console.log(insertRow)
	let newData = {
		amount: insertRow.children[1].children['input_amount'].value,
		tags: insertRow.children[2].children['input_tags'].value.trim().split(","),
		note: insertRow.children[3].children['input_note'].value,
	}
	let newDataJSON = JSON.stringify(newData)
	let request = new Request(`${cfg.config.apiBase}/transaction/create`, {
		headers: new Headers({
			"X-Clacks-Overhead": "GNU Terry Pratchett"
		}),
		body: newDataJSON,
	})
	fetch(request).then(function (resp) {
		resp.json().then(function (transactions) {
			// Create a `success`-toast
			// Remove the `temporary`-flag from the newly inserted TR
			// or redraw table
			console.log("I guess it works!")
		}).catch(function (err) {
			// Create a `failure`-toast
			// Remove the newly inserted entry to the table (class `temporary`)
			console.log(err)
			console.log("Couldn't convert API data to JSON")
		})
	}).catch(function (err) {
		console.log("Error calling API")
	})
}
