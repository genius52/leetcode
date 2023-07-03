/**
 * @param {integer} init
 * @return { increment: Function, decrement: Function, reset: Function }
 */
var createCounter = function(init) {
    this.origin = init;
    this.cur = init;
    this.increment = function(){
        this.cur++;
        return this.cur;
    }
    this.reset = function(){
        this.cur = this.origin;
        return this.cur;
    }
    this.decrement = function(){
        this.cur--;
        return this.cur;
    }
    return this;
};

/**
 * const counter = createCounter(5)
 * counter.increment(); // 6
 * counter.reset(); // 5
 * counter.decrement(); // 4
 */