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
function fetchData(vm) {
  let req = new Request(`${conf.apiBaseUrl}/transaction/read`)
  fetch(req)
    .then(resp => {
      resp
        .json()
        .then(ts => {
          vm.transactions = ts
        })
        .catch(err => {
          alert("#1 response not valid JSON")
        })
    })
    .catch(err => {
      alert("#2 cannot talk to api server")
    })
}

Vue.component("transaction-table-head", {
  template: `
    <tr>
      <th>date</th>
      <th>amount</th>
      <th>note</th>
      <th>name</th>
      <th>tags</th>
    </tr>
  `,
})

Vue.component("transaction-item", {
  props: {
    trans: Array,
  },
  filters: {
    dataReadable: function(date) {
      return new Date(date).toLocaleDateString("de-DE")
    },
    userName: function(user_id) {
      return user_id ? "Kerstin" : "Stephan"
    },
    getTags: function(raw_tags) {
      let mytags = ""
      if (raw_tags !== null) {
        raw_tags.forEach(function(raw_tag) {
          mytags.concat("asf", `<span class="tag">${raw_tag.name}</span>`)
        })
        console.log(mytags)
        return mytags
      }
    },
  },
  template: `
    <tr>
      <td>{{ trans.created | dataReadable }}</td>
      <td>{{ trans.amount }},-</td>
      <td><b>{{ trans.note }}</b></td>
      <td>{{ trans.user_id | userName }}</td>
      <td class="tags">{{ trans.tags | getTags }}</td>
    </tr>`,
})

// Create the application
var vm = new Vue({
  el: "#app",
  data: {
    transactions: [],
  },
  created: function() {
    fetchData(this)
  },
})
