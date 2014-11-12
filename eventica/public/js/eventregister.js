$('#submitbutton').click(function() {
  var memberno = 0;
 if($("#subject").val()=="Skysparks" || $("#subject").val()=="Multirover" || $("#subject").val()=="Cruise Control"  || $("#subject").val()=="Hoverush") {   
        memberno = 5;
    } else if($("#subject").val()=="Embedded") {   
        memberno = 5;
    } else if($("#subject").val()=="IORC") {   
        memberno = 1;
    } else if($("#subject").val()=="Electromania") {   
        memberno = 4;
    } else if($("#subject").val() == "FPGA" || ("#subject").val() == "Impulse") {   
        memberno = 4;
    } else if($("#subject").val()=="Electrade") {   
        memberno = 3;
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
    var evname = $("#subject").val();
    var memberno = 0;
    var comment = "";
    console.log($("#subject").val());
    if($("#subject").val()=="Skysparks" || $("#subject").val()=="Multirover" || $("#subject").val()=="Cruise Control"  || $("#subject").val()=="Hoverush") {   
        memberno = 5;
        comment = "";
    } else if($("#subject").val()=="Embedded") {   
        memberno = 5;
        comment = "Only for Undergraduate students";
    } else if($("#subject").val()=="IORC") {   
        memberno = 1;
        comment = "";
    } else if($("#subject").val()=="Electromania") {   
        memberno = 4;
        comment = "Only for Undergraduate students";
    } else if($("#subject").val() == "FPGA" || ("#subject").val() == "Impulse") {   
        memberno = 4;
        comment = "";
    } else if($("#subject").val()=="Electrade") {   
        memberno = 3;
        comment = "Only for Undergraduate students";
    } 

    $("#comment").html(comment) ;

     for(var i=0;i<memberno;i++)
    {
        /*var table = document.getElementById("mytable");
        var row = table.insertRow(i);
        var cell1 = row.insertCell(0);
        var cell2 = row.insertCell(1);*/
        //cell1.innerHTML = "NEW CELL1";
       // cell2.innerHTML = "NEW CELL2";
     //  if(i==0) {
       // $('#sport').append('<tr><td width="150px">Captain</td><td><input type="text" id="members'+(i+1)+'" class="demo-default" name="members" required></td><td width="180px">Contact Number</td><td><input type="number" id="connum'+(i+1)+'" class="demo-default" name="connum" required></td> </tr>');

       //}
       //else if(i==1) {
       // $('#sport').append('<tr><td width="150px">Vice Captain</td><td><input type="text" id="members'+(i+1)+'" class="demo-default" name="members" required></td><td width="180px">Contact Number</td><td><input type="number" id="connum'+(i+1)+'" class="demo-default" name="connum" required></td> </tr>');
       //}
       //else {
        $('#sport').append('<tr><td width="150px">Member '+(i+1)+'</td><td><input type="text" id="members'+(i+1)+'" class="demo-default" name="members"></td><td width="180px">Contact Number</td><td><input type="number" id="connum'+(i+1)+'" class="demo-default" name="connum"></td> </tr>');
   // }
    }
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