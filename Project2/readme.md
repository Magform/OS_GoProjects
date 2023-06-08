# Project2
## Main
This version of the code uses a structure that contains the model name, the number of vehicles of that model currently in use, and a mutex to ensure that the number of vehicles in use is not modified by multiple goroutines simultaneously.

In the main function, a different goroutine is called for each client, and a random vehicle is assigned to them.

## Main2
This version of the code uses a structure that contains the model name and the number of vehicles of that model currently in use. To ensure that the number of vehicles used per model is thread safe, an atomic increment is used.

In the main function, a different goroutine is called for each client, and a random vehicle is assigned to them.