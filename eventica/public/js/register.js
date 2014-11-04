var host = "http://localhost:8080/"
function Register() {
$.ajax({
  url: '/api/user/register',
  type: 'POST',
  contentType: 'application/json; charset=utf-8',
  data: JSON.stringify({'name': $('#name').val(), 'email': $('#email').val(), 'number':$('#number').val(), 'password':$('#password').val(), 'password2':$('#password2').val()}),
  success: function(data) {
  			console.log(data);
        if(data.success){
          toastr.success("Signup was successful, You can now login Using the same Credentials");
          //document.location = host+"app"
        } else {
          toastr.error(data.flashes.Error.message);
        }
          //success message mybe...
     }
});
}
function validate()
{
 
   if( document.myForm.name.value == "" )
   {
     toastr.error( "Please provide your name!" );
     document.myForm.name.focus() ;
     return false;
   }
   if( document.myForm.Email.value == "" )
   {
     toastr.error( "Please provide your Email!" );
     document.myForm.Email.focus() ;
     return false;
   }
   if( document.myForm.number.value == "" ||
           isNaN( document.myForm.number.value ) ||
           document.myForm.number.value.length != 10 )
   {
     toastr.error( "Please provide Phone Number of length 10." );
     document.myForm.number.focus() ;
     return false;
   }
   return( true );
}
function validateEmail()
{
 
   var emailID = document.myForm.Email.value;
   atpos = emailID.indexOf("@");
   dotpos = emailID.lastIndexOf(".");
   if (atpos < 1 || ( dotpos - atpos < 2 )) 
   {
       toastr.error("Please enter correct email ID")
       document.myForm.Email.focus() ;
       return false;
   }
   return( true );
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
	if(validate()  && validateEmail() )
        Register();
});

