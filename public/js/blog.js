//let namaSiswa1 ="tamaro";
//let namaSiswa2 ="Robby";
//let namaSiswa3 ="puja";

//console.log(namaSiswa1);
//console.log(namaSiswa2);
//console.log(namaSiswa3);

//array
//let namaSiswa =["tamaro", "Robby", "Puja", "Rendi"];
//console.log(namaSiswa);
//console.log(namaSiswa[0])

//object
//let nama ="Badriana";
//let alamat ="Tangerang Selatan";
//let umur =24;

//console.log(nama);
//console.log(alamat);
//console.log(umur);

//let dataPersonal1 ={
 // nama: "Badriana",
  //alamat: "Tangerang Selatan",
  //umur: 24,
//};
//console.log(dataPersonal1);

//let dataPersonal2 ={
  //nama: "Puan Maharani",
  //alamat: "Jakarta Barat",
  //umur: 48,
//};
//console.log(dataPersonal2);

//let dataPersonal3 ={
  //nama:"Joko Widodo",
  //alamat: "Jakarta Timur",
  //umur: 44,
//};
//console.log(dataPersonal3);

//let dataPersonal4 ={
  //nama:"Megawati",
  //alamat: "Jakarta Timur",
  //umur: 44,
//};
//

//console.log(dataPersonal4);

//js-blog

let dataBlog = [];

function addBlog(event) {
  event.preventDefault();

  let projectName = document.getElementById("inputTitle").value;
  let startDate = new Date(document.getElementById("startDate").value);
  let endDate = new Date(document.getElementById("endDate").value);
  let description = document.getElementById("description").value;
  let image = document.getElementById("inputImage").files;

  


  let distance = endDate - startDate; 
  console.log(distance);

  let milisecond = 1000; // milisecond
  let secondInHours = 3600; // 1 jam 3600 detik
  let hoursInDays = 24; // 1 hari 24 jam
  let daysInMonth = 30; // 30 hari sebulan
  let monthInYear = 12; // 12 bulan setahun

  let distanceDay = Math.floor(
    distance / (milisecond * secondInHours * hoursInDays)
  ); 
  let distanceMonth = Math.floor(
    distance / (milisecond * secondInHours * hoursInDays * daysInMonth)
  );
  let distanceYear = Math.floor(
    distance /
      (milisecond * secondInHours * hoursInDays * daysInMonth * monthInYear)
  );

  //untuk menampilkan gambar
  image = URL.createObjectURL(image[0]);
  console.log(image);

  let blog = {
    projectName,
    distanceDay,
    distanceMonth,
    distanceYear,
    description,
    distance,
    realTime : new Date(),
    image,
  };

  dataBlog.push(blog);
  console.log(dataBlog);

  renderBlog();
}

function renderBlog() {
  //untuk menghilangkan post yang pertama
  document.getElementById("contents").innerHTML="";

  //menambahkan post
  for (let index = 0; index < dataBlog.length; index++) {
    document.getElementById("contents").innerHTML += `
    <div class="col">
    <div class="card">
        <img src="${dataBlog[index].image}" class="card-img-top" alt="">
        <div class="card-body">
            <h5 class="card-title">${dataBlog[index].projectName}</h5>
            <div class="mb-2 fw-light">${dataBlog[index].distanceDay} Day , ${dataBlog[index].distanceMonth} Month , ${dataBlog[index].distanceYear} Year </div>
            <p class="card-text">
              ${dataBlog[index].description}
              </p>
            <div class="d-flex gap-2" style="width: 100%;">
              <button type="button" class="btn btn-edit" style="width:100%;">edit</button>
              <button type="button" class="btn btn-delete" style="width: 100%;">Delete</button>
            </div>
        </div>
      </div>
      </div>
              <div style="float: left; margin: 10px">
              <p style="font-size: 20px; color: blue">${getRealTime(dataBlog[index].realTime)}</p>
            </div>
            </div>
          </div>
    `;
  }
}

function getFullTime(time) {
  let monthName = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Desc",
  ];

  let date = time.getDate();
  let monthIndex = time.getMonth();
  let year = time.getFullYear();
  let hours = time.getHours();
  let minutes = time.getMinutes();

  if (hours <= 9) {
    hours = "0" + hours;
  } else if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `${date} ${monthName[monthIndex]} ${year} ${hours} : ${minutes} WIB`;
}

function getRealTime(time) {
  let timeNow = new Date();
  let timePost = time;

  // waktu sekarang - waktu post
  let realTime = timeNow - timePost; // hasilnya milidetik
  console.log(realTime);

  let milisecond = 1000; // milisecond
  let secondInHours = 3600; // 1 jam 3600 detik
  let hoursInDays = 24; // 1 hari 24 jam

  let realTimeDay = Math.floor(
    realTime / (milisecond * secondInHours * hoursInDays)
  ); // 1/86400000
  let realTimeHours = Math.floor(realTime / (milisecond * 60 * 60)); // 1/3600000
  let realTimeMinutes = Math.floor(realTime / (milisecond * 60)); // 1/60000
  let realTimeSecond = Math.floor(realTime / milisecond); // 1/1000

  if (realTimeDay > 0) {
    return `${realTimeDay} Day Ago`;
  } else if (realTimeHours > 0) {
    return `${realTimeHours} Hours Ago`;
  } else if (realTimeMinutes > 0) {
    return `${realTimeMinutes} Minutes Ago`;
  } else {
    return `${realTimeSecond} Seconds Ago`;
  }
}

setInterval(function () {
  renderBlog();
}, 20000);

