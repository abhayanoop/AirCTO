#AirCTO Changelog


* 4th may - 2018

Usage of actual database instead of global maps.  
Postgres DB functions implemented.    

Users also stored in database. User addition removed from structs.  
New users can now be added from "dbfuncs/db.go".    

Issue ID is now auto generated (integers) in database and hence create and update request structures no longer requires Issue ID. 