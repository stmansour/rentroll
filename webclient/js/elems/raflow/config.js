/* global
    RACompConfig, SliderContentDivLength
*/
"use strict";

// right side slider content div length
window.SliderContentDivLength = 400;

// RACompConfig for each section
window.RACompConfig = {
    "dates": {
        loader: "loadRADatesForm",
        w2uiComp: "RADatesForm",
        sliderWidth: 0
    },
    "people": {
        loader: "loadRAPeopleGrid",
        w2uiComp: "RAPeopleGrid",
        sliderWidth: 600
    },
    "pets": {
        loader: "loadRAPetsGrid",
        w2uiComp: "RAPetsGrid",
        sliderWidth: 850
    },
    "vehicles": {
        loader: "loadRAVehiclesGrid",
        w2uiComp: "RAVehiclesGrid",
        sliderWidth: 850
    },
    "rentables": {
        loader: "loadRARentablesGrid",
        w2uiComp: "RARentablesGrid",
        sliderWidth: 850
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
        loader: "loadFinalSection",
        w2uiComp: "",
        sliderWidth: 0
    }
};
