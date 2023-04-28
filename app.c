#include <stdio.h>
#include <stdlib.h>
#define GREEN "\033[32m"  // * green
#define CYAN "\033[36m"   // * cyan
#define RED "\033[31m"    // * RED
#define RESET "\033[0m"   // * default

// * function prototype
// * defines the function name, return types, and parameters
void instruction(void);
int userChoice();
int isOneOfTheChoices(int);

int main() {
  int choice;

  // * application description
  printf(GREEN);
  printf(" ----------------------------------------------------------   \n");
  printf("|                         TEMTER                           |  \n");
  printf("|                                                          |  \n");
  printf("| A simple CLI temperature converter application that will |  \n");
  printf("| allow you to convert different unit of temperature.      |  \n");
  printf("|__________________________________________________________|  \n");
  printf(GREEN RESET);

  // * instructions
  instruction();
  
  // * choices
  choice = userChoice();

  // * validate
  if (choice == 0) {
    do {
      system("clear");
      printf(RED);
      printf(" -----------------------------------------------------------    \n");
      printf("|                                                           |   \n");
      printf("|         YOUR INPUT IS INVALID. PLEASE TRY AGAIN.          |   \n");
      printf("|___________________________________________________________|   \n");
      printf(RED RESET);

      choice = userChoice();
    } while(choice == 0);
  }

  return 0;
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
  int isValidChoice;
  char str[100];

  printf("\n");
  printf(CYAN);
  printf("Your choice (1 | 2 | 3 | 4 | 5 | 6): ");
  scanf("%d\n", &choice);
  fgets(str, 100, stdin);
  printf(CYAN RESET);
  printf("\n");

  isValidChoice = isOneOfTheChoices(choice);


  return isValidChoice;
}

int isOneOfTheChoices(int choice) {
  int lowestChoice = 1;
  int highestChoice = 6;
    
  if (lowestChoice <= choice && choice <= highestChoice) {
    return 1;
  } else {
    return 0;
  }
}

