(function(){
      if(localStorage.fname && localStorage.fname != "null"){
        $('#side-name').html(localStorage.firstname);
        $('#top-name').html(localStorage.fname);
      } 

      if(localStorage.fpic && localStorage.fpic != "null"){
      var url ="//graph.facebook.com/"+localStorage.fpic+"/picture";
        $('#side-pic').attr('src',url);
        $('#top-pic').attr('src', url);
      } 
})()

$( "#user-info" ).click(function() {
$( this ).toggleClass( "open" );
}); 


setTimeout(
  function() 
  {
	if(localStorage.Tid && localStorage.Tid != "null"){
        $('#techid').html(localStorage.Tid);
      } 
  }, 5000);
