int optionSelectionPrompt()
{
  int choice;
  bool validOption = true;

  do
  {
    applicationDescription();
    choicesDescription();
    cout << endl;

    if (!validOption)
    {
      cout << RED;
      cout << "Input is not valid. Please try again." << endl;
      cout << endl;
      cout << RED << RESET;
    }

    // * inputs
    cout << "Option (1 | 2 | 3 | 4 | 5 | 6): ";
    cin >> choice;

    if (cin.fail())
    {
      validOption = false;
    }
    else
    {
      cout << choice;

      if (
          choice == 1 ||
          choice == 2 ||
          choice == 3 ||
          choice == 4 ||
          choice == 5 ||
          choice == 6)
      {
        validOption = true;
      }
      else
      {
        validOption = false;
      }
    }

    clearScreen();
    cleanInputBuffer();
  } while (!validOption);

  return choice;
}
