# Exercise 1: Quiz Game

Create a program that will read in a quiz provided via a CSV file and will give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is right or wrong, the next question should be asked immediatley after.

The CSV file should default to ```problems.csv```, but the user should be able to customize the filename via a flag.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

## Usage

First build the go file using ```go build```.

Then run using the default ```problems.csv``` by going ```./quiz```. Or specify the program to run a .csv file to run by including a flag like so: ```./quiz -f "asdf.csv"```.
