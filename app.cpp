#include <iostream>
#include <limits>
#define GREEN "\033[32m"  // * green
#define CYAN "\033[36m"   // * cyan
#define RED "\033[31m"    // * RED
#define RESET "\033[0m"   // * default
using namespace std;

// * function prototype
void clearScreen(void);
void applicationDescription(void);
void choicesDescription(void);

int main() {
  clearScreen();

  int choice;
  int conversionInput;
  int conversionOutput;

  // * descriptions
  applicationDescription();
  choicesDescription();
  cout << endl;

  // * inputs
  cout << "Option (1 | 2 | 3 | 4 | 5 | 6): ";
  cin >> choice;

  if (cin.fail()) {
    cout << "Something is wrong";
    cin.clear();
    cin.ignore(10000, '\n');
  }

  return 0;
}

void clearScreen() {
  cout << "\033c";
}

void applicationDescription() {
  cout << GREEN;
  cout << " ---------------------------------------------------------- " << endl;
  cout << "|                         TEMTER                           |" << endl;
  cout << "| A simple CLI temperature converter application that will |" << endl;
  cout << "| allow you to convert different unit of temperature.      |" << endl;
  cout << "|__________________________________________________________|" << endl;
  cout << GREEN RESET << endl;
}

void choicesDescription() {
  cout << CYAN;
  cout << "[1] Celsius to Fahrenheit" << endl;
  cout << "[2] Celsius to Kelvin" << endl;
  cout << "[3] Fahrenheit to Celsius" << endl;
  cout << "[4] Fahrenheit to Kelvin" << endl;
  cout << "[5] Kelvin to Celsius" << endl;
  cout << "[6] Kelvin to Fahrenheit" << endl;
  cout << CYAN RESET << endl;
}

