var host = "http://localhost:8080/"
/*$('#registerbutton').click(function() {
$.ajax({
  url: '/api/user/register',
  type: 'POST',
  contentType: 'application/json; charset=utf-8',
  data: JSON.stringify({'password': $('#password').val(), 'email': $('#email').val(), 'email2':$('#email2').val()}),
  success: function(data) {
  			console.log(data);
        if(data.success){
          toastr.success("Signup was successful, You can now login Using the same Credentials");
          //document.location = host+"app"
        } else {
          toastr.error(data.flashes.general.message);
        }
          //success message mybe...
     }
});
});*/
function Register() {
$.ajax({
  url: '/api/user/register',
  type: 'POST',
  contentType: 'application/json; charset=utf-8',
  data: JSON.stringify({'password': $('#password').val(), 'email': $('#email').val(), 'email2':$('#email2').val()}),
  success: function(data) {
  			console.log(data);
        if(data.success){
          toastr.success("Signup was successful. Check your email to activate your Account. It may take few minutes for the email to reach your inbox.");
          //document.location = host+"app"
        } else {
          toastr.error(data.flashes.general.message);
        }
          //success message mybe...
     }
});
}

$('#email').keypress(function(e) {
console.log(e);
if (e.which == '13') {
         Register();
   }
});
$('#email2').keypress(function(e) {
console.log(e);
if (e.which == '13') {
         Register();
   }
});
$('#password').keypress(function(e){
console.log(e);
if(e.which=='13'){
    Register();
	}
});

$('#registerbutton').click(function(){
    Register();
});

