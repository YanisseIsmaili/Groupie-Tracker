fetch('https://groupietrackers.herokuapp.com/api/locations')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error(error));
