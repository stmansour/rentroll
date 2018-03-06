# Automated UI Tests via Cypress.IO

## AIR Receipt application (/rhome)

List of covered use cases:
- Assertion of application title
- Left side node selection(Tendered Payment Receipts)
- Grid rendering after selection of Tendered Payment Receipts node
- Assertion of cell value in the receipt grid with API response(/v1/receipts/1) 
- Checking visibility and value of input fileds for detail record. (We're clicking on first record in grid and performing tests on detail record form.) Except fields: `Resident Address(ERentableName)`
- Checking visibility and default value of input fields for adding new record. 
- Buttons visibility in add new record form / detail record form
- Visibility and class of Unallocated section in detail record from
- Print Receipt UI PopUp
- Close the form and checking that form is closed properly.

## AIR Roller application (/home)

List of covered use cases for Assess Charges, Tendered Receipt Payment, Chart of accounts, Payment Types, Deposit Methods, Deposit Accounts, Account Rules:
- Assertion of application title
- Left side node selection
- Grid rendering after selection of node
- Assertion of cell value in the  grid with API response(/v1/<grid_name>/1) 
- Checking visibility and value of input fileds for detail record. (We're clicking on first record in grid and performing tests on detail record form.) 
- Checking visibility and default value of input fields for adding new record.
- Buttons visibility in add new record form / detail record form
- Visibility and class of Unallocated section(IF Exists) in detail record from
- Close the form and checking that form is closed properly.