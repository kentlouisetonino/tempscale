#include <iostream>
#include "headers/clearScreen.hpp"
#include "headers/cleanInputBuffer.hpp"
#include "headers/applicationDescription.hpp"
#include "headers/choicesDescription.hpp"
#include "headers/conversionTitle.hpp"
#include "headers/conversionCalculation.hpp"
#include "headers/optionSelectionPrompt.hpp"
#include "headers/constants.hpp"
using namespace std;

int main() {
  clearScreen();

  int choice;
  float input;
  float output;
  bool validOption = true;

  choice = optionSelectionPrompt();

  do {
    cout << endl;
    conversionTitle(choice);
    cout << endl << endl;

    if (!validOption) {
      cout << RED;
      cout << "Input is not valid. Please try again.";
      cout << endl << endl;
      cout << RED << RESET;
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

