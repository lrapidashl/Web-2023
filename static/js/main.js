const menuButton = document.getElementById("heading__menu");
const menu = document.getElementById("heading__burger");

function toggleBurgerMenu(event) {
    if (event.target === menuButton) {
        menu.classList.remove("heading__burger_hidden");
    }
    else if (event.target !== menu) {
        menu.classList.add("heading__burger_hidden");
    }
}

document.addEventListener("click", toggleBurgerMenu);