/**
 * @param {Array} arr
 * @param {number} size
 * @return {Array[]}
 */
var chunk = function(arr, size) {
    var l = arr.length;
    var res = [];
    var i = 0;
    while(i < l){
        var cur = []
        var j = 0
        while(i < l && j < size){
            cur.push(arr[i])
            i++;
            j++;
        }
        res.push(cur);
    }
    return res;
};