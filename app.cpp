#include <iostream>
#include "headers/clearScreen.h"
#include "headers/cleanInputBuffer.h"
#include "headers/applicationDescription.h"
#include "headers/choicesDescription.h"
#include "headers/conversionTitle.h"
#define GREEN "\033[32m"
#define CYAN "\033[36m"
#define RED "\033[31m"
#define YELLOW "\033[31m"
#define RESET "\033[0m"
using namespace std;

float conversionProcess(int, float);

int main() {
  clearScreen();

  int choice;
  float conversionInput;
  float conversionOutput;
  bool validOption = true;

  do {
    applicationDescription();
    choicesDescription();
    cout << endl;
    
    if (!validOption) {
      cout << RED;
      cout << "Input is not valid. Please try again." << endl;
      cout << endl;
      cout << RED RESET;
    }
  
    // * inputs
    cout << "Option (1 | 2 | 3 | 4 | 5 | 6): ";
    cin >> choice;

    if (cin.fail()) {
      validOption = false;
    } else {
      cout << choice;

      if (
        choice == 1 ||
        choice == 2 ||
        choice == 3 ||
        choice == 4 ||
        choice == 5 ||
        choice  == 6
      ) {
        validOption = true;
      } else {
        validOption = false;
      }
    }

    clearScreen();
    cleanInputBuffer();
  } while (!validOption);

  do {
    cout << endl;
    conversionTitle(choice);
    cout << endl << endl;

    if (!validOption) {
      cout << RED;
      cout << "Input is not valid. Please try again.";
      cout << endl << endl;
      cout << RED RESET;
    }

    cout << "Please input the value to convert: ";
    cin >> conversionInput;
    
    if (cin.fail()) {
      validOption = false;
    } else {
      conversionOutput = conversionProcess(choice, conversionInput);
      validOption = true;
    }

    cout << conversionInput;
    cout << endl;
  
    clearScreen();
    cleanInputBuffer();
  } while (!validOption);

  return 0;
}

float conversionProcess(int choice, float conversionInput) {
  float result = 0;
  
  // * celsius -> fahrenheit
  if (choice == 1) {
    result = ((9.0 * conversionInput) / 5) + 32;
  }
  
  // * celsius -> kelvin
  if (choice == 2) {
    result = conversionInput + 273.15;
  }
  
  // * fahrenheit -> celsius
  if (choice == 3) {
    result = ((5.0 * conversionInput) - (5.0 * 32.0)) / 9.0;
  }
  
  // * fahrenheit -> kelvin
  if (choice == 4) {
    result = (((conversionInput * 5.0) - (32 * 5)) / 9.0) + 273.15;
  }
  
  // * kelvin -> celsius
  if (choice == 5) {
    result = conversionInput - 273.15;
  }
  
  // * kelvin -> fahrenheit
  if (choice == 6) {
    result = ((conversionInput * 9.0) - (273.15 * 9) / 5) + 32;
  }

  return result;
}
