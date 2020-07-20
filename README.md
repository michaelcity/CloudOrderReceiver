### CloudOrderReceiver
CloudOrderReceiver is the Golang base application for getting Cloud order from ordering cloud server. This programmer running every 3 second to collect the data
from the cloud then print the receipt on connected EPSON printer. 
___
### prerequisite
1.The server installed Golang including go-sql-driver plugin

---

### install
1. Create the folder and copy the source to destination folder
2. Build the program by go command `go build`
3. run the file, for example `./CloudOrderReceiver`
