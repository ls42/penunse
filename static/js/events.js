import * as dm from "./domManipulator.js"

export function mouseEnterTransactionTR(event) {
	dm.inserEditButtonsToTR(event.target)
}

export function mouseOutTransactionTR(event) {
	dm.removeEditButtonsFromTR(event.target)
}
