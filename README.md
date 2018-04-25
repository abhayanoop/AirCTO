# AirCTO
Code for issue tracking test provided by AirCTO

Requirements:

* An SMTP mail server
* Postman tool or any alternate API testing tool

To Do:

* Submit the SMTP mail server credentials in the file "structs.go" for the variables provided
  
  smtpServerPort, smtpFrom, smtpAuth
  
* Create new legit users in code in the file "structs.go" to test email service.
  Three dummy users have been hardcoded in code for reference. These users can also be utilised for API calls.
  
Working:

* Server is hosted in local host port :8080. 
* All API's require a request header - "Authorization" with value accesstoken of logged in user
* API's are provided below,

* /issue/read?id=(issue-id)
  
* /issue/create
  
      raw application/JSON body
  
      request body example = {"ID":"issue1","Title":"Issue 1","Description":"first issue", "AssignedTo":"user1","CreatedBy":"user2","Status":"Open"}
  
  
* /issue/update?id=(issue-id)
  
      raw application/JSON body
  
      request body example = {"ID":"issue1","Title":"Issue 1 update","Description":"updated issue", "AssignedTo":"user3","CreatedBy":"user2","Status":"Open"}
  
  
* /issue/delete?id=(issue-id)
 
* /issue/list 
  
  
  
  
  Example request in postman:
  
  GET localhost:8080/issue/create
  
  Header "Authorization" = "tokenuser2"
  
  Body raw application/JSON => 
  
  {"ID":"issue1","Title":"Issue 1","Description":"first issue", "AssignedTo":"user1","CreatedBy":"user2","Status":"Open"}
  
  
  
