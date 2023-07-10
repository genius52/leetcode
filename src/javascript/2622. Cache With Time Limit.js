var TimeLimitedCache = function() {
    this.key_val = {};
    this.time = new Date().getTime();
    //var key_duration = {};
};

/** 
 * @param {number} key
 * @param {number} value
 * @param {number} duration time until expiration in ms
 * @return {boolean} if un-expired key already existed
 */
TimeLimitedCache.prototype.set = function(key, value, duration) {
    var ret = false;
    cur_time = new Date().getTime();
    if(key in this.key_val){
        clearTimeout(this.key_val[key].timeout);
        ret = true;
    }
    //this.key_val[key] = {"val":value,"expire":cur_time + duration};
    timeout = setTimeout(() => {
        if(key in this.key_val){
            delete this.key_val[key];
        }
    }, duration);
    this.key_val[key] = {value,timeout};
    return ret;
};

/** 
 * @param {number} key
 * @return {number} value associated with key
 */
TimeLimitedCache.prototype.get = function(key) {
    if(key in this.key_val){
        return this.key_val[key].value;
    }else{
        return -1;
    }
};

/** 
 * @return {number} count of non-expired keys
 */
TimeLimitedCache.prototype.count = function() {
    return Object.keys(this.key_val).length;
};

/**
 * Your TimeLimitedCache object will be instantiated and called as such:
 * var obj = new TimeLimitedCache()
 * obj.set(1, 42, 1000); // false
 * obj.get(1) // 42
 * obj.count() // 1
 */