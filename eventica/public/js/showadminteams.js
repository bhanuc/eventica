
$.ajax({
  url: '/team/adminall',
  type: 'GET',
  success: function(data) {
  			console.log(data);
  			if(data.success && data.teams){
          window.teams =data.teams;
          var str="<thead><tr><th>name</th><th>Created By</th><th>Event</th><th>Submit Status</th><th>Approval Status</th><th>Members</th><th>Approve</th><th>Comment</th><th>Add Comment</th><th>Decline</th><th>PrintPage</th></tr></thead><tbody>";
          for (var i = data.teams.length - 1; i >= 0; i--) {
            str = str+ '<tr><td>'+data.teams[i].name+ '</td><td><a onclick=showprofile("' + data.teams[i].createdby + '")>' + data.teams[i].createdby + '</a></td><td> '+data.teams[i].event+ '</td><td> '+data.teams[i].requestmod+ '</td><td> '+data.teams[i].approved+ '</td><td> '+ data.teams[i].members+'</td><td> <a href="/team/adminsapprove?name='+ data.teams[i].name+'">Approve</a></td><td> '+data.teams[i].comments+ '</td><td><form action="/team/adminscomment" type="GET"><input type="text" name="comments"><input type="hidden" name="name" value='+data.teams[i].name+'><button type ="submit">Submit</button></form></td><td> <a href="/team/adminsdissapprove?name='+ data.teams[i].name+'">DisApprove</a></td><td> <a href="/print#'+ data.teams[i].name+'">Print</a></td></tr>';
          };
          str += '</tbody>'
          $('#teams').html(str);
  			} else {
          $('#teams').html("No teams have submitted yet");
  			}
          //success message mybe...
     }
});

function sortevent(){
$.ajax({
  url: '/team/adminevent',
  type: 'POST',
  data: JSON.stringify({'event': document.getElementById('mevent').value}),
  success: function(data) {
        console.log(data);
        if(data.success && data.teams){
                    window.teams =data.teams;
          var str="<thead><tr><th>name</th><th>Created By</th><th>Event</th><th>Submit Status</th><th>Approval Status</th><th>Members</th><th>Approve</th><th>Comment</th><th>Add Comment</th><th>Decline</th><th>PrintPage</th></tr></thead><tbody>";
          for (var i = data.teams.length - 1; i >= 0; i--) {
            str = str+ '<tr><td>'+data.teams[i].name+ '</td><td><a onclick=showprofile("' + data.teams[i].createdby + '")>' + data.teams[i].createdby + '</a></td><td> '+data.teams[i].event+ '</td><td> '+data.teams[i].requestmod+ '</td><td> '+data.teams[i].approved+ '</td><td> '+ data.teams[i].members+'</td><td> <a href="/team/adminsapprove?name='+ data.teams[i].name+'">Approve</a></td><td> '+data.teams[i].comments+ '</td><td><form action="/team/adminscomment" type="GET"><input type="text" name="comments"><input type="hidden" name="name" value='+data.teams[i].name+'><button type ="submit">Submit</button></form></td><td> <a href="/team/adminsdissapprove?name='+ data.teams[i].name+'">DisApprove</a></td><td> <a href="/print#'+ data.teams[i].name+'">Print</a></td></tr>';
          };
          str += '</tbody>'
          $('#teams').html(str);
        } else {
          $('#teams').html("No teams have submitted yet");
        }
          //success message mybe...
     }
});
}

var showprofile = function(url) {
    $.ajax({
        url: '/user/vprofile?id=' + url,
        type: 'GET',
        success: function(data) {
            console.log(data);
            if (data.success) {
                var json = data.success;
                $('#name').html(json.name);
                $('#number').html(json.number);
                $('#college').html(json.college);
                $('#email').html(json.email);
                $('#alternatenumber').html(json.alternatenumber);
                $('#ambassador').html(json.ambassador);
                $('#sex').html(json.sex);
                $('#branch').html(json.branch);
                $('#bookingid').html(json.bookingid);
                $('#year').html(json.year);
                $('#techid').html(json.techid);
                $('#modal').modal('show');
            } else {
                alert('Some error has occured. Contact support.')
            }
            //success message mybe...
        }
    });
}

var getallemails = function() {
    if (window.$team && window.$team > 1) {
        var emails = [];
        for (var i = window.$teams.length - 1; i >= 0; i--) {
            if (window.$teams[i] != '') {
                $.ajax({
                        url: '/user/tek-profile?id=' + window.$teams[i],
                        type: 'GET',
                        success: function(data) {
                            console.log(data);
                            // emails.push[data.email]
                        }
                    });
                }
            }

        }
    }


var Techinfo = function(){
 var tid = document.getElementById('techid').value;
 if (tid){
        $.ajax({
        url: '/user/tek-profile?id=' + tid,
        type: 'GET',
        success: function(data) {
            console.log(data);
            if (data.success) {
                var json = data.user;
                $('#name').html(json.name);
                $('#number').html(json.number);
                $('#college').html(json.college);
                $('#email').html(json.email);
                $('#alternatenumber').html(json.alternatenumber);
                $('#ambassador').html(json.ambassador);
                $('#sex').html(json.sex);
                $('#branch').html(json.branch);
                $('#bookingid').html(json.bookingid);
                $('#year').html(json.year);
                $('#techid').html(json.techid);
                $('#modal').modal('show');
            } else {
                alert('Some error has occured. Contact support.')
            }
            //success message mybe...
        }
    });    

 }
}

var Excel = function() {
    if (window.teams) {
        var csvRows = [];
        var teams = window.teams;
        var csvString = 'name,createdby,Event,requestmod,approved,members,comments,id,inactivesince%0A';
        for (var i = teams.length - 1; i >= 0; i--) {
            var re = /,/gi;
            var memberz = str.replace(re, teams[i].members);
            var row = teams[i].name+','+teams[i].createdby+','+teams[i].event+','+teams[i].requestmod+','+teams[i].approved+','+memberz+','+teams[i].comments+','+teams[i].id+','+teams[i].inactivesince;
            csvRows.push(row);
        };

        csvString += csvRows.join("%0A");
        var a = document.createElement('a');
        a.href = 'data:attachment/csv,' + csvString;
        a.target = '_blank';
        a.download = 'myFile.csv';

        document.body.appendChild(a);
        a.click();
    } else {
        alert('Wait you nigger! Have some patience :p. Try Again Later')
    }

}