function showModal(e) {
    console.log("eid: ", e.getAttribute("data-ID"))
    let eDetails = `
        <div class="box">
            <div>
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" stroke="#ffffff"
                    class="left-arrow" onclick="hideModal()">
                    <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                    <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                    <g id="SVGRepo_iconCarrier">
                        <path
                            d="M9.66088 8.53078C9.95402 8.23813 9.95442 7.76326 9.66178 7.47012C9.36913 7.17698 8.89426 7.17658 8.60112 7.46922L9.66088 8.53078ZM4.47012 11.5932C4.17698 11.8859 4.17658 12.3607 4.46922 12.6539C4.76187 12.947 5.23674 12.9474 5.52988 12.6548L4.47012 11.5932ZM5.51318 11.5771C5.21111 11.2936 4.73648 11.3088 4.45306 11.6108C4.16964 11.9129 4.18475 12.3875 4.48682 12.6709L5.51318 11.5771ZM8.61782 16.5469C8.91989 16.8304 9.39452 16.8152 9.67794 16.5132C9.96136 16.2111 9.94625 15.7365 9.64418 15.4531L8.61782 16.5469ZM5 11.374C4.58579 11.374 4.25 11.7098 4.25 12.124C4.25 12.5382 4.58579 12.874 5 12.874V11.374ZM15.37 12.124V12.874L15.3723 12.874L15.37 12.124ZM17.9326 13.1766L18.4614 12.6447V12.6447L17.9326 13.1766ZM18.25 15.7351C18.2511 16.1493 18.5879 16.4841 19.0021 16.483C19.4163 16.4819 19.7511 16.1451 19.75 15.7309L18.25 15.7351ZM8.60112 7.46922L4.47012 11.5932L5.52988 12.6548L9.66088 8.53078L8.60112 7.46922ZM4.48682 12.6709L8.61782 16.5469L9.64418 15.4531L5.51318 11.5771L4.48682 12.6709ZM5 12.874H15.37V11.374H5V12.874ZM15.3723 12.874C16.1333 12.8717 16.8641 13.1718 17.4038 13.7084L18.4614 12.6447C17.6395 11.8276 16.5267 11.3705 15.3677 11.374L15.3723 12.874ZM17.4038 13.7084C17.9435 14.245 18.2479 14.974 18.25 15.7351L19.75 15.7309C19.7468 14.572 19.2833 13.4618 18.4614 12.6447L17.4038 13.7084Z"
                            fill="#ffffff"></path>
                    </g>
                </svg>
            </div>
            <form action="/dashboard/" method="POST">
            <div class="input-box" hidden>
                <input type=text name=id value="${e.getAttribute("data-id")}">
                <label>ID</label>
            </div>
            <div class="input-box">
                <input type=text name=name value="${e.getAttribute("data-name")}" required>
                <label>Name</label>
            </div>
            <div class="input-box">
                <input type=text name=username value=${e.getAttribute("data-username")} required>
                <label>Username</label>
            </div>
            <div class="input-box">
                <input type=text name=password value=${e.getAttribute("data-password")} required>
                <label>Password</label>
            </div>
            <div class="input-box">
                <input type=text name=active value=${e.getAttribute("data-active")} required>
                <label>Active</label>
            </div>
            <div>
                <label class="switch">
                    <input type="checkbox">
                    <span class="slider"></span>
                </label>
            </div>
            <input id="submit" type="submit" name="" value="Update" disabled>
            </form>
        </div>
    `;
    let modal = document.getElementById("modal-cont");
    modal.innerHTML = eDetails;
    modal.removeAttribute("hidden")

    document.querySelectorAll(".input-box input").forEach((value) => {
        value.addEventListener("change", () => {
            document.getElementById("submit").removeAttribute("disabled")
        });
    })
}

function hideModal() {
    document.getElementById("modal-cont").setAttribute("hidden", "hidden");
}
