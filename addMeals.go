package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Meal struct {
	Name string `json:"name"`
	Cost float64 `json:"cost"`
	Calories int `json:"calories"`
}
var Database struct{
	Breakfast []*Meal `json:"breakfast"`
	Lunch []*Meal `json:"lunch"`
	Dinner []*Meal `json:"dinner"`
	Dessert []*Meal `json:"dessert"`
	Snack []*Meal `json:"snack"`
}
func init(){
	bytes, err := ioutil.ReadFile("mealDatabase.json")
	if err != nil{
		fmt.Println("Your meal database doesn't exist! I'll make one for you!")
		bytes, _ := json.Marshal(&Database)
		ioutil.WriteFile("mealDatabase.json", bytes, os.ModePerm)
	}
	json.Unmarshal(bytes, &Database)
}
func consoleInput(message string) string{
	customImport := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	valueToReturn, _ := customImport.ReadString('\n')
	valueToReturn = strings.TrimSuffix(valueToReturn, "\n")
	return valueToReturn
}
func main(){
	defer func() {
		bytes, _ := json.Marshal(&Database)
		ioutil.WriteFile("mealDatabase.json", bytes, os.ModePerm)
	}()
	choice := consoleInput("What type do you want to add (please add numerical value in text box)? \n1) Breakfast \n2) Lunch \n3) Dinner \n4) Dessert \n5) Snack \n6) Exit\nInput here: ")
	switch choice {
	case "1":
		name := consoleInput("What is the breakfast item's name: ")
		preCost := consoleInput("What is the cost of the breakfast item: ")
		preCalories := consoleInput("What is the calorie count for your "+name+": ")
		cost, err := strconv.ParseFloat(preCost,64)
		if err != nil{
			panic("You specified a string for a cost. They are naked numbers, with no currency mark.")
		}
		calories, err := strconv.Atoi(preCalories)
		if err != nil{
			panic("You specified a string for the calorie count. Calories are numbers with no decimal places.")
		}
		Database.Breakfast = append(Database.Breakfast, &Meal{name, cost, calories})
	case "2":
		name := consoleInput("What is the lunch item's name: ")
		preCost := consoleInput("What is the cost of the lunch item: ")
		preCalories := consoleInput("What is the calorie count for your "+name+": ")
		cost, err := strconv.ParseFloat(preCost,64)
		if err != nil{
			panic("You specified a string for a cost. They are naked numbers, with no currency mark.")
		}
		calories, err := strconv.Atoi(preCalories)
		if err != nil{
			panic("You specified a string for the calorie count. Calories are numbers with no decimal places.")
		}
		Database.Lunch = append(Database.Lunch, &Meal{name, cost, calories})
	case "3":
		name := consoleInput("What is the dinner item's name: ")
		preCost := consoleInput("What is the cost of the dinner item: ")
		preCalories := consoleInput("What is the calorie count for your "+name+": ")
		cost, err := strconv.ParseFloat(preCost,64)
		if err != nil{
			panic("You specified a string for a cost. They are naked numbers, with no currency mark.")
		}
		calories, err := strconv.Atoi(preCalories)
		if err != nil{
			panic("You specified a string for the calorie count. Calories are numbers with no decimal places.")
		}
		Database.Dinner = append(Database.Dinner, &Meal{name, cost, calories})
	case "4":
		name := consoleInput("What is the dessert item's name: ")
		preCost := consoleInput("What is the cost of the dessert item: ")
		preCalories := consoleInput("What is the calorie count for your "+name+": ")
		cost, err := strconv.ParseFloat(preCost,64)
		if err != nil{
			panic("You specified a string for a cost. They are naked numbers, with no currency mark.")
		}
		calories, err := strconv.Atoi(preCalories)
		if err != nil{
			panic("You specified a string for the calorie count. Calories are numbers with no decimal places.")
		}
		Database.Dessert = append(Database.Dessert, &Meal{name, cost, calories})
	case "5":
		name := consoleInput("What is the snack's name: ")
		preCost := consoleInput("What is the cost of the snack: ")
		preCalories := consoleInput("What is the calorie count for the "+name+": ")
		cost, err := strconv.ParseFloat(preCost,64)
		if err != nil{
			panic("You specified a string for a cost. They are naked numbers, with no currency mark.")
		}
		calories, err := strconv.Atoi(preCalories)
		if err != nil{
			panic("You specified a string for the calorie count. Calories are numbers with no decimal places.")
		}
		Database.Snack = append(Database.Snack, &Meal{name, cost, calories})
	case "6":
		fmt.Println("\nShutting down.")
		os.Exit(0)
	default:
		panic("Yeah, you screwed up. You need to specify a number 1-6 that attributes to each meal in the menu above.")
	}
	fmt.Println("\nSuccess!")
}

