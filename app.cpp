#include <iostream>
#define GREEN "\033[32m"  // * green
#define CYAN "\033[36m"   // * cyan
#define RED "\033[31m"    // * RED
#define RESET "\033[0m"   // * default
using namespace std;

// * function prototype
// * defines the function name, return types, and parameters
void description(void);
void instruction(void);
int userChoice();
int isOneOfTheChoices(int);

int main() {
  // * clear first the screen
  cout << "\033c";

  // * application description
  description();

  return 0;
}

void description() {
  int proceed;

  cout << GREEN;
  cout << " ---------------------------------------------------------- " << endl;
  cout << "|                         TEMTER                           |" << endl;
  cout << "| A simple CLI temperature converter application that will |" << endl;
  cout << "| allow you to convert different unit of temperature.      |" << endl;
  cout << "|__________________________________________________________|" << endl;
  cout << GREEN RESET << endl;
}

