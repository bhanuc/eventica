(function(){
      if(localStorage.fname){
        $('#side-name').html(localStorage.firstname);
        $('#top-name').html(localStorage.fname);
      } 

      if(localStorage.fpic){
      var url ="//graph.facebook.com/"+localStorage.fpic+"/picture";
        $('#side-pic').attr('src',url);
        $('#top-pic').attr('src', url);
      } 
})()

$( "#user-info" ).click(function() {
$( this ).toggleClass( "open" );
}); 