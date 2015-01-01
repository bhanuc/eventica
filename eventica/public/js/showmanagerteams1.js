
$.ajax({
  url: '/team/event',
  type: 'POST',
  success: function(data) {
        console.log(data);
        if(data.success && data.teams){
          var str="<thead><tr><th>name</th><th>Created By</th><th>Event</th><th>Submit Status</th><th>Approval Status</th><th>Members</th><th>Approve</th><th>Comment</th><th>Add Comment</th><th>PrintPage</th></tr></thead><tbody>";
          for (var i = data.teams.length - 1; i >= 0; i--) {
            str = str+ '<tr><td>'+data.teams[i].name+ '</td><td><a onclick=showprofile('+data.teams[i].createdby+ ')">'+data.teams[i].createdby+'</a></td><td> '+data.teams[i].event+ '</td><td> '+data.teams[i].requestmod+ '</td><td> '+data.teams[i].approved+ '</td><td> '+ data.teams[i].members+'</td><td> '+data.teams[i].comments+ '</td><td><form action="/team/managerscomment" type="GET"><input type="text" name="comments"><input type="hidden" name="name" value='+data.teams[i].name+'><button type ="submit">Submit</button></form></td><td> <a href="/team/adminsdissapprove?name='+ data.teams[i].name+'">DisApprove</a></td><td> <a href="/print#'+ data.teams[i].name+'">Print</a></td></tr>';
          };
          str += '</tbody>'
          $('#teams').html(str);
        } else {
          $('#teams').html("No teams have submitted yet");
        }
        var array = [];
        var count = 0;
        array[count]="";
        var flag = false;
        for(var counter = 0; counter<data.teams.length; counter++){
          var str = data.teams[counter].members;
          for(var i = 0; i<str.length; i++){
            parseInt(str[i])
            console.log(str[i]);
          }
          for(var i = 0; i<str.length; i++){
            if(!isNaN(str[i])){
              array[count] = array[count]+str[i];
              flag = false;
            }
            else{
              if(!flag){
                flag = true;
                count++;
                array[count] = "";
              }
            }
          }
          if(!flag){
            flag = true;
            count++;
            array[count] = "";
          }
        }
        window.$team = array;

          //success message mybe...
     }
});

var showprofile =function(url){
        $.ajax({
  url: '/user/mprofile?id='+url,
  type: 'POST',
  success: function(data) {
        console.log(data);
        if(data.success ){
          var json = data.success;
        $('#name').innerHTML = json.name;
        $('#number').innerHTML = json.number;
        $('#college').innerHTML = json.college;
        $('#email').innerHTML = json.email;
        $('#alternatenumber').innerHTML = json.alternatenumber;
        $('#ambassador').innerHTML = json.ambassador;
        $('#sex').innerHTML = json.sex;
        $('#branch').innerHTML = json.branch;
        $('#bookingid').innerHTML = json.bookingid;
        $('#year').innerHTML = json.year;
        $('#techid').innerHTML = json.techid;
        $('#modal').modal('show');
        } else {
          alert('Some error has occured. Contact support.')
        }
          //success message mybe...
     }
    }); 
 }
