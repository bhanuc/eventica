$('#submitbutton').click(function() {
  if($("#subject").val()=="Football" || $("#subject").val()=="Cricket"){
        memberno=15;
      } else if($("#subject").val()=="Table Tennis (B)" || $("#subject").val()=="Table Tennis (G)"){
        memberno=2;
      } else {
        memberno=0;
      }
  var memberscontact = "";
  for(var i=0;i<memberno;i++){
     memberscontact = memberscontact + $("#members"+(i+1)).val() + ':'+ $("#connum"+(i+1)).val() + ',';
   }
$.ajax({
  url: '/team/create',
  type: 'POST',
  contentType: 'application/json; charset=utf-8',
  data: JSON.stringify({'name': $('#teamname').val(), 'sport': $('#subject').val(), 'gender': $('#gender').val(), 'members': memberscontact }),
  success: function(data) {
if(data.success ){
  				toastr.success("Team created");
  				//document.location = host+"app"
  			} else {
          toastr.error( data.flashes.Error.message);
  			}
     }
});
});
 
$( "#subject" ).change(function () {
    var sportno = $("#subject").val();
    console.log(sportno);
});

/** checking should be done not in the callback the on click      

  console.log(data);
        if($("#members1").val().length == 0) {
          toastr.error("Name of Captain required")
        }
        else if($("#members2").val().length == 0) {
          toastr.error("Name of Vice-Captain required")
        }
        
        else if($("#connum1").val().length != 10 && $("#connum1").val().length != 11 ) {
          toastr.error("Provide a valid contact no. of Captain")
        }
         else if($("#connum2").val().length != 10 && $("#connum2").val().length != 11 ) {
          toastr.error("Provide a valid contact no. of Vice-Captain")
        }
        else 
        **/