/**
 * @param {string} val
 * @return {Object}
 */
var expect = function(val) {
    function toBe(v) {
         if (val === v){return true} else{ throw Error("Not Equal")};
    }
    function notToBe(v) {
        if (val !== v){return true} else{ throw Error("Equal")};
    }
    return {toBe, notToBe}  
};

/**
 * expect(5).toBe(5); // true
 * expect(5).notToBe(5); // throws "Equal"
 */
