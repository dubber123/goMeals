package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
		panic("The database does not exist. Please run addMeals.go and add at least 2 meals per course before trying to manipulate them here!")
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
	choice := consoleInput("How would you like this meal to be generated?\n 1) Select Meal Name \n 2) Have Computer Generate a Meal Based on Current Time \n 3) Exit \n Input here: ")
	switch choice{
	case "1":
		choice2 := consoleInput("\nWhat meal are you about to eat?\n 1) Breakfast \n 2) Lunch \n 3) Dinner \n 4) Snack \n 5) Exit \n Input here:")
		switch choice2 {
		case "1":
			rand.Seed(time.Now().UnixNano())
			randBreakfast := Database.Breakfast[rand.Intn(len(Database.Breakfast))]
			fmt.Println("Your randomly chosen breakfast for the day is "+randBreakfast.Name)
			totalCost := strconv.FormatFloat(randBreakfast.Cost, 'f',2,64)
			totalCalories := strconv.Itoa(randBreakfast.Calories)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)


		case "2":
			rand.Seed(time.Now().UnixNano())
			randLunch := Database.Lunch[rand.Intn(len(Database.Lunch))]
			randDessert := Database.Dessert[rand.Intn(len(Database.Dessert))]
			fmt.Println("Your randomly chosen lunch for the afternoon is "+randLunch.Name)
			fmt.Println("Once you are done with that, you can enjoy "+randDessert.Name)
			fmt.Println("Make sure to pop a breathmint after lunch, too!")
			totalCostUno := randLunch.Cost + randDessert.Cost
			totalCaloriesUno := randLunch.Calories + randDessert.Calories
			totalCost := strconv.FormatFloat(totalCostUno, 'f',2,64)
			totalCalories := strconv.Itoa(totalCaloriesUno)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost+"\n...and a breathmint!")

		case "3":
			rand.Seed(time.Now().UnixNano())
			randDinner := Database.Dinner[rand.Intn(len(Database.Dinner)-1)]
			randDessert := Database.Dessert[rand.Intn(len(Database.Dessert)-1)]
			fmt.Println("Your randomly chosen dinner for the night is "+randDinner.Name)
			fmt.Println("Once you are done with that, you can enjoy "+randDessert.Name)
			totalCostUno := randDinner.Cost + randDessert.Cost
			totalCaloriesUno := randDinner.Calories + randDessert.Calories
			totalCost := strconv.FormatFloat(totalCostUno, 'f',2,64)
			totalCalories := strconv.Itoa(totalCaloriesUno)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)

		case "4":
			rand.Seed(time.Now().UnixNano())
			randSnack := Database.Snack[rand.Intn(len(Database.Snack)-1)]
			fmt.Println("Your random snack is "+ randSnack.Name+", don't get too fat!")
			totalCost := strconv.FormatFloat(randSnack.Cost, 'f',6,64)
			totalCalories := strconv.Itoa(randSnack.Calories)
			fmt.Println(" Your Stat Line:\n\nCalories Consumed: "+totalCalories+"\n Cost of Meal: "+totalCost)

		case "5":
			fmt.Println("\nShutting down.")
			os.Exit(0)
		default:
			panic("Please use a number when accessing elements in dubber's menus. Thx")
		}
	case "2":
		j := time.Now().Format("154")
		fmt.Println(j)
		hour2, _ := strconv.Atoi(j)
		if hour2 >= 500 && hour2 <= 800 || hour2 >= 50 && hour2<=80 {
			rand.Seed(time.Now().UnixNano())
			randBreakfast := Database.Breakfast[rand.Intn(len(Database.Breakfast)-1)]
			fmt.Println("Found your time to be between 5:00 am and 8:00 am, so... \nYour randomly chosen breakfast for the day is "+randBreakfast.Name)
			totalCost := strconv.FormatFloat(randBreakfast.Cost, 'f',2,64)
			totalCalories := strconv.Itoa(randBreakfast.Calories)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)
		} else if hour2 >= 801 && hour2 <= 1030 || hour2 >= 81 && hour2 <= 103{
			rand.Seed(time.Now().UnixNano())
			randSnack := Database.Snack[rand.Intn(len(Database.Snack)-1)]
			fmt.Println("Found your time to be between 8:00 am and 10:30 am, so... \nYour random snack is "+ randSnack.Name+", don't get too fat!")
			totalCost := strconv.FormatFloat(randSnack.Cost, 'f',2,64)
			totalCalories := strconv.Itoa(randSnack.Calories)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)
		} else if hour2 >= 1031 && hour2 <= 1415|| hour2 >= 130 && hour2 <= 140{
			rand.Seed(time.Now().UnixNano())
			randLunch := Database.Lunch[rand.Intn(len(Database.Lunch)-1)]
			randDessert := Database.Dessert[rand.Intn(len(Database.Dessert)-1)]
			fmt.Println("Found your time to be between 10:30 am and 2:15 pm, so... \nYour randomly chosen lunch for the afternoon is "+randLunch.Name)
			fmt.Println("Once you are done with that, you can enjoy "+randDessert.Name)
			fmt.Println("Make sure to pop a breathmint after lunch, too!")
			totalCostUno := randLunch.Cost + randDessert.Cost
			totalCaloriesUno := randLunch.Calories + randDessert.Calories
			totalCost := strconv.FormatFloat(totalCostUno, 'f',2,64)
			totalCalories := strconv.Itoa(totalCaloriesUno)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost+"\n...and a breathmint!")

		} else if hour2 >= 1416 && hour2 <= 1700 || hour2 >= 1416 && hour2<=17{
			rand.Seed(time.Now().UnixNano())
			randSnack := Database.Snack[rand.Intn(len(Database.Snack)-1)]
			fmt.Println("Found your time to be between 2:15 pm and 5:00 pm, so... \nYour random snack is "+ randSnack.Name+", don't get too fat!")
			totalCost := strconv.FormatFloat(randSnack.Cost, 'f',2,64)
			totalCalories := strconv.Itoa(randSnack.Calories)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)
		} else if hour2 >= 1701 && hour2 <= 1920{
			rand.Seed(time.Now().UnixNano())
			randDinner := Database.Dinner[rand.Intn(len(Database.Dinner)-1)]
			randDessert := Database.Dessert[rand.Intn(len(Database.Dessert)-1)]
			fmt.Println("Found your time to be between 5:00 pm and 7:20 pm, so... \nYour randomly chosen dinner for the night is "+randDinner.Name)
			fmt.Println("Once you are done with that, you can enjoy "+randDessert.Name)
			totalCostUno := randDinner.Cost + randDessert.Cost
			totalCaloriesUno := randDinner.Calories + randDessert.Calories
			totalCost := strconv.FormatFloat(totalCostUno, 'f',2,64)
			totalCalories := strconv.Itoa(totalCaloriesUno)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)
		} else if hour2 >= 1921 && hour2 <= 2215{
			rand.Seed(time.Now().UnixNano())
			randSnack := Database.Snack[rand.Intn(len(Database.Snack)-1)]
			fmt.Println("Found your time to be between 7:20 pm and 10:15 pm, so... \nYour random snack is "+ randSnack.Name+", don't get too fat!")
			totalCost := strconv.FormatFloat(randSnack.Cost, 'f',2,64)
			totalCalories := strconv.Itoa(randSnack.Calories)
			fmt.Println(" Your Stat Line:\nCalories Consumed: "+totalCalories+"\nCost of Meal: "+totalCost)
		} else {
			fmt.Println("It isn't eating time, chief.")
		}
	case "3":
		fmt.Println("\n Shutting down.")
		os.Exit(0)
	default:
		panic("You entered something other than 1,2,3. Please read menu again for more info.")
	}
}
