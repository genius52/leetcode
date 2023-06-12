
Array.prototype.last = function() {
    let l = this.length;
    if (l == 0){
        return -1;
    }
    return this[l - 1];
};

/**
 * const arr = [1, 2, 3];
 * arr.last(); // 3
 */