# Project3
## Main
A fairly straightforward implementation that uses semaphores (implemented through channels) to manage the slots where cakes can be placed, with a separate goroutine for each pastry chef.

## Main verbose
This implementation coincides with the previous one except that it prints the actions that are performed to provide greater readability and facilitate bugfixing

## AI-humanFixed
I decided to include this implementation as well because the AI was unable to create a working implementation, and I made significant modifications to it while keeping only its skeleton.

In the main function, a goroutine is generated for the first pastry chef for each cake. Once a pastry chef finishes their task, they call the next pastry chef, and so on.
To prevent a pastry chef from performing multiple tasks simultaneously channels are used. Channel are used also to manage available spaces.