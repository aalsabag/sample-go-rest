###This is a project for testing maxed CPU usage in an application###

To test run 
`go build`
`./sample-go-rest -cpus=25 -time=10`

There are two endpoints, one for specifying your own number of desired cpus
`/execute/{cpus}/{timeInSecs}`
Another for just maxing out what you have
`/max/{timeInSecs}`