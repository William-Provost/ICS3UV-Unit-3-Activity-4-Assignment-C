/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-17
 * @fileoverview This program uses nested if statements to determine a user's life stage.
 */

// variables
let nameInput: string = "";
let ageInput: number = 0;
let studentInput: boolean = false;

// get name from user
nameInput = prompt("Enter your name: ") || "";

// get age from user
ageInput = parseInt(prompt("Enter your age: ") || "0");

// get student status
studentInput = (prompt("Are you a student (true/false): ") || "false") === "true";

// nested if checks
if (studentInput) {
  if (ageInput >= 13 && ageInput <= 19) {
    console.log(`Student ${nameInput} is a teenager.`);
  } else {
    if (ageInput >= 5 && ageInput <= 12) {
      console.log(`Student ${nameInput} is a child.`);
    } else {
      console.log(`Student ${nameInput} is in a different life stage.`);
    }
  }
} else {
  if (ageInput >= 20 && ageInput <= 30) {
    console.log(`${nameInput} is a young adult.`);
  } else {
    console.log(`${nameInput} is in a different life stage.`);
  }
}

console.log("\nDone.");
