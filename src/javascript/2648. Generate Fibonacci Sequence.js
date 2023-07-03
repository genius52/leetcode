



var fibGenerator = function*() {
    var pre1 = 0;
    var pre2 = 1;
    while(true){
        cur = pre1 + pre2;
        yield cur;
        pre1 = pre2;
        pre2 = cur;
    };
};