const emailInput = document.getElementById("login-form__email-input");
const passwordInput = document.getElementById("login-form__password-input");
const passwordEye = document.querySelector(".login-form__password-eye");
const loginButton = document.querySelector(".login-form__submit-button");

const emailError = document.getElementById("login-form__email-error");
const passwordError = document.getElementById("login-form__password-error");
const formError = document.querySelector(".logination__error-message");

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

function setTextInputClass(event) {
    el = event.target;
    if (!el.value) {
        el.classList.remove("login-form__text-input_filled");
    }
    else { 
        el.classList.add("login-form__text-input_filled");
    }
}


async function logIn() {
    emailError.classList.add("hidden");
    passwordError.classList.add("hidden");
    emailError.innerHTML = "";
    passwordError.innerHTML = "";
    for (i of formError.children) {
        i.classList.add("hidden");
    } 
    formError.classList.add("logination__error-message_hidden");
    formError.lastChild.previousSibling.innerHTML = "";
    emailInput.classList.remove("login-form__text-input_error");
    passwordInput.classList.remove("login-form__text-input_error");
    let inputError = false;
    const reg = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

    if (emailInput.value === "") {
        emailError.classList.remove("hidden");
        emailError.innerHTML = "Email is required.";
        emailInput.classList.add("login-form__text-input_error");
        inputError = true;
    } 
    else if (!emailInput.value.match(reg)) {
        emailError.classList.remove("hidden");
        emailError.innerHTML = "Incorrect email format. Correct format is ****@**.***";
        emailInput.classList.add("login-form__text-input_error");
        inputError = true;
    }

    if (passwordInput.value === "") {
        passwordError.classList.remove("hidden");
        passwordError.innerHTML = "Password is required.";
        passwordInput.classList.add("login-form__text-input_error");
        inputError = true;
    }

    if (inputError) {
        formError.lastChild.previousSibling.innerHTML = "A-Ah! Check all fields.";
        formError.classList.remove("logination__error-message_hidden");
        for (i of formError.children) {
            i.classList.remove("hidden");
        } 
    }
    /*else {
        const response = await fetch("/api/auth", {
            method: "POST",
            headers: {
                "Content-Type": "application/json;charset=utf-8"
            },
            body: inputValues
        });
        if (response.ok) {
            let authError = response.json().check;
            if (authError === "no") {
                formError.classList.remove("hidden");
                formError.lastChild.previousSibling.innerHTML = "Email or password is incorrect.";
            }
            else {
                window.location.href = "/admin"
            }
        }
        else {
            alert("Ошибка HTTP: " + response.status);
        }
    }*/

}

function initEventListeners() {
    emailInput.addEventListener("input", setTextInputClass);
    passwordInput.addEventListener("input", setTextInputClass);
    passwordEye.addEventListener("click", showPassword);
    loginButton.addEventListener("click", logIn);
}

initEventListeners();