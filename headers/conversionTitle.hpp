#define CYAN "\033[36m"
#define RESET "\033[0m"

void conversionTitle(int choice) {
  if (choice == 1) {
    cout << CYAN;
    cout << " -------------------------------------------------------- " << endl;
    cout << "|                                                        |" << endl;
    cout << "|             Celsius (C) -> Fahrenheit (F)              |" << endl;
    cout << "|________________________________________________________|" << endl;
    cout << CYAN RESET;
  }

  if (choice == 2) {
    cout << CYAN;
    cout << " -------------------------------------------------------- " << endl;
    cout << "|                                                        |" << endl;
    cout << "|             Celsius (C) -> Kelvin (K)                  |" << endl;
    cout << "|________________________________________________________|" << endl;
    cout << CYAN RESET;
  }

  if (choice == 3) {
    cout << CYAN;
    cout << " -------------------------------------------------------- " << endl;
    cout << "|                                                        |" << endl;
    cout << "|             Fahrenheit (F) -> Celsius (C)              |" << endl;
    cout << "|________________________________________________________|" << endl;
    cout << CYAN RESET;
  } 

  if (choice == 4) {
    cout << CYAN;
    cout << " -------------------------------------------------------- " << endl;
    cout << "|                                                        |" << endl;
    cout << "|             Fahrenheit (F) -> Kelvin (K)               |" << endl;
    cout << "|________________________________________________________|" << endl;
    cout << CYAN RESET;
  }

  if (choice == 5) {
    cout << CYAN;
    cout << " -------------------------------------------------------- " << endl;
    cout << "|                                                        |" << endl;
    cout << "|             Kelvin (K) -> Celsius (C)                  |" << endl;
    cout << "|________________________________________________________|" << endl;
    cout << CYAN RESET;
  }

  if (choice == 6) {
    cout << CYAN;
    cout << " -------------------------------------------------------- " << endl;
    cout << "|                                                        |" << endl;
    cout << "|             Kelvin (F) -> Fahrenheit (F)               |" << endl;
    cout << "|________________________________________________________|" << endl;
    cout << CYAN RESET;
  }
}

