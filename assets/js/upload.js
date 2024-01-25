
function submitForm() {
    document.addEventListener('htmx:afterRequest', function(event) {
        var xhr = event.detail.xhr;
        var errorMessageElement = document.getElementById('error-message');

        if (xhr.status !== 200) {
            errorMessageElement.innerHTML = xhr.responseText;

            setTimeout(function() {
                errorMessageElement.innerHTML = '';
            }, 3000);
        }
    });

    return false;
}

function toggleUpload(input) {
    var uploadButton = document.getElementById("upload-button");
    uploadButton.style.display = input.value ? "block" : "none";
}
