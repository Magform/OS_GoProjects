# Project 1
## Main V1
This version of the code use a sync.WaitGroup to synchronize the goroutines and a channel to ensure that the variable containing the character count is not modified concurrently by multiple goroutines. The main function iterates over each character in the where string, launches a goroutine to check the character, and adds it to the wait group.

After launching all the goroutines, the main function waits for all goroutines to finish using wg.Wait(). Finally, it retrieves the final count from the results channel and prints it as the number of times the desired character appears in the string.

### Problems and considerations
The main problem of this code is that the variable containing the number of characters can only be used by one goroutine at a time thus slowing down the execution of the code

## Main V2
This version of the code use a sync.WaitGroup to synchronize the goroutines and a mutex to ensure that the variable containing the character count is not modified concurrently by multiple goroutines. The main function iterates over each character in the where string, launches a goroutine to check the character, and adds it to the wait group.

After launching all the goroutines, the main function waits for all goroutines to finish using wg.Wait(). Finally, it retrieves the final count from the results channel and prints it as the number of times the desired character appears in the string.

### Problems and considerations
The main problem of this code is that the variable containing the number of characters can only be used by one goroutine at a time thus slowing down the execution of the code

## Main V3
Like the previous versions, this code snippet uses a sync.WaitGroup to synchronize the goroutines and a mutex to ensure that the variable containing the character count is not concurrently modified by multiple goroutines.

In this version, we have introduced three separate instances of the function that checks the character, each with its own variable to store the repetitions. The string is split into three parts, and we check each part individually. Afterwards, we sum up the counts from the three variables.

We did this to decrease the average waiting time to access the variable that counts the number of repetitions

After launching all the goroutines, the main function waits for all goroutines to finish using wg.Wait(). Finally, it retrieves the final count from the repetitions variable and prints it as the number of times the desired character appears in the string.

### Problems and considerations 
Now we are going to use three different instances of the function but hypothetically if we were going to use more the code would become even faster up to a maximum limit after which the sum of the various variables becomes more time expensive than the time gained

## Main V4
This version of the code snippet uses a sync.WaitGroup to synchronize the goroutines and a mutex to ensure that the variable containing the character count is not concurrently modified by multiple goroutines.

In this version, we have introduced multiple separate instances of the function,  each responsible for checking a specific character in the string.
An array is used to store the results, with each element corresponding to a character in the string. The elements are set to 1 if the character matches the desired character and 0 otherwise. By summing up all the elements, the total count of the desired character is obtained.

We did this to decrease the average waiting time to access the variable that counts the number of repetitions

After launching all the goroutines, the main function waits for all goroutines to finish using wg.Wait(). Finally, it retrieves the final count from the repetitions variable and prints it as the number of times the desired character appears in the string.

### Problems and considerations 
As previously highlighted, here we go beyond the maximum limit where having more variables for various characters is advantageous since the sum of the various variables becomes more expensive than the time gained by not using mutexes