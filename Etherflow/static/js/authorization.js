const connectButton = document.getElementById('connectButton')
const profileButton = document.getElementById('profileButton')

connectButton.addEventListener('click', async () => {


    try {
        // Проверяем, доступен ли MetaMask
        if (typeof window.ethereum !== 'undefined') {
            // Запрашиваем доступ к кошельку
            await window.ethereum.request({ method: 'eth_requestAccounts' });

            // Теперь у нас есть доступ к кошельку и аккаунту пользователя
            const accounts = await window.ethereum.request({ method: 'eth_accounts' });
            const userAddress = accounts[0];

            localStorage.setItem("userAddress", userAddress)

            connectButton.classList.add("hidden")

            profileButton.classList.remove("hidden")

            alert(`Connected to MetaMask!\nYour address: ${userAddress}`);

            const data = { address: userAddress };

            const response = await fetch('http://localhost:1234/send-metamask-data', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();
            console.log('Server response:', result);
        } else {
            alert('MetaMask is not available.');
        }
    } catch (error) {
        console.error(error);
        alert('An error occurred while connecting to MetaMask.');
    }
});