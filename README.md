# AirCTO
Code for issue tracking test provided by AirCTO

Requirements:

* An SMTP mail server
* Postman tool or any alternate API testing tool
* Install Postgres database. If not installed,  

```
  ubuntu/linux - https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04

  mac osx - https://www.codementor.io/engineerapart/getting-started-with-postgresql-on-mac-osx-are8jcopb
```
    

To Do:

* By default, postgres installation creates user 'postgres' and database 'postgres'.  
  These credentials have been used in this program. If you require to change database or user or password(optional), you can change anytime in "structs/structs.go" database credentials.  

* Submit the SMTP mail server credentials in the file "structs.go" for the variables provided
  
  smtpServerPort, smtpFrom, smtpAuth
  
* Create new legit users in code in the file "dbfuncs/db.go" to test email service.
  Three dummy users have been hardcoded in code for reference. These users can also be utilised for API calls.
  
Working:

* Server is hosted in local host port :8080. 
* All API's require a request header - "Authorization" with value accesstoken of logged in user
* API's are provided below,

* /issue/read?id=(issue-id)
  
* /issue/create
  
      raw application/JSON body
  
      request body example = {"Title":"Issue 1","Description":"first issue", "AssignedTo":"user1","CreatedBy":"user2","Status":"Open"}
  
  
* /issue/update?id=(issue-id)
  
      raw application/JSON body
  
      request body example = {"Title":"Issue 1 update","Description":"updated issue", "AssignedTo":"user3","CreatedBy":"user2","Status":"Open"}
  
  
* /issue/delete?id=(issue-id)
 
* /issue/list 
  
  
  
  
  Example request in postman:
  
  GET localhost:8080/issue/create
  
  Header "Authorization" = "tokenuser2"
  
  Body raw application/JSON => 
  
  {"Title":"Issue 1","Description":"first issue", "AssignedTo":"user1","CreatedBy":"user2","Status":"Open"}
  
  
  
