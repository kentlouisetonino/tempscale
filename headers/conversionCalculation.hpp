float conversionCalculation(int choice, float input) {
  float result = 0;
  
  // * celsius -> fahrenheit
  if (choice == 1) {
    result = ((9.0 * input) / 5) + 32;
  }
  
  // * celsius -> kelvin
  if (choice == 2) {
    result = input + 273.15;
  }
  
  // * fahrenheit -> celsius
  if (choice == 3) {
    result = ((5.0 * input) - (5.0 * 32.0)) / 9.0;
  }
  
  // * fahrenheit -> kelvin
  if (choice == 4) {
    result = (((input * 5.0) - (32 * 5)) / 9.0) + 273.15;
  }
  
  // * kelvin -> celsius
  if (choice == 5) {
    result = input - 273.15;
  }
  
  // * kelvin -> fahrenheit
  if (choice == 6) {
    result = ((input * 9.0) - (273.15 * 9) / 5) + 32;
  }

  return result;
}

