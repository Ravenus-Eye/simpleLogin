clickElem(document.getElementById("submit"));
function changePass() {
    data = {
        "username": document.getElementById("user-name").value,
        "password1": document.getElementById("user-pass1").value,
        "password2": document.getElementById('user-pass2').value
    }
    fetch('/api/changePassword', {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded", // Adjust the content type if needed
        },
        body: new URLSearchParams(data), // Serialize the data
    })
        .then(response => response.json())
        .then(data => {
            console.log("data: ", data)
            if (data["status"] == "error") {
                var error = document.getElementById("error-msg");
                error.innerText = data["err_msg"];
                error.removeAttribute("hidden", "hidden");
            } else {
                window.location.href = "/login/"
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}