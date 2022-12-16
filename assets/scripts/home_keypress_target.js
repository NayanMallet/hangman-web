document.addEventListener('keypress', (event) => {
  const keyName = event.key;
  if (keyName === 'Enter' && (document.querySelector('#user_name_input_js').value !== "")) {
    document.getElementById("home_form").submit()
  } else if (keyName === 'Backspace') {
    document.querySelector('#user_name_input_js').value = document.querySelector('#user_name_input_js').value.slice(0, document.querySelector('#user_name_input_js').value.length-1)
  } else {
    document.querySelector('#user_name_input_js').value += keyName
  }
});

