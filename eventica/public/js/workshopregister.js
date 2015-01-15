$('#submitbutton').one("click", function() {
    createteam();
});

var workshopname = { '3-D CAD & Printing': { 'no': 1, "comment": ""},  'Advanced AutoMobile Mechanics and Technology': { 'no': 1, "comment": ""}, 'Android Application Development': { 'no': 1, "comment": ""}, 'Application Forensics': { 'no': 1, "comment": ""}, 'Automobile Mechanics & IC Engine': { 'no': 1, "comment": ""}, 'Automobile Sketching': { 'no': 1, "comment": ""}, 'CAN-SAT Satellite Designing': { 'no': 6, "comment": ""}, 'Cloud Computing': { 'no': 1, "comment": ""}, 'Digital VLSI Design': { 'no': 1, "comment": ""}, 'Embedded Systems': { 'no': 1, "comment": ""}, 'Ethical Hacking & Cyber Forensics': { 'no': 1, "comment": ""}, 'Mobile Hacking': { 'no': 1, "comment": ""}, 'Multi-Touch and Augmented Reality': { 'no': 1, "comment": ""}, 'Propeller Clock': { 'no': 4, "comment": ""}, 'Python': { 'no': 1, "comment": ""}, 'Quadcoptor': { 'no': 6, "comment": ""}, 'Raspberry Pi': { 'no': 3, "comment": ""}, 'Share Market and Investment': { 'no': 1, "comment": ""}, 'Speech Robo': { 'no': 4, "comment": ""}, 'Touch Screen Robotics': { 'no': 4, "comment": ""}, 'Whatshack': { 'no': 1, "comment": ""}
    };

var createteam = function() {
    var evname = $("#subject").val();
    if( workshopname.hasOwnProperty(evname)){
        var memberno = workshopname[evname]['no'];
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

        var teamname = $('#teamname').val();
        var rezcheck = /[a-zA-Z0-9]+/
        if (rezcheck.exec(teamname)[0] != teamname) {
            toastr.error("Please Enter team name in correct format");
        } else {
            $.ajax({
                url: '/workshop/create',
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
                    $('#submitbutton').one("click", function() {
                        createteam();
                    });
                }
            });
        }
    } else {
        console.log('Go Home .. The Developer was drunk :p')
    }
};

$("#subject").change(function() {
    $("#sport").empty();
    var evname = $("#subject").val();
    var memberno = workshopname[evname]['no'];
    var comment = workshopname[evname]["comment"];
    $('#submitbutton').show();
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