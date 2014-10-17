var host = "http://localhost:8080/"

function submitLogin() {
$.ajax({
  url: '/api/user/authenticate',
  type: 'POST',
  contentType: 'application/json; charset=utf-8',
  data: JSON.stringify({'email': $('#email').val(), 'password': $('#password').val(), 'rememberme': $('#rememberme').val()}),
  success: function(data) {
        console.log(data);
        if(data.valid){
          toastr.success("Login was successful, You will be redirected to Dashboard");
          document.location = "/profile";
        } else {
          toastr.error(data.flashes.Error.message)
        }
          //success message mybe...
     }
});
}
function validate()
{
 
   
   if( document.loginForm.email.value == "" )
   {
     toastr.error( "Please provide your Email!" );
     document.loginForm.email.focus() ;
     return false;
   }
   if( document.loginForm.password.value == "" )
   {
     toastr.error( "Please provide your Password" );
     document.loginForm.password.focus() ;
     return false;
   }
   return( true );
}
function validateEmail()
{
 
   var emailID = document.loginForm.email.value;
   atpos = emailID.indexOf("@");
   dotpos = emailID.lastIndexOf(".");
   if (atpos < 1 || ( dotpos - atpos < 2 )) 
   {
       toastr.error("Please enter correct email ID")
       document.loginForm.email.focus() ;
       return false;
   }
   return( true );
}
$('#email').keypress(function(e) {
console.log(e);
if (e.which == '13') {
         submitLogin();
   }
});

$('#password').keypress(function(e){
console.log(e);
if(e.which=='13'){
    submitLogin();
  }
});

$('#loginbutton').click(function(){
if(validate() && validateEmail())
    submitLogin();
});
