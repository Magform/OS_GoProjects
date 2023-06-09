# Testing

## Frist test
In the first test I'll execute both codes 10000000 times in series and I'm going to visualize the total execution time, the string is always the same, that is: "aaaaaaaaaaaaabbbbbbbbcccccddddccccccffff" and we're going to look for the character "c"
### Results

AI      -> 14m57.262297797s 
Mine    -> 8m25.950812538s
Mine V2 -> 7m29.492556143s
Mine V3 -> 6m50.881159914s
Mine V4 -> 7m54.707332454s

## Second test
Now I try to run the code 10000 times with a random string of 100000 characters generated each time and always look for the letter "c"
### Results

AI      -> 12m48.602838387s 
Mine    -> 18m51.754294776s  
Mine V2 -> 18m18.363812306s
Mine V3 -> 3m57.126076473s
Main V4 -> 18m57.336762635s

### Comment
This test is what made it most difficult for me to achieve a better time than the AI. In fact, I was only able to beat it with V3, which uses the trick of dividing the string into three parts and analyzing them in three different goroutines (see project1.md for more information).

## Third test
For the last test I run the code only 100 times but with a random string of 10000000 characters

### Results

AI      -> 22m42.162630503s
Mine    -> 17m22.22877994s 
Mine V2 -> 17m5.153488573s 
Mine V3 -> 3m33.77881213s  
Mine V4 -> 17m27.640429873s

### Comment 
Looking at this test and the previous ones we can see that the V3 code is much faster in processing very long strings than in processing many strings, this is precisely due to the fact that V3 can write string data into 3 different variables which has an advantage that is not too significant in small strings but very significant in large strings.