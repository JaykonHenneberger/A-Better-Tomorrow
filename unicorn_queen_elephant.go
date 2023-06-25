package main
 
import "fmt"
 
func main() { 
	// Print welcome message 
	fmt.Println("Welcome to our charity organization!")

	// Prompt user for their donation 
	var donation float64 
	fmt.Print("Please enter your donation: ") 
	fmt.Scan(&donation)
 
	// Thank user for donation 
	fmt.Println("Thank you for your generous donation of", donation, "dollars!") 


	// Keep track of donations 
	var totalDonations float64 
	totalDonations = totalDonations + donation 
 
	// Ask user if they would like to continue donating 
	var answer string 
	fmt.Print("Would you like to make another donation (yes/no)? ") 
	fmt.Scan(&answer) 

	// Loop to keep track of donations until user says "no" 
	for answer == "yes" { 
		fmt.Print("Please enter your donation: ") 
		fmt.Scan(&donation) 
		totalDonations = totalDonations + donation 
		fmt.Print("Would you like to make another donation (yes/no)? ") 
		fmt.Scan(&answer) 
	} 

	// Print total amount donated 
	fmt.Println("Thank you for your donations totaling", totalDonations, "dollars.") 
 
	// Calculate donations for charity project 
	percentage := .90 
	projectDonations := totalDonations * percentage 
	fmt.Println("We will use", projectDonations, "dollars for our charity project.") 
 
	// Calculate leftover donations 
	remainder := totalDonations - projectDonations 
	fmt.Println("We will place the remaining", remainder, "dollars in our reserve fund.") 
 
	// Print thank you message 
	fmt.Println("Thank you for making a difference in the world!") 
}