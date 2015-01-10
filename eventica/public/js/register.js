var host = "http://portal.techkriti.org/"


function Register() {
  var opts = {
  lines: 13, // The number of lines to draw
  length: 26, // The length of each line
  width: 10, // The line thickness
  radius: 57, // The radius of the inner circle
  corners: 1, // Corner roundness (0..1)
  rotate: 19, // The rotation offset
  direction: 1, // 1: clockwise, -1: counterclockwise
  color: '#000', // #rgb or #rrggbb or array of colors
  speed: 1.6, // Rounds per second
  trail: 89, // Afterglow percentage
  shadow: false, // Whether to render a shadow
  hwaccel: false, // Whether to use hardware acceleration
  className: 'spinner', // The CSS class to assign to the spinner
  zIndex: 2e9, // The z-index (defaults to 2000000000)
  top: '50%', // Top position relative to parent
  left: '50%' // Left position relative to parent
};
  var target = document.getElementById('container-page');
var spinner = new Spinner(opts).spin(target);
$.ajax({
  url: '/api/user/register',
  type: 'POST',
  contentType: 'application/json; charset=utf-8',
  data: JSON.stringify({'name': $('#name').val(), 'email': $('#email').val(), 'number':$('#number').val(), 'password':$('#password').val(), 'password2':$('#password2').val()}),
  success: function(data) {
  			console.log(data);
        if(data.success){
         toastr.success("Your registration is pending. Please check your inbox/spam to activate your account. Email might take 5 min. for delivery. ");
          //document.location = host+"/actsuccess"
        } else {
          toastr.error(data.flashes.Error.message);
        }
        spinner.stop()
        $( "#registerbutton" ).one( "click", function() {
if(validate()  && validateEmail() ){
          Register();
}
});
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

$( "#registerbutton" ).one( "click", function() {
if(validate()  && validateEmail() ){
          Register();

}
});


