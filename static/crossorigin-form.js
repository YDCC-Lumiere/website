const handleSubmit = async (event) => {
  event.preventDefault();
  const target = event.target;
  console.log(target);
  const data = new FormData(target);
  const res = await fetch(target.action, {
    method: target.method,
    body: data
  }).then(res => res.json());

  target.getElementsByTagName('label')[0].innerText = ""
  target.getElementsByTagName('label')[1].innerText = ""

  if (res.success) {
    target.getElementsByTagName('label')[0].innerText = res.msg
  }
  else {
    target.getElementsByTagName('label')[1].innerText = res.msg
  }
}

for (let el of document.getElementsByTagName('form')) {
  el.onsubmit = handleSubmit;
}
