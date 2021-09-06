# GOLANG ASSESMENT

## Task
---
- [x] Read the raw.json file 
- [x] Loop through the customers slice and send a POST request with the customerId to the web service
- [x] Save the response to a log file and print out the response

## Installation 
---
### Prerequesites 
- Ensure you have the latest version of Go installed
  
### Steps
1. clone the repo down from the main branch
2. run the appropriate script depending on your operating system IE. Windows or linux
3. ```./linux_script.sh ``` or  ``` sh window_script.sh```
4. A golang_json executable file should now be in the root directory of the project
5. run the executable with ```./golang_json```

### Alternative Run
1. Alternatively if you would like to run the go code without creating an executable, run ```go run main.go customers.go orders.go```


--- 
### Notes
- Really enjoyed working on this project, it was a good wait to learn a new language with a goal in mind
- Had opportunities to work with structs and function receivers which was quite new to me.
- I added the raw.json file as an embedded file instead of reading it so that the executable can be run from anywhere (would love to know if that was a common use case)
- Would have liked to use Go routines and channels on this, but was slightly unsure whether this was a common use case to use routines for looping through and calling APIS.
- Had some issues in ensuring the executable worked properly on a windows machine. Temporarily solved it with the embeded raw.json file.