# wmd
 
I wrote "wmd" as part of my bachelor thesis. The application is currently just an prototype and needs some improvement, but nevertheless I published it on github. 

WMD is an knowledge management database, which was mostly written in GO. Please note that as database I used Redis with the default properties.

The following functionalities are implemented:

- Running webserver
- Registration with Email verfication
- No accessing any page without logging in
- Two kinds of accounts with different access rights and functionalities
   - "PROFESSOR" : can see all Uploads (PDFs) and can also Upload new informations
   - "STUDENTS" : can only see all uploads
- With each login new Sessions will be activated and terminated after 2 hours of inactivity
- An formular for new Uploads which allow just PDF files
- Each Uploads will be presented at the index page
- Clicking on any of those entries the PDF file will be displayed

I am still learing and trying to improve my skills. Therefore I would appreciate feedback from skilled developer. Thanks :D
