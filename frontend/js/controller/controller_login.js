import { postWithToken } from "https://jscroot.github.io/api/croot.js";
import { PostLogin, ResponseLogin } from "../config/config.js";
import { UrlLogin } from "../template/template.js";
import { token } from '../template/template.js';

document.addEventListener("DOMContentLoaded", function() {
  const form = document.getElementById("login-form");
  form.addEventListener("submit", function(event) {
    event.preventDefault();
    let data = PostLogin();
    postWithToken(UrlLogin, 'Authorization', 'Bearer ' + token, data, ResponseLogin);
  });
});