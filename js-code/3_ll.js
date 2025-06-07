class Node {
	constructor(value) {
		this.value = value;
		this.next = null;
	}
}

class LinkedList {
	head = null;
	tail = null;
	length = 0;
	constructor(value) {
		const newNode = new Node(value);
		this.head = newNode;
		this.tail = this.head;
		this.length = 1;
	}

	push(value) {
		const newNode = new Node(value);
		if (this.head === null) {
			this.head = newNode;
			this.tail = this.head;
			this.length = 1;
			return;
		}

		this.tail.next = newNode;
		this.tail = newNode;
		this.length += 1;
		return;
	}

	pop() {
		if (!this.head) {
			return;
		}

		if (!this.head.next) {
			this.head = null;
			this.tail = null;
			this.length = 0;
			return;
		}


		let curr = this.head;
		let prev = null;

		while (curr.next) {
			prev = curr;
			curr = curr.next;

			// console.log(`curr: ${JSON.stringify(curr)}`);
			// console.log(`prev: ${JSON.stringify(prev)}`);
			// console.log(``);
		}
		prev.next = null;
		this.tail = prev;
		this.length -= 1;
		return;
	}

	unshift(value) {
		if (!this.head) {
			const newNode = new Node(value);
			this.head = newNode;
			this.tail = this.head;
			this.length += 1;
			return;
		}


		const newNode = new Node(value);
		newNode.next = this.head;
		this.head = newNode;
		this.length += 1;
		return;
	}

	shift() {
		if (!this.head) {
			return;
		}
		// we have only one element which is head and tail both
		if (!this.head.next) {
			this.head = null;
			this.tail = null;
			this.length -= 1;
			return;
		}

		const secondEle = this.head.next;
		this.head.next = null;
		this.head = secondEle;
		this.length -= 1;
		return;
	}


	get(index) {
		if (!this.head) {
			return null;
		}
		if (index === 0) {
			return this.head;
		}

		let counter = 0;
		let currNode = this.head;
		let returnNode = null;
		while (currNode.next) {
			if (counter === index) {
				returnNode = currNode;
				break;
			}
			// console.log(`index counter ${counter}`);
			currNode = currNode.next;
			counter += 1;
		}

		// edge case if missing the last element
		if (counter === index) {
			return currNode;
		}
		return returnNode;
	}

	set(index, value) {
		const element = this.get(index);
		if (!element) {
			return;
		}
		element.value = value;
		return;
	}

	getValueOfNode(node) {
		return node ? node.value : null;
	}

	insert(index, value) {
		// index is not within the range
		if (index < 0 || index > this.length) return;

		// inserting at head
		if (index === 0) {
			const newNode = new Node(value);
			newNode.next = this.head;
			this.head = newNode;
			this.length += 1;
			return;
		}

		// inserting at tail
		if (index === this.length) {
			const newNode = new Node(value);
			newNode.next = null;
			this.tail.next = newNode;
			this.tail = newNode;
			this.length += 1;
			return;
		}

		// inserting in the middle
		let prevNode = null;
		let currNode = this.head;
		let counter = 0;
		while (currNode.next) {
			if (index === counter) {
				break;
			}
			counter += 1;
			prevNode = currNode;
			currNode = currNode.next;
		}
		
		const newNode = new Node(value);
		newNode.next = currNode;
		prevNode.next = newNode;
		this.length += 1;
		return;
	}

	remove(index) {
		if (index < 0 || index >= this.length) return;
		
		// remove head
		if (index === 0) {
			let currHeadNode = this.head;
			this.head = this.head.next;
			currHeadNode.next = null;
			this.length -= 1;
			return;
		}

		let prevNode = null;
		let currNode = this.head;
		let counter = 0;
		while (currNode.next) {
			if (index === counter) break;
			counter += 1;
			prevNode = currNode;
			currNode = currNode.next;
		}

		// if the node to remove is tail
		if (index === (this.length - 1)) {
			prevNode.next = null;
			this.tail = prevNode;
			this.length -= 1;
			return;
		}

		// last case removing something in the middle
		prevNode.next = currNode.next;
		currNode.next = null;
		this.length -= 1;
		return;
	}


	reverse() {
		if (!this.head || !this.head.next) return this;

		let currentNode = this.head;
		this.tail = this.head;
		let prevNode = null;
		while (currentNode) {
			let nextNode = currentNode.next;
			currentNode.next = prevNode;
			prevNode = currentNode;
			currentNode = nextNode;
		}
		this.head = prevNode;
		return;
	}

	reverseBetween(m, n) {
		if (!this.head || !this.head.next || m === n) return;

		let dummy = new Node(0);
		dummy.next = this.head;

		let counter = 0;
		let preReverse = dummy;
		while (counter < m) {
			preReverse = preReverse.next;
			counter++;
		}

		console.log(`counter: ${counter}   preReverse: ${JSON.stringify(preReverse)}`);

		let currentNode = preReverse.next;
		let prevNode = null;
		let nextNode = null;
		let tailOfReversed = currentNode;

		counter++; // move into the reversal loop
		while (counter <= n) {
			nextNode = currentNode.next;
			currentNode.next = prevNode;
			prevNode = currentNode;
			currentNode = nextNode;
			counter++;
		}

		// Stitch the reversed sublist back into the list
		preReverse.next = prevNode;
		tailOfReversed.next = currentNode;

		// Update head in case m === 0
		this.head = dummy.next;

		// update tail
		currentNode = this.head;
		while (currentNode.next) {
			currentNode = currentNode.next;
		}
		this.tail = currentNode;
		return;
	}


	reversePairs() {
		if (!this.head || !this.head.next) return;
		let dummy = new Node();
		dummy.next = this.head;

		let prevNode = dummy;
		while (prevNode && prevNode.next && prevNode.next.next) {
			let firstNode = prevNode.next;
			let secondNode = prevNode.next.next;

			// swap 
			firstNode.next = secondNode.next;
			secondNode.next = firstNode;
			prevNode.next = secondNode;

			// move cursor
			prevNode = firstNode;
		}

		this.head = dummy.next;
		return;
	}


	findMiddle() {
		if (!this.head) return null;
		let slow = this.head;
		let fast = this.head;

		while(fast && fast.next) {
			slow = slow.next;
			fast = fast.next.next;
		}
		return slow.value;
	}

	hasLoop() {
		if (!this.head) return false;
		let slow = this.head;
		let fast = this.head;

		while(fast && fast.next) {
			slow = slow.next;
			fast = fast.next.next;

			if (slow === fast) return true;
		}
		return false;
	}

	getKthElementFromEnd(index) {
		if (index <= 0 || index > this.length) return null; 
		if (!this.head) return;
		
		let first = this.head;
		let second = this.head;

		for (let i = 0; i < index; i++) {
			first = first.next;
		}

		while (first) {
			first = first.next;
			second = second.next;
		}
		return second.value;
	}

	removeKthElementFromEnd(index) {
		if (index <= 0 || index > this.length) return null; 
		if (!this.head) return;

		
		let first = this.head;
		let second = this.head;
		let prev = null;
		for (let i = 0; i < index; i++) {
			first = first.next;
		}
        
		while (first) {
			first = first.next;
			prev = second;
			second = second.next;
		}
		// we are removing head, hence prev is not moved forward
		if (!prev) {
			
			this.head = second.next;
			second.next = null;
			this.length -= 1;
			return;	
		}
		prev.next = second.next;
		second.next = null;

		// special case where tail is outdated in LL
		if (index === 1) {
			this.tail = prev;
		}
		this.length -= 1;
		return;
	}

	removeDuplicates() {
		if (!this.head) return;
		let currentNode = this.head;

		while (currentNode) {
			let runnerNode = currentNode;
			while (runnerNode.next) {
				if (runnerNode.next.value === currentNode.value) {
					runnerNode.next = runnerNode.next.next;
					this.length -= 1;
				} else {
					runnerNode = runnerNode.next;
				}

			}
			this.tail = currentNode;
			currentNode = currentNode.next;
			
		}
	}

	partitionList(x) {
		let dummyOne = new Node(0);
		let dummyTwo = new Node(0);
		let prevOne = dummyOne;
		let prevTwo = dummyTwo;
		let currentNode = this.head;
		
		while (currentNode) {
			let newNode = new Node(currentNode.value);
			if (currentNode.value < x) {
				prevOne.next = newNode;
				prevOne = prevOne.next;
			} else {
				prevTwo.next = newNode;
				prevTwo = prevTwo.next;
			}
			currentNode = currentNode.next;
		}
		
		// join two lists
		prevOne.next = dummyTwo.next; // prevTwo.next is 0 set in the beginings as dummy

		// lets create a new LL
		const newLL = new LinkedList();
		newLL.length = 0;// no node added in above initialization
		newLL.head = dummyOne.next; // dummyOne 0th is 0 set above, hence ignoring that
		currentNode = dummyOne.next;
		while (currentNode) {
			newLL.push(currentNode.value);
			currentNode = currentNode.next;
		}
		return newLL;
	}
	


}

// const newLl1 = new LinkedList(4);

// console.log(JSON.stringify(newLl1));
// console.log("AFTER ADDING MANY ELEMENTS");
// newLl1.push(5);
// newLl1.push(9);
// newLl1.push(7);
// newLl1.push(6);
// console.log(JSON.stringify(newLl1));
// console.log("AFTER POPPING EN ELEMENT");
// newLl1.pop();
// console.log(JSON.stringify(newLl1));

//******  testing unshifting with zero elements */
// newLl1.pop();
// console.log(JSON.stringify(newLl1));
//******  testing unshifting with zero elements */

// newLl1.unshift(22);
// newLl1.unshift(44);
// console.log("AFTER UNSHIFTING EN ELEMENT");
// console.log(JSON.stringify(newLl1));




// //******  testing shifting with zero elements */
// newLl1.shift();
// console.log("AFTER SHIFTING EN ELEMENT");
// console.log(JSON.stringify(newLl1));
// //******  testing shifting with zero elements */

// console.log("\n\n\n\nTESTING SHIFTING WITH CASES");
// // 1 with zero elem
// const newLl2 = new LinkedList(2);
// console.log(JSON.stringify(newLl2));
// newLl2.pop();
// console.log(JSON.stringify(newLl2));
// newLl2.shift();
// console.log(JSON.stringify(newLl2));
// newLl2.push(2);
// console.log("adding one element");
// console.log(JSON.stringify(newLl2));
// newLl2.shift();
// console.log("after shifting one element");
// console.log(JSON.stringify(newLl2));
// //******  testing shifting with zero elements */

// console.log(`index 0: ${newLl1.get(0) ? newLl1.get(0).value : null}`);
// console.log(`index 0: ${newLl1.getValueOfNode(newLl1.get(0))}`);
// console.log(`index 1: ${newLl1.getValueOfNode(newLl1.get(1))}`);
// console.log(`index 4: ${newLl1.getValueOfNode(newLl1.get(4))}`);
// console.log(`index 5: ${newLl1.getValueOfNode(newLl1.get(5))}`);
// console.log(`index 10: ${newLl1.getValueOfNode(newLl1.get(10))}`);




// newLl1.set(5, 55);
// newLl1.set(0, 0);
// newLl1.set(1, 11);
// newLl1.set(99, 99);
// console.log(`AFTER SETTING ELEMENT`);
// console.log(JSON.stringify(newLl1));




// //******  testing inserting and removing elements */
// console.log("inserting 22 at 1st index");
// newLl1.insert(1, 22);
// console.log("inserting 224 at 4th index");
// newLl1.insert(4, 224);
// console.log(`AFTER INSERTING ELEMENTS`);
// console.log(JSON.stringify(newLl1));

// console.log("inserting 1122 at 0th index");
// newLl1.insert(0, 1122);
// console.log("inserting 2211 at 9th index");
// newLl1.insert(9, 2211);
// console.log(`AFTER INSERTING oth and 9th ELEMENTS`);
// console.log(JSON.stringify(newLl1));



// // removing 0th element
// newLl1.remove(0);
// console.log(JSON.stringify(newLl1));
// newLl1.remove(newLl1.length-1);
// console.log(JSON.stringify(newLl1));

// newLl1.remove(2);
// console.log(JSON.stringify(newLl1));


// //******  testing inserting and removing elements */


// //******  testing reversing elements */
// console.log("\n\n\n");
// let newLl3 = new LinkedList(1);
// newLl3.reverse();
// console.log(JSON.stringify(newLl3));
// newLl3.push(2);
// newLl3.push(3);
// newLl3.push(4);
// newLl3.push(5);
// newLl3.push(6);
// newLl3.push(7);
// console.log("BEFORE REVERSE");
// console.log(JSON.stringify(newLl3));
// newLl3.reverse();
// console.log("AFTER REVERSE");
// console.log(JSON.stringify(newLl3));

// console.log(newLl3.findMiddle());
// newLl3.pop();
// console.log(newLl3.findMiddle());
// console.log("\n\n\n");
// //******  testing reversing elements */


// /****** TESTING CYCLE IN A LL */
// console.log("\n\n\n");
// let newLl4 = new LinkedList(1);
// newLl4.push(2);
// newLl4.push(3);
// newLl4.push(4);
// newLl4.push(5);
// const nodeAtGivenIndex = newLl4.get(2);
// newLl4.tail.next = nodeAtGivenIndex;
// // console.log(`${JSON.stringify(newLl4)}`); // should produce error
// console.log(newLl4.hasLoop());
// console.log("\n\n\n");
// /****** TESTING CYCLE IN A LL */

// // /****** TESTING getKthElementFromEnd IN A LL */
// console.log("\n\n\n");
// let newLl4 = new LinkedList(1);
// newLl4.push(2);
// newLl4.push(3);
// newLl4.push(4);
// newLl4.push(5);
// console.log(JSON.stringify(newLl4));
// console.log(newLl4.getKthElementFromEnd(5));
// console.log(newLl4.getKthElementFromEnd(4));
// console.log(newLl4.getKthElementFromEnd(3));
// console.log(newLl4.getKthElementFromEnd(2));
// console.log(newLl4.getKthElementFromEnd(1));
// console.log(newLl4.getKthElementFromEnd(6));
// console.log(newLl4.getKthElementFromEnd(0));
// console.log("\n\n\n");
// // /****** TESTING getKthElementFromEnd IN A LL */



// // // /****** TESTING REMOVEKTHFROMEND_DUMMY IN A LL */
// console.log("\n\n\n");
// let newLl4 = new LinkedList(1);
// newLl4.push(2);
// newLl4.push(3);
// newLl4.push(4);
// newLl4.push(5);
// console.log(JSON.stringify(newLl4));
// console.log("removing");
// newLl4.removeKthElementFromEnd(1);
// // newLl4.removeKthElementFromEnd(3);
// console.log(JSON.stringify(newLl4));
// console.log("\n\n\n");
// // // /****** TESTING REMOVEKTHFROMEND_DUMMY IN A LL */


// // // /****** TESTING REMOVE_DUPLICATES IN A LL */
// console.log("\n\n\n");
// let newLl4 = new LinkedList(1);
// newLl4.push(2);
// newLl4.push(3);
// newLl4.push(2);
// newLl4.push(4);
// newLl4.push(3);
// newLl4.push(2);
// newLl4.push(5);
// newLl4.push(3);
// newLl4.push(3);
// newLl4.push(2);
// newLl4.push(3);
// newLl4.push(3);
// console.log(JSON.stringify(newLl4));
// console.log("AFTER REMOVING THE DUPLICATES");
// newLl4.removeDuplicates();
// console.log(JSON.stringify(newLl4));
// // // /****** TESTING REMOVE_DUPLICATES IN A LL */




// /****** TESTING PARTITION_LIST IN A LL */

// let newLl4 = new LinkedList(3);
// newLl4.push(8);
// newLl4.push(5);
// newLl4.push(10);
// newLl4.push(2);
// newLl4.push(1);
// console.log(JSON.stringify(newLl4));
// console.log("AFTER PARTITION_LIST");
// newLl4 = newLl4.partitionList(5);
// console.log(JSON.stringify(newLl4));
// /****** TESTING PARTITION_LIST IN A LL */


// /****** TESTING REVERSE_BETWEEN IN A LL */
// console.log("\n\n\n");
// let newLl4 = new LinkedList(1);
// [2,3,4,5,6,7,8,9,10,11].forEach(x => newLl4.push(x));
// console.log(JSON.stringify(newLl4));
// newLl4.reverseBetween(0,10);
// console.log("\n");
// console.log(JSON.stringify(newLl4));
// /****** TESTING REVERSE_BETWEEN IN A LL */


/****** TESTING REVERSE_BETWEEN IN A LL */
console.log("\n\n\n");
let newLl4 = new LinkedList(1);
[2,3,4,5,6,7,8].forEach(x => newLl4.push(x));
console.log(JSON.stringify(newLl4));
newLl4.reversePairs();
console.log("\n");
console.log(JSON.stringify(newLl4));
/****** TESTING REVERSE_BETWEEN IN A LL */