let obj1 = {
    value : 5
}
let obj2 = obj1
console.log(`obj1 ${obj1.value}`) // 5 
console.log(`obj2 ${obj2.value}`) // 5

obj1.value = 10
console.log(`obj1 ${obj1.value}`) // 10
console.log(`obj2 ${obj2.value}`) // 10