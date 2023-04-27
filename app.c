#include <stdio.h>
#define GREEN "\033[32m"  // * green
#define RESET "\033[0m"   // * default

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
  printf("\n");
  printf("Choose an option:           \n");
  printf("[1] Celsius to Fahrenheit   \n");
  printf("[2] Celsius to Kelvin       \n");
  printf("[3] Fahrenheit to Celsius   \n");
  printf("[4] Fahrenheit to Kelvin    \n");
  printf("[5] Kelvin to Celsius       \n");
  printf("[6] Kelvin to Fahrenheit    \n");
  
  // * choice
  printf("\n");
  printf("Your choice (1 | 2 | 3 | 4 | 5 | 6): ");
  scanf("%d", &choice);
  printf("\n");

  return 0;
}

