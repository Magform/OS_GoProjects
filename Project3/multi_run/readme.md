# Testing

To execute the tests and avoid extremely long execution times, I deliberately decided to reduce the times to 1%.

## Frist test
I will execute the code 1000 times with a total of 5 cakes.

### Results

AI-HumanFixed   ->  7m33.537925105s
Mine            ->  7m33.469785456s
MineVerbose     ->  7m33.525985719s

## Second test
I will execute the code 1 times with a total of 1000 cakes.

### Results

AI-HumanFixed   ->  13m28.688896726s
Mine            ->  13m28.112161707s
MineVerbose     ->  13m28.474006934s

## Conclusions

The three executions have almost identical times in both tests, which is likely due to the fact that the most time-consuming part of the code is the delay caused by the cake processing done by the pastry chefs. To try to achieve better results, I attempted to reduce the waiting times of the pastry chefs to 0.


# Testing V2

## Frist test V2
I will execute the code 1000000 times with a total of 5 cakes, setting the waiting times of the pastry chefs to 0.

### Results

AI-HumanFixed   ->  2m10.597270361s
Mine            ->  2m19.53836021s
MineVerbose     ->  4m58.583345802s

### Conclusion
We can first notice that not having any output allows for a significant boost in speed. The execution time goes from 4m8s with output to 2m19s without output.

Furthermore, it is evident that the AI's approach is faster compared to the one I devised. This could be due to the fact that my code includes a for loop inside each pastry chef, which continues until all the cakes are finished. It repeatedly calls the semaphore function to check if there are any occupied spots and only proceeds to process the cake if there are available spots. Otherwise, it repeats the loop.

On the other hand, the AI's approach doesn't require this check because there is one instance of a pastry chef for each cake. If there are no available cakes, there are no pastry chefs either. However, it's important to note that multiple pastry chefs with the same role cannot work together, thanks to the use of a channel.

Another possible reason why my idea is slower could be that the use of a dedicated structure, including a constructor and functions, further slows down the code. However, this trade-off provides improved code readability.

## Second test V2
I will execute the code 10000 times with a total of 1000 cakes, setting the waiting times of the pastry chefs to 0.

### Results

AI-HumanFixed   ->  4m9.042528791s
Mine            ->  2m5.277549909s
MineVerbose     ->  10m43.959604551s

### Conclusion
Here, not having any output allows for an even greater boost in speed compared to the previous case, most likely due to the larger number of outputs. Infact the execution time goes from 10m43s with output to 2m5swithout output.

The code with output, AI-HumanFixed and MineVerbose, doesn't experience significantly greater slowdowns. In fact, the code I developed is consistently around 2.4 times slower, which is likely due to the same issues mentioned earlier.