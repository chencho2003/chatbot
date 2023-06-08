function login(){
    console.log("oks")
    var data = {
        email : document.getElementById("email").value,
        password:document.getElementById("password").value
     
    }
    fetch("/login",{
        method:"POST",
        body:JSON.stringify(data),
        headers: { "content-type": "application/json; charset=UTF-8" }
    }).then(response =>{
        if (response.status==200){
            alert("login successful")
            window.location.href = "../chat.html"
            // window.open("../chat.html","_self")
        }else if (response.status==401){
            alert("invalid login")
        }else{
            throw new Error(response)
        }
    })
}

function Logout(){
    fetch("/logout")
    .then(res => {
        if (res.ok){
            window.open("index.html","_self")
        }else{
            throw new Error(res.statusText)
        }
    }).catch(e =>{
        alert(e)
    })
}