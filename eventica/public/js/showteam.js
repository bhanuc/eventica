
$.ajax({
  url: '/team/all',
  type: 'GET',
  success: function(data) {
  			console.log(data);
  			if(data.success){
          var str="<thead><tr><th>name</th><th>Game</th><th>Submitted Team for approval</th><th>Members</th><th>Submit your team</th><th>Approval Status</th><th>Comments</th></tr></thead><tbody>";
          for (var i = data.teams.length - 1; i >= 0; i--) {
            str = str+ '<tr><td>'+data.teams[i].name+ '</td><td> '+data.teams[i].requestmod+ '</td><td> '+ data.teams[i].members+'</td><td> <a href="/team/register?name='+ data.teams[i].name+'&createdby='+data.teams[i].createdby+'">Submit</a></td><td> '+data.teams[i].approved+ '</td><td> '+data.teams[i].comments+ '</td></tr>';
          };
          str += '</tbody>'
          $('#teams').html(str);
  			} else {
  			}
          //success message mybe...
     }
});
