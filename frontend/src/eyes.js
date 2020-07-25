var w = window.innerWidth;
var h = window.innerHeight;
const numCircles = 60;

let ground;
let wall1;
let wall2;

const content = document.querySelector('.eyeball-physics');

// Matter.js module aliases
let elements = [];
let eyeballs = [];

window.addEventListener('resize', (e) => {
  w = window.innerWidth;
  h = window.innerHeight;
  engine.render.canvas.width = w;
  engine.render.canvas.height = h;

  Matter.Body.setPosition(wall2, Matter.Vector.create(w + 30, h * .5));
  Matter.Body.setPosition(ground, Matter.Vector.create(w * .5, h + 30));

});

// create a Matter.js engine
var engine = Matter.Engine.create(content, {
  render: {
    options: {
      width: w,
      height: h,
      wireframes: false,
      background: "#000000"
    }
  }
});

window.engine = engine;

var mouseConstraint = Matter.MouseConstraint.create(engine, {
  constraint: {
    render: {
      visible: false
    },
    stiffness:0.1
  }
});

let spritesArea = document.querySelector('.eyeball-sprites');

class Eyeball {
  constructor() {
    var x = Math.random() * w;
    var y = Math.random() * - h;
    var base = w / 30;
    if(base < 5) base = 5;
    if(base > 10) base = 10;
    var multiplier = w / 10;
    if(multiplier < 30) multiplier = 30;
    if(multiplier > 100) multiplier = 100;

    this.radius = base + (Math.random() * multiplier);
    this.body = Matter.Bodies.circle(x, y, this.radius,
      {render: {
         fillStyle: 'black'
      }});

    this.element = document.createElement('div');
    this.element.className = 'eyeball ' + 'eyeball--' + Math.floor(Math.random() * 5);
    this.element.style.width = this.radius * 2 + 'px';
    this.element.style.height = this.radius * 2 + 'px';
    this.cornea = document.createElement('div');
    this.element.appendChild(this.cornea)
    spritesArea.appendChild(this.element);
  }

  update() {
    this.pos = {x: this.body.position.x, y: this.body.position.y}
    this.element.style.transform = `translate(${this.pos.x - this.radius - 8}px, ${this.pos.y - this.radius - 8}px)`;
  }

  lookAt(pos) {
    let diff = { x: pos.x - this.pos.x, y: pos.y - this.pos.y };
    let polar = [
      Math.sqrt(
        diff.x * diff.x + diff.y * diff.y
      ),
      Math.atan2(diff.y, diff.x)
    ];
    let dist = polar[0] < this.radius * .5 ? polar[0] : this.radius * .5;
    this.cornea.style.transform = `translate(${Math.cos(polar[1]) * dist}px, ${Math.sin(polar[1]) * dist}px)`;

    window.cornea = `translate(${Math.cos(polar[1]) * dist}px, ${Math.sin(polar[1]) * dist}px)`;
    window.polar = polar
  }
}

let mousepos = {x: 0, y: 0};

window.addEventListener('pointermove', (e) => {
  mousepos = {x: e.clientX, y: e.clientY};
});

// create two boxes and a ground
for(var i = 0; i < numCircles; i++)
{
  eyeballs.push(new Eyeball);
}
ground = Matter.Bodies.rectangle(w/2, h+30, 50000., 60, { isStatic: true });
wall1 = Matter.Bodies.rectangle(-30, h/2, 60, h*2, { isStatic: true });
wall2 = Matter.Bodies.rectangle(w+30, h/2, 60, h*2, { isStatic: true });
window.wall2 = wall2;
elements.push(ground);
elements.push(wall1);
elements.push(wall2);

// add all of the bodies to the world
console.log(eyeballs.map(eyeball => eyeball.body).concat(elements));
Matter.World.add(engine.world, eyeballs.map(eyeball => eyeball.body).concat(elements));
Matter.World.add(engine.world, mouseConstraint);

// run the engine
Matter.Engine.run(engine);

Matter.Events.on(engine, "afterUpdate", () => {
  eyeballs.forEach((eye) => {
    eye.update();
    eye.lookAt(mousepos);
  });
});