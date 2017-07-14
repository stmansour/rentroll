function requireAll(r) {
    r.keys().forEach(r);
}

// @param1: Location of javascript file
// @param2: Include sub Directory
// @param3: Regex for use the all .js file
requireAll(require.context('./src/', true, /\.js$/));
