var host = "http://portal.techkriti.org/"

function submitreset() {
    $.ajax({
        url: '/api/user/resetrequest',
        type: 'POST',
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            'email': $('#email').val()
        }),
        success: function(data) {
            console.log(data);
            if (data.flashes && data.flashes.success) {
                toastr.success(data.flashes.success.message);
                // / document.location = "/profile";
            } else {
                toastr.error(data.flashes.Error.message)
            }
            $('#ResetButton').one("click", function() {
                if (validate() && validateEmail()) {
                    submitreset();
                }
            });
        }
    });
}

function validateEmail() {

    var emailID = $('#email').val();
    atpos = emailID.indexOf("@");
    dotpos = emailID.lastIndexOf(".");
    if (atpos < 1 || (dotpos - atpos < 2)) {
        toastr.error("Please enter a correct email ID")
        document.loginForm.email.focus();
        return false;
    }
    return (true);
}
$('#email').keypress(function(e) {
    console.log(e);
    if (e.which == '13') {
        submitreset();
    }
});

$('#password').keypress(function(e) {
    console.log(e);
    if (e.which == '13') {
        submitreset();
    }
});

$('#ResetButton').one("click", function() {
    if (validateEmail()) {
        submitreset();
    }
});