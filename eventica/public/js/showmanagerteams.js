
$.ajax({
  url: '/team/event',
  type: 'GET',
  success: function(data) {
  			console.log(data);
  			if(data.success && data.teams){
          var str="<thead><tr><th>name</th><th>Created By</th><th>Event</th><th>Submit Status</th><th>Approval Status</th><th>Members</th><th>Approve</th><th>Comment</th><th>Add Comment</th><th>Decline</th><th>PrintPage</th></tr></thead><tbody>";
          for (var i = data.teams.length - 1; i >= 0; i--) {
            str = str+ '<tr><td>'+data.teams[i].name+ '</td><td><a href="/user/vprofile?id='+data.teams[i].createdby+ '">'+data.teams[i].createdby+'</a></td><td> '+data.teams[i].event+ '</td><td> '+data.teams[i].requestmod+ '</td><td> '+data.teams[i].approved+ '</td><td> '+ data.teams[i].members+'</td><td> <a href="/team/adminsapprove?name='+ data.teams[i].name+'">Approve</a></td><td> '+data.teams[i].comments+ '</td><td><form action="/team/adminscomment" type="GET"><input type="text" name="comments"><input type="hidden" name="name" value='+data.teams[i].name+'><button type ="submit">Submit</button></form></td><td> <a href="/team/adminsdissapprove?name='+ data.teams[i].name+'">DisApprove</a></td><td> <a href="/print#'+ data.teams[i].name+'">Print</a></td></tr>';
          };
          str += '</tbody>'
          $('#teams').html(str);
  			} else {
          $('#teams').html("No teams have submitted yet");
  			}
          //success message mybe...
     }
});

