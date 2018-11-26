import * as dm from "./domManipulator.js"
import * as sc from "./serverConnector.js"

export function mouseEnterTransactionTR(event) {
	dm.inserEditButtonsToTR(event.target)
}

export function mouseOutTransactionTR(event) {
	dm.removeEditButtonsFromTR(event.target)
}

export function clickEditButton(event) {
	let transaction_id = event.target.dataset.transactionId
}

export function clickDeleteButton(event) {
	let transaction_id = event.target.dataset.transactionId
	sc.sendDeleteTransaction(transaction_id)
}
