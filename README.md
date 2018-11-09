# goMeals
Simple meal maker program to learn how to manipulate json databases using console inputs in Go.
This program has two files: one to add meals (also create a DB on first run) and one to get a random meal with statistics.
This is to learn how to use json marshalling and ioutil's WriteFile in GoLang.
You need minimum 2 meals per category before trying to manipulate the database in getMeals.go.

Installation:
1. Run addMeals.go
2. Follow directions using console inputs, adding custom meals to any course of your choosing.
3. Once you have added your first meal, a  json file should be created in the directory. Don't mess with it, as it will goof.
4. Once you have 2 meals in each category, you can run getMeals.go.
5. In getMeals.go, there will be two ways of accessing the data: manual and automatic.
 -manual is where you select what course you want and it picks a random course
 -automatic is where it finds the current time of day and matches you up with a meal
6. Enjoy!

If you have any questions, feel free to contact me on discord. My name is dubber#7428.

Big thanks to Discoli for teaching me the basics of GoLang (structs, json marshalling, golang funcs, defer funcs, ioutil, strings, etc etc).
