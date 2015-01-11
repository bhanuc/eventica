$.ajax({
    url: '/team/all2',
    type: 'GET',
    success: function(data) {
        console.log(data);
        if (data.success) {
            var str = '<thead><tr><th>Name</th><th>Event</th><th>Submitted Team for approval</th><th>Members</th><th>Submit your team</th><th>Approval Status</th><th>Comments</th><th>Edit</th></tr></thead><tbody>';
            for (var i = data.teams.length - 1; i >= 0; i--) {
                str = str + '<tr><td>' + data.teams[i].name + '</td><td> ' + data.teams[i].event + '</td><td> ' + (function() {
                    if (data.teams[i].requestmod) return "Submitted";
                    else return "Not Submitted"
                })() + '</td><td> ' + data.teams[i].members + '</td>'+(function(){if(data.teams[i].requestmod){ return "<td>Submitted</td>";
        } else{
            return '<td> <a href="/team/register?name=' + data.teams[i].name + '&createdby=' + data.teams[i].createdby + '">Click to Submit</a></td>'
        }})()+ '<td> ' + data.teams[i].approved + '</td><td> ' + data.teams[i].comments + '</td><td> <button onclick=editmyteam("' + data.teams[i].id + '") type="button" class="btn btn-primary btn-lg">Edit</button></td></tr>';
            };
            str += '</tbody>'
            $('#teams').html(str);
        } else {}
    }
});

function editmyteam(id) {

    $.ajax({
        url: '/team/single',
        type: 'POST',
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            'id': id
        }),
        success: function(data) {
            console.log(data, data.event);
            $('#myModal').modal('show');
            document.getElementById('mevent').value = data.team.event;
            document.getElementById('mmembers').value = data.team.members;
            document.getElementById('mname').value = data.team.name;
            document.getElementById('mid').value = data.team.id;
        }
    });



}

function submiteventname() {
    var check = document.getElementById('mmembers').value;
    var k = check.split(',');
    for (var i = k.length - 1; i >= 0; i--) {
      if (!k[i]) {
        k.splice(i,1);
      }
    };
    check = _.uniq(k, true).toString();
    if (k.length == 0) {
        toastr.error("You cannot submit empty team");
    } else {
        $.ajax({
            url: '/team/update',
            type: 'POST',
            contentType: 'application/json; charset=utf-8',
            data: JSON.stringify({
                'id': document.getElementById('mid').value,
                "name": document.getElementById('mname').value,
                event: document.getElementById('mevent').value,
                members: check
            }),
            success: function(data) {

                if (data.success) {
                    window.location = "/showteam";
                }
                $('#myModal').modal('hide');
            }
        });

    }
}