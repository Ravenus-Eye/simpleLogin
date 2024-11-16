clickElem(document.getElementById("submit"));
function createAccount() {
    data = {
        "username": document.getElementById("user-name").value,
        "name": document.getElementById("name").value,
        "password": document.getElementById('user-pass').value
    }
    console.log("data: ", data)
    fetch('/api/create_account', {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded", // Adjust the content type if needed
        },
        body: new URLSearchParams(data), // Serialize the data
    })
        .then(response => response.json())
        .then(data => {
            console.log("data: ", data)
            var error = document.getElementById("error-msg");
            if (data["status"] == "error") {
                error.innerText = data["err_msg"];
                error.removeAttribute("hidden", "hidden");
            } else {
                error.innerText = data["err_msg"];
                error.removeAttribute("hidden", "hidden");
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}