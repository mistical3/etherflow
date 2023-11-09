const editProfileBtn = document.getElementById('profileButton');

const modal = document.getElementById('profileModal');

const nameInput = document.getElementById('nameInput');
const avatarInput = document.getElementById('avatarInput');

editProfileBtn.addEventListener('click', () => {
    // Откройте модальное окно
    modal.style.display = 'block';
});

// Слушайте событие отправки формы редактирования профиля
document.getElementById('profileForm').addEventListener('submit', (e) => {
    e.preventDefault();

    // Получите данные из формы
    const newName = nameInput.value;
    const newAvatar = avatarInput.value;

    // Обновите данные профиля (например, отправьте их на сервер)
    // Закройте модальное окно
    modal.style.display = 'none';

    // Обновите отображение имени и аватара на странице
    updateProfile(newName, newAvatar);
});

// Функция для обновления данных профиля на странице
function updateProfile(name, avatar) {
    // Вставьте новое имя и аватар на страницу
    // Например:
    // document.getElementById('profileName').textContent = name;
    // document.getElementById('profileAvatar').src = avatar;
}