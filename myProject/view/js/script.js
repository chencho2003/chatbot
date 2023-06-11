
function teach() {
  console.log('bug');
  //reading the form data
  var data = {
      question: document.getElementById("question").value,
      answer: document.getElementById("answer").value
  }
  console.log(data);
  fetch('/teach', {
      method: "POST",
      body: JSON.stringify(data),
      headers: { "content-type": "application/json; charset=UTF-8" }
  }).then(response => {
      if (response.status == 201) {
          alert(`Successful Added`)
          // fetch('/signup' + uname)
          //     .then(response => response.text())
          location.reload();
      } else {
          throw new Error(response.statusText)

         
      }
  }).catch(e => {
      alert(e)
  })
}



function addBotToTable(id, question, answer) {
  var botTableBody = document.querySelector("#botTable tbody");

  var row = document.createElement("tr");
  row.innerHTML = `
    <td>${id}</td>
    <td>${question}</td>
    <td>${answer}</td>
    <td>
      <button onclick="editBot(${id})">Edit</button>
      <button onclick="deleteBot(${id})">Delete</button>
    </td>
  `;

  botTableBody.appendChild(row);
}


function deleteBot(id) {
  if (confirm("Are you sure you want to delete this bot?")) {
    fetch('/delete', {
      method: "DELETE",
      body: JSON.stringify({ id: id }),
      headers: { "Content-Type": "application/json" }
    })
    .then(response => {
      if (response.status === 201) {
        alert("Successfully Deleted");
        location.reload();
        
      } else {
        throw new Error(response.statusText);
      }
    })
    .catch(error => {
      console.error(error);
      alert("An error occurred while deleting the bot");
    });
  }
}



function editBot(id) {
  console.log('test')
  console.log(id)

    var newQuestion = prompt("Enter the new question:");
    var newAnswer = prompt("Enter the new answer:");

    if (newQuestion !== null && newAnswer !== null) {
      fetch('/updateqna', {
        method: "PUT",
        body: JSON.stringify({ id: id, question: newQuestion, answer: newAnswer }),
        headers: { "Content-Type": "application/json" }
      })
      .then(response => {
        if (response.status === 200) {
          alert("Bot successfully updated");
          location.reload();
        } else {
          throw new Error(response.statusText);
        }
      })
      .catch(error => {
        console.error(error);
        alert("An error occurred while updating the bot");
      });
    }
  }

// Fetch existing bots when the page loads
document.addEventListener("DOMContentLoaded", function() {
  fetchBots();
});

function fetchBots() {
  fetch("/getall")
    .then(response => {
      if (response.status === 200) {
        return response.json();
      } else {
        throw new Error(response.statusText);
      }
    })
    .then(data => {
      var botTableBody = document.querySelector("#botTable tbody");
      botTableBody.innerHTML = "";

      data.forEach(bot => {
        addBotToTable(bot.id, bot.question, bot.answer);
      });
    })
    .catch(error => {
      console.error(error);
      alert("An error occurred while fetching the bots");
    });
}