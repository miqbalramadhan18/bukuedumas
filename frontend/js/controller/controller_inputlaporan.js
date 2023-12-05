// masih error

import { postWithToken } from "https://jscroot.github.io/api/croot.js";
import { GetDataReport } from "../config/config.js";
import { UrlSubmitReport, token } from "../template/template.js";
import { ResponsePost } from "../config/config.js";

document.addEventListener("DOMContentLoaded", function () {
  const form = document.querySelector("form"); // Sesuaikan dengan struktur HTML Anda
  form.addEventListener("submit", function (event) {
    event.preventDefault();

    const data = GetDataReport();

    // Menggunakan fungsi postWithToken untuk mengirim data ke API dengan token
    postWithToken(UrlSubmitReport, 'Authorization', 'Bearer ' + token, data, ResponsePost);
  });
});
