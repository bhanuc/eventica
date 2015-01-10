$('#submitbutton').click(function() {
    var evname = $("#subject").val();
    var memberno = 0;

    if (evname == "Hoverush" || evname == "IDEAS" || evname == "Manoeuvre" || evname == "Shuffle" || evname == "Mix Bowl Quiz") {
        memberno = 5;
    } else if (evname == "FPGA" || evname == "Embedded" || evname == "electromania" || evname == "Design 3D") {
        memberno = 5;
    } else if (evname == "IORC"  || evname == "Zonals" || evname == "Be the Tycoon" || evname == "29 States" || evname == "crypto" || evname == "Scimatex" || evname == "Stocksim") {
        memberno = 1;
    } else if (evname == "Impulse" || evname == "Hackathon" || evname == "Iarc") {
        memberno = 4;
    } else if (evname == "Electrade") {
        memberno = 4;
    } else if (evname == "Skysparks" || evname == "Soccon" || evname == "Cruise Control" || evname == "Multirover" || evname == "concatenate" || evname == "Bridge Design Challenge") {
        memberno = 6;
    } else if (evname == "Crime Run" || evname == "AISRC" || evname == "Scientoon") {
        memberno = 4;
    } else if (evname == "TGP") {
        memberno = 8;
    } else if (evname == "IOPC" || evname == "What's Up" || evname == "Astro Quiz" || evname == "Innovation in Manufactruning Processes" || evname == "Battlefield" || evname == "Do your due" || evname == "IHPC" || evname == "Marketing Villa" || evname == "Chaos" || evname == "Battlecity" || evname == "Finquest" || evname == "Mark-Ops") {
        memberno = 3;
    } else if (evname == "Wild soccer") {
        memberno = 10;
    } else if (evname == "Astro Treasure") {
        memberno = 7;
    }



    var memberscontact = "";
    for (var i = 0; i < memberno; i++) {
        if (i == 0) {
            memberscontact = $("#members" + (i + 1)).val();
        } else {
            if ($("#members" + (i + 1)).val() != '') {
                memberscontact = memberscontact + ',' + $("#members" + (i + 1)).val();
            }
        }
    }
    console.log(memberscontact);
    var l = memberscontact.split(',');
    memberscontact = _.uniq(l, true).toString();

    $.ajax({
        url: '/team/create',
        type: 'POST',
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            'name': $('#teamname').val(),
            'event': $('#subject').val(),
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
        $('#submitbutton').show();
    } else if (evname == "FPGA" || evname == "Embedded") {
        memberno = 5;
        comment = "Only for Undergraduate students";
        $('#submitbutton').show();
    } else if (evname == "Battlefield") {
        memberno = 3;
        comment = "Each team will have 2-3 members. Registration for this event will open on 15th January and will close on 10th February.";
        $('#submitbutton').hide();
    } else if (evname == "Marketing Villa") {
        memberno = 3;
        comment = "Each team will have 2-3 members. Registration for this event will close on 31st January.";
        $('#submitbutton').show();
    } else if (evname == "Do your due") {
        memberno = 3;
        comment = "Each team will have 2-3 members. Registration for this event will open on 15th January and will close on 4th February.";
        $('#submitbutton').hide();
    } else if (evname == "IORC" || evname == "29 States" || evname == "Stocksim") {
        memberno = 1;
        comment = "";
        $('#submitbutton').show();
    } else if (evname == "Be the Tycoon") {
        memberno = 1;
        comment = "Registration for this event will open on 25th January and will close on 15th February.";
        $('#submitbutton').hide();
    } else if (evname == "electromania") {
        memberno = 5;
        comment = "Only for Undergraduate students";
        $('#submitbutton').show();
    } else if (evname == "Impulse" || evname == "Hackathon" || evname == "Iarc") {
        memberno = 4;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Electrade") {
        memberno = 4;
        $('#submitbutton').show();
        comment = "Only for Undergraduate students";
    } else if (evname == "Skysparks" || evname == "Soccon" || evname == "Cruise Control" || evname == "Multirover" || evname == "concatenate") {
        memberno = 6;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Crime Run") {
        memberno = 4;
        $('#submitbutton').show();
        comment = "No personal devices or eqipments are allowed";
    } else if (evname == "TGP") {
        memberno = 8;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "IOPC" || evname == "IHPC" || evname == "Chaos" || evname == "Battlecity") {
        memberno = 3;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Wild soccer") {
        memberno = 10;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "AISRC") {
        memberno = 4;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Finquest") {
        memberno = 3;
        $('#submitbutton').show();
        comment = "2-3 Members are allowed.Registration will start from 1st January and closes on 24th January.";
    } else if (evname == "Mark-Ops") {
        memberno = 3;
        $('#submitbutton').show();
        comment = "Each team will have 2-3 members.Registration will start from 1st February and closes on 24th February.";
    } else if (evname == "crypto") {
        memberno = 1;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "What's Up") {
        memberno = 3;
        comment = "Each team will have 1-3 members.";
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Astro Quiz") {
        memberno = 3;
        comment = "Each team will have 1-3 members.";
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Astro Treasure") {
        comment = "Each team will have 3-5 members.";
        memberno = 5;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Scimatex" || evname == "Zonals") {
        memberno = 1;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Mix Bowl Quiz") {
        memberno = 5;
        comment = "Each team will have 1-5 members.";
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Innovation in Manufactruning Processes") {
        memberno = 3;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Scientoon") {
        memberno = 4;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Design 3D") {
        memberno = 5;
        $('#submitbutton').show();
        comment = "";
    } else if (evname == "Bridge Design Challenge") {
        memberno = 6;
        $('#submitbutton').show();
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