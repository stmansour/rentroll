/* global
    GetFeeGridColumns, GetRAFlowCompLocalData,
    BuildRAFinalRentablesFeesGrid,
    BuildRAFinalPetsFeesGrid,
    BuildRAFinalVehiclesFeesGrid,
    AssignRAFinalRentablesFeesGridRecords,
    AssignRAFinalPetsFeesGridRecords,
    AssignRAFinalVehiclesFeesGridRecords,
    RenderFeesGridSummary, GetVehicleIdentity,
*/


"use strict";

// -----------------------------------------------------------------------------
// BuildRAFinalRentablesFeesGrid
//      initialize the fees grid for all rentables only in final screen summary
// -----------------------------------------------------------------------------
window.BuildRAFinalRentablesFeesGrid = function() {

    // if exists then return
    if (w2ui.hasOwnProperty("RAFinalRentablesFeesGrid")) {
        return true;
    }

    // first get grid columns for fees
    var cols = GetFeeGridColumns();

    // prepend RID, RentableName column in columns
    var prependCols = [
        {
            field: "RID",
            caption: "RID",
            hidden: true
        },
        {
            field: "Rentable",
            caption: "Rentable",
            size: "100px"
        }
    ];
    var gridCols = prependCols.concat(cols);

    // initialize rentable fees grid for final section
    $().w2grid({
        name: "RAFinalRentablesFeesGrid",
        header: "<strong>Rentable Fees</strong>",
        show: {
            toolbar:    false,
            header:     true,
            footer:     false,
        },
        multiSelect: false,
        style: 'display: block; background-color: transparent;',
        columns: gridCols,
        onSelect: function (event) {
            event.preventDefault(); // Prevent selection of row
        }
    });
};

// -----------------------------------------------------------------------------
// BuildRAFinalPetsFeesGrid
//      initialize the fees grid for all pets only in final screen summary
// -----------------------------------------------------------------------------
window.BuildRAFinalPetsFeesGrid = function() {

    // if exists then return
    if (w2ui.hasOwnProperty("BuildRAFinalPetsFeesGrid")) {
        return true;
    }

    // first get grid columns for fees
    var cols = GetFeeGridColumns();

    // prepend TMPPETID, PetName column in columns
    var prependCols = [
        {
            field: "TMPPETID",
            caption: "TMPPETID",
            hidden: true
        },
        {
            field: "Pet",
            caption: "Pet",
            size: "100px"
        }
    ];
    var gridCols = prependCols.concat(cols);

    // initialize pet fees grid for final section
    $().w2grid({
        name: "RAFinalPetsFeesGrid",
        header: "<strong>Pets Fees</strong>",
        show: {
            toolbar:    false,
            header:     true,
            footer:     false
        },
        multiSelect: false,
        style: 'display: block;',
        columns: gridCols,
        onSelect: function (event) {
            event.preventDefault(); // Prevent selection of row
        }
    });
};

// -----------------------------------------------------------------------------
// BuildRAFinalVehiclesFeesGrid
//      initialize the fees grid for all vehicles only in final screen summary
// -----------------------------------------------------------------------------
window.BuildRAFinalVehiclesFeesGrid = function() {

    // if exists then return
    if (w2ui.hasOwnProperty("RAFinalVehiclesFeesGrid")) {
        return true;
    }

    // first get grid columns for fees
    var cols = GetFeeGridColumns();

    // prepend TMPVID, VehicleName column in columns
    var prependCols = [
        {
            field: "TMPVID",
            caption: "TMPVID",
            hidden: true
        },
        {
            field: "Vehicle",
            caption: "Vehicle",
            size: "100px"
        }
    ];
    var gridCols = prependCols.concat(cols);

    // initialize vehicles fees grid for final section
    $().w2grid({
        name: "RAFinalVehiclesFeesGrid",
        header: "<strong>Vehicles Fees</strong>",
        show: {
            toolbar:    false,
            header:     true,
            footer:     false
        },
        multiSelect: false,
        style: 'display: block;',
        columns: gridCols,
        onSelect: function (event) {
            event.preventDefault(); // Prevent selection of row
        }
    });
};

// -----------------------------------------------------------------------------
// loadFinalSection
//      to load summary of fees grid in final section
//      maybe more content will be there
// -----------------------------------------------------------------------------
window.loadFinalSection = function() {

    // if components are not loaded then load
    BuildRAFinalRentablesFeesGrid();
    BuildRAFinalPetsFeesGrid();
    BuildRAFinalVehiclesFeesGrid();

    // load all fees grid in div
    $('#ra-form #final .container #rentable-fees').w2render(w2ui.RAFinalRentablesFeesGrid);
    $('#ra-form #final .container #pet-fees').w2render(w2ui.RAFinalPetsFeesGrid);
    $('#ra-form #final .container #vehicle-fees').w2render(w2ui.RAFinalVehiclesFeesGrid);

    setTimeout(function() {
        AssignRAFinalRentablesFeesGridRecords();
        AssignRAFinalPetsFeesGridRecords();
        AssignRAFinalVehiclesFeesGridRecords();
    }, 500);
};

// -----------------------------------------------------------------------------
// AssignRAFinalRentablesFeesGridRecords
//      render the grid records from local data of rentables section
// -----------------------------------------------------------------------------
window.AssignRAFinalRentablesFeesGridRecords = function() {
    var grid = w2ui.RAFinalRentablesFeesGrid;

    // clear the grid
    grid.clear();

    // get rentables local data
    var compData = GetRAFlowCompLocalData("rentables") || [];

    // loop over all rentable and collect fees
    compData.forEach(function(rentable) {

        // loop over fees collection and append in grid records
        rentable.Fees.forEach(function(feeItem) {
            // take a clone of local fee record
            var rentableFee = $.extend(true, {}, feeItem);
            rentableFee.Rentable = rentable.RentableName;
            rentableFee.RID = rentable.RID;
            grid.records.push(rentableFee);
        });

    });

    // render fees amount summary
    RenderFeesGridSummary(grid, grid.records);

    // assign record in grid
    reassignGridRecids(grid.name);
};

// -----------------------------------------------------------------------------
// AssignRAFinalPetsFeesGridRecords
//      render the grid records from local data of pets section
// -----------------------------------------------------------------------------
window.AssignRAFinalPetsFeesGridRecords = function() {
    var grid = w2ui.RAFinalPetsFeesGrid;

    // clear the grid
    grid.clear();

    // get pets local data
    var compData = GetRAFlowCompLocalData("pets") || [];

    // loop over all pets and collect fees
    compData.forEach(function(pet) {

        // loop over fees collection and append in grid records
        pet.Fees.forEach(function(feeItem) {
            // take a clone of local fee record
            var petFee = $.extend(true, {}, feeItem);
            petFee.Pet = pet.Name;
            petFee.TMPPETID = pet.TMPPETID;
            grid.records.push(petFee);
        });

    });

    // render fees amount summary
    RenderFeesGridSummary(grid, grid.records);

    // assign record in grid
    reassignGridRecids(grid.name);
};

// -----------------------------------------------------------------------------
// AssignRAFinalVehiclesFeesGridRecords
//      render the grid records from local data of vehicles section
// -----------------------------------------------------------------------------
window.AssignRAFinalVehiclesFeesGridRecords = function() {
    var grid = w2ui.RAFinalVehiclesFeesGrid;

    // clear the grid
    grid.clear();

    // get vehicles local data
    var compData = GetRAFlowCompLocalData("vehicles") || [];

    // loop over all vehicles and collect fees
    compData.forEach(function(vehicle) {

        // loop over fees collection and append in grid records
        vehicle.Fees.forEach(function(feeItem) {
            // take a clone of local fee record
            var vehicleFee = $.extend(true, {}, feeItem);

            vehicleFee.Vehicle = GetVehicleIdentity(vehicle);
            vehicleFee.TMPVID = vehicle.TMPVID;
            grid.records.push(vehicleFee);
        });

    });

    // render fees amount summary
    RenderFeesGridSummary(grid, grid.records);

    // assign record in grid
    reassignGridRecids(grid.name);
};
