function deleteItem(val) {
    let token = document.getElementById('token').value;
    fetch(`/delete/${val}/${token}`, {
        method: 'POST',
    }).then((_res) => {
        window.location.href = '/';
    }).catch((err) => {
        console.log(err);
    })
}