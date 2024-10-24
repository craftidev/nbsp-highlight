function copyToClipboard() {
    const copyButton = document.getElementById('copyButton');
    const copySuccessText = copyButton.getAttribute('data-copy-success');
    const copyButtonDefaultText = copyButton.getAttribute('data-copy-default');

    const outputText = document.getElementById('output').innerText;

    navigator.clipboard.writeText(outputText).then(() => {
        copyButton.innerHTML = copySuccessText;
        copyButton.disabled = true;

        setTimeout(() => {
            copyButton.innerHTML = copyButtonDefaultText;
            copyButton.disabled = false;
        }, 2000);
    }).catch(err => {
        alert('Failed to copy text: ' + err);
    });
}
