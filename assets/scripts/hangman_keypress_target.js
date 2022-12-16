document.addEventListener('keypress', (event) => {
    const keyName = event.key;
    if (keyName === 'Enter' && (document.querySelector('#letter_input').value !== "")) {
        document.getElementById("input-form").submit()
    } else if (keyName === 'Backspace') {
        document.querySelector('#letter_input').value = document.querySelector('#letter_input').value.slice(0, document.querySelector('#letter_input').value.length-1)
    } else {
        document.querySelector('#letter_input').value += keyName
    }
});

