/**
 * @param {number} n
 * @return {Function} counter
 */
var cur_num = 0;
var createCounter = function(n) {
    cur_num = n;
    return function() {
        return n++;
    };
};

/** 
 * const counter = createCounter(10)
 * counter() // 10
 * counter() // 11
 * counter() // 12
 */