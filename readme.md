# OS_GoProject
## Idea
During my university course on Operating Systems, I was assigned four projects to be implemented using the Go programming language.

## Projects

### Project 1 - Character Count
Write a program in Go that counts the number of times a specific character "x" appears in a string. The program should utilize concurrency by launching a goroutine for each character in the string and checking if the character matches the desired character.

Example: If the string is "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff" and the character to search for is 'c', the program should launch a goroutine for each character in the string. It should use synchronization mechanisms like a WaitGroup and a channel to keep track of the total count of matching characters.

Initialize a test string and the character to search for in the main function, e.g.:
stringa := "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
character := 'c'

At the end of the process, the program should print the final count of matching characters. In our example, the final count would be 11 because the character 'c' appears 11 times in the string.

### Project 2 - Car Rental Agency

Write a program in Go that simulates a car rental agency managing bookings for 10 customers. Each customer rents a vehicle from the available options: Sedan, SUV, or Station Wagon.

• Create a structure named "Customer" with the field "name".
• Create a structure named "Vehicle" with the field "type".
• Create a function named "rent" that takes a customer as input and randomly books a vehicle for them. This function should also print a message stating that customer x has rented vehicle y.
• Create a function named "print" that, at the end of the process, prints the number of Sedans, SUVs, and Station Wagons rented.
• Each customer can rent a vehicle concurrently with others.

Note that besides the two required functions described above, additional functions can be created to solve the problem.

### Project 3 - Cake Production

Write a program in Go that simulates the production of 5 cakes by 3 pastry chefs. The production of each cake consists of 3 stages that must occur in order: first, the cake is baked, then garnished, and finally decorated.

The first pastry chef is responsible only for baking the cakes and takes 1 second for each cake. This pastry chef has 2 spaces available to place the cakes once they finish baking. If there are free spaces, they can start baking the next cake without waiting for the second pastry chef to become available to garnish the cake just baked.

The second pastry chef is responsible only for garnishing the cakes and takes 4 seconds for each cake. This pastry chef also has 2 spaces available to place the cakes once they finish garnishing them.

The third pastry chef is responsible only for decorating the cakes and takes 8 seconds for each cake.

The three pastry chefs work simultaneously.

### Project 4 - Currency Trading Simulation

Write a program in Go that simulates currency trading in a fictional market.

The program should simulate three currency pairs using concurrency: EUR/USD, GBP/USD, and JPY/USD, and simulate buying and selling operations in parallel.

Create a function named "simulateMarketData" that simulates the price of currency pairs and sends the simulated data on a channel. In particular:
• The price of the EUR/USD pair randomly varies between 1.0 and 1.5.
• The price of the GBP/USD pair randomly varies between 1.0 and 1.5.
• The price of the JPY/USD pair randomly varies between 0.006 and 0.009.

The prices are generated and sent on the corresponding channel at regular intervals, specifically every second.

Create a function named "selectPair" that uses a "select" statement to handle buying and selling operations based on the specified conditions. In particular:
• If the EUR/USD price exceeds 1.20, it should sell EUR/USD. Simulate the sale with a time delay of 4 seconds, i.e., insert a 4-second delay before confirming the sale.
• If the GBP/USD price drops below 1.35, it should buy GBP/USD. Simulate the purchase with a time delay of 3 seconds, i.e., insert a 3-second delay before confirming the purchase.
• If the JPY/USD price drops below 0.0085, it should buy JPY/USD. Simulate the purchase with a time delay of 3 seconds, i.e., insert a 3-second delay before confirming the purchase.

The program should execute the trading cycle for one minute and terminate at the end of the cycle.

## Testing
In addition to completing them myself, I also decided to challenge an artificial intelligence (ChatGPT-3.5) to solve the same projects, aiming to achieve better execution times with my own software.

### Test execution methods
To run the various tests I went to create special go programs, found in the "multi_run" folder of each project.
These programs run the code multiple times in series, allowing for accurate performance evaluation. Additionally, I have developed different test versions with minor modifications to identify cases where one approach outperforms another. 
In the same folder there is also a readme where there are the results of the tests on my machine and some considerations.

### Test command
To streamline the execution of tests, I have created a convenient bash file named "grepTime.sh". This file allows you to run all the tests effortlessly.
To use it, simply make "grepTime.sh" executable by running chmod +x grepTime.sh, and then execute it to obtain the test times. As the tests can be time-consuming, I recommend saving the output to a file and running the script in the background using the command ./grepTime.sh > out.txt &. 

### Testing Machine
For all tests, I am using my private server, a DELL Precision 5810 with the following specifications:
|  |  |
|--------------|-----------|
| CPU | Intel Xeon E5-1650 v4 Prozessor (3,6GHz, 6C, 15MB, 140W, 4GHz Turbo) |
| RAM | 32 GB ECC DDR4 (8*4GB) |
| GPU | NVIDIA Quadro M4000 8GB | 

The server runs Ubuntu as its operating system, and I access it via SSH. In idle state, it typically has around 0.2% CPU usage.