const baseURL = 'https://eb05-2001-fb1-13f-10b2-515-73a7-e299-5eae.ngrok.io/course';
fetch(baseURL)
    .then(resp => resp.json())
    .then(data => displayData(data))
    .catch(err => console.log('error: ' +err));
function displayData(data){
    //document.querySelector("pre").innerHTML = JSON.stringify(data,null,2);
var mainContainer = document.getElementById("myData");
    for(var i = 0 ; i < data.length;i++){
        var div = document.createElement("div");
        div.innerHTML = "CourseID: " + data[i].ID + ' ' + data[i].Name + ' '+ data[i].Price + ' '+ data[i].Instructor ;
        mainContainer.appendChild(div);
    }
}