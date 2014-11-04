
$.ajax({
  url: '/team/view?name='+document.location.hash.split('#')[1],
  type: 'GET',
  success: function(data) {
  			console.log(data);
  			if(data.success){
          var str="Team Name :"+data.success.name+'<br>';
          str += "Game :"+data.success.game+'<br>';
          str += "Members :"+data.success.members+'<br>';
          str += "Approved Status :"+data.success.approved+'<br>';
          $('#team').html(str);
          user(data.success.createdby);
  			} else {
          $('#teams').html("No teams Found !!!");
  			}
          //success message mybe...
     }
});

function user(id){	
$.ajax({
  url: 'user/vprofile?id='+id,
  type: 'GET',
  success: function(data) {
  			console.log(data);
  			if(data.success ){
 var str="Applicant Name :"+data.success.name+'<br>';
          str += "College Name :"+data.success.college+'<br>';
          str += "Email :"+data.success.email+'<br>';
          str += "Number :"+data.success.number+'<br>';
          str += "Alternate Number :"+data.success.alternatenumber+'<br>';
          str += "Sport Incharge :"+data.success.sportincharge+'<br>';
          $('#user').html(str);       
  			} else {
          $('#user').html("No User have submitted yet");
  			}
          //success message mybe...
     }
});

}