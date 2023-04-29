#include <stdio.h>
#include <stdlib.h>
#define GREEN "\033[32m"  // * green
#define CYAN "\033[36m"   // * cyan
#define RED "\033[31m"    // * RED
#define RESET "\033[0m"   // * default

// * function prototype
// * defines the function name, return types, and parameters
void description(void);
void instruction(void);
int userChoice();
int isOneOfTheChoices(int);

int main() {
  /// * clear first the screen
  system("clear");

  int choice;

  // * application description
  description();

  // * instructions
  instruction();
  
  // * choices
  choice = userChoice();

  printf("%d", choice);

  return 0;
}

void description() {
  printf(GREEN);
  printf(" ----------------------------------------------------------   \n");
  printf("|                         TEMTER                           |  \n");
  printf("| A simple CLI temperature converter application that will |  \n");
  printf("| allow you to convert different unit of temperature.      |  \n");
  printf("|__________________________________________________________|  \n");
  printf(GREEN RESET);
}

void instruction() {
  printf("\n");
  printf("Choose an option:           \n");
  printf("[1] Celsius to Fahrenheit   \n");
  printf("[2] Celsius to Kelvin       \n");
  printf("[3] Fahrenheit to Celsius   \n");
  printf("[4] Fahrenheit to Kelvin    \n");
  printf("[5] Kelvin to Celsius       \n");
  printf("[6] Kelvin to Fahrenheit    \n");
}

int userChoice() {
  int choice;

  printf("\n");
  printf(CYAN);
  printf("Your choice (1 | 2 | 3 | 4 | 5 | 6): ");
  scanf("%d", &choice);
  printf(CYAN RESET);
  printf("\n");
  
  if (choice == 1) {
    return 1;
  } else if (choice == 2) {
    return 2;
  } else {
    return 0;
  }
}

