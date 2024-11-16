clickElem(document.getElementById("submit"));
function login() {
    data = {
        "username": document.getElementById("user-name").value,
        "password": document.getElementById("user-pass").value
    }
    fetch('/api/login', {
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
                location.reload();
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
