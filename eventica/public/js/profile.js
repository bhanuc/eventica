var global = Object.create(null);

var getprofile = function getprofile () {
  $.ajax({
  url: '/api/user/profile',
  type: 'GET',
  success: function(data) {
    console.log(data)
    updatename();
    updatepic();
  			if(data.profile){
          localStorage.setItem('name',data.profile.name);
          global.profile = data.profile;
          var profile = Object.keys(data.profile);
          global.keys = profile;
          var string = "";
          for (var i = 1; i <= profile.length - 1; i++) {
            console.log(profile[i])
            switch (profile[i]) {
                case "alternatenumber":
            string = string+'<div class="row"><div class="col-md-6">Alternate person of contact</div>'+'<div class="col-md-6">not written</div>'+'</div>';
               break;
                case "apnr":
                string = string+'<div class="row"><div class="col-md-6">Arrival PNR</div><div class="col-md-6">'+(data.profile[profile[i]]|| ' ')+'</div></div>';
               break;
                case "adate":
            string = string+'<div class="row"><div class="col-md-6">Arrival Date</div><div class="col-md-6">'+(data.profile[profile[i]]|| ' ')+'</div></div>';
               break;
                case "dpnr":
            string = string+'<div class="row"><div class="col-md-6">Departure PNR</div><div class="col-md-6">'+(data.profile[profile[i]]|| ' ')+'</div></div>';
               break;
                case "ddate":
            string = string+'<div class="row"><div class="col-md-6">Departure Date</div><div class="col-md-6">'+(data.profile[profile[i]]|| ' ')+'</div></div>';
               break;
              default:
                    string = string+'<div class="row"><div class="col-md-6">'+profile[i]+'</div><div class="col-md-6">'+(data.profile[profile[i]]|| ' ')+'</div></div>';
              }
            };
	
          $('#profileedit').hide();
        //  $('#SubmitButton').hide();
          $('#profileview').show().html(string)
  			} else {
  				toastr.error("Your Session seems to expired. Please Login again.")
  			}
     }
});
}();

$('#EditButton').click(function editbutton () {
  console.log(global);
    if(global && global.profile){
      var form = '<form>';
      for (var i = 0; i <= global.keys.length - 1; i++) {
	if(global.keys[i] == 'id'){
	continue;
	}
  if(global.keys[i] == 'name' || global.keys[i] == 'email') {
      form = form+'<div class="row"><div class="col-md-6">'+global.keys[i]+'</div><div class="col-md-6"><input type="text" value="'+(global.profile[global.keys[i]] || ' ')+'" id="'+global.keys[i]+'" name="profileForm" disabled ></div></div>';
            } else if(global.keys[i] == 'alternatenumber'){
            if(global.profile[global.keys[i]]){
           var input =  global.profile[global.keys[i]].split('$$');
                    form = form+'<div class="row"><div class="col-md-6"> Alternate person of contact: </div><div class="col-md-6"><input type="text" value="'+input[0]+'" id="alternatenumber1" name="profileForm"  >  <input type="text" value="'+(input[1] || 'Not submitted')+'" id="alternatenumber2" name="profileForm"  > <input type="text" value="'+(input[2] || "Not submitted")+'" id="alternatenumber3" name="profileForm"  ></div></div>';
            } else {
                    form = form+'<div class="row"><div class="col-md-6">Alternate person of contact:</div><div class="col-md-6"><input type="text" value="Name" id="alternatenumber1" name="profileForm"  >  <input type="text" value="Email" id="alternatenumber2" name="profileForm"  ><input type="text" value="Contact No" id="alternatenumber3" name="profileForm"  ></div></div>';   
                  } 
                }  else {
              form = form+'<div class="row"><div class="col-md-6">'+global.keys[i]+'</div><div class="col-md-6"><input type="text" value="'+( global.profile[global.keys[i]] || ' ')+'" id="'+global.keys[i]+'" name="profileForm"></div></div>';
              } 
           }
             form = form + '</form>';
                        $('#profileview').hide();
                        $('#EditButton').hide();
                       $('#profileedit').show();
                       $('#profileform').html(form);
    } else {
                  toastr.error("Your Session seems to expired. Please Login again.")
    }

})

$('#SubmitButton').click(function submitbutton () {
    if(global && global.profile){
      var value = Object.create(null);
      var index = 1;
      for (var i = 0; i <= global.keys.length - 1; i++) {
        if(global.keys[i] == 'alternatenumber'){
            value[global.keys[i]] = $('#alternatenumber1').val()+'$$'+ $('#alternatenumber2').val()+'$$'+$('#alternatenumber3').val();
        }  else {
       value[global.keys[i]] = $('#'+global.keys[i]).val();
     }
       console.log(global.keys[i],$('#'+global.keys[i]).val() );
       if(global.keys[i] == "number" ){
          var regex = /\d/; 
          var num = $('#'+global.keys[i]).val()
          console.log(new RegExp(regex).test(num), num.length );
            if( !(new RegExp(regex).test(num)) && num.length != 10){
                  toastr.error("Please enter a correct phone number");
                  index = 0;
                  break;
            }
       }
       if(global.keys[i] == "sportincharge" ){
          var sic = $('#'+global.keys[i]).val();
            if(sic==" "){
                  toastr.error("Please enter sport in charge");
                  index = 0;
                  break;
            }
       }
             };
              if(index == 1){
                $.ajax({
                url: '/api/user/profile',
                type: 'POST',
                data: JSON.stringify(value),
                contentType: 'application/json; charset=utf-8',
                success: function(data) {
                          if(data.success){
                         toastr.success(data.flashes.profile_updated.message);
                         getprofile(); 
        } else {
          toastr.error("Your Session seems to expired. Please Login again.")
        }
     }
        });
              } else {
                          toastr.error("Form not submitted. Please check stuff")
                        }
    } else { 
                  toastr.error("Your Session seems to expired. Please Login again.")
    }
});

var updatename = function(){
      if(localStorage.name){
        $('#side-name').html(localStorage.name);
        $('#top-name').html(localStorage.name);
      } 
};
var updatepic = function(){
      if(localStorage.fpic){
      var url ="//graph.facebook.com/"+sarfraz.anees+"/picture";
        $('#side-pic').attr('src',localStorage.fpic);
        $('#top-pic').attr('src',localStorage.fpic);
      } 
};