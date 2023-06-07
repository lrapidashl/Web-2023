const authorAvatarInput = document.getElementById("upload-avatar__file-field");
const authorAvatarPreview = document.querySelector(".upload-avatar__author-avatar-img");
const authorAvatarPreviewPost = document.querySelector(".card-area__author-avatar");
const authorAvatarInscription = document.querySelector(".upload-avatar__inscription");
const authorAvatarButtons = document.getElementById("add-post-form__uploaded-avatar-buttons");
let base64AuthorAvatar = "";

const pageImgInput = document.getElementById("upload-page-img__file-field");
const pageImgPreview = document.getElementById("upload-page-img__button");
const pageImgPreviewPost = document.querySelector(".page-area__img");
const pageImgPreviewCard = document.getElementById("card-area__img");
const pageImgInscription = document.getElementById("add-post-form__page-inscription");
const pageImgButtons = document.getElementById("add-post-form__uploaded-page-img-buttons");
let base64PageImg = "";

const formError = document.querySelector(".new-post__error-message");
const formSuccess = document.querySelector(".new-post__success-message");
const titleError = document.getElementById("add-post-form__title-error");
const subtitleError = document.getElementById("add-post-form__subtitle-error");
const authorNameError = document.getElementById("add-post-form__author-name-error");
const authorAvatarError = document.getElementById("add-post-form__author-avatar-error");
const publishDateError = document.getElementById("add-post-form__publish-date-error");
const pageImgError = document.getElementById("add-post-form__page-img-error");
const contentError = document.querySelector(".content__error");
const contentData = document.querySelector(".content__text-area");

const titleInput = document.getElementById("add-post-form__title-input-field");
const pageAreaTitle = document.querySelector(".page-area__title");
const cardAreaTitle = document.querySelector(".card-area__title");

const subtitleInput = document.getElementById("add-post-form__subtitle-input-field");
const pageAreaSubtitle = document.querySelector(".page-area__subtitle");
const cardAreaSubtitle = document.querySelector(".card-area__subtitle");

const authorNameInput = document.getElementById("add-post-form__author-name-input-field");
const cardAreaAuthorName = document.querySelector(".card-area__author-name");

const publishDateInput = document.getElementById("add-post-form__publish-date-input-field");
const cardAreaPublishDate = document.querySelector(".card-area__publish-date");

const authorAvatarRemoveButton = document.getElementById("uploaded-avatar-buttons__remove");
const pageImgRemoveButton = document.getElementById("uploaded-page-buttons__remove");
const publishButton = document.querySelector(".new-post__publish-button");
const logoutButton = document.querySelector(".heading__log-out");

function previewAuthorAvatar() {
    const file = authorAvatarInput.files[0];
    
    if (file) {
        const reader = new FileReader();
        reader.onloadend = function () {
            base64AuthorAvatar = reader.result.replace("data:", "").replace(/^.+,/, "")
            authorAvatarPreview.style.backgroundImage = "url(" + reader.result + ")";
            authorAvatarPreviewPost.style.backgroundImage = "url(" + reader.result + ")";
        }
        authorAvatarInscription.classList.add("hidden");
        authorAvatarButtons.classList.remove("hidden");
        for (i of authorAvatarPreview.children) {
            i.classList.add("hidden");
        }
        authorAvatarPreview.style.backgroundColor = "rgba(0, 0, 0, 0)";
        authorAvatarPreviewPost.style.backgroundColor = "rgba(0, 0, 0, 0)";
        authorAvatarPreview.style.border = "none";
        reader.readAsDataURL(file); 
    }
}

function removeAuthorAvatar() {
    authorAvatarInput.value = null;
    for (i of authorAvatarPreview.children) {
        i.classList.remove("hidden");
    }
    authorAvatarPreview.style.backgroundImage = "";
    authorAvatarPreviewPost.style.backgroundImage = "";
    authorAvatarPreview.style.backgroundColor = "#F7F7F7";
    authorAvatarPreviewPost.style.backgroundColor = "#F7F7F7";
    authorAvatarPreview.style.border = "1px dashed #D3D3D3";
    authorAvatarInscription.classList.remove("hidden");
    authorAvatarButtons.classList.add("hidden");
}

function previewPageImg() {
    const file = pageImgInput.files[0];
    
    if (file) {
        const reader = new FileReader();
        reader.onloadend = function () {
            base64PageImg = reader.result.replace("data:", "").replace(/^.+,/, "");
            pageImgPreview.style.backgroundImage = "url(" + reader.result + ")";
            pageImgPreviewPost.style.backgroundImage = "url(" + reader.result + ")";
            pageImgPreviewCard.style.backgroundImage = "url(" + reader.result + ")";
        }
        pageImgInscription.classList.add("hidden");
        pageImgButtons.classList.remove("hidden");
        for (i of pageImgPreview.children) {
            i.classList.add("hidden");
        }
        pageImgPreview.style.backgroundColor = "rgba(0, 0, 0, 0)";
        pageImgPreviewPost.style.backgroundColor = "rgba(0, 0, 0, 0)";
        pageImgPreviewCard.style.backgroundColor = "rgba(0, 0, 0, 0)";
        pageImgPreview.style.border = "none";
        reader.readAsDataURL(file); 
    } 
}

function removePageImg() {
    pageImgInput.value = null;
    pageImgPreview.style.backgroundImage = "";
    pageImgPreviewPost.style.backgroundImage = "";
    pageImgPreviewCard.style.backgroundImage = "";
    for (i of pageImgPreview.children) {
        i.classList.remove("hidden");
    }
    pageImgPreview.style.backgroundColor = "#F7F7F7";
    pageImgPreviewPost.style.backgroundColor = "#F7F7F7";
    pageImgPreviewCard.style.backgroundColor = "#F7F7F7";
    pageImgPreview.style.border = "1px dashed #D3D3D3";
    pageImgInscription.classList.remove("hidden");
    pageImgButtons.classList.add("hidden");
}

function previewTitle() {
    setTextInputClass(titleInput);

    pageAreaTitle.innerHTML = titleInput.value;
    cardAreaTitle.innerHTML = titleInput.value;
}

function previewSubtitle() {
    setTextInputClass(subtitleInput);

    pageAreaSubtitle.innerHTML = subtitleInput.value;
    cardAreaSubtitle.innerHTML = subtitleInput.value;
}

function previewAuthorName() {
    setTextInputClass(authorNameInput);

    cardAreaAuthorName.innerHTML = authorNameInput.value;
}

function previewPublishDate() {
    setTextInputClass(publishDateInput);

    cardAreaPublishDate.innerHTML = publishDateInput.value;
}

function setTextInputClass(el) {
    if (!el.value) {
        el.classList.remove("add-post-form__text-input_filled");
    }
    else { 
        el.classList.add("add-post-form__text-input_filled");
    }
}

function checkInputs() {
    titleError.classList.add("hidden");
    subtitleError.classList.add("hidden");
    authorNameError.classList.add("hidden");
    authorAvatarError.classList.add("hidden");
    publishDateError.classList.add("hidden");
    pageImgError.classList.add("hidden");
    contentError.classList.add("hidden");
    for (i of formError.children) {
        i.classList.add("hidden");
    } 
    for (i of formSuccess.children) {
        i.classList.add("hidden");
    } 
    formError.classList.add("new-post__error-message_hidden");
    formSuccess.classList.add("new-post__success-message_hidden");
    titleInput.classList.remove("add-post-form__text-input_error");
    subtitleInput.classList.remove("add-post-form__text-input_error");
    authorNameInput.classList.remove("add-post-form__text-input_error");
    publishDateInput.classList.remove("add-post-form__text-input_error");
    let inputError = false;

    if (titleInput.value === "") {
        titleError.classList.remove("hidden");
        titleInput.classList.add("add-post-form__text-input_error");
        inputError = true;
    } 

    if (subtitleInput.value === "") {
        subtitleError.classList.remove("hidden");
        subtitleInput.classList.add("add-post-form__text-input_error");
        inputError = true;
    }

    if (authorNameInput.value === "") {
        authorNameError.classList.remove("hidden");
        authorNameInput.classList.add("add-post-form__text-input_error");
        inputError = true;
    }

    if (publishDateInput.value === "") {
        publishDateError.classList.remove("hidden");
        publishDateInput.classList.add("add-post-form__text-input_error");
        inputError = true;
    }

    if (contentData.value === "") {
        contentError.classList.remove("hidden");
        inputError = true;
    }

    if (!authorAvatarInput.files[0]) {
        authorAvatarError.classList.remove("hidden");
        inputError = true;
    }

    if (!pageImgInput.files[0]) {
        pageImgError.classList.remove("hidden");
        inputError = true;
    }

    if (inputError) {
        formError.classList.remove("new-post__error-message_hidden");
        for (i of formError.children) {
            i.classList.remove("hidden");
        } 
    }
    else {
        formSuccess.classList.remove("new-post__success-message_hidden");
        for (i of formSuccess.children) {
            i.classList.remove("hidden");
        }  
    }
    return inputError;
}

async function createPost() {
    if (!checkInputs()) {
        const authorAvatarFile = authorAvatarInput.files[0];
        const pageImgFile = pageImgInput.files[0];
        const cardImgFile = cardImgInput.files[0];
        const authorAvatar = authorAvatarFile ? authorAvatarFile.name : "";
        const pageImg = pageImgFile ? pageImgFile.name : "";
        const cardImg = cardImgFile ? cardImgFile.name : "";
        const formData = {
            title: titleInput.value,
            subtitle: subtitleInput.value,
            author_name: authorNameInput.value,
            author_avatar: authorAvatar,
            author_avatar_file: base64AuthorAvatar,
            publish_date: publishDateInput.value,
            page_image: pageImg,
            page_image_file: base64PageImg,
            card_image: cardImg,
            card_image_file: base64CardImg,
            content: contentData.value
        }
        const json = JSON.stringify(formData);
        console.log(json);

        const response = await fetch("/api/post", {
            method: "POST",
            headers: {
                "Content-Type": "application/json;charset=utf-8"
            },
            body: json
        });
        if (!response.ok) { 
            alert("Ошибка HTTP: " + response.status);
        }
    }
} 

async function logOut() {
    const response = await fetch("/api/logout");
    if (response.ok) { 
        window.location.href = "/login";
    }
    else {
        alert("Ошибка HTTP: " + response.status);
    }
}

function initEventListeners() {
    titleInput.addEventListener("input", previewTitle);
    subtitleInput.addEventListener("input", previewSubtitle);
    authorNameInput.addEventListener("input", previewAuthorName);
    publishDateInput.addEventListener("input", previewPublishDate);
    authorAvatarInput.addEventListener("input", previewAuthorAvatar);
    authorAvatarRemoveButton.addEventListener('click', removeAuthorAvatar);
    pageImgInput.addEventListener("input", previewPageImg);
    pageImgRemoveButton.addEventListener('click', removePageImg);
    publishButton.addEventListener("click", createPost);
    logoutButton.addEventListener("click", logOut);
}

initEventListeners();