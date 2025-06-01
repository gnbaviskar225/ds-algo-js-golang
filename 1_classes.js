class Coockie {
    constructor(color) {
        this.color = color
    }
    getColor() {
        return this.color
    }
    setColor(color) {
        this.color = color
    }

}

let coockieOne = new Coockie("green")
console.log(coockieOne.color) // green

coockieOne.setColor("yellow")
console.log(coockieOne.getColor()) // yellow

