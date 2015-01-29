$('#submitbutton').one("click", function() {
    createteam();
});

var zonals = {'Amlapuram': { 'date': '28/1//2015','institute':'Bonam Venkata Chalamayya Institute of Technogy and Science'}, 'Bhopal': { 'date': '12/2/2015','institute':'Oriental College of Technology'}, 'Bhubneshwar': { 'date': '18/1/2015','institute':'Kalinga Institute of Industrial Technology '}, 'Faridabad': { 'date': '30/1//2015','institute': 'YMCA institute of science and Technology'}, 'Gwalior': { 'date': '27/2/2015','institute':'ITM UNiversity (Jivaji in case ITM not confirmed)'}, 'Indore': { 'date': '30/1//2015','institute':'Indore Institute of Science and Technology'},'Jaipur': {'date':'11/2/2015', 'institute':'Amity University'}, 'Jaipur(c)': { 'date': '12/3/2015','institute':'Vivekananda Institute Of Technology'}, 'Jhansi': { 'date': '8/3/2015','institute':'SR Group of institutions'}, 'Lucknow': { 'date': '25/2/2015','institute':'Babu Banarasi Das University '}, 'Meerut': { 'date': '4/3/2015','institute':'IIMT group of colleges'}, 'Mumbai': { 'date': '17/1/2015','institute': 'Thakur college of Engineering and Technology'}, 'Nagpur': { 'date': '1/3/2015','institute':'Rajeev Gandhi College of Engineering and Research'}, 'Noida': { 'date': '23/2/2015','institute':'IEC University'}, 'Raipur': { 'date': '21/2/2015','institute':'MATS University RAIPUR'}, 'Solan': { 'date': '12/3/2015','institute':'IEC university'}};

var workshops = {'IPDC':{'no':1, 'comment':""},'Android':{'no':1, 'comment':""},'Ethical Hacking':{'no':1, 'comment':""},"Rubik's CUbe":{'no':1, 'comment':""}};

var createteam = function() {
    var evname = $("#subject").val();
    if( workshops.hasOwnProperty(evname)){
        var memberno = workshops[evname]['no'];
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
                    $('#submitbutton').one("click", function() {
    createteam();
});
            toastr.error("Please Enter team name in correct format");
        } else {
            var k = $('#subject').val()+"-"+$('#zcities').val();
            $.ajax({
                url: '/team/wcreate',
                type: 'POST',
                contentType: 'application/json; charset=utf-8',
                data: JSON.stringify({
                    'name': $('#teamname').val(),
                    'event': k,
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
        toastr.error("Select a Workshop");
        $('#submitbutton').one("click", function() {
    createteam();
});
    }
};


$("#subject").change(function() {
    $("#sport").empty();
    var evname = $("#subject").val();
    var memberno = workshops[evname]['no'];
    var comment = workshops[evname]["comment"];
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