const handleSubmit = async (event) => {
  event.preventDefault();
  const target = event.target;
  console.log(target);
  const data = new FormData(target);
  const res = await fetch(target.action, {
    method: target.method,
    body: data
  }).then(res => res.text());

  if (res === "OK") {
    target.getElementsByTagName('label')[0].innerText = "File is uploaded"
  }
}

for (let el of document.getElementsByTagName('form')) {
  el.onsubmit = handleSubmit;
}
