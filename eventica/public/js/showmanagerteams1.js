$.ajax({
    url: '/team/event',
    type: 'POST',
    success: function(data) {
        console.log(data);
        if (data.success && data.teams) {
            window.teams = data.teams;
            $('#nots').html(data.teams.length);
            var str = "<thead><tr><th>name</th><th>Created By</th><th>Event</th><th>Submit Status</th><th>Approval Status</th><th>Members</th><th>Approve</th><th>Comment</th><th>Add Comment</th><th>PrintPage</th></tr></thead><tbody>";
            for (var i = data.teams.length - 1; i >= 0; i--) {
                str = str + '<tr><td>' + data.teams[i].name + '</td><td><a onclick=showprofile("' + data.teams[i].createdby + '")>' + data.teams[i].createdby + '</a></td><td> ' + data.teams[i].event + '</td><td> ' + data.teams[i].requestmod + '</td><td> ' + data.teams[i].approved + '</td><td onclick=showprofile(' + data.teams[i].createdby + ')> ' + data.teams[i].members + '</td><td> <a href="/team/managersapprove?name=' + data.teams[i].name + '">Approve</a></td><td> ' + data.teams[i].comments + '</td><td><form action="/team/managerscomment" type="GET"><input type="text" name="comments"><input type="hidden" name="name" value=' + data.teams[i].name + '><button type ="submit">Submit</button></form></td><td> <td> <a href="/print#' + data.teams[i].name + '">Print</a></td></tr>';
            };
            str += '</tbody>'
            $('#teams').html(str);
        } else {
            $('#teams').html("No teams have submitted yet");
        }

        var array = [];
        var count = 0;
        array[count] = "";
        var flag = false;
        for (var counter = 0; counter < data.teams.length; counter++) {
            var str = data.teams[counter].members;
            for (var i = 0; i < str.length; i++) {
                parseInt(str[i])
                console.log(str[i]);

            }
            for (var i = 0; i < str.length; i++) {
                if (!isNaN(str[i])) {
                    array[count] = array[count] + str[i];
                    flag = false;
                } else {
                    if (!flag) {
                        flag = true;
                        count++;
                        array[count] = "";
                    }
                }
            }
            if (!flag) {
                flag = true;
                count++;
                array[count] = "";
            }
        }
        window.$team = array;


        //success message mybe...
    }
});

var showprofile = function(url) {
    $.ajax({
        url: '/user/mprofile?id=' + url,
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


var Techinfo = function() {
    var tid = document.getElementById('techid').value;
    if (tid) {
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
        for (var i , teams.length - 1; i >= 0; i--) {
            var row = teams[i].name+','+teams[i].createdby+','+teams[i].Event+','+teams[i].requestmod+','+teams[i].approved+','+teams[i].members+','+teams[i].comments+','+teams[i].id+','+teams[i].inactivesince;
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