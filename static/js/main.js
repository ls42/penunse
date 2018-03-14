//
// Initial stuff
//

// [Stolen](https://stackoverflow.com/a/15252131)
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

console.log("Hey, this is debug stuff. You should get out while you can!!")
