const form = document.querySelector(".admin-form"); 
const textInputs = document.querySelectorAll(".add-post-form__text-input");
const authorAvatarInput = document.getElementById("upload-avatar__file-field");
const pageImgInput = document.getElementById("upload-page-img__file-field");
const cardImgInput = document.getElementById("upload-card-img__file-field");
const pageAreaTitle = document.querySelector(".page-area__title");
const cardAreaTitle = document.querySelector(".card-area__title");
const pageAreaSubtitle = document.querySelector(".page-area__subtitle");
const cardAreaSubtitle = document.querySelector(".card-area__subtitle");
const cardAreaAuthorName = document.querySelector(".card-area__author-name");
const cardAreaPublishDate = document.querySelector(".card-area__publish-date");

function previewAuthorAvatar() {
    const file = document.querySelector(".upload-avatar__file-field").files[0];
    const preview = document.querySelector(".upload-avatar__author-avatar");
    const previewPost = document.querySelector(".card-area__author-avatar");
    const reader = new FileReader();
    
    if (file) {
        const replacableInscription = document.querySelector(".upload-avatar__inscription");
        reader.onloadend = function () {
            preview.style.backgroundImage = "url(" + reader.result + ")";
            previewPost.style.backgroundImage = "url(" + reader.result + ")";
        }
        drawUploadNewAndRemove(replacableInscription, "upload-avatar__file-field");
        replacableInscription.remove();
        preview.style.border = "none";
        preview.innerHTML = "";
        preview.style.backgroundColor = "rgba(0, 0, 0, 0)";
        previewPost.style.backgroundColor = "rgba(0, 0, 0, 0)";
        reader.readAsDataURL(file); 
    } 
    else {
        previewPost.style.backgroundImage = "";
        previewPost.style.backgroundColor = "#F7F7F7";
    }
}

function previewPageImg() {
    const file = document.querySelectorAll(".upload-post-img__file-field")[0].files[0];
    const preview = document.querySelectorAll(".upload-post-img__button")[0];
    const previewPost = document.querySelector(".page-area__img");
    const reader = new FileReader();
    
    if (file) {
        const replacableInscription = preview.parentNode.nextSibling.nextSibling;
        reader.onloadend = function () {
            preview.style.backgroundImage = "url(" + reader.result + ")";
            previewPost.style.backgroundImage = "url(" + reader.result + ")";
        }
        drawUploadNewAndRemove(replacableInscription, "upload-page-img__file-field");
        replacableInscription.remove();
        preview.style.border = "none";
        preview.innerHTML = "";
        preview.style.backgroundColor = "rgba(0, 0, 0, 0)";
        previewPost.style.backgroundColor = "rgba(0, 0, 0, 0)";
        reader.readAsDataURL(file);
    } 
    else {
        previewPost.style.backgroundImage = "";
        previewPost.style.backgroundColor = "#F7F7F7";
    }
}

function previewCardImg() {
    const file = document.querySelectorAll(".upload-post-img__file-field")[1].files[0];
    const preview = document.querySelectorAll(".upload-post-img__button")[1];
    const previewPost = document.querySelector(".card-area__img");
    const reader = new FileReader();
    
    if (file) {
        const replacableInscription = preview.parentNode.nextSibling.nextSibling;
        reader.onloadend = function () {
            preview.style.backgroundImage = "url(" + reader.result + ")";
            previewPost.style.backgroundImage = "url(" + reader.result + ")";
        }
        drawUploadNewAndRemove(replacableInscription, "upload-card-img__file-field");
        replacableInscription.remove();
        preview.style.border = "none";
        preview.innerHTML = "";
        preview.style.backgroundColor = "rgba(0, 0, 0, 0)";
        previewPost.style.backgroundColor = "rgba(0, 0, 0, 0)";
        reader.readAsDataURL(file);
    } 
    else {
        previewPost.style.backgroundImage = "";
        previewPost.style.backgroundColor = "#F7F7F7";
    }
}

function drawUploadNewAndRemove(el, fileID){
    const buttons = document.createElement("div");
    buttons.classList.add("add-post-form__uploaded-img-buttons");
    buttons.innerHTML = `
        <label for="` + fileID + `" class="uploaded-img-buttons__upload-new">
            <img src="/static/img/upload_image.svg" class="upload-new__img" />
            <span class="upload-new__inscription">Upload New</span>
        </label>
        <div class="auploaded-img-buttons__remove" onclick="removeImg()">
            <img src="/static/img/remove_image.svg" class="remove__img" />
            <span class="remove__inscription">Remove</span>
        </div>
    `;
    el.after(buttons);
}

function removeImg() {
    
}

function previewText() {
    setTextInputClass(textInputs);

    pageAreaTitle.innerHTML = textInputs[0].value;
    cardAreaTitle.innerHTML = textInputs[0].value;
    pageAreaSubtitle.innerHTML = textInputs[1].value;
    cardAreaSubtitle.innerHTML = textInputs[1].value;
    cardAreaAuthorName.innerHTML = textInputs[2].value;
    cardAreaPublishDate.innerHTML = textInputs[3].value;
}

function setTextInputClass(arr) {
    for (let el of arr) {
        if (el.value === "" && el.classList.contains("add-post-form__text-input_filled")) {
            el.classList.remove("add-post-form__text-input_filled");
        }
        if (el.value !== "" && !el.classList.contains("add-post-form__text-input_filled")) {
            el.classList.add("add-post-form__text-input_filled");
        }
    }
}

function printJson() {
    const authorAvatarFile = authorAvatarInput.files[0];
    const pageImgFile = pageImgInput.files[0];
    const cardImgFile = cardImgInput.files[0];
    const authorAvatar = (authorAvatarFile) ? authorAvatarFile.name : "";
    const pageImg = pageImgFile ? pageImgFile.name : "";
    const cardImg = cardImgFile ? cardImgFile.name : "";
    const formData = {
        title: textInputs[0].value,
        subtitle: textInputs[1].value,
        author_name: textInputs[2].value,
        author_avatar: authorAvatar,
        publish_date: textInputs[3].value,
        page_image: pageImg,
        card_image: cardImg
    }
    console.log(JSON.stringify(formData, null, "\t"));
}

function initEventListeners() {
    for (let i of textInputs) {
        i.addEventListener("input", previewText);
    }
    form.addEventListener("click", printJson);
}

initEventListeners();
setTextInputClass(textInputs);