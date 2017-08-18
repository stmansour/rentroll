/*global
  console, QUnit, number_format,
*/
"use strict";
QUnit.log(function QunitLOG( details ) {
  if ( details.result ) {
    console.log("Test Module: "+ details.module ,"-- Routine name: "+ details.name, "-- Testcase Result: "+ details.result, "-- Detail: " + details.message);
    return;
  }
  var loc = details.module + ": " + details.name + ": ",
    output = "FAILED: " + loc + ( details.message ? details.message + ", " : "" );
  console.log( output );
});

QUnit.done(function QunitDone( details ) {
  console.log("-----------------------------------------------------");
  console.log( "Total: ", details.total, "\nFailed: ", details.failed, "\nPassed: ", details.passed, "\nRuntime: ", details.runtime);
  console.log("-----------------------------------------------------");
});

QUnit.module("Plural Test");
QUnit.test('plural test', function pluralTest(assert, details){
    assert.ok(plural("accord") == "accords", "accord");
    assert.ok(plural("a") == "as", "a");
});

QUnit.module("Int to bool Test");
QUnit.test('int_to_bool', function intToBoolTest(assert){
    assert.ok(int_to_bool(-1) === false, "-1");
    assert.ok(int_to_bool(3) === true, "3");
});

QUnit.module("dateTodayStr Test");
QUnit.test('dateTodayStr', function dateTodayStrTest(assert){
    var today = new Date();
    var dd = today.getDate();
    var mm = today.getMonth() + 1; //January is 0!
    var yyyy = today.getFullYear();
    var today_date =  mm + '/' + dd + '/' + yyyy;
    assert.ok(dateTodayStr() == today_date, "dateTodayStr");
});

QUnit.module("Number Format Test");
QUnit.test('number_format', function numberFormatTest(assert){
    assert.ok(number_format(1000) == "1,000", "1000");
    assert.ok(number_format(1234.56) == "1,235", "1234.56");
    assert.ok(number_format(1234.56, 2, ',', ' ') == "1 234,56", "1234.56, 2, ',', ' '");
    assert.ok(number_format(1234.5678, 2, '.', '') == "1234.57", "1234.5678, 2, '.', ''");
    assert.ok(number_format(67, 2, ',', '.') == "67,00", "67, 2, ',', '.'");
    assert.ok(number_format(1000.55, 1) == "1,000.6", "1000.55, 1");
    assert.ok(number_format(67000, 5, ',', '.') == "67.000,00000", "67000, 5, ',', '.'");
    assert.ok(number_format(0.9, 0) == "1", "0.9, 0");
    assert.ok(number_format('1.20', 4) == "1.2000", "'1.20', 4");
    assert.ok(number_format('1.2000', 3) == "1.200", "'1.2000', 3");
});

