const util = require('util');


class Node {
    value = null;
    next = null;
    prev = null;
    constructor(value = null) {
        this.value = value;
        this.prev = null;
        this.next = null;
    }
}


class DoublyLinkedList {
    head = null;
    tail = null;
    length = 0;
    constructor(value) {
        const newNode = new Node(value);
        this.head = newNode;
        this.tail = newNode;
        this.length = 1;
    }

    printDoublyLinkedList(appendText="") {
        const values = [];
        let currentNode = this.head;
        while(currentNode) {
            values.push(currentNode.value);
            currentNode = currentNode.next;
        }
        if (appendText !== "") {
            console.log(appendText);
        }
        console.log(`list  :   ${values}`);
        console.log(`head  :   ${this.head.value}`);
        console.log(`tail  :   ${this.tail.value}`);
        return;
    }

    printDoublyLinkedListStringified(appendText="") {
        if (appendText !== "") {
            console.log(`\n\n\n${appendText}`);
        }
        console.dir(this, { depth: null, colors: true });
        return;
    }

    push(value) {
        // create a new node
        const newNode = new Node(value);
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
            this.length = 1;
            return;
        }
        this.tail.next = newNode;
        newNode.prev = this.tail;
        this.tail = newNode;
        this.length += 1;
        return;
    }

    pop() {
        if (!this.head) return;
        if (!this.head.next) { // length 1
            this.head = null;
            this.tail = null;
            this.length -= 1;
            return;
        }
        
        const prevOfTailNode = this.tail.prev;
        prevOfTailNode.next = null;
        this.tail.prev = null;
        this.tail = prevOfTailNode;
        this.length -= 1;
        return;
    }

    unshift(value) {
        const newNode = new Node(value);
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
            this.length += 1;
            return;
        }

        newNode.next = this.head;
        this.head.prev = newNode;
        this.head = newNode;
        this.length += 1;
        return;
    }

    shift() {
        if (!this.head) return;
        if (!this.head.next) { // length 1
            this.head = null;
            this.tail = null;
            this.length -= 1;
            return;
        }

        const newHead = this.head.next;
        newHead.prev = null;
        this.head.next = null;
        this.head = newHead;
        this.length -= 1;
        return;
    }

    // considering 0 based indices
    get(index) {
        if (!this.head || index < 0 || index > (this.length - 1)) return undefined;

        let currentNode = this.head;
        let counter = 0;
        while (currentNode) {
            if (index === counter) return currentNode.value;
            currentNode = currentNode.next;
            counter += 1;
        }
        return undefined;
    }

    // considering 0 based indices
    // more efficient code
    get2(index) {
        if (!this.head || index < 0 || index > (this.length - 1)) return null;
        const half = Math.floor(this.length / 2);
        let currentNode = this.head;
        let counter = 0;
        if (index < half) {
            while (counter < index) {
                currentNode = currentNode.next;
                counter += 1;
            }
            return currentNode;
        }

        counter = this.length - 1;
        currentNode = this.tail;
        while (counter > index) {
            currentNode = currentNode.prev;
            counter -= 1;
        }
        return currentNode;
    }

    get2values(index) {
        const temp = this.get2(index);
        return temp?.value ? temp.value : null;
    }

    set(index, value) {
        const found = this.get2(index);
        if (found) {
            found.value = value;
            return true;
        }
        return false;
    }

    insert(index, value) {
        if (!this.head || index < 0 || index > (this.length - 1)) return false;
        const newNode = new Node(value);

        if (index === 0) {
            this.unshift(value);
            return;
        } else if (index === this.length) {
            this.push(value);
            return;
        }

        const half = Math.floor(this.length / 2);
        let counter = 0;
        let currentNode = this.head;
        if (index < half) {
            while (counter < index) {
                currentNode = currentNode.next;
                counter += 1;
            }
        } else {
            currentNode = this.tail;
            counter = this.length - 1;
            while (counter > index) {
                currentNode = currentNode.prev;
                counter -= 1;
            }
        }
        // console.log(`current node: ${JSON.stringify(currentNode.value)}`);
        const prevNode = currentNode.prev;
        prevNode.next = newNode;
        newNode.prev = prevNode;
        newNode.next = currentNode;
        currentNode.prev = newNode;
        this.length += 1;
        return true;
    }


    remove(index) {
       if (!this.head || index < 0 || index > (this.length - 1)) return false;
       if (index === 0) { // remove head
            const tempNode = this.head;
            if (this.length === 1) {
                this.head = null;
                this.tail = null;
            } else {
                this.head = tempNode.next;
                this.head.prev = null;
                tempNode.prev = null;
                tempNode.next = null;
            }
            this.length -= 1;
            return true; 
       }
       if (index === this.length - 1) { // remove tail
            const tempNode = this.tail;
            this.tail = this.tail.prev;
            this.tail.next = null;
            tempNode.next = null;
            tempNode.prev = null;
            this.length -= 1;
            return true;
       }
       let removeNode;
       const half = Math.floor(this.length / 2);
       let counter;
       if (index < half) {
            removeNode = this.head;
            counter = 0;
            while (counter < index) {
                removeNode = removeNode.next;
                counter += 1;
            }
       } else {
            counter = this.length - 1;
            removeNode = this.tail;
            while (counter > index) {
                removeNode = removeNode.prev;
                counter -= 1;
            }
       }



       const prevNode = removeNode.prev;
       const nextNode = removeNode.next;

       prevNode.next = nextNode;
       nextNode.prev = prevNode;

       removeNode.next = null;
       removeNode.prev = null;
       this.length -= 1;
       return false;
    }


}

// const dll1 = new DoublyLinkedList(1);
// dll1.printDoublyLinkedListStringified();

// // testing push
// [2,3,4,5,6].forEach(x => dll1.push(x));
// dll1.printDoublyLinkedListStringified("After pushing new nodes");

// IMP
// uncomment above code to test below until "*THIS_COMMENT"


// // testing pop
// dll1.pop();
// dll1.printDoublyLinkedListStringified("After pop");
// // testing pop



// // testing unshift
// dll1.unshift(11);
// dll1.printDoublyLinkedListStringified("After unshift");
// // testing unshift


// /*  
// / testing shift
// */
// dll1.shift();
// dll1.printDoublyLinkedListStringified("After shift");
// /*  
// / testing shift
// */


/*  
/ testing get
*/
// console.log(dll1.get(4));
/*  
/ testing get
*/

// /*  
// / testing get2 efficient
// */
// console.log(`0th value ${dll1.get2values(0)}`);
// console.log(`1th value ${dll1.get2values(1)}`);
// console.log(`2th value ${dll1.get2values(2)}`);
// console.log(`3th value ${dll1.get2values(3)}`);
// console.log(`4th value ${dll1.get2values(4)}`);
// console.log(`5th value ${dll1.get2values(5)}`);
// console.log(`6th value ${dll1.get2values(6)}`);
// console.log(`66th value ${dll1.get2values(66)}`);
// /*  
// / testing get2 efficient
// */

// THIS_COMMENT



// /*  
// / testing set method
// */
// let dll = new DoublyLinkedList(11);
// [2,3,44,55,66].forEach(x => dll.push(x));
// dll.printDoublyLinkedListStringified("BEFORE SET");
// dll.set(0,1);
// dll.set(3,4);
// dll.set(4, 5);
// dll.set(5,6);
// dll.printDoublyLinkedListStringified("AFTER SET");
// /*  
// / testing set method
// */



// /*  
// / testing insert method
// */
// let dll = new DoublyLinkedList(2);
// [3,5,7].forEach(x => dll.push(x));
// dll.printDoublyLinkedListStringified("BEFORE INSERT");
// // console.log(dll.get2values(0));
// // console.log(dll.get2values(1));
// // console.log(dll.get2values(2));
// // console.log(dll.get2values(3));
// dll.insert(0, 1);
// dll.insert(3, 4);


// dll.printDoublyLinkedListStringified("AFTER INSERT");
// console.log("5th index holds value 7 currently, insert there 6");
// dll.insert(5, 6); // 5th index holds value 7 currently, insert there 6

// dll.printDoublyLinkedListStringified("AFTER INSERT LAST");
// /*  
// / testing insert method
// */



/*  
/ testing remove method
*/
let dll = new DoublyLinkedList(1);
[2,3,4,5,6,7].forEach(x => dll.push(x));
dll.printDoublyLinkedListStringified("BEFORE REMOVE");
dll.remove(0);
dll.printDoublyLinkedListStringified("AFTER REMOVE HEAD");
console.log("lets remove last/tail");
dll.remove(5);
dll.printDoublyLinkedListStringified("AFTER REMOVE LAST/TAI");

console.log("lets remove in between");
dll.remove(3);
dll.printDoublyLinkedListStringified("AFTER REMOVE IN BETWEEN");
// dll.printDoublyLinkedListStringified("AFTER INSERT LAST");
/*  
/ testing remove method
*/
