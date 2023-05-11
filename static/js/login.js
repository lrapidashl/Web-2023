const passwordInput = document.querySelectorAll('.login-form__text-input')[1];
const passwordEye = document.querySelector(".login-form__password-eye");
function showPassword() {
    if (passwordInput.type === "password") {
        passwordInput.type = "text";
        passwordEye.src = "/static/img/password_eye_off.svg";
    }
    else {
        passwordInput.type = "password";
        passwordEye.src = "/static/img/password_eye.svg";
    }
}

function initEventListeners() {
    passwordEye.addEventListener("click", showPassword);
}
initEventListeners();