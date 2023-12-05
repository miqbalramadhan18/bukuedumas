import { deleteCookie } from 'https://jscroot.github.io/cookie/croot.js';

function logout() {
  deleteCookie('Login');
  window.location.href = '/';
}

document.getElementById('logoutButton').addEventListener('click', logout);