function getIn() {
    var data2 = {
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
        password2 : document.getElementById("password2").value
    }

    var mail = data2.email;
    if (mail == '') {
        alert("Please can you add your email")
        return
    } else if (data2.password == '') {
        alert('password is required')
        return
    } else if (data2.password2 == !data2.password){
        alert("password doest match!")
    }
    
    fetch('/login', {
        method: "POST",
        body: JSON.stringify(data2),
        headers: { "content-type": "application/json; charset=UTF-8" }
    }).then(response => {
        if (response.status == 200) {
            // fetch('/signup' + uname)
            //     .then(response => response.text())
            window.location.href = "chat.html"
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        alert(e)
    })

}
