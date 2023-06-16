//console.log("hello student");
//alert("good morning");
//variabel
//var
//let
//const
//bisa di deklarasikan ulang dan di ambil value nya saja
//var gelas="air putih";
//console.log("gelas");
function submitData() {
  let name= document.getElementById("name").value;
  let email= document.getElementById("email").value;
  let phone= document.getElementById("Phone").value;
  let subject= document.getElementById("subject").value;
  let message= document.getElementById("message").value;



  if (name=="") {
    return alert("nama tidak boleh kosong");
  }
  if (email=="") {
    return alert("email tidak boleh kosong");
  }
  if (phone=="") {
    return alert("no hp tidak boleh kosong");
  }
  if (subject=="") {
    return alert("subject harus di pilih");
  }
  if (message=="") {
    return alert("pesan harus di isi");
  }
let emailreceiver="payumiswari@gmail.com";
let a = document.createElement("a");
a.href=`mailto:${emailreceiver}?subject=${subject}&body=halo, nama saya ${name}, ${message}.silahkan kontak saya di nomor ${phone}`;
a.click();

console.log(name);
console.log(email);
console.log(phone);
console.log(subject);
console.log(message);

let emailer = {
  name,
  email,
  phone,
  subject,
  message,
};
console.log(emailer);

}