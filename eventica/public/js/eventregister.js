$('#submitbutton').click(function() {
    var evname = $("#subject").val();
    var memberno = 0;

    if (evname == "Hoverush" ||  evname == "IDEAS" ||  evname == "Manoeuvre" || evname == "Shuffle") {
        memberno = 5;
    } else if (evname == "FPGA" || evname == "Embedded" || evname == "electromania") {
        memberno = 5;
    } else if (evname == "IORC" || evname == "Be the Tycoon" || evname == "29 States" ) {
        memberno = 1;
    } else if (evname == "Impulse" || evname == "Hackathon" || evname == "Iarc") {
        memberno = 4;
    } else if (evname == "Electrade") {
        memberno = 4;
    } else if (evname == "Skysparks" || evname == "Soccon" || evname == "Cruise Control" || evname == "Multirover" || evname == "concatenate") {
        memberno = 6;
    } else if (evname == "Crime Run") {
        memberno = 4;
    } else if (evname == "TGP") {
        memberno = 8;
    } else if (evname == "IOPC" || evname == "Battlefield" || evname == "Do your due" ||  evname == "IHPC" || evname == "Marketing Villa" || evname == "Chaos" || evname == "Battlecity") {
        memberno = 3;
    } else if (evname == "Wild soccer") {
        memberno = 10;
    }
    var memberscontact = "";
    for (var i = 0; i < memberno; i++) {
        if (i == 0) {
            memberscontact = $("#members" + (i + 1)).val();
        } else {
            memberscontact = memberscontact + ',' + $("#members" + (i + 1)).val();
        }
    }
    $.ajax({
        url: '/team/create',
        type: 'POST',
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            'name': $('#teamname').val(),
            'sport': $('#subject').val(),
            'gender': $('#gender').val(),
            'members': memberscontact
        }),
        success: function(data) {
            if (data.success) {
                toastr.success("Team created");
                //document.location = host+"app"
            } else {
                toastr.error(data.flashes.Error.message);
            }
        }
    });
});

$("#subject").change(function() {
    $("#sport").empty();
    var evname = $("#subject").val();
    var memberno = 0;
    var comment = "";

    if (evname == "Hoverush" || evname == "IDEAS" || evname == "Manoeuvre" || evname == "Shuffle") {
        memberno = 5;
        comment = "";
    } else if (evname == "FPGA" || evname == "Embedded") {
        memberno = 5;
        comment = "Only for Undergraduate students";
    } else if (evname == "Marketing Villa" || evname == "Do your due" || evname == "Battlefield") {
        memberno = 3;
        comment = "2 Members are also allowed";
    } else if (evname == "IORC" || evname == "29 States") {
        memberno = 1;
        comment = "";
    } else if (evname == "Be the Tycoon" ) {
        memberno = 1;
        comment = "";
    } else if (evname == "electromania") {
        memberno = 5;
        comment = "Only for Undergraduate students";
    } else if (evname == "Impulse" || evname == "Hackathon" || evname == "Iarc") {
        memberno = 4;
        comment = "";
    } else if (evname == "Electrade") {
        memberno = 4;
        comment = "Only for Undergraduate students";
    } else if (evname == "Skysparks" || evname == "Soccon" || evname == "Cruise Control" || evname == "Multirover" || evname == "concatenate") {
        memberno = 6;
        comment = "";
    } else if (evname == "Crime Run") {
        memberno = 4;
        comment = "No personal devices or eqipments are allowed";
    } else if (evname == "TGP") {
        memberno = 8;
        comment = "";
    } else if (evname == "IOPC" || evname == "IHPC" || evname == "Chaos" || evname == "Battlecity") {
        memberno = 3;
        comment = "";
    } else if (evname == "Wild soccer") {
        memberno = 10;
        comment = "";
    }

    $("#comment").html(comment);

    for (var i = 0; i < memberno; i++) {
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
        $('#sport').append('<tr><td width="150px">Member ' + (i + 1) + '</td><td><input type="text" id="members' + (i + 1) + '" class="demo-default" name="members"></td></tr>');
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