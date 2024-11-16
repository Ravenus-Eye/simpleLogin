function setCookie(name, value, days) {
    let expires = "";
    if (days) {
        const date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + encodeURIComponent(value) + expires + "; path=/";
}

function getCookie(name) {
    const nameEQ = name + "=";
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        let cookie = cookies[i].trim();
        if (cookie.indexOf(nameEQ) === 0) {
            return decodeURIComponent(cookie.substring(nameEQ.length));
        }
    }
    return null;
}

function clickElem(e) {
    document.addEventListener("keydown", function (event) {
        if (event.key === "Enter" || event.keyCode === 13) {
            e.click();
        }
    });
}

const passwordElement = document.querySelectorAll('.password');
passwordElement.forEach(value => {
    var eyeClose = true;
    value.addEventListener('click', (event) => {
        // Get the element's position and size
        const rect = value.getBoundingClientRect();

        // Calculate the area of the ::after pseudo-element based on its CSS properties
        const afterLeft = rect.right - 26; // 25px width + 1px right
        const afterTop = rect.top + 10;    // 10px top offset
        const afterWidth = 25;
        const afterHeight = 25;

        // Check if the click is within the ::after area
        if (
            event.clientX >= afterLeft &&
            event.clientX <= afterLeft + afterWidth &&
            event.clientY >= afterTop &&
            event.clientY <= afterTop + afterHeight
        ) {
            // Example: toggle the background image
            var newImage;
            const currentImage = value.style.getPropertyValue('--background-image');
            if (eyeClose) {
                newImage = "url('/static/images/eyepass.svg')";
                eyeClose = false;
            } else {
                newImage = currentImage.includes("eyeclose.svg")
                    ? "url('/static/images/eyepass.svg')"
                    : "url('/static/images/eyeclose.svg')";
            }
            if (newImage == "url('/static/images/eyepass.svg')") {
                value.querySelector("input").setAttribute("type", "text");
            } else {
                value.querySelector("input").setAttribute("type", "password");
            }

            value.style.setProperty('--background-image', newImage);
        }
    });
});

function prevPage() {
    window.history.back();
}