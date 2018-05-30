"use strict";

// right side slider content div length
var sliderContentDivLength = 400;

// RACompConfig for each section
var RACompConfig = {
    "dates": {
        loader: "loadRADatesForm",
        w2uiComp: "RADatesForm",
        sliderWidth: 0
    },
    "people": {
        loader: "loadRAPeopleForm",
        w2uiComp: "RAPeopleForm",
        sliderWidth: 600
    },
    "pets": {
        loader: "loadRAPetsGrid",
        w2uiComp: "RAPetsGrid",
        sliderWidth: 400
    },
    "vehicles": {
        loader: "loadRAVehiclesGrid",
        w2uiComp: "RAVehiclesGrid",
        sliderWidth: 400
    },
    "rentables": {
        loader: "loadRARentablesGrid",
        w2uiComp: "RARentablesGrid",
        sliderWidth: 800
    },
    "parentchild": {
        loader: "loadRAPeopleChildSection",
        w2uiComp: "RAParentChildGrid",
        sliderWidth: 0
    },
    "tie": {
        loader: "loadRATieSection",
        w2uiComp: "",
        sliderWidth: 0
    },
    "final": {
        loader: "",
        w2uiComp: "",
        sliderWidth: 0
    }
};
