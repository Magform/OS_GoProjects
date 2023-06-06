# Project4

## Main 1
This implementation launches two goroutines, one that takes care of simulating market data and the other instead that takes care of buying or selling Pairs in case of need.
Both goroutines handle each different currency pair with a respective goroutine.
After 60 seconds the program is terminated.

## Main 2
This implementation coincides with the previous one except that it prints the actions that are performed to provide greater readability and facilitate bugfixing

## Consideration
Considering that this test requires execution for 60 seconds it loses some sense to do speed tests, I still decided to go and generate the code from the artificial intelligence to see a further possible implementations.