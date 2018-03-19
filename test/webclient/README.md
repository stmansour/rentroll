# Automated UI Tests via Cypress.IO

## AIR Receipt application (/rhome)

List of covered use cases:
- An assertion of application title
- Left side node selection(Tendered Payment Receipts)
- Grid rendering after selection of Tendered Payment Receipts node
- An assertion of cell value in the receipt grid with API response(/v1/receipts/1) 
- Checking visibility and value of input fields for detail record. (We're clicking on the first record in the grid and performing tests on detail record form.) Except fields: `Resident Address(ERentableName)`
- Checking visibility and the default value of input fields for adding a new record. 
- Buttons visibility in add new record form/detail record form
- Visibility and class of Unallocated section in detail record from
- Print Receipt UI PopUp
- Close the form and checking that form is closed properly.

## AIR Roller application (/home)

List of covered use cases for Assess Charges, Tendered Receipt Payment, 
Chart of accounts, Payment Types, 
Deposit Methods, Deposit Accounts, 
Account Rules, Expense, 
Deposit:
- Left side node selection
- Grid rendering after selection of a node
- Assertion of cell value in the  grid with API response(/v1/<grid_name>/1) 
- Checking visibility and value of input fields for detail record. (We're clicking on the first record in a grid and performing tests on detail record form.) 
- Checking visibility and a default value of input fields for adding a new record.
- Buttons visibility in add new record form/detail record form
- Visibility and class of Unallocated section(IF Exists) in detail record from
- Close the form and checking that form is closed properly.

Rent Roll, RA Statements,
Payor Staments:
- Left side node selection
- Grid rendering after selection of a node
- Assertion of cell value in the  grid with API response(/v1/<grid_name>/1) 
- Checking visibility and value of input fields/Grids in form/tabbed form for detail record. (We're clicking on the first record in a grid and performing tests on detail record form.)
- Close the form and checking that form is closed properly.  

## Deposit section
- Deposit list grid test in detailed record form