/**
 * @param {Array} arr
 * @return {Generator}
 */

function recrusive_visit(arr,data){
    var l = arr.length;
    for(var i = 0;i < l;i++){
        if(typeof(arr[i]) == "number"){
            data.push(arr[i]);
        }else{
            recrusive_visit(arr[i],data)
        }
    }
}

var inorderTraversal = function*(arr) {
    var data = [];
    recrusive_visit(arr,data);
    var l = data.length;
    var idx = 0;
    while(idx < l){
        yield data[idx];
        idx++;
    }
};

arr = [[[6]],[1,3],[]];
const generator = inorderTraversal(arr);
generator.next().value; // 6
generator.next().value; // 1
/**
 * const gen = inorderTraversal([1, [2, 3]]);
 * gen.next().value; // 1
 * gen.next().value; // 2
 * gen.next().value; // 3
 */