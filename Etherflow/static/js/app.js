VANTA.NET({
        el: "#wrp",
    mouseControls: true,
    touchControls: true,
    gyroControls: false,
    minHeight: 200.00,
    minWidth: 200.00,
    scale: 1.00,
    scaleMobile: 1.00,
    color: 0xc53fff,
    maxDistance: 25.00,
    spacing: 17.00
})



window.onload = () => {
    localStorage.clear()
    if (localStorage.getItem("userAddress")) {
        const connectButton = document.getElementById('connectButton')
        connectButton.classList.add("hidden")
    } else {
        const profileButton = document.getElementById('profileButton')
        profileButton.classList.add("hidden")
    }
}

// window.addEventListener("close", () => {
//     localStorage.clear()
// })
