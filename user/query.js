db.users.update({},{$set : {"userprofile.bookingid": "", "userprofile.year": "" }}, true, true);
