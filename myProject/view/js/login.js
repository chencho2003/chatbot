document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Prevent form submission
  
    console.log("oks");
    var data = {
      email: document.getElementById("email").value,
      password: document.getElementById("password").value
    };
  
    fetch("/login", {
      method: "POST",
      body: JSON.stringify(data),
      headers: { "content-type": "application/json; charset=UTF-8" }
    })
      .then(function(response) {
        if (response.status == 200) {
          
          window.location.href = "../chat.html";
        } else if (response.status == 480) {
          alert("admin login successful");
          window.location.href = "../admin.html"
        }
        else if (response.status == 401) {
          alert("Invalid login");
        } else {
          throw new Error(response.statusText);
        }
      })
      .catch(function(error) {
        alert(error);
      });
  });
  