tableType = document.getElementById("list-type");
list = document.getElementById("list");
listAPI = tableType.querySelector(".list-api");
detailBox = document.getElementById("detail-box");
detail = document.getElementById("detail");
closeDetailBtn = document.getElementById("close-btn");
var types;
const arrAPIs = []; // to check status api

tbody = listAPI.getElementsByTagName("tbody");

let api = `https://be-monitor.gear5.guru/system/list-api`;
let datajson;
fetch(api)
  .then((response) => response.json())
  .then((data) => {
    datajson = data.data.apis;
    for (let i = 0; i < datajson.length; i++) {
      api = datajson[i];
      arrAPIs.push(api.url);
      let tmpAPI = `
          <tr class="api-ele" onclick=detailAPI(${api.id})>
            <td class="c-table__cell">${api.serverName}</td>
            <td class="c-table__cell">${api.server}</td>
            <td class="c-table__cell">${api.type}</td>
            <td class="c-table__cell last-update"></td>
            <td class="c-table__cell update"></td>
            <td class="c-table__cell insert"></td>
            <td class="c-table__cell duration">${api.duration}</td>
            <td class="c-table__cell status"></td>
          </tr>
          `;
      tbody[0].innerHTML += tmpAPI;
    }

    statusTags = tbody[0].getElementsByClassName("status");
    lastUpdateTags = tbody[0].getElementsByClassName("last-update");
    updateTags = tbody[0].getElementsByClassName("update");
    insertTags = tbody[0].getElementsByClassName("insert");
    listTr = tbody[0].getElementsByClassName("api-ele");

    (function checkStatus() {
      for (let i = 0; i < arrAPIs.length; i++) {
        fetch(arrAPIs[i])
          .then((response) => response.json())
          .then((response) => {
            if (arrAPIs[i]==="https://dev-be.client.gear5.guru") {
              console.log(response)
            }
            if (response.code === "200" || response.code === "B.200") {
              let time = new Date()
              // console.log(Date.parse(response.data.LastUpdateTime))
              // console.log("time")
              // console.log(time.getTime())
              // console.log(Date.parse(response.data.LastUpdateTime))

              // ms = parseInt(Date.parse(time.getTime())/1000 - Date.parse(response.data.LastUpdateTime)/1000)
              statusTags[i].innerHTML = `<span style='color: #1abc9c; font-size: 15px;'> Active </td>`;
              lastUpdateTags[i].innerHTML = 
              // `<span style='font-size: 15px;'>${response.data.LastUpdateTime}</td>`;
              // `<span style='font-size: 15px;'>${msToTime(ms)}</td>`;
              `<span style='font-size: 15px;'>${response.data.LastUpdateTime.slice(11,16) + " / " + response.data.LastUpdateTime.slice(5,10)}</td>`;
              updateTags[i].innerHTML = 
              `<span style='font-size: 15px;'>${Intl.NumberFormat().format(response.data.Update)}</td>`;
              insertTags[i].innerHTML = 
              `<span style='font-size: 15px;'>${Intl.NumberFormat().format(response.data.Insert)}</td>`;
            } else {
              statusTags[i].innerHTML = `<span style='color: #e74c3c; font-size: 15px;'> Error </td>`;
            }
          });
      }
      setTimeout(checkStatus, 30000);
    })();
  });

function removeAPI(i) {
  fetch(`https://be-monitor.gear5.guru/system/api?id=${i}`, {
    method: "DELETE",
  })
    .then((response) => response.json())
    .then((response) => {
      if (response.code === "B.200") {
        alert("Delete API successfully!");
        location.reload();
      } else {
        alert("Delete API failed!");
      }
    });
}


function msToTime(ms) {
  let seconds = (ms / 1000).toFixed(1);
  let minutes = (ms / (1000 * 60)).toFixed(1);
  let hours = (ms / (1000 * 60 * 60)).toFixed(1);
  let days = (ms / (1000 * 60 * 60 * 24)).toFixed(1);
  if (seconds < 60) return seconds + " Sec";
  else if (minutes < 60) return minutes + " Min";
  else if (hours < 24) return hours + " Hrs";
  else return days + " Days"
}

