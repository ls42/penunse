import * as dm from "./domManipulator.js"

export function mouseEnterTransactionTR(event) {
	dm.inserEditButtonsToTR(event.target)
}

export function mouseOutTransactionTR(event) {
	dm.removeEditButtonsFromTR(event.target)
}

export function clickEditButton(event) {
	console.log(`edit transaction #${event.target.dataset.transactionId}`)
}

export function clickDeleteButton(event) {
	console.log(`delete transaction #${event.target.dataset.transactionId}`)
}
