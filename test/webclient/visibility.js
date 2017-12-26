var isVisibleInViewPort = function(el){

    //Check element exists in DOM
    if(el){
        var rect = el.getBoundingClientRect(), top = rect.top, height = rect.height,
            el = el.parentNode;
        do {
            rect = el.getBoundingClientRect();
            if (top <= rect.bottom === false) return false;
            // Check if the element is out of view due to a container scrolling
            if ((top + height) <= rect.top) return false;
            el = el.parentNode;
        } while (el !== document.body);
        // Check its within the document viewport
        return top <= document.documentElement.clientHeight;
    }
    else {
        return false;
    }
};