#include <iostream>
#include "headers/clearScreen.h"
#include "headers/cleanInputBuffer.h"
#include "headers/applicationDescription.h"
#include "headers/choicesDescription.h"
#include "headers/conversionTitle.h"
#include "headers/conversionCalculation.h"
#define GREEN "\033[32m"
#define CYAN "\033[36m"
#define RED "\033[31m"
#define YELLOW "\033[31m"
#define RESET "\033[0m"
using namespace std;

int main() {
  clearScreen();

  int choice;
  float input;
  float output;
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
    cin >> input;
    
    if (cin.fail()) {
      validOption = false;
    } else {
      output = conversionCalculation(choice, input);
      validOption = true;
    }
  
    clearScreen();
    cleanInputBuffer();
  } while (!validOption);

  return 0;
}

