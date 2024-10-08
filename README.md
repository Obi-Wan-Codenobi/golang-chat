# golang-chat
This is a personal test project to get a little more familiar with golang 

The client creates a simple tcp connection to the server and receives an echo reply. 

While the client's connection thread is slept for 15 seconds, the server prints the number of connctions every 5 seconds which is updated via go channels. 
