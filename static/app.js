// WORKFLOW
// ====================================================================
// 1. Fetch all transactions via native FetchAPI
// 2. Do Vue.js stuff with the transactions
//    - I could solve pagination here
//    - Searching
//    - Editing maybe?
//    - Create a table with pagination
//    - How to check if there are new entries?

//
// Temporary global variables
let transactions

//
// Basic configuration
const conf = {
  apiBaseUrl: "http://localhost:4202/api",
  headers: ["Date", "Amount", "Tags", "Note", "User", "Action"],
  users: [{id: 0, name: "Stephan"}, {id: 1, name: "Kerstin"}],
}

//
// Fetch data from API server
function fetchData(app) {
  let req = new Request(`${conf.apiBaseUrl}/transaction/read`)
  fetch(req)
    .then(resp => {
      resp
        .json()
        .then(ts => {
          app.transactions = ts
        })
        .catch(err => {
          alert("#1 response not valid JSON")
        })
    })
    .catch(err => {
      alert("#2 cannot talk to api server")
    })
}

// Create the application
var app = new Vue({
  el: "#app",
  data: {
    transactions: null,
  },
  filters: {
    dateReadable(date) {
      return new Date(date).toLocaleDateString("de-DE")
    },
  },
  mounted() {
    fetchData(this)
  },
})
